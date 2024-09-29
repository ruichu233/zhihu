package logic

import (
	"context"

	"zhihu/app/user/internal/svc"
	"zhihu/app/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFollowerListLogic {
	return &GetUserFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFollowerListLogic) GetUserFollowerList(in *user.FollowerListRequest) (*user.FollowerListResponse, error) {

	return &user.FollowerListResponse{}, nil
}
