package logic

import (
	"context"

	"zhihu/app/follow/internal/svc"
	"zhihu/app/follow/pb/follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckServiceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckServiceLogic {
	return &CheckServiceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckServiceLogic) CheckService(in *follow.PingRequest) (*follow.PingResponse, error) {
	// todo: add your logic here and delete this line

	return &follow.PingResponse{}, nil
}
