package logic

import (
	"context"
	"fmt"
	"zhihu/app/user/model"
	"zhihu/pkg/utils"

	"zhihu/app/user/internal/svc"
	"zhihu/app/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoUpdateLogic {
	return &UserInfoUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoUpdateLogic) UserInfoUpdate(in *user.UserInfoUpdateRequest) (*user.UserInfoUpdateResponse, error) {
	// 1、获取用户数据
	var _user model.User
	if err := l.svcCtx.DB.Model(&model.User{}).Where("id=?", in.UserId).Limit(1).Find(&_user).Error; err != nil {
		return nil, err
	}
	if in.Username != "" {
		_user.Username = in.Username
	}
	if in.Avatar != "" {
		_user.Avatar = in.Avatar
	}
	if in.Signature != "" {
		_user.Signature = in.Signature
	}
	if in.Password != "" {
		crypt := utils.Md5Crypt(in.OldPassword)
		if crypt != _user.Password {
			return nil, fmt.Errorf("密码错误")
		}
		_user.Password = utils.Md5Crypt(in.Password)
	}
	if err := l.svcCtx.DB.Model(&model.User{}).Where("id=?", in.UserId).Updates(_user).Error; err != nil {
		return nil, err
	}
	return &user.UserInfoUpdateResponse{
		Message: "更新成功",
	}, nil
}
