// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/novel/userevent/v1/user_event.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserEventClient is the client API for UserEvent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserEventClient interface {
	// 分页查询UserEvent
	GetUserEventPage(ctx context.Context, in *UserEventPageReq, opts ...grpc.CallOption) (*UserEventPageReply, error)
	// 获取UserEvent
	GetUserEvent(ctx context.Context, in *UserEventReq, opts ...grpc.CallOption) (*UserEventReply, error)
	// 更新UserEvent
	UpdateUserEvent(ctx context.Context, in *UserEventUpdateReq, opts ...grpc.CallOption) (*UserEventUpdateReply, error)
	// 创建UserEvent
	CreateUserEvent(ctx context.Context, in *UserEventCreateReq, opts ...grpc.CallOption) (*UserEventCreateReply, error)
	// 删除UserEvent
	DeleteUserEvent(ctx context.Context, in *UserEventDeleteReq, opts ...grpc.CallOption) (*UserEventDeleteReply, error)
	// 批量删除UserEvent
	BatchDeleteUserEvent(ctx context.Context, in *UserEventBatchDeleteReq, opts ...grpc.CallOption) (*UserEventDeleteReply, error)
}

type userEventClient struct {
	cc grpc.ClientConnInterface
}

func NewUserEventClient(cc grpc.ClientConnInterface) UserEventClient {
	return &userEventClient{cc}
}

