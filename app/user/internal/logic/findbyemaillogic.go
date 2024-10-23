package logic

import (
	"context"
	"fmt"
	"strings"
	"zhihu/app/user/model"

	"zhihu/app/user/internal/svc"
	"zhihu/app/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByEmailLogic {
	return &FindByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByEmailLogic) FindByEmail(in *user.FindByEmailRequest) (*user.FindByEmailResponse, error) {
	in.Email = strings.TrimSpace(in.Email)
	if len(in.Email) == 0 {
		return nil, fmt.Errorf("邮箱不能为空")
	}
	var users model.Users
	if err := l.svcCtx.DB.Model(&model.Users{}).Limit(1).Find(&users, "email = ?", in.Email).Error; err != nil {
		return nil, err
	}
	return &user.FindByEmailResponse{
		UserId:   users.Id,
		Email:    users.Email,
		Username: users.UserName,
		Password: users.Password,
	}, nil
}
