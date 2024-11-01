package logic

import (
	"context"

	"zhihu/app/feed/internal/svc"
	"zhihu/app/feed/pb/feed"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnfollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnfollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowLogic {
	return &UnfollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取消关注
func (l *UnfollowLogic) Unfollow(in *feed.UnfollowRequest) (*feed.UnfollowResponse, error) {
	// todo: add your logic here and delete this line

	return &feed.UnfollowResponse{}, nil
}
