// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/novel/usermsg/v1/user_msg.proto

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

// UserMsgClient is the client API for UserMsg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserMsgClient interface {
	// 分页查询UserMsg
	GetUserMsgPage(ctx context.Context, in *UserMsgPageReq, opts ...grpc.CallOption) (*UserMsgPageReply, error)
	// 获取UserMsg
	GetUserMsg(ctx context.Context, in *UserMsgReq, opts ...grpc.CallOption) (*UserMsgReply, error)
	// 更新UserMsg
	UpdateUserMsg(ctx context.Context, in *UserMsgUpdateReq, opts ...grpc.CallOption) (*UserMsgUpdateReply, error)
	// 创建UserMsg
	CreateUserMsg(ctx context.Context, in *UserMsgCreateReq, opts ...grpc.CallOption) (*UserMsgCreateReply, error)
	// 删除UserMsg
	DeleteUserMsg(ctx context.Context, in *UserMsgDeleteReq, opts ...grpc.CallOption) (*UserMsgDeleteReply, error)
	// 批量删除UserMsg
	BatchDeleteUserMsg(ctx context.Context, in *UserMsgBatchDeleteReq, opts ...grpc.CallOption) (*UserMsgDeleteReply, error)
}

type userMsgClient struct {
	cc grpc.ClientConnInterface
}

func NewUserMsgClient(cc grpc.ClientConnInterface) UserMsgClient {
	return &userMsgClient{cc}
}

