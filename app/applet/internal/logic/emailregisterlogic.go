package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"
	"zhihu/app/user/userclient"
)

type EmailRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailRegisterLogic {
	return &EmailRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailRegisterLogic) EmailRegister(req *types.EmailRegisterRequest) (resp *types.EmailRegisterResponse, err error) {
	if req.Code = strings.TrimSpace(req.Code); len(req.Code) == 0 {
		return nil, errors.New("验证码不能为空")
	}
	if req.Email = strings.TrimSpace(req.Email); len(req.Email) == 0 {
		return nil, errors.New("邮箱不能为空")
	}
	if req.Password = strings.TrimSpace(req.Password); len(req.Password) == 0 {
		return nil, errors.New("密码不能为空")
	}
	regResp, err := l.svcCtx.UserRPC.Register(l.ctx, &userclient.RegisterRequest{
		Username: strings.Split(req.Email, "@")[0],
		Email:    req.Email,
		Password: req.Password,
		Code:     req.Code,
	})
	if err != nil {
		return nil, err
	}

	return &types.EmailRegisterResponse{
		AccessToken: regResp.AccessToken,
		UserId:      regResp.UserId,
	}, nil
}
