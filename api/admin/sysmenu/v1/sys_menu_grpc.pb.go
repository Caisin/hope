// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/admin/sysmenu/v1/sys_menu.proto

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

// SysMenuClient is the client API for SysMenu service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysMenuClient interface {
	// 分页查询SysMenu
	GetSysMenuPage(ctx context.Context, in *SysMenuPageReq, opts ...grpc.CallOption) (*SysMenuPageReply, error)
	// 获取SysMenu
	GetSysMenu(ctx context.Context, in *SysMenuReq, opts ...grpc.CallOption) (*SysMenuReply, error)
	// 更新SysMenu
	UpdateSysMenu(ctx context.Context, in *SysMenuUpdateReq, opts ...grpc.CallOption) (*SysMenuUpdateReply, error)
	// 创建SysMenu
	CreateSysMenu(ctx context.Context, in *SysMenuCreateReq, opts ...grpc.CallOption) (*SysMenuCreateReply, error)
	// 删除SysMenu
	DeleteSysMenu(ctx context.Context, in *SysMenuDeleteReq, opts ...grpc.CallOption) (*SysMenuDeleteReply, error)
	// 批量删除SysMenu
	BatchDeleteSysMenu(ctx context.Context, in *SysMenuBatchDeleteReq, opts ...grpc.CallOption) (*SysMenuDeleteReply, error)
}

type sysMenuClient struct {
	cc grpc.ClientConnInterface
}

func NewSysMenuClient(cc grpc.ClientConnInterface) SysMenuClient {
	return &sysMenuClient{cc}
}

