package logic

import (
	"context"
	"zhihu/app/like/model"

	"zhihu/app/like/internal/svc"
	"zhihu/app/like/pb/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLikesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLikesLogic {
	return &GetUserLikesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询某个用户的点赞列表
func (l *GetUserLikesLogic) GetUserLikes(in *like.GetUserLikesRequest) (*like.GetUserLikesResponse, error) {
	// 1、查询缓存
	key := model.GetLikeRecordKey("short_video", in.UserId)

	return &like.GetUserLikesResponse{}, nil
}
