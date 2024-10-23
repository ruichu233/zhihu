package logic

import (
	"context"
	"zhihu/app/user/model"
	"zhihu/pkg/model/follow"

	"zhihu/app/user/internal/svc"
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
	var u model.Users
	if err := l.svcCtx.DB.Model(&model.Users{}).Limit(1).Find(&u, in.UserId).Error; err != nil {
		return nil, err
	}
	// 2、根据用户id查询用户关注数和粉丝数
	var followCount int64
	if err := l.svcCtx.DB.Model(&follow.Follows{}).Where("follow_id = ?", in.UserId).Count(&followCount).Error; err != nil {
		return nil, err
	}
	return &user.UserInfoResponse{
		Username: u.UserName,
		Email:    u.Email,
	}, nil
}
