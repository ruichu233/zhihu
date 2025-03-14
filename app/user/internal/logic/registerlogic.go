package logic

import (
	"context"
	"errors"
	client "github.com/gorse-io/gorse-go"
	"github.com/yitter/idgenerator-go/idgen"
	"gorm.io/gorm"
	"strconv"
	"zhihu/app/user/model"
	"zhihu/pkg/token"
	"zhihu/pkg/utils"

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
	var u model.User
	err := l.svcCtx.DB.Model(model.User{}).Limit(1).Find(&u, "email = ?", in.Email).Error
	if err != nil {
		return nil, err
	}
	if u.Id > 0 {
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
	l.svcCtx.RDB.Del(l.ctx, in.Email)
	// 4、注册用户
	u.Id = idgen.NextId()
	u.Email = in.Email
	u.Password = utils.Md5Crypt(in.Password)
	u.Username = in.Username
	u.Avatar = "https://avatars.githubusercontent.com/u/89794548?v=4"
	// 5、开启事务保存用户信息
	l.svcCtx.DB.Session(&gorm.Session{}).Transaction(func(tx *gorm.DB) error {
		err = l.svcCtx.DB.Model(&model.User{}).Create(&u).Error
		if err != nil {
			return err
		}
		rowAffected, err := l.svcCtx.Gorse.InsertUser(l.ctx, client.User{
			UserId:    strconv.FormatInt(u.Id, 10),
			Labels:    []string{},
			Subscribe: []string{},
			Comment:   "",
		})
		if err != nil {
			return err
		}
		if rowAffected.RowAffected <= 0 {
			logx.Errorf("用户 %d 推送到gorse失败", u.Id)
		}
		return nil
	})
	l.svcCtx.RDB.Del(l.ctx, in.Email)
	// 5、返回token
	tokenString, err := token.Sign(strconv.FormatInt(u.Id, 10))
	if err != nil {
		return nil, err
	}
	return &user.RegisterResponse{
		UserId:      u.Id,
		AccessToken: tokenString,
	}, nil
}
