// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"zhihu/app/user/internal/logic"
	"zhihu/app/user/internal/svc"
	"zhihu/app/user/pb/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Ping(ctx context.Context, in *user.Request) (*user.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) Register(ctx context.Context, in *user.RegisterRequest) (*user.RegisterResponse, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) GetUserInfo(ctx context.Context, in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UserServer) GetUserInfoList(ctx context.Context, in *user.UserInfoListRequest) (*user.UserInfoListResponse, error) {
	l := logic.NewGetUserInfoListLogic(ctx, s.svcCtx)
	return l.GetUserInfoList(in)
}

func (s *UserServer) SendVerifyCode(ctx context.Context, in *user.SendVerifyCodeRequest) (*user.SendVerifyCodeResponse, error) {
	l := logic.NewSendVerifyCodeLogic(ctx, s.svcCtx)
	return l.SendVerifyCode(in)
}

func (s *UserServer) UserInfoUpdate(ctx context.Context, in *user.UserInfoUpdateRequest) (*user.UserInfoUpdateResponse, error) {
	l := logic.NewUserInfoUpdateLogic(ctx, s.svcCtx)
	return l.UserInfoUpdate(in)
}

func (s *UserServer) GetAVatarUrl(ctx context.Context, in *user.GetAvatarRequest) (*user.GetAvatarResponse, error) {
	l := logic.NewGetAVatarUrlLogic(ctx, s.svcCtx)
	return l.GetAvatarUrl(in)
}
