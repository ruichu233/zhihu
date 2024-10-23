// Code generated by goctl. DO NOT EDIT.
// Source: feed.proto

package server

import (
	"context"

	"zhihu/app/feed/feed"
	"zhihu/app/feed/internal/logic"
	"zhihu/app/feed/internal/svc"
)

type FeedServer struct {
	svcCtx *svc.ServiceContext
	feed.UnimplementedFeedServer
}

func NewFeedServer(svcCtx *svc.ServiceContext) *FeedServer {
	return &FeedServer{
		svcCtx: svcCtx,
	}
}

func (s *FeedServer) Ping(ctx context.Context, in *feed.Request) (*feed.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}