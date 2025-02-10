package logic

import (
	"context"
	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"
	"zhihu/app/user/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(userId int64) (resp *types.UserInfoResponse, err error) {
	userInfo, err := l.svcCtx.UserRPC.GetUserInfo(l.ctx, &userclient.UserInfoRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResponse{
		UserId:        userInfo.Id,
		UserName:      userInfo.Username,
		Email:         userInfo.Email,
		Avatar:        userInfo.Avatar,
		Signature:     userInfo.Signature,
		FollowerCount: userInfo.FollowerCount,
		FollowedCount: userInfo.FollowedCount,
	}, nil
}
