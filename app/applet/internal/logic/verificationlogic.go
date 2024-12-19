package logic

import (
	"context"
	"zhihu/app/user/userclient"

	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerificationLogic {
	return &VerificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerificationLogic) Verification(req *types.VerificationRequest) (resp *types.VerificationResponse, err error) {
	_, err = l.svcCtx.UserRPC.SendVerifyCode(l.ctx, &userclient.SendVerifyCodeRequest{
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &types.VerificationResponse{}, nil
}
