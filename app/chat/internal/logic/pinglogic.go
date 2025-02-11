package logic

import (
	"context"

	"zhihu/app/chat/internal/svc"
	"zhihu/app/chat/pb/chat"

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

func (l *PingLogic) Ping(in *chat.Request) (*chat.Response, error) {
	// todo: add your logic here and delete this line

	return &chat.Response{}, nil
}