func (c *sysMenuClient) GetSysMenuPage(ctx context.Context, in *SysMenuPageReq, opts ...grpc.CallOption) (*SysMenuPageReply, error) {
	out := new(SysMenuPageReply)
	err := c.cc.Invoke(ctx, "/sysmenu.v1.SysMenu/GetSysMenuPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysMenuClient) GetSysMenu(ctx context.Context, in *SysMenuReq, opts ...grpc.CallOption) (*SysMenuReply, error) {
	out := new(SysMenuReply)
	err := c.cc.Invoke(ctx, "/sysmenu.v1.SysMenu/GetSysMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysMenuClient) UpdateSysMenu(ctx context.Context, in *SysMenuUpdateReq, opts ...grpc.CallOption) (*SysMenuUpdateReply, error) {
	out := new(SysMenuUpdateReply)
	err := c.cc.Invoke(ctx, "/sysmenu.v1.SysMenu/UpdateSysMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysMenuClient) CreateSysMenu(ctx context.Context, in *SysMenuCreateReq, opts ...grpc.CallOption) (*SysMenuCreateReply, error) {
	out := new(SysMenuCreateReply)
	err := c.cc.Invoke(ctx, "/sysmenu.v1.SysMenu/CreateSysMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysMenuClient) DeleteSysMenu(ctx context.Context, in *SysMenuDeleteReq, opts ...grpc.CallOption) (*SysMenuDeleteReply, error) {
	out := new(SysMenuDeleteReply)
	err := c.cc.Invoke(ctx, "/sysmenu.v1.SysMenu/DeleteSysMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysMenuClient) BatchDeleteSysMenu(ctx context.Context, in *SysMenuBatchDeleteReq, opts ...grpc.CallOption) (*SysMenuDeleteReply, error) {
	out := new(SysMenuDeleteReply)
	err := c.cc.Invoke(ctx, "/sysmenu.v1.SysMenu/BatchDeleteSysMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysMenuServer is the server API for SysMenu service.
// All implementations must embed UnimplementedSysMenuServer
// for forward compatibility
type SysMenuServer interface {
	// 分页查询SysMenu
	GetSysMenuPage(context.Context, *SysMenuPageReq) (*SysMenuPageReply, error)
	// 获取SysMenu
	GetSysMenu(context.Context, *SysMenuReq) (*SysMenuReply, error)
	// 更新SysMenu
	UpdateSysMenu(context.Context, *SysMenuUpdateReq) (*SysMenuUpdateReply, error)
	// 创建SysMenu
	CreateSysMenu(context.Context, *SysMenuCreateReq) (*SysMenuCreateReply, error)
	// 删除SysMenu
	DeleteSysMenu(context.Context, *SysMenuDeleteReq) (*SysMenuDeleteReply, error)
	// 批量删除SysMenu
	BatchDeleteSysMenu(context.Context, *SysMenuBatchDeleteReq) (*SysMenuDeleteReply, error)
	mustEmbedUnimplementedSysMenuServer()
}

// UnimplementedSysMenuServer must be embedded to have forward compatible implementations.
type UnimplementedSysMenuServer struct {
}

func (UnimplementedSysMenuServer) GetSysMenuPage(context.Context, *SysMenuPageReq) (*SysMenuPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSysMenuPage not implemented")
}
func (UnimplementedSysMenuServer) GetSysMenu(context.Context, *SysMenuReq) (*SysMenuReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSysMenu not implemented")
}
func (UnimplementedSysMenuServer) UpdateSysMenu(context.Context, *SysMenuUpdateReq) (*SysMenuUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSysMenu not implemented")
}
func (UnimplementedSysMenuServer) CreateSysMenu(context.Context, *SysMenuCreateReq) (*SysMenuCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSysMenu not implemented")
}
func (UnimplementedSysMenuServer) DeleteSysMenu(context.Context, *SysMenuDeleteReq) (*SysMenuDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSysMenu not implemented")
}
func (UnimplementedSysMenuServer) BatchDeleteSysMenu(context.Context, *SysMenuBatchDeleteReq) (*SysMenuDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteSysMenu not implemented")
}
func (UnimplementedSysMenuServer) mustEmbedUnimplementedSysMenuServer() {}

// UnsafeSysMenuServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysMenuServer will
// result in compilation errors.
type UnsafeSysMenuServer interface {
	mustEmbedUnimplementedSysMenuServer()
}

func RegisterSysMenuServer(s grpc.ServiceRegistrar, srv SysMenuServer) {
	s.RegisterService(&SysMenu_ServiceDesc, srv)
}

func _SysMenu_GetSysMenuPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysMenuPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysMenuServer).GetSysMenuPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmenu.v1.SysMenu/GetSysMenuPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysMenuServer).GetSysMenuPage(ctx, req.(*SysMenuPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysMenu_GetSysMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysMenuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysMenuServer).GetSysMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmenu.v1.SysMenu/GetSysMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysMenuServer).GetSysMenu(ctx, req.(*SysMenuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysMenu_UpdateSysMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysMenuUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysMenuServer).UpdateSysMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmenu.v1.SysMenu/UpdateSysMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysMenuServer).UpdateSysMenu(ctx, req.(*SysMenuUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysMenu_CreateSysMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysMenuCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysMenuServer).CreateSysMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmenu.v1.SysMenu/CreateSysMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysMenuServer).CreateSysMenu(ctx, req.(*SysMenuCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysMenu_DeleteSysMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysMenuDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysMenuServer).DeleteSysMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmenu.v1.SysMenu/DeleteSysMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysMenuServer).DeleteSysMenu(ctx, req.(*SysMenuDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysMenu_BatchDeleteSysMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysMenuBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysMenuServer).BatchDeleteSysMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmenu.v1.SysMenu/BatchDeleteSysMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysMenuServer).BatchDeleteSysMenu(ctx, req.(*SysMenuBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SysMenu_ServiceDesc is the grpc.ServiceDesc for SysMenu service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SysMenu_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sysmenu.v1.SysMenu",
	HandlerType: (*SysMenuServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSysMenuPage",
			Handler:    _SysMenu_GetSysMenuPage_Handler,
		},
		{
			MethodName: "GetSysMenu",
			Handler:    _SysMenu_GetSysMenu_Handler,
		},
		{
			MethodName: "UpdateSysMenu",
			Handler:    _SysMenu_UpdateSysMenu_Handler,
		},
		{
			MethodName: "CreateSysMenu",
			Handler:    _SysMenu_CreateSysMenu_Handler,
		},
		{
			MethodName: "DeleteSysMenu",
			Handler:    _SysMenu_DeleteSysMenu_Handler,
		},
		{
			MethodName: "BatchDeleteSysMenu",
			Handler:    _SysMenu_BatchDeleteSysMenu_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/admin/sysmenu/v1/sys_menu.proto",
}
