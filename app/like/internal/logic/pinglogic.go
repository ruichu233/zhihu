package logic

import (
	"context"

	"zhihu/app/like/internal/svc"
	"zhihu/app/like/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *like.Request) (*like.Response, error) {
	// todo: add your logic here and delete this line

	return &like.Response{}, nil
}
