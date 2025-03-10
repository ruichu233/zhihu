// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0--rc3
// source: video.proto

package video

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Video_Ping_FullMethodName         = "/video.video/Ping"
	Video_GetUploadURL_FullMethodName = "/video.video/GetUploadURL"
	Video_PublishVideo_FullMethodName = "/video.video/PublishVideo"
	Video_DetailVideo_FullMethodName  = "/video.video/DetailVideo"
	Video_DetailList_FullMethodName   = "/video.video/DetailList"
	Video_WorkList_FullMethodName     = "/video.video/WorkList"
	Video_LikeList_FullMethodName     = "/video.video/LikeList"
)

// VideoClient is the client API for Video service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// 获取视频上传的预签名 URL
	GetUploadURL(ctx context.Context, in *GetUploadURLRequest, opts ...grpc.CallOption) (*GetUploadURLResponse, error)
	// 发布视频
	PublishVideo(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	// 根据Id获取视频详情
	DetailVideo(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error)
	// 根据IdList获取视频详情列表
	DetailList(ctx context.Context, in *DetailListRequest, opts ...grpc.CallOption) (*DetailListResponse, error)
	// 根据userId获取作品列表
	WorkList(ctx context.Context, in *WorkListRequest, opts ...grpc.CallOption) (*WorkListResponse, error)
	// 根据userId获取喜欢列表
	LikeList(ctx context.Context, in *LikeListRequest, opts ...grpc.CallOption) (*LikeListResponse, error)
}

type videoClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoClient(cc grpc.ClientConnInterface) VideoClient {
	return &videoClient{cc}
}

func (c *videoClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Video_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) GetUploadURL(ctx context.Context, in *GetUploadURLRequest, opts ...grpc.CallOption) (*GetUploadURLResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUploadURLResponse)
	err := c.cc.Invoke(ctx, Video_GetUploadURL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) PublishVideo(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, Video_PublishVideo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) DetailVideo(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DetailResponse)
	err := c.cc.Invoke(ctx, Video_DetailVideo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) DetailList(ctx context.Context, in *DetailListRequest, opts ...grpc.CallOption) (*DetailListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DetailListResponse)
	err := c.cc.Invoke(ctx, Video_DetailList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) WorkList(ctx context.Context, in *WorkListRequest, opts ...grpc.CallOption) (*WorkListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WorkListResponse)
	err := c.cc.Invoke(ctx, Video_WorkList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) LikeList(ctx context.Context, in *LikeListRequest, opts ...grpc.CallOption) (*LikeListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LikeListResponse)
	err := c.cc.Invoke(ctx, Video_LikeList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServer is the server API for Video service.
// All implementations must embed UnimplementedVideoServer
// for forward compatibility
type VideoServer interface {
	Ping(context.Context, *Request) (*Response, error)
	// 获取视频上传的预签名 URL
	GetUploadURL(context.Context, *GetUploadURLRequest) (*GetUploadURLResponse, error)
	// 发布视频
	PublishVideo(context.Context, *PublishRequest) (*PublishResponse, error)
	// 根据Id获取视频详情
	DetailVideo(context.Context, *DetailRequest) (*DetailResponse, error)
	// 根据IdList获取视频详情列表
	DetailList(context.Context, *DetailListRequest) (*DetailListResponse, error)
	// 根据userId获取作品列表
	WorkList(context.Context, *WorkListRequest) (*WorkListResponse, error)
	// 根据userId获取喜欢列表
	LikeList(context.Context, *LikeListRequest) (*LikeListResponse, error)
	mustEmbedUnimplementedVideoServer()
}

// UnimplementedVideoServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServer struct {
}

func (UnimplementedVideoServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedVideoServer) GetUploadURL(context.Context, *GetUploadURLRequest) (*GetUploadURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUploadURL not implemented")
}
func (UnimplementedVideoServer) PublishVideo(context.Context, *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishVideo not implemented")
}
func (UnimplementedVideoServer) DetailVideo(context.Context, *DetailRequest) (*DetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailVideo not implemented")
}
func (UnimplementedVideoServer) DetailList(context.Context, *DetailListRequest) (*DetailListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailList not implemented")
}
func (UnimplementedVideoServer) WorkList(context.Context, *WorkListRequest) (*WorkListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WorkList not implemented")
}
func (UnimplementedVideoServer) LikeList(context.Context, *LikeListRequest) (*LikeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeList not implemented")
}
func (UnimplementedVideoServer) mustEmbedUnimplementedVideoServer() {}

// UnsafeVideoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServer will
// result in compilation errors.
type UnsafeVideoServer interface {
	mustEmbedUnimplementedVideoServer()
}

func RegisterVideoServer(s grpc.ServiceRegistrar, srv VideoServer) {
	s.RegisterService(&Video_ServiceDesc, srv)
}

func _Video_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_GetUploadURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUploadURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).GetUploadURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_GetUploadURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).GetUploadURL(ctx, req.(*GetUploadURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_PublishVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).PublishVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_PublishVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).PublishVideo(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_DetailVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).DetailVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_DetailVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).DetailVideo(ctx, req.(*DetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_DetailList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).DetailList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_DetailList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).DetailList(ctx, req.(*DetailListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_WorkList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).WorkList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_WorkList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).WorkList(ctx, req.(*WorkListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_LikeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).LikeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_LikeList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).LikeList(ctx, req.(*LikeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Video_ServiceDesc is the grpc.ServiceDesc for Video service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Video_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "video.video",
	HandlerType: (*VideoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Video_Ping_Handler,
		},
		{
			MethodName: "GetUploadURL",
			Handler:    _Video_GetUploadURL_Handler,
		},
		{
			MethodName: "PublishVideo",
			Handler:    _Video_PublishVideo_Handler,
		},
		{
			MethodName: "DetailVideo",
			Handler:    _Video_DetailVideo_Handler,
		},
		{
			MethodName: "DetailList",
			Handler:    _Video_DetailList_Handler,
		},
		{
			MethodName: "WorkList",
			Handler:    _Video_WorkList_Handler,
		},
		{
			MethodName: "LikeList",
			Handler:    _Video_LikeList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}
