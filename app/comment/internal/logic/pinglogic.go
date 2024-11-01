package logic

import (
	"context"

	"zhihu/app/comment/comment"
	"zhihu/app/comment/internal/svc"

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

func (l *PingLogic) Ping(in *comment.Request) (*comment.Response, error) {
	// todo: add your logic here and delete this line

	return &comment.Response{}, nil
}
