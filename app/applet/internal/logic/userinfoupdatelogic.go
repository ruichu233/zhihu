package logic

import (
	"context"
	"zhihu/app/user/pb/user"

	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoUpdateLogic {
	return &UserInfoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoUpdateLogic) UserInfoUpdate(req *types.UserInfoUpdateRequest) (resp *types.UserInfoUpdateResponse, err error) {

	userInfoUpdateResponse, err := l.svcCtx.UserRPC.UserInfoUpdate(l.ctx, &user.UserInfoUpdateRequest{
		UserId:      req.UserId,
		Avatar:      req.Avatar,
		OldPassword: req.OldPassword,
		Password:    req.Password,
		Signature:   req.Signature,
		Username:    req.UserName,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.UserInfoUpdateResponse{
		Status: userInfoUpdateResponse.Message,
	}
	return
}
