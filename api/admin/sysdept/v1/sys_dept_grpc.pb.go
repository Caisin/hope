// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/admin/sysdept/v1/sys_dept.proto

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

// SysDeptClient is the client API for SysDept service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysDeptClient interface {
	// 分页查询SysDept
	GetSysDeptPage(ctx context.Context, in *SysDeptPageReq, opts ...grpc.CallOption) (*SysDeptPageReply, error)
	// 获取SysDept
	GetSysDept(ctx context.Context, in *SysDeptReq, opts ...grpc.CallOption) (*SysDeptReply, error)
	// 更新SysDept
	UpdateSysDept(ctx context.Context, in *SysDeptUpdateReq, opts ...grpc.CallOption) (*SysDeptUpdateReply, error)
	// 创建SysDept
	CreateSysDept(ctx context.Context, in *SysDeptCreateReq, opts ...grpc.CallOption) (*SysDeptCreateReply, error)
	// 删除SysDept
	DeleteSysDept(ctx context.Context, in *SysDeptDeleteReq, opts ...grpc.CallOption) (*SysDeptDeleteReply, error)
	// 批量删除SysDept
	BatchDeleteSysDept(ctx context.Context, in *SysDeptBatchDeleteReq, opts ...grpc.CallOption) (*SysDeptDeleteReply, error)
}

type sysDeptClient struct {
	cc grpc.ClientConnInterface
}

func NewSysDeptClient(cc grpc.ClientConnInterface) SysDeptClient {
	return &sysDeptClient{cc}
}

