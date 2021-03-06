// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/novel/vipuser/v1/vip_user.proto

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

// VipUserClient is the client API for VipUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VipUserClient interface {
	// 分页查询VipUser
	GetVipUserPage(ctx context.Context, in *VipUserPageReq, opts ...grpc.CallOption) (*VipUserPageReply, error)
	// 获取VipUser
	GetVipUser(ctx context.Context, in *VipUserReq, opts ...grpc.CallOption) (*VipUserReply, error)
	// 更新VipUser
	UpdateVipUser(ctx context.Context, in *VipUserUpdateReq, opts ...grpc.CallOption) (*VipUserUpdateReply, error)
	// 创建VipUser
	CreateVipUser(ctx context.Context, in *VipUserCreateReq, opts ...grpc.CallOption) (*VipUserCreateReply, error)
	// 删除VipUser
	DeleteVipUser(ctx context.Context, in *VipUserDeleteReq, opts ...grpc.CallOption) (*VipUserDeleteReply, error)
	// 批量删除VipUser
	BatchDeleteVipUser(ctx context.Context, in *VipUserBatchDeleteReq, opts ...grpc.CallOption) (*VipUserDeleteReply, error)
}

type vipUserClient struct {
	cc grpc.ClientConnInterface
}

func NewVipUserClient(cc grpc.ClientConnInterface) VipUserClient {
	return &vipUserClient{cc}
}

