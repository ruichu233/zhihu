package logic

import (
	"context"
	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailLoginLogic {
	return &EmailLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailLoginLogic) EmailLogin(req *types.EmailLoginRequest) (resp *types.EmailLoginResponse, err error) {
	resp = new(types.EmailLoginResponse)
	// 查询用户是否存在
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)

	resp = &types.EmailLoginResponse{}
	return
}
