package logic

import (
	"context"
	"strconv"
	"time"

	"zhihu/app/like/internal/svc"
	"zhihu/app/like/model"
	"zhihu/app/like/pb/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostLikeCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostLikeCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostLikeCountLogic {
	return &GetPostLikeCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询某个稿件的点赞数
func (l *GetPostLikeCountLogic) GetPostLikeCount(in *like.GetPostLikeCountRequest) (*like.GetPostLikeCountResponse, error) {
	// 先查缓存
	key := model.GetLikeCountKey(in.BizId, in.ObjId)
	result, _ := l.svcCtx.RDB.Get(l.ctx, key).Result()
	if result != "" {
		if result == "-1" {
			return &like.GetPostLikeCountResponse{
				Count: 0,
			}, nil
		} else {
			count, _ := strconv.ParseInt(result, 10, 64)
			return &like.GetPostLikeCountResponse{
				Count: count,
			}, nil
		}
	}
	// 缓存中没有数据，查数据库
	var likeCount model.LikeCount
	if err := l.svcCtx.DB.Model(&model.LikeCount{}).Where("biz_id = ? and obj_id = ?", in.BizId, in.ObjId).Limit(1).Find(&likeCount).Error; err != nil {
		return nil, err
	}
	// 缓存中更新数据
	if likeCount.Id == 0 {
		l.svcCtx.RDB.Set(l.ctx, key, "-1", 60*time.Minute)
	} else {
		l.svcCtx.RDB.Set(l.ctx, key, likeCount.LikeNum, 60*time.Minute)
	}
	return &like.GetPostLikeCountResponse{
		Count: likeCount.LikeNum,
	}, nil
}
