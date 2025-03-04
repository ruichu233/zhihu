// Code generated by goctl. DO NOT EDIT.
// Source: follow.proto

package followclient

import (
	"context"

	"zhihu/app/follow/pb/follow"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FollowActionRequest     = follow.FollowActionRequest
	FollowActionResponse    = follow.FollowActionResponse
	FollowItem              = follow.FollowItem
	GetFollowListRequest    = follow.GetFollowListRequest
	GetFollowListResponse   = follow.GetFollowListResponse
	GetFollowerListRequest  = follow.GetFollowerListRequest
	GetFollowerListResponse = follow.GetFollowerListResponse
	GetFriendListRequest    = follow.GetFriendListRequest
	GetFriendListResponse   = follow.GetFriendListResponse
	IsFollowRequest         = follow.IsFollowRequest
	IsFollowResponse        = follow.IsFollowResponse
	PingRequest             = follow.PingRequest
	PingResponse            = follow.PingResponse

	Follow interface {
		CheckService(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
		FollowAction(ctx context.Context, in *FollowActionRequest, opts ...grpc.CallOption) (*FollowActionResponse, error)
		ListFollowing(ctx context.Context, in *GetFollowListRequest, opts ...grpc.CallOption) (*GetFollowListResponse, error)
		ListFollowers(ctx context.Context, in *GetFollowerListRequest, opts ...grpc.CallOption) (*GetFollowerListResponse, error)
		ListFriends(ctx context.Context, in *GetFriendListRequest, opts ...grpc.CallOption) (*GetFriendListResponse, error)
		IsFollow(ctx context.Context, in *IsFollowRequest, opts ...grpc.CallOption) (*IsFollowResponse, error)
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

func (m *defaultFollow) CheckService(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	client := follow.NewFollowClient(m.cli.Conn())
	return client.CheckService(ctx, in, opts...)
}

func (m *defaultFollow) FollowAction(ctx context.Context, in *FollowActionRequest, opts ...grpc.CallOption) (*FollowActionResponse, error) {
	client := follow.NewFollowClient(m.cli.Conn())
	return client.FollowAction(ctx, in, opts...)
}

func (m *defaultFollow) ListFollowing(ctx context.Context, in *GetFollowListRequest, opts ...grpc.CallOption) (*GetFollowListResponse, error) {
	client := follow.NewFollowClient(m.cli.Conn())
	return client.ListFollowing(ctx, in, opts...)
}

func (m *defaultFollow) ListFollowers(ctx context.Context, in *GetFollowerListRequest, opts ...grpc.CallOption) (*GetFollowerListResponse, error) {
	client := follow.NewFollowClient(m.cli.Conn())
	return client.ListFollowers(ctx, in, opts...)
}

func (m *defaultFollow) ListFriends(ctx context.Context, in *GetFriendListRequest, opts ...grpc.CallOption) (*GetFriendListResponse, error) {
	client := follow.NewFollowClient(m.cli.Conn())
	return client.ListFriends(ctx, in, opts...)
}

func (m *defaultFollow) IsFollow(ctx context.Context, in *IsFollowRequest, opts ...grpc.CallOption) (*IsFollowResponse, error) {
	client := follow.NewFollowClient(m.cli.Conn())
	return client.IsFollow(ctx, in, opts...)
}
