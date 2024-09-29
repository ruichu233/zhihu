package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"strconv"
	user_model "zhihu/pkg/model/user"
	"zhihu/pkg/token"

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
	var u user_model.Users
	err := l.svcCtx.DB.Model(&user_model.Users{}).Find(&u, "email = ?", in.Email).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在")
	}
	// 2、检查密码是否正确
	if u.Password != in.Password {
		return nil, errors.New("密码错误")
	}
	// 3、生成token
	accessToken, err := token.Sign(strconv.FormatInt(u.Id, 10))
	if err != nil {
		return nil, err
	}
	return &user.LoginResponse{
		Token: accessToken,
	}, nil
}
