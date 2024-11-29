package logic

import (
	"context"
	"errors"
	"log"
	"strconv"
	"zhihu/app/user/internal/svc"
	"zhihu/app/user/model"
	"zhihu/app/user/pb/user"
	"zhihu/pkg/token"
	"zhihu/pkg/utils"

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
		log.Printf("err:%v", err)
		return nil, err
	}

	// 2、检查密码是否正确
	if u.Password != utils.Md5Crypt(in.Password) {
		log.Printf("密码错误")
		log.Println(in.Password)
		return nil, errors.New("密码错误")
	}
	// 3、生成token
	accessToken, err := token.Sign(strconv.FormatInt(u.Id, 10))
	if err != nil {
		log.Printf("err:%v", err)
		return nil, err
	}
	log.Printf("登录成功")
	return &user.LoginResponse{
		UserId:      u.Id,
		AccessToken: accessToken,
	}, nil
}
