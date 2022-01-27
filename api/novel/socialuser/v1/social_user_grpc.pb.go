// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/novel/socialuser/v1/social_user.proto

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

// SocialUserClient is the client API for SocialUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SocialUserClient interface {
	// 分页查询SocialUser
	GetSocialUserPage(ctx context.Context, in *SocialUserPageReq, opts ...grpc.CallOption) (*SocialUserPageReply, error)
	// 获取SocialUser
	GetSocialUser(ctx context.Context, in *SocialUserReq, opts ...grpc.CallOption) (*SocialUserReply, error)
	// 更新SocialUser
	UpdateSocialUser(ctx context.Context, in *SocialUserUpdateReq, opts ...grpc.CallOption) (*SocialUserUpdateReply, error)
	// 创建SocialUser
	CreateSocialUser(ctx context.Context, in *SocialUserCreateReq, opts ...grpc.CallOption) (*SocialUserCreateReply, error)
	// 删除SocialUser
	DeleteSocialUser(ctx context.Context, in *SocialUserDeleteReq, opts ...grpc.CallOption) (*SocialUserDeleteReply, error)
	// 批量删除SocialUser
	BatchDeleteSocialUser(ctx context.Context, in *SocialUserBatchDeleteReq, opts ...grpc.CallOption) (*SocialUserDeleteReply, error)
}

type socialUserClient struct {
	cc grpc.ClientConnInterface
}

func NewSocialUserClient(cc grpc.ClientConnInterface) SocialUserClient {
	return &socialUserClient{cc}
}

