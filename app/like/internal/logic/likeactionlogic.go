package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
	"zhihu/app/like/internal/svc"
	"zhihu/app/like/model"
	"zhihu/app/like/pb/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const LikeActionKey = "HASH_LIKE_RECORD"

func NewLikeActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeActionLogic {
	return &LikeActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikeActionLogic) LikeAction(in *like.LikeActionRequest) (*like.LikeActionResponse, error) {
	// 1、获取当前用户对目标是否点赞
	isLike, err := IsLike(l.ctx, l.svcCtx.RDB, l.svcCtx.DB, in.BizId, in.UserId, in.ObjId)
	if err != nil {
		return nil, err
	}
	// 2、根据actionType判断是点赞还是取消点赞
	switch in.ActionType {
	case like.LikeActionRequest_LIKE:
		if isLike {
			return nil, errors.New("已经点赞过了")
		} else {
			// 2、更新数据库
			likeRecord := model.LikeRecord{
				BizId:  in.BizId,
				ObjId:  in.ObjId,
				UserId: in.UserId,
			}
			err = l.svcCtx.DB.Model(&model.LikeRecord{}).Create(&likeRecord).Error
			if err != nil {
				return nil, err
			}
			// 3、更新缓存
			l.svcCtx.RDB.ZAdd(l.ctx, model.GetLikeRecordKey(in.BizId, in.UserId), redis.Z{
				Member: in.ObjId,
				Score:  float64(time.Now().Unix()),
			})
		}
	case like.LikeActionRequest_UNLIKE:
		if !isLike {
			return nil, errors.New("没有点赞过")
		} else {
			// 2、更新数据库
			err = l.svcCtx.DB.Model(&model.LikeRecord{}).Where("obj_id = ? and biz_id = ? and user_id = ?", in.ObjId, in.BizId, in.UserId).Unscoped().Delete(&model.LikeRecord{}).Error
			if err != nil {
				return nil, err
			}
			// 3、更新缓存
		}
	}
	return &like.LikeActionResponse{
		LikeCount: 0,
	}, nil
}

// IsLike 查询是否对单个obj点过赞
func IsLike(ctx context.Context, rdb *redis.Client, db *gorm.DB, bizId string, userId, objId int64) (bool, error) {
	// 1、从缓存在获取当前用户对目标是否点赞
	key := model.GetLikeRecordKey(bizId, userId)
	member := fmt.Sprintf("%d", objId)
	exists, err := ExistsInSortedSet(ctx, rdb, key, member)
	if err != nil {
		return false, err
	}
	if exists {
		return true, nil
	}
	// 2、查询数据库获取当前用户对目标是否点赞
	var likeRecord model.LikeRecord
	err = db.Model(&model.LikeRecord{}).Unscoped().Where("obj_id = ? and biz_id = ? and user_id = ?", objId, bizId, userId).Limit(1).Find(&likeRecord).Error
	if err != nil {
		return false, err
	}
	if likeRecord.Id == 0 {
		return false, nil
	}
	if likeRecord.DeletedAt.Valid {
		return false, nil
	}
	// 3、更新缓存
	rdb.ZAdd(ctx, key, redis.Z{
		Member: member,
		Score:  float64(time.Now().Unix()),
	})
	return true, nil
}

func ExistsInSortedSet(ctx context.Context, rdb *redis.Client, key string, member string) (bool, error) {
	_, err := rdb.ZScore(ctx, key, member).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