func (c *sysDeptClient) GetSysDeptPage(ctx context.Context, in *SysDeptPageReq, opts ...grpc.CallOption) (*SysDeptPageReply, error) {
	out := new(SysDeptPageReply)
	err := c.cc.Invoke(ctx, "/sysdept.v1.SysDept/GetSysDeptPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDeptClient) GetSysDept(ctx context.Context, in *SysDeptReq, opts ...grpc.CallOption) (*SysDeptReply, error) {
	out := new(SysDeptReply)
	err := c.cc.Invoke(ctx, "/sysdept.v1.SysDept/GetSysDept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDeptClient) UpdateSysDept(ctx context.Context, in *SysDeptUpdateReq, opts ...grpc.CallOption) (*SysDeptUpdateReply, error) {
	out := new(SysDeptUpdateReply)
	err := c.cc.Invoke(ctx, "/sysdept.v1.SysDept/UpdateSysDept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDeptClient) CreateSysDept(ctx context.Context, in *SysDeptCreateReq, opts ...grpc.CallOption) (*SysDeptCreateReply, error) {
	out := new(SysDeptCreateReply)
	err := c.cc.Invoke(ctx, "/sysdept.v1.SysDept/CreateSysDept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDeptClient) DeleteSysDept(ctx context.Context, in *SysDeptDeleteReq, opts ...grpc.CallOption) (*SysDeptDeleteReply, error) {
	out := new(SysDeptDeleteReply)
	err := c.cc.Invoke(ctx, "/sysdept.v1.SysDept/DeleteSysDept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDeptClient) BatchDeleteSysDept(ctx context.Context, in *SysDeptBatchDeleteReq, opts ...grpc.CallOption) (*SysDeptDeleteReply, error) {
	out := new(SysDeptDeleteReply)
	err := c.cc.Invoke(ctx, "/sysdept.v1.SysDept/BatchDeleteSysDept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysDeptServer is the server API for SysDept service.
// All implementations must embed UnimplementedSysDeptServer
// for forward compatibility
type SysDeptServer interface {
	// 分页查询SysDept
	GetSysDeptPage(context.Context, *SysDeptPageReq) (*SysDeptPageReply, error)
	// 获取SysDept
	GetSysDept(context.Context, *SysDeptReq) (*SysDeptReply, error)
	// 更新SysDept
	UpdateSysDept(context.Context, *SysDeptUpdateReq) (*SysDeptUpdateReply, error)
	// 创建SysDept
	CreateSysDept(context.Context, *SysDeptCreateReq) (*SysDeptCreateReply, error)
	// 删除SysDept
	DeleteSysDept(context.Context, *SysDeptDeleteReq) (*SysDeptDeleteReply, error)
	// 批量删除SysDept
	BatchDeleteSysDept(context.Context, *SysDeptBatchDeleteReq) (*SysDeptDeleteReply, error)
	mustEmbedUnimplementedSysDeptServer()
}

// UnimplementedSysDeptServer must be embedded to have forward compatible implementations.
type UnimplementedSysDeptServer struct {
}

func (UnimplementedSysDeptServer) GetSysDeptPage(context.Context, *SysDeptPageReq) (*SysDeptPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSysDeptPage not implemented")
}
func (UnimplementedSysDeptServer) GetSysDept(context.Context, *SysDeptReq) (*SysDeptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSysDept not implemented")
}
func (UnimplementedSysDeptServer) UpdateSysDept(context.Context, *SysDeptUpdateReq) (*SysDeptUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSysDept not implemented")
}
func (UnimplementedSysDeptServer) CreateSysDept(context.Context, *SysDeptCreateReq) (*SysDeptCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSysDept not implemented")
}
func (UnimplementedSysDeptServer) DeleteSysDept(context.Context, *SysDeptDeleteReq) (*SysDeptDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSysDept not implemented")
}
func (UnimplementedSysDeptServer) BatchDeleteSysDept(context.Context, *SysDeptBatchDeleteReq) (*SysDeptDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteSysDept not implemented")
}
func (UnimplementedSysDeptServer) mustEmbedUnimplementedSysDeptServer() {}

// UnsafeSysDeptServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysDeptServer will
// result in compilation errors.
type UnsafeSysDeptServer interface {
	mustEmbedUnimplementedSysDeptServer()
}

func RegisterSysDeptServer(s grpc.ServiceRegistrar, srv SysDeptServer) {
	s.RegisterService(&SysDept_ServiceDesc, srv)
}

func _SysDept_GetSysDeptPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDeptPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDeptServer).GetSysDeptPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysdept.v1.SysDept/GetSysDeptPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDeptServer).GetSysDeptPage(ctx, req.(*SysDeptPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDept_GetSysDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDeptReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDeptServer).GetSysDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysdept.v1.SysDept/GetSysDept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDeptServer).GetSysDept(ctx, req.(*SysDeptReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDept_UpdateSysDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDeptUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDeptServer).UpdateSysDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysdept.v1.SysDept/UpdateSysDept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDeptServer).UpdateSysDept(ctx, req.(*SysDeptUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDept_CreateSysDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDeptCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDeptServer).CreateSysDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysdept.v1.SysDept/CreateSysDept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDeptServer).CreateSysDept(ctx, req.(*SysDeptCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDept_DeleteSysDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDeptDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDeptServer).DeleteSysDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysdept.v1.SysDept/DeleteSysDept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDeptServer).DeleteSysDept(ctx, req.(*SysDeptDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDept_BatchDeleteSysDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDeptBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDeptServer).BatchDeleteSysDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysdept.v1.SysDept/BatchDeleteSysDept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDeptServer).BatchDeleteSysDept(ctx, req.(*SysDeptBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SysDept_ServiceDesc is the grpc.ServiceDesc for SysDept service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SysDept_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sysdept.v1.SysDept",
	HandlerType: (*SysDeptServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSysDeptPage",
			Handler:    _SysDept_GetSysDeptPage_Handler,
		},
		{
			MethodName: "GetSysDept",
			Handler:    _SysDept_GetSysDept_Handler,
		},
		{
			MethodName: "UpdateSysDept",
			Handler:    _SysDept_UpdateSysDept_Handler,
		},
		{
			MethodName: "CreateSysDept",
			Handler:    _SysDept_CreateSysDept_Handler,
		},
		{
			MethodName: "DeleteSysDept",
			Handler:    _SysDept_DeleteSysDept_Handler,
		},
		{
			MethodName: "BatchDeleteSysDept",
			Handler:    _SysDept_BatchDeleteSysDept_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/admin/sysdept/v1/sys_dept.proto",
}
