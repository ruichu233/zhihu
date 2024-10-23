package logic

import (
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"strconv"
	"zhihu/app/user/model"
	"zhihu/pkg/token"

	"zhihu/app/user/internal/svc"
	"zhihu/app/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 1、检测邮箱是否注册
	var u model.Users
	err := l.svcCtx.DB.Model(model.Users{}).Limit(1).Find(&u, "email = ?", in.Email).Error
	if err != nil {
		return nil, err
	}
	if u.Id != 0 {
		return nil, errors.New("邮箱已注册")
	}
	// 2、检查验证码是否过期
	result, err := l.svcCtx.RDB.Exists(l.ctx, in.Email).Result()
	if err != nil {
		return nil, err
	}
	if result == 0 {
		return nil, errors.New("验证码已过期")
	}

	// 3、检查验证码是否正确
	code, err := l.svcCtx.RDB.Get(l.ctx, in.Email).Result()
	if err != nil {
		return nil, err
	}
	if code != in.Code {
		return nil, errors.New("验证码不正确")
	}
	// 4、注册用户
	u.Id = idgen.NextId()
	u.Email = in.Email
	u.Password = in.Password
	u.UserName = in.Username
	err = l.svcCtx.DB.Model(&model.Users{}).Create(&u).Error
	if err != nil {
		return nil, err
	}
	l.svcCtx.RDB.Del(l.ctx, in.Email)
	// 5、返回token
	tokenString, err := token.Sign(strconv.FormatInt(u.Id, 10))
	if err != nil {
		return nil, err
	}
	return &user.RegisterResponse{
		AccessToken: tokenString,
	}, nil
}
