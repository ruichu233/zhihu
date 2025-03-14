// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: chat.proto

package chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Chat_Ping_FullMethodName        = "/chat.Chat/Ping"
	Chat_GetMassages_FullMethodName = "/chat.Chat/GetMassages"
	Chat_SendMassage_FullMethodName = "/chat.Chat/SendMassage"
)

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 聊天服务
type ChatClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetMassages(ctx context.Context, in *GetMassagesRequest, opts ...grpc.CallOption) (*GetMassagesResponse, error)
	SendMassage(ctx context.Context, in *SendMassageRequest, opts ...grpc.CallOption) (*SendMassageResponse, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Chat_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetMassages(ctx context.Context, in *GetMassagesRequest, opts ...grpc.CallOption) (*GetMassagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMassagesResponse)
	err := c.cc.Invoke(ctx, Chat_GetMassages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) SendMassage(ctx context.Context, in *SendMassageRequest, opts ...grpc.CallOption) (*SendMassageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendMassageResponse)
	err := c.cc.Invoke(ctx, Chat_SendMassage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility.
//
// 聊天服务
type ChatServer interface {
	Ping(context.Context, *Request) (*Response, error)
	GetMassages(context.Context, *GetMassagesRequest) (*GetMassagesResponse, error)
	SendMassage(context.Context, *SendMassageRequest) (*SendMassageResponse, error)
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChatServer struct{}

func (UnimplementedChatServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedChatServer) GetMassages(context.Context, *GetMassagesRequest) (*GetMassagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMassages not implemented")
}
func (UnimplementedChatServer) SendMassage(context.Context, *SendMassageRequest) (*SendMassageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMassage not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}
func (UnimplementedChatServer) testEmbeddedByValue()              {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	// If the following call pancis, it indicates UnimplementedChatServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetMassages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMassagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetMassages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_GetMassages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetMassages(ctx, req.(*GetMassagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_SendMassage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMassageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SendMassage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_SendMassage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SendMassage(ctx, req.(*SendMassageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Chat_Ping_Handler,
		},
		{
			MethodName: "GetMassages",
			Handler:    _Chat_GetMassages_Handler,
		},
		{
			MethodName: "SendMassage",
			Handler:    _Chat_SendMassage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
