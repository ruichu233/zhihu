package logic

import (
	"context"
	"errors"
	"strconv"
	"zhihu/app/user/model"
	"zhihu/pkg/token"
	"zhihu/pkg/utils"

	"zhihu/app/user/internal/svc"
	"zhihu/app/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// 1、检查用户是否存在
	var u model.User
	err := l.svcCtx.DB.Model(&model.User{}).Limit(1).Find(&u, "email = ?", in.Email).Error
	if err != nil {
		return nil, err
	}

	// 2、检查密码是否正确
	if utils.Md5Crypt(u.Password) != in.Password {
		return nil, errors.New("密码错误")
	}
	// 3、生成token
	accessToken, err := token.Sign(strconv.FormatInt(u.Id, 10))
	if err != nil {
		return nil, err
	}
	return &user.LoginResponse{
		UserId:      u.Id,
		AccessToken: accessToken,
	}, nil
}
