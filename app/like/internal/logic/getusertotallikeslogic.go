package logic

import (
	"context"

	"zhihu/app/like/internal/svc"
	"zhihu/app/like/pb/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserTotalLikesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserTotalLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserTotalLikesLogic {
	return &GetUserTotalLikesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户收到的总点赞数
func (l *GetUserTotalLikesLogic) GetUserTotalLikes(in *like.GetUserTotalLikesRequest) (*like.GetUserTotalLikesResponse, error) {
	// todo: add your logic here and delete this line

	return &like.GetUserTotalLikesResponse{}, nil
}
