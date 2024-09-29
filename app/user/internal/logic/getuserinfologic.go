package logic

import (
	"context"
	user_model "zhihu/pkg/model/user"

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
	var u user_model.Users
	if err := l.svcCtx.DB.Model(&user_model.Users{}).Find(&u, "id = ?", in.UserId).Error; err != nil {
		return nil, err
	}
	// 2、根据用户id查询用户关注数和粉丝数
	var followCount int64
	l.svcCtx.DB.Model(&user_model.Follows{}).Where("follow_id = ?", in.UserId).Count(&followCount)
	return &user.UserInfoResponse{
		Username: u.UserName,
		Email:    u.Email,
	}, nil
}
