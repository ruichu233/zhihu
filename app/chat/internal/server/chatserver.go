// Code generated by goctl. DO NOT EDIT.
// Source: chat.proto

package server

import (
	"context"

	"zhihu/app/chat/internal/logic"
	"zhihu/app/chat/internal/svc"
	"zhihu/app/chat/pb/chat"
)

type ChatServer struct {
	svcCtx *svc.ServiceContext
	chat.UnimplementedChatServer
}

func NewChatServer(svcCtx *svc.ServiceContext) *ChatServer {
	return &ChatServer{
		svcCtx: svcCtx,
	}
}

func (s *ChatServer) Ping(ctx context.Context, in *chat.Request) (*chat.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *ChatServer) GetMassages(ctx context.Context, in *chat.GetMassagesRequest) (*chat.GetMassagesResponse, error) {
	l := logic.NewGetMassagesLogic(ctx, s.svcCtx)
	return l.GetMassages(in)
}

func (s *ChatServer) SendMassage(ctx context.Context, in *chat.SendMassageRequest) (*chat.SendMassageResponse, error) {
	l := logic.NewSendMassageLogic(ctx, s.svcCtx)
	return l.SendMassage(in)
}
