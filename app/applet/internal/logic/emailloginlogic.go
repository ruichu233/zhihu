package logic

import (
	"context"
	"fmt"
	"strings"
	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"
	"zhihu/app/user/userclient"
	"zhihu/pkg/utils"

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
	if req.Email = strings.TrimSpace(req.Email); len(req.Email) == 0 {
		return nil, fmt.Errorf("邮箱不能为空")
	}
	if req.Password = strings.TrimSpace(req.Password); len(req.Password) == 0 {
		return nil, fmt.Errorf("密码不能为空")
	}
	enPassword := utils.Md5Crypt(req.Password)
	loginResp, err := l.svcCtx.UserRPC.Login(l.ctx, &userclient.LoginRequest{
		Email:    req.Email,
		Password: enPassword,
	})
	if err != nil {
		return nil, err
	}
	return &types.EmailLoginResponse{
		UserId:      loginResp.UserId,
		AccessToken: loginResp.AccessToken,
	}, nil
}