func (c *userMsgClient) GetUserMsgPage(ctx context.Context, in *UserMsgPageReq, opts ...grpc.CallOption) (*UserMsgPageReply, error) {
	out := new(UserMsgPageReply)
	err := c.cc.Invoke(ctx, "/usermsg.v1.UserMsg/GetUserMsgPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMsgClient) GetUserMsg(ctx context.Context, in *UserMsgReq, opts ...grpc.CallOption) (*UserMsgReply, error) {
	out := new(UserMsgReply)
	err := c.cc.Invoke(ctx, "/usermsg.v1.UserMsg/GetUserMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMsgClient) UpdateUserMsg(ctx context.Context, in *UserMsgUpdateReq, opts ...grpc.CallOption) (*UserMsgUpdateReply, error) {
	out := new(UserMsgUpdateReply)
	err := c.cc.Invoke(ctx, "/usermsg.v1.UserMsg/UpdateUserMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMsgClient) CreateUserMsg(ctx context.Context, in *UserMsgCreateReq, opts ...grpc.CallOption) (*UserMsgCreateReply, error) {
	out := new(UserMsgCreateReply)
	err := c.cc.Invoke(ctx, "/usermsg.v1.UserMsg/CreateUserMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMsgClient) DeleteUserMsg(ctx context.Context, in *UserMsgDeleteReq, opts ...grpc.CallOption) (*UserMsgDeleteReply, error) {
	out := new(UserMsgDeleteReply)
	err := c.cc.Invoke(ctx, "/usermsg.v1.UserMsg/DeleteUserMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMsgClient) BatchDeleteUserMsg(ctx context.Context, in *UserMsgBatchDeleteReq, opts ...grpc.CallOption) (*UserMsgDeleteReply, error) {
	out := new(UserMsgDeleteReply)
	err := c.cc.Invoke(ctx, "/usermsg.v1.UserMsg/BatchDeleteUserMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserMsgServer is the server API for UserMsg service.
// All implementations must embed UnimplementedUserMsgServer
// for forward compatibility
type UserMsgServer interface {
	// 分页查询UserMsg
	GetUserMsgPage(context.Context, *UserMsgPageReq) (*UserMsgPageReply, error)
	// 获取UserMsg
	GetUserMsg(context.Context, *UserMsgReq) (*UserMsgReply, error)
	// 更新UserMsg
	UpdateUserMsg(context.Context, *UserMsgUpdateReq) (*UserMsgUpdateReply, error)
	// 创建UserMsg
	CreateUserMsg(context.Context, *UserMsgCreateReq) (*UserMsgCreateReply, error)
	// 删除UserMsg
	DeleteUserMsg(context.Context, *UserMsgDeleteReq) (*UserMsgDeleteReply, error)
	// 批量删除UserMsg
	BatchDeleteUserMsg(context.Context, *UserMsgBatchDeleteReq) (*UserMsgDeleteReply, error)
	mustEmbedUnimplementedUserMsgServer()
}

// UnimplementedUserMsgServer must be embedded to have forward compatible implementations.
type UnimplementedUserMsgServer struct {
}

func (UnimplementedUserMsgServer) GetUserMsgPage(context.Context, *UserMsgPageReq) (*UserMsgPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMsgPage not implemented")
}
func (UnimplementedUserMsgServer) GetUserMsg(context.Context, *UserMsgReq) (*UserMsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMsg not implemented")
}
func (UnimplementedUserMsgServer) UpdateUserMsg(context.Context, *UserMsgUpdateReq) (*UserMsgUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserMsg not implemented")
}
func (UnimplementedUserMsgServer) CreateUserMsg(context.Context, *UserMsgCreateReq) (*UserMsgCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserMsg not implemented")
}
func (UnimplementedUserMsgServer) DeleteUserMsg(context.Context, *UserMsgDeleteReq) (*UserMsgDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserMsg not implemented")
}
func (UnimplementedUserMsgServer) BatchDeleteUserMsg(context.Context, *UserMsgBatchDeleteReq) (*UserMsgDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteUserMsg not implemented")
}
func (UnimplementedUserMsgServer) mustEmbedUnimplementedUserMsgServer() {}

// UnsafeUserMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserMsgServer will
// result in compilation errors.
type UnsafeUserMsgServer interface {
	mustEmbedUnimplementedUserMsgServer()
}

func RegisterUserMsgServer(s grpc.ServiceRegistrar, srv UserMsgServer) {
	s.RegisterService(&UserMsg_ServiceDesc, srv)
}

func _UserMsg_GetUserMsgPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMsgPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMsgServer).GetUserMsgPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermsg.v1.UserMsg/GetUserMsgPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMsgServer).GetUserMsgPage(ctx, req.(*UserMsgPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMsg_GetUserMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMsgServer).GetUserMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermsg.v1.UserMsg/GetUserMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMsgServer).GetUserMsg(ctx, req.(*UserMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMsg_UpdateUserMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMsgUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMsgServer).UpdateUserMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermsg.v1.UserMsg/UpdateUserMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMsgServer).UpdateUserMsg(ctx, req.(*UserMsgUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMsg_CreateUserMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMsgCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMsgServer).CreateUserMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermsg.v1.UserMsg/CreateUserMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMsgServer).CreateUserMsg(ctx, req.(*UserMsgCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMsg_DeleteUserMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMsgDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMsgServer).DeleteUserMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermsg.v1.UserMsg/DeleteUserMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMsgServer).DeleteUserMsg(ctx, req.(*UserMsgDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMsg_BatchDeleteUserMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMsgBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMsgServer).BatchDeleteUserMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermsg.v1.UserMsg/BatchDeleteUserMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMsgServer).BatchDeleteUserMsg(ctx, req.(*UserMsgBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserMsg_ServiceDesc is the grpc.ServiceDesc for UserMsg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserMsg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usermsg.v1.UserMsg",
	HandlerType: (*UserMsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserMsgPage",
			Handler:    _UserMsg_GetUserMsgPage_Handler,
		},
		{
			MethodName: "GetUserMsg",
			Handler:    _UserMsg_GetUserMsg_Handler,
		},
		{
			MethodName: "UpdateUserMsg",
			Handler:    _UserMsg_UpdateUserMsg_Handler,
		},
		{
			MethodName: "CreateUserMsg",
			Handler:    _UserMsg_CreateUserMsg_Handler,
		},
		{
			MethodName: "DeleteUserMsg",
			Handler:    _UserMsg_DeleteUserMsg_Handler,
		},
		{
			MethodName: "BatchDeleteUserMsg",
			Handler:    _UserMsg_BatchDeleteUserMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/novel/usermsg/v1/user_msg.proto",
}
