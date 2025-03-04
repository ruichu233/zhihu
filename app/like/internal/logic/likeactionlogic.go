package logic

import (
	"context"
	"fmt"
	"sync"
	"zhihu/app/like/internal/svc"
	"zhihu/app/like/local_queue"
	"zhihu/app/like/model"
	"zhihu/app/like/pb/like"
	"zhihu/app/like/types"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var lock sync.Mutex

const LikeActionKey = "HASH_LIKE_RECORD"

func NewLikeActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeActionLogic {
	return &LikeActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikeActionLogic) LikeAction(in *like.LikeActionRequest) (*like.LikeActionResponse, error) {
	resp := &like.LikeActionResponse{}
	// 1、获取当前用户对目标是否点赞
	isLike, err := IsLike(l.ctx, l.svcCtx.RDB, l.svcCtx.DB, in.BizId, in.UserId, in.ObjId)
	if err != nil {
		return nil, err
	}
	// 2、判断异常情况
	switch in.ActionType {
	case like.LikeActionRequest_LIKE:
		if isLike {
			return resp, nil
		}
	case like.LikeActionRequest_UNLIKE:
		if !isLike {
			return resp, nil
		}
	}
	// 3、投入处理队列
	value := types.LikeAction{
		UserId:     in.UserId,
		BizId:      in.BizId,
		ObjId:      in.ObjId,
		ActionType: in.ActionType,
	}
	local_queue.LikeQueue.Push(&value)
	return &like.LikeActionResponse{
		LikeCount: 0,
	}, nil
}

// IsLike 查询是否对单个obj点过赞
func IsLike(ctx context.Context, rdb *redis.Client, db *gorm.DB, bizId string, userId, objId int64) (bool, error) {
	key := model.GetLikeRecordKey(bizId, userId)
	member := fmt.Sprintf("%d", objId)
	score, _ := rdb.ZScore(ctx, key, member).Result()
	if score != 0 {
		if score == -1 {
			return false, nil
		}
		return true, nil
	}
	var likeRecord model.LikeRecord
	if err := db.Model(&model.LikeRecord{}).Where("biz_id = ? and obj_id = ? and user_id = ?", bizId, objId, userId).Limit(1).Find(&likeRecord).Error; err != nil {
		return false, err
	}
	var isLike bool
	if likeRecord.Id != 0 {
		isLike = true
	}
	if isLike {
		_ = rdb.ZAdd(ctx, key, redis.Z{
			Score:  float64(likeRecord.UpdatedAt.Unix()),
			Member: member,
		}).Err()
	} else {
		_ = rdb.ZAdd(ctx, key, redis.Z{
			Score:  -1,
			Member: member,
		}).Err()
	}
	return true, nil
}
