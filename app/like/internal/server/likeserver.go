// Code generated by goctl. DO NOT EDIT.
// Source: like.proto

package server

import (
	"context"

	"zhihu/app/like/internal/logic"
	"zhihu/app/like/internal/svc"
	"zhihu/app/like/like"
)

type LikeServer struct {
	svcCtx *svc.ServiceContext
	like.UnimplementedLikeServer
}

func NewLikeServer(svcCtx *svc.ServiceContext) *LikeServer {
	return &LikeServer{
		svcCtx: svcCtx,
	}
}

func (s *LikeServer) Ping(ctx context.Context, in *like.Request) (*like.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}