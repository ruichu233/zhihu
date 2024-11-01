package logic

import (
	"context"
	"strconv"

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
	result, err := l.svcCtx.RDB.Get(l.ctx, key).Result()
	if err == nil {
		// 缓存中有数据，直接返回
		count, err := strconv.ParseInt(result, 10, 64)
		if err != nil {
			return nil, err
		}
		return &like.GetPostLikeCountResponse{
			Count: count,
		}, nil
	}
	// 缓存中没有数据，查数据库
	var likeCount model.LikeCount
	tx := l.svcCtx.DB.Model(&model.LikeCount{}).Where("biz_id = ? and obj_id = ?", in.BizId, in.ObjId).First(&likeCount)
	if tx.Error != nil {
		return nil, err
	}
	return &like.GetPostLikeCountResponse{
		Count: likeCount,
	}, nil
}