func (c *socialUserClient) GetSocialUserPage(ctx context.Context, in *SocialUserPageReq, opts ...grpc.CallOption) (*SocialUserPageReply, error) {
	out := new(SocialUserPageReply)
	err := c.cc.Invoke(ctx, "/socialuser.v1.SocialUser/GetSocialUserPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialUserClient) GetSocialUser(ctx context.Context, in *SocialUserReq, opts ...grpc.CallOption) (*SocialUserReply, error) {
	out := new(SocialUserReply)
	err := c.cc.Invoke(ctx, "/socialuser.v1.SocialUser/GetSocialUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialUserClient) UpdateSocialUser(ctx context.Context, in *SocialUserUpdateReq, opts ...grpc.CallOption) (*SocialUserUpdateReply, error) {
	out := new(SocialUserUpdateReply)
	err := c.cc.Invoke(ctx, "/socialuser.v1.SocialUser/UpdateSocialUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialUserClient) CreateSocialUser(ctx context.Context, in *SocialUserCreateReq, opts ...grpc.CallOption) (*SocialUserCreateReply, error) {
	out := new(SocialUserCreateReply)
	err := c.cc.Invoke(ctx, "/socialuser.v1.SocialUser/CreateSocialUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialUserClient) DeleteSocialUser(ctx context.Context, in *SocialUserDeleteReq, opts ...grpc.CallOption) (*SocialUserDeleteReply, error) {
	out := new(SocialUserDeleteReply)
	err := c.cc.Invoke(ctx, "/socialuser.v1.SocialUser/DeleteSocialUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialUserClient) BatchDeleteSocialUser(ctx context.Context, in *SocialUserBatchDeleteReq, opts ...grpc.CallOption) (*SocialUserDeleteReply, error) {
	out := new(SocialUserDeleteReply)
	err := c.cc.Invoke(ctx, "/socialuser.v1.SocialUser/BatchDeleteSocialUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SocialUserServer is the server API for SocialUser service.
// All implementations must embed UnimplementedSocialUserServer
// for forward compatibility
type SocialUserServer interface {
	// 分页查询SocialUser
	GetSocialUserPage(context.Context, *SocialUserPageReq) (*SocialUserPageReply, error)
	// 获取SocialUser
	GetSocialUser(context.Context, *SocialUserReq) (*SocialUserReply, error)
	// 更新SocialUser
	UpdateSocialUser(context.Context, *SocialUserUpdateReq) (*SocialUserUpdateReply, error)
	// 创建SocialUser
	CreateSocialUser(context.Context, *SocialUserCreateReq) (*SocialUserCreateReply, error)
	// 删除SocialUser
	DeleteSocialUser(context.Context, *SocialUserDeleteReq) (*SocialUserDeleteReply, error)
	// 批量删除SocialUser
	BatchDeleteSocialUser(context.Context, *SocialUserBatchDeleteReq) (*SocialUserDeleteReply, error)
	mustEmbedUnimplementedSocialUserServer()
}

// UnimplementedSocialUserServer must be embedded to have forward compatible implementations.
type UnimplementedSocialUserServer struct {
}

func (UnimplementedSocialUserServer) GetSocialUserPage(context.Context, *SocialUserPageReq) (*SocialUserPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSocialUserPage not implemented")
}
func (UnimplementedSocialUserServer) GetSocialUser(context.Context, *SocialUserReq) (*SocialUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSocialUser not implemented")
}
func (UnimplementedSocialUserServer) UpdateSocialUser(context.Context, *SocialUserUpdateReq) (*SocialUserUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSocialUser not implemented")
}
func (UnimplementedSocialUserServer) CreateSocialUser(context.Context, *SocialUserCreateReq) (*SocialUserCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSocialUser not implemented")
}
func (UnimplementedSocialUserServer) DeleteSocialUser(context.Context, *SocialUserDeleteReq) (*SocialUserDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSocialUser not implemented")
}
func (UnimplementedSocialUserServer) BatchDeleteSocialUser(context.Context, *SocialUserBatchDeleteReq) (*SocialUserDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteSocialUser not implemented")
}
func (UnimplementedSocialUserServer) mustEmbedUnimplementedSocialUserServer() {}

// UnsafeSocialUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SocialUserServer will
// result in compilation errors.
type UnsafeSocialUserServer interface {
	mustEmbedUnimplementedSocialUserServer()
}

func RegisterSocialUserServer(s grpc.ServiceRegistrar, srv SocialUserServer) {
	s.RegisterService(&SocialUser_ServiceDesc, srv)
}

func _SocialUser_GetSocialUserPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SocialUserPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialUserServer).GetSocialUserPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/socialuser.v1.SocialUser/GetSocialUserPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialUserServer).GetSocialUserPage(ctx, req.(*SocialUserPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocialUser_GetSocialUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SocialUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialUserServer).GetSocialUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/socialuser.v1.SocialUser/GetSocialUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialUserServer).GetSocialUser(ctx, req.(*SocialUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocialUser_UpdateSocialUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SocialUserUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialUserServer).UpdateSocialUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/socialuser.v1.SocialUser/UpdateSocialUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialUserServer).UpdateSocialUser(ctx, req.(*SocialUserUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocialUser_CreateSocialUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SocialUserCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialUserServer).CreateSocialUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/socialuser.v1.SocialUser/CreateSocialUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialUserServer).CreateSocialUser(ctx, req.(*SocialUserCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocialUser_DeleteSocialUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SocialUserDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialUserServer).DeleteSocialUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/socialuser.v1.SocialUser/DeleteSocialUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialUserServer).DeleteSocialUser(ctx, req.(*SocialUserDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocialUser_BatchDeleteSocialUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SocialUserBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialUserServer).BatchDeleteSocialUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/socialuser.v1.SocialUser/BatchDeleteSocialUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialUserServer).BatchDeleteSocialUser(ctx, req.(*SocialUserBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SocialUser_ServiceDesc is the grpc.ServiceDesc for SocialUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SocialUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "socialuser.v1.SocialUser",
	HandlerType: (*SocialUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSocialUserPage",
			Handler:    _SocialUser_GetSocialUserPage_Handler,
		},
		{
			MethodName: "GetSocialUser",
			Handler:    _SocialUser_GetSocialUser_Handler,
		},
		{
			MethodName: "UpdateSocialUser",
			Handler:    _SocialUser_UpdateSocialUser_Handler,
		},
		{
			MethodName: "CreateSocialUser",
			Handler:    _SocialUser_CreateSocialUser_Handler,
		},
		{
			MethodName: "DeleteSocialUser",
			Handler:    _SocialUser_DeleteSocialUser_Handler,
		},
		{
			MethodName: "BatchDeleteSocialUser",
			Handler:    _SocialUser_BatchDeleteSocialUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/novel/socialuser/v1/social_user.proto",
}
