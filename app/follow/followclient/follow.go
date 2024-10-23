// Code generated by goctl. DO NOT EDIT.
// Source: follow.proto

package followclient

import (
	"context"

	"zhihu/app/follow/follow"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = follow.Request
	Response = follow.Response

	Follow interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultFollow struct {
		cli zrpc.Client
	}
)

func NewFollow(cli zrpc.Client) Follow {
	return &defaultFollow{
		cli: cli,
	}
}

func (m *defaultFollow) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := follow.NewFollowClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
