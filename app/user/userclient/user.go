// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"zhihu/app/user/pb/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FindByEmailRequest     = user.FindByEmailRequest
	FindByEmailResponse    = user.FindByEmailResponse
	FollowedListRequest    = user.FollowedListRequest
	FollowedListResponse   = user.FollowedListResponse
	FollowerListRequest    = user.FollowerListRequest
	FollowerListResponse   = user.FollowerListResponse
	LoginRequest           = user.LoginRequest
	LoginResponse          = user.LoginResponse
	RegisterRequest        = user.RegisterRequest
	RegisterResponse       = user.RegisterResponse
	Request                = user.Request
	Response               = user.Response
	SendVerifyCodeRequest  = user.SendVerifyCodeRequest
	SendVerifyCodeResponse = user.SendVerifyCodeResponse
	UserInfoRequest        = user.UserInfoRequest
	UserInfoResponse       = user.UserInfoResponse

	User interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
		GetUserFollowerList(ctx context.Context, in *FollowerListRequest, opts ...grpc.CallOption) (*FollowerListResponse, error)
		GetUserFollowedList(ctx context.Context, in *FollowedListRequest, opts ...grpc.CallOption) (*FollowerListResponse, error)
		SendVerifyCode(ctx context.Context, in *SendVerifyCodeRequest, opts ...grpc.CallOption) (*SendVerifyCodeResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultUser) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUser) GetUserFollowerList(ctx context.Context, in *FollowerListRequest, opts ...grpc.CallOption) (*FollowerListResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserFollowerList(ctx, in, opts...)
}

func (m *defaultUser) GetUserFollowedList(ctx context.Context, in *FollowedListRequest, opts ...grpc.CallOption) (*FollowerListResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserFollowedList(ctx, in, opts...)
}

func (m *defaultUser) SendVerifyCode(ctx context.Context, in *SendVerifyCodeRequest, opts ...grpc.CallOption) (*SendVerifyCodeResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.SendVerifyCode(ctx, in, opts...)
}
