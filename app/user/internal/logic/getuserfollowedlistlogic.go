package logic

import (
	"context"

	"zhihu/app/user/internal/svc"
	"zhihu/app/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFollowedListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFollowedListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFollowedListLogic {
	return &GetUserFollowedListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFollowedListLogic) GetUserFollowedList(in *user.FollowedListRequest) (*user.FollowerListResponse, error) {
	// todo: add your logic here and delete this line

	return &user.FollowerListResponse{}, nil
}
