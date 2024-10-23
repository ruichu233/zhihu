// Code generated by goctl. DO NOT EDIT.
// Source: like.proto

package likeclient

import (
	"context"

	"zhihu/app/like/like"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = like.Request
	Response = like.Response

	Like interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultLike struct {
		cli zrpc.Client
	}
)

func NewLike(cli zrpc.Client) Like {
	return &defaultLike{
		cli: cli,
	}
}

func (m *defaultLike) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := like.NewLikeClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}