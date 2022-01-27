// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/param/userconsume/v1/user_consume.proto

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

// UserConsumeClient is the client API for UserConsume service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserConsumeClient interface {
	// 分页查询UserConsume
	GetUserConsumePage(ctx context.Context, in *UserConsumePageReq, opts ...grpc.CallOption) (*UserConsumePageReply, error)
	// 获取UserConsume
	GetUserConsume(ctx context.Context, in *UserConsumeReq, opts ...grpc.CallOption) (*UserConsumeReply, error)
	// 更新UserConsume
	UpdateUserConsume(ctx context.Context, in *UserConsumeUpdateReq, opts ...grpc.CallOption) (*UserConsumeUpdateReply, error)
	// 创建UserConsume
	CreateUserConsume(ctx context.Context, in *UserConsumeCreateReq, opts ...grpc.CallOption) (*UserConsumeCreateReply, error)
	// 删除UserConsume
	DeleteUserConsume(ctx context.Context, in *UserConsumeDeleteReq, opts ...grpc.CallOption) (*UserConsumeDeleteReply, error)
	// 批量删除UserConsume
	BatchDeleteUserConsume(ctx context.Context, in *UserConsumeBatchDeleteReq, opts ...grpc.CallOption) (*UserConsumeDeleteReply, error)
}

type userConsumeClient struct {
	cc grpc.ClientConnInterface
}

func NewUserConsumeClient(cc grpc.ClientConnInterface) UserConsumeClient {
	return &userConsumeClient{cc}
}