func (c *vipUserClient) GetVipUserPage(ctx context.Context, in *VipUserPageReq, opts ...grpc.CallOption) (*VipUserPageReply, error) {
	out := new(VipUserPageReply)
	err := c.cc.Invoke(ctx, "/vipuser.v1.VipUser/GetVipUserPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vipUserClient) GetVipUser(ctx context.Context, in *VipUserReq, opts ...grpc.CallOption) (*VipUserReply, error) {
	out := new(VipUserReply)
	err := c.cc.Invoke(ctx, "/vipuser.v1.VipUser/GetVipUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vipUserClient) UpdateVipUser(ctx context.Context, in *VipUserUpdateReq, opts ...grpc.CallOption) (*VipUserUpdateReply, error) {
	out := new(VipUserUpdateReply)
	err := c.cc.Invoke(ctx, "/vipuser.v1.VipUser/UpdateVipUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vipUserClient) CreateVipUser(ctx context.Context, in *VipUserCreateReq, opts ...grpc.CallOption) (*VipUserCreateReply, error) {
	out := new(VipUserCreateReply)
	err := c.cc.Invoke(ctx, "/vipuser.v1.VipUser/CreateVipUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vipUserClient) DeleteVipUser(ctx context.Context, in *VipUserDeleteReq, opts ...grpc.CallOption) (*VipUserDeleteReply, error) {
	out := new(VipUserDeleteReply)
	err := c.cc.Invoke(ctx, "/vipuser.v1.VipUser/DeleteVipUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vipUserClient) BatchDeleteVipUser(ctx context.Context, in *VipUserBatchDeleteReq, opts ...grpc.CallOption) (*VipUserDeleteReply, error) {
	out := new(VipUserDeleteReply)
	err := c.cc.Invoke(ctx, "/vipuser.v1.VipUser/BatchDeleteVipUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VipUserServer is the server API for VipUser service.
// All implementations must embed UnimplementedVipUserServer
// for forward compatibility
type VipUserServer interface {
	// 分页查询VipUser
	GetVipUserPage(context.Context, *VipUserPageReq) (*VipUserPageReply, error)
	// 获取VipUser
	GetVipUser(context.Context, *VipUserReq) (*VipUserReply, error)
	// 更新VipUser
	UpdateVipUser(context.Context, *VipUserUpdateReq) (*VipUserUpdateReply, error)
	// 创建VipUser
	CreateVipUser(context.Context, *VipUserCreateReq) (*VipUserCreateReply, error)
	// 删除VipUser
	DeleteVipUser(context.Context, *VipUserDeleteReq) (*VipUserDeleteReply, error)
	// 批量删除VipUser
	BatchDeleteVipUser(context.Context, *VipUserBatchDeleteReq) (*VipUserDeleteReply, error)
	mustEmbedUnimplementedVipUserServer()
}

// UnimplementedVipUserServer must be embedded to have forward compatible implementations.
type UnimplementedVipUserServer struct {
}

func (UnimplementedVipUserServer) GetVipUserPage(context.Context, *VipUserPageReq) (*VipUserPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVipUserPage not implemented")
}
func (UnimplementedVipUserServer) GetVipUser(context.Context, *VipUserReq) (*VipUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVipUser not implemented")
}
func (UnimplementedVipUserServer) UpdateVipUser(context.Context, *VipUserUpdateReq) (*VipUserUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVipUser not implemented")
}
func (UnimplementedVipUserServer) CreateVipUser(context.Context, *VipUserCreateReq) (*VipUserCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVipUser not implemented")
}
func (UnimplementedVipUserServer) DeleteVipUser(context.Context, *VipUserDeleteReq) (*VipUserDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVipUser not implemented")
}
func (UnimplementedVipUserServer) BatchDeleteVipUser(context.Context, *VipUserBatchDeleteReq) (*VipUserDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteVipUser not implemented")
}
func (UnimplementedVipUserServer) mustEmbedUnimplementedVipUserServer() {}

// UnsafeVipUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VipUserServer will
// result in compilation errors.
type UnsafeVipUserServer interface {
	mustEmbedUnimplementedVipUserServer()
}

func RegisterVipUserServer(s grpc.ServiceRegistrar, srv VipUserServer) {
	s.RegisterService(&VipUser_ServiceDesc, srv)
}

func _VipUser_GetVipUserPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VipUserPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VipUserServer).GetVipUserPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vipuser.v1.VipUser/GetVipUserPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VipUserServer).GetVipUserPage(ctx, req.(*VipUserPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VipUser_GetVipUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VipUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VipUserServer).GetVipUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vipuser.v1.VipUser/GetVipUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VipUserServer).GetVipUser(ctx, req.(*VipUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VipUser_UpdateVipUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VipUserUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VipUserServer).UpdateVipUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vipuser.v1.VipUser/UpdateVipUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VipUserServer).UpdateVipUser(ctx, req.(*VipUserUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VipUser_CreateVipUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VipUserCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VipUserServer).CreateVipUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vipuser.v1.VipUser/CreateVipUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VipUserServer).CreateVipUser(ctx, req.(*VipUserCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VipUser_DeleteVipUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VipUserDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VipUserServer).DeleteVipUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vipuser.v1.VipUser/DeleteVipUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VipUserServer).DeleteVipUser(ctx, req.(*VipUserDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VipUser_BatchDeleteVipUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VipUserBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VipUserServer).BatchDeleteVipUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vipuser.v1.VipUser/BatchDeleteVipUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VipUserServer).BatchDeleteVipUser(ctx, req.(*VipUserBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VipUser_ServiceDesc is the grpc.ServiceDesc for VipUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VipUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vipuser.v1.VipUser",
	HandlerType: (*VipUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVipUserPage",
			Handler:    _VipUser_GetVipUserPage_Handler,
		},
		{
			MethodName: "GetVipUser",
			Handler:    _VipUser_GetVipUser_Handler,
		},
		{
			MethodName: "UpdateVipUser",
			Handler:    _VipUser_UpdateVipUser_Handler,
		},
		{
			MethodName: "CreateVipUser",
			Handler:    _VipUser_CreateVipUser_Handler,
		},
		{
			MethodName: "DeleteVipUser",
			Handler:    _VipUser_DeleteVipUser_Handler,
		},
		{
			MethodName: "BatchDeleteVipUser",
			Handler:    _VipUser_BatchDeleteVipUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/novel/vipuser/v1/vip_user.proto",
}
