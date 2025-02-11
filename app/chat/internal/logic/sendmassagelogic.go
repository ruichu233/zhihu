package logic

import (
	"context"

	"zhihu/app/chat/internal/svc"
	"zhihu/app/chat/pb/chat"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMassageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMassageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMassageLogic {
	return &SendMassageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMassageLogic) SendMassage(in *chat.SendMassageRequest) (*chat.SendMassageResponse, error) {

	return &chat.SendMassageResponse{}, nil
}
