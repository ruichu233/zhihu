package logic

import (
	"context"
	"zhihu/app/user/internal/svc"
	"zhihu/app/user/model"
	"zhihu/app/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// 1、查询用户信息
	var u model.User
	if err := l.svcCtx.DB.Model(&model.User{}).Limit(1).Find(&u, in.UserId).Error; err != nil {
		return nil, err
	}
	// 2、根据用户id查询用户关注数和粉丝数
	return &user.UserInfoResponse{
		Id:            u.Id,
		Username:      u.Username,
		Email:         u.Email,
		Avatar:        u.Avatar,
		Signature:     u.Signature,
		FollowerCount: 0,
		FollowedCount: 0,
	}, nil
}
