package logic

import (
	"context"

	"zhihu/app/chat/internal/svc"
	"zhihu/app/chat/pb/chat"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMassagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMassagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMassagesLogic {
	return &GetMassagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMassagesLogic) GetMassages(in *chat.GetMassagesRequest) (*chat.GetMassagesResponse, error) {
	
	return &chat.GetMassagesResponse{}, nil
}