func (c *userConsumeClient) GetUserConsumePage(ctx context.Context, in *UserConsumePageReq, opts ...grpc.CallOption) (*UserConsumePageReply, error) {
	out := new(UserConsumePageReply)
	err := c.cc.Invoke(ctx, "/userconsume.v1.UserConsume/GetUserConsumePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConsumeClient) GetUserConsume(ctx context.Context, in *UserConsumeReq, opts ...grpc.CallOption) (*UserConsumeReply, error) {
	out := new(UserConsumeReply)
	err := c.cc.Invoke(ctx, "/userconsume.v1.UserConsume/GetUserConsume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConsumeClient) UpdateUserConsume(ctx context.Context, in *UserConsumeUpdateReq, opts ...grpc.CallOption) (*UserConsumeUpdateReply, error) {
	out := new(UserConsumeUpdateReply)
	err := c.cc.Invoke(ctx, "/userconsume.v1.UserConsume/UpdateUserConsume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConsumeClient) CreateUserConsume(ctx context.Context, in *UserConsumeCreateReq, opts ...grpc.CallOption) (*UserConsumeCreateReply, error) {
	out := new(UserConsumeCreateReply)
	err := c.cc.Invoke(ctx, "/userconsume.v1.UserConsume/CreateUserConsume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConsumeClient) DeleteUserConsume(ctx context.Context, in *UserConsumeDeleteReq, opts ...grpc.CallOption) (*UserConsumeDeleteReply, error) {
	out := new(UserConsumeDeleteReply)
	err := c.cc.Invoke(ctx, "/userconsume.v1.UserConsume/DeleteUserConsume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConsumeClient) BatchDeleteUserConsume(ctx context.Context, in *UserConsumeBatchDeleteReq, opts ...grpc.CallOption) (*UserConsumeDeleteReply, error) {
	out := new(UserConsumeDeleteReply)
	err := c.cc.Invoke(ctx, "/userconsume.v1.UserConsume/BatchDeleteUserConsume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserConsumeServer is the server API for UserConsume service.
// All implementations must embed UnimplementedUserConsumeServer
// for forward compatibility
type UserConsumeServer interface {
	// 分页查询UserConsume
	GetUserConsumePage(context.Context, *UserConsumePageReq) (*UserConsumePageReply, error)
	// 获取UserConsume
	GetUserConsume(context.Context, *UserConsumeReq) (*UserConsumeReply, error)
	// 更新UserConsume
	UpdateUserConsume(context.Context, *UserConsumeUpdateReq) (*UserConsumeUpdateReply, error)
	// 创建UserConsume
	CreateUserConsume(context.Context, *UserConsumeCreateReq) (*UserConsumeCreateReply, error)
	// 删除UserConsume
	DeleteUserConsume(context.Context, *UserConsumeDeleteReq) (*UserConsumeDeleteReply, error)
	// 批量删除UserConsume
	BatchDeleteUserConsume(context.Context, *UserConsumeBatchDeleteReq) (*UserConsumeDeleteReply, error)
	mustEmbedUnimplementedUserConsumeServer()
}

// UnimplementedUserConsumeServer must be embedded to have forward compatible implementations.
type UnimplementedUserConsumeServer struct {
}

func (UnimplementedUserConsumeServer) GetUserConsumePage(context.Context, *UserConsumePageReq) (*UserConsumePageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserConsumePage not implemented")
}
func (UnimplementedUserConsumeServer) GetUserConsume(context.Context, *UserConsumeReq) (*UserConsumeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserConsume not implemented")
}
func (UnimplementedUserConsumeServer) UpdateUserConsume(context.Context, *UserConsumeUpdateReq) (*UserConsumeUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserConsume not implemented")
}
func (UnimplementedUserConsumeServer) CreateUserConsume(context.Context, *UserConsumeCreateReq) (*UserConsumeCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserConsume not implemented")
}
func (UnimplementedUserConsumeServer) DeleteUserConsume(context.Context, *UserConsumeDeleteReq) (*UserConsumeDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserConsume not implemented")
}
func (UnimplementedUserConsumeServer) BatchDeleteUserConsume(context.Context, *UserConsumeBatchDeleteReq) (*UserConsumeDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteUserConsume not implemented")
}
func (UnimplementedUserConsumeServer) mustEmbedUnimplementedUserConsumeServer() {}

// UnsafeUserConsumeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserConsumeServer will
// result in compilation errors.
type UnsafeUserConsumeServer interface {
	mustEmbedUnimplementedUserConsumeServer()
}

func RegisterUserConsumeServer(s grpc.ServiceRegistrar, srv UserConsumeServer) {
	s.RegisterService(&UserConsume_ServiceDesc, srv)
}

func _UserConsume_GetUserConsumePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConsumePageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConsumeServer).GetUserConsumePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userconsume.v1.UserConsume/GetUserConsumePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConsumeServer).GetUserConsumePage(ctx, req.(*UserConsumePageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConsume_GetUserConsume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConsumeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConsumeServer).GetUserConsume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userconsume.v1.UserConsume/GetUserConsume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConsumeServer).GetUserConsume(ctx, req.(*UserConsumeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConsume_UpdateUserConsume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConsumeUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConsumeServer).UpdateUserConsume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userconsume.v1.UserConsume/UpdateUserConsume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConsumeServer).UpdateUserConsume(ctx, req.(*UserConsumeUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConsume_CreateUserConsume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConsumeCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConsumeServer).CreateUserConsume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userconsume.v1.UserConsume/CreateUserConsume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConsumeServer).CreateUserConsume(ctx, req.(*UserConsumeCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConsume_DeleteUserConsume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConsumeDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConsumeServer).DeleteUserConsume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userconsume.v1.UserConsume/DeleteUserConsume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConsumeServer).DeleteUserConsume(ctx, req.(*UserConsumeDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConsume_BatchDeleteUserConsume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConsumeBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConsumeServer).BatchDeleteUserConsume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userconsume.v1.UserConsume/BatchDeleteUserConsume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConsumeServer).BatchDeleteUserConsume(ctx, req.(*UserConsumeBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserConsume_ServiceDesc is the grpc.ServiceDesc for UserConsume service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserConsume_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userconsume.v1.UserConsume",
	HandlerType: (*UserConsumeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserConsumePage",
			Handler:    _UserConsume_GetUserConsumePage_Handler,
		},
		{
			MethodName: "GetUserConsume",
			Handler:    _UserConsume_GetUserConsume_Handler,
		},
		{
			MethodName: "UpdateUserConsume",
			Handler:    _UserConsume_UpdateUserConsume_Handler,
		},
		{
			MethodName: "CreateUserConsume",
			Handler:    _UserConsume_CreateUserConsume_Handler,
		},
		{
			MethodName: "DeleteUserConsume",
			Handler:    _UserConsume_DeleteUserConsume_Handler,
		},
		{
			MethodName: "BatchDeleteUserConsume",
			Handler:    _UserConsume_BatchDeleteUserConsume_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/param/userconsume/v1/user_consume.proto",
}
