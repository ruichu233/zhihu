package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"
	"zhihu/app/user/userclient"
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
	resp = &types.EmailLoginResponse{}
	if req.Email = strings.TrimSpace(req.Email); len(req.Email) == 0 {
		return nil, fmt.Errorf("邮箱不能为空")
	}
	if req.Password = strings.TrimSpace(req.Password); len(req.Password) == 0 {
		return nil, fmt.Errorf("密码不能为空")
	}
	loginResp, err := l.svcCtx.UserRPC.Login(l.ctx, &userclient.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	resp.AccessToken = loginResp.AccessToken
	resp.UserId = loginResp.UserId
	return resp, nil
}