func (c *userEventClient) GetUserEventPage(ctx context.Context, in *UserEventPageReq, opts ...grpc.CallOption) (*UserEventPageReply, error) {
	out := new(UserEventPageReply)
	err := c.cc.Invoke(ctx, "/userevent.v1.UserEvent/GetUserEventPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userEventClient) GetUserEvent(ctx context.Context, in *UserEventReq, opts ...grpc.CallOption) (*UserEventReply, error) {
	out := new(UserEventReply)
	err := c.cc.Invoke(ctx, "/userevent.v1.UserEvent/GetUserEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userEventClient) UpdateUserEvent(ctx context.Context, in *UserEventUpdateReq, opts ...grpc.CallOption) (*UserEventUpdateReply, error) {
	out := new(UserEventUpdateReply)
	err := c.cc.Invoke(ctx, "/userevent.v1.UserEvent/UpdateUserEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userEventClient) CreateUserEvent(ctx context.Context, in *UserEventCreateReq, opts ...grpc.CallOption) (*UserEventCreateReply, error) {
	out := new(UserEventCreateReply)
	err := c.cc.Invoke(ctx, "/userevent.v1.UserEvent/CreateUserEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userEventClient) DeleteUserEvent(ctx context.Context, in *UserEventDeleteReq, opts ...grpc.CallOption) (*UserEventDeleteReply, error) {
	out := new(UserEventDeleteReply)
	err := c.cc.Invoke(ctx, "/userevent.v1.UserEvent/DeleteUserEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userEventClient) BatchDeleteUserEvent(ctx context.Context, in *UserEventBatchDeleteReq, opts ...grpc.CallOption) (*UserEventDeleteReply, error) {
	out := new(UserEventDeleteReply)
	err := c.cc.Invoke(ctx, "/userevent.v1.UserEvent/BatchDeleteUserEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserEventServer is the server API for UserEvent service.
// All implementations must embed UnimplementedUserEventServer
// for forward compatibility
type UserEventServer interface {
	// 分页查询UserEvent
	GetUserEventPage(context.Context, *UserEventPageReq) (*UserEventPageReply, error)
	// 获取UserEvent
	GetUserEvent(context.Context, *UserEventReq) (*UserEventReply, error)
	// 更新UserEvent
	UpdateUserEvent(context.Context, *UserEventUpdateReq) (*UserEventUpdateReply, error)
	// 创建UserEvent
	CreateUserEvent(context.Context, *UserEventCreateReq) (*UserEventCreateReply, error)
	// 删除UserEvent
	DeleteUserEvent(context.Context, *UserEventDeleteReq) (*UserEventDeleteReply, error)
	// 批量删除UserEvent
	BatchDeleteUserEvent(context.Context, *UserEventBatchDeleteReq) (*UserEventDeleteReply, error)
	mustEmbedUnimplementedUserEventServer()
}

// UnimplementedUserEventServer must be embedded to have forward compatible implementations.
type UnimplementedUserEventServer struct {
}

func (UnimplementedUserEventServer) GetUserEventPage(context.Context, *UserEventPageReq) (*UserEventPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserEventPage not implemented")
}
func (UnimplementedUserEventServer) GetUserEvent(context.Context, *UserEventReq) (*UserEventReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserEvent not implemented")
}
func (UnimplementedUserEventServer) UpdateUserEvent(context.Context, *UserEventUpdateReq) (*UserEventUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserEvent not implemented")
}
func (UnimplementedUserEventServer) CreateUserEvent(context.Context, *UserEventCreateReq) (*UserEventCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserEvent not implemented")
}
func (UnimplementedUserEventServer) DeleteUserEvent(context.Context, *UserEventDeleteReq) (*UserEventDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserEvent not implemented")
}
func (UnimplementedUserEventServer) BatchDeleteUserEvent(context.Context, *UserEventBatchDeleteReq) (*UserEventDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteUserEvent not implemented")
}
func (UnimplementedUserEventServer) mustEmbedUnimplementedUserEventServer() {}

// UnsafeUserEventServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserEventServer will
// result in compilation errors.
type UnsafeUserEventServer interface {
	mustEmbedUnimplementedUserEventServer()
}

func RegisterUserEventServer(s grpc.ServiceRegistrar, srv UserEventServer) {
	s.RegisterService(&UserEvent_ServiceDesc, srv)
}

func _UserEvent_GetUserEventPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEventPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserEventServer).GetUserEventPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userevent.v1.UserEvent/GetUserEventPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserEventServer).GetUserEventPage(ctx, req.(*UserEventPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserEvent_GetUserEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserEventServer).GetUserEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userevent.v1.UserEvent/GetUserEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserEventServer).GetUserEvent(ctx, req.(*UserEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserEvent_UpdateUserEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEventUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserEventServer).UpdateUserEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userevent.v1.UserEvent/UpdateUserEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserEventServer).UpdateUserEvent(ctx, req.(*UserEventUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserEvent_CreateUserEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEventCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserEventServer).CreateUserEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userevent.v1.UserEvent/CreateUserEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserEventServer).CreateUserEvent(ctx, req.(*UserEventCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserEvent_DeleteUserEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEventDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserEventServer).DeleteUserEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userevent.v1.UserEvent/DeleteUserEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserEventServer).DeleteUserEvent(ctx, req.(*UserEventDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserEvent_BatchDeleteUserEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEventBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserEventServer).BatchDeleteUserEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userevent.v1.UserEvent/BatchDeleteUserEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserEventServer).BatchDeleteUserEvent(ctx, req.(*UserEventBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserEvent_ServiceDesc is the grpc.ServiceDesc for UserEvent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserEvent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userevent.v1.UserEvent",
	HandlerType: (*UserEventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserEventPage",
			Handler:    _UserEvent_GetUserEventPage_Handler,
		},
		{
			MethodName: "GetUserEvent",
			Handler:    _UserEvent_GetUserEvent_Handler,
		},
		{
			MethodName: "UpdateUserEvent",
			Handler:    _UserEvent_UpdateUserEvent_Handler,
		},
		{
			MethodName: "CreateUserEvent",
			Handler:    _UserEvent_CreateUserEvent_Handler,
		},
		{
			MethodName: "DeleteUserEvent",
			Handler:    _UserEvent_DeleteUserEvent_Handler,
		},
		{
			MethodName: "BatchDeleteUserEvent",
			Handler:    _UserEvent_BatchDeleteUserEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/novel/userevent/v1/user_event.proto",
}
