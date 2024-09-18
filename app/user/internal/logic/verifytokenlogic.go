package logic

import (
	"context"

	"zhihu/app/user/internal/svc"
	"zhihu/app/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenLogic {
	return &VerifyTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyTokenLogic) VerifyToken(in *user.VerifyTokenRequest) (*user.VerifyTokenResponse, error) {
	// todo: add your logic here and delete this line

	return &user.VerifyTokenResponse{}, nil
}
