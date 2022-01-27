// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/novel/adchangelog/v1/ad_change_log.proto

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

// AdChangeLogClient is the client API for AdChangeLog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdChangeLogClient interface {
	// 分页查询AdChangeLog
	GetAdChangeLogPage(ctx context.Context, in *AdChangeLogPageReq, opts ...grpc.CallOption) (*AdChangeLogPageReply, error)
	// 获取AdChangeLog
	GetAdChangeLog(ctx context.Context, in *AdChangeLogReq, opts ...grpc.CallOption) (*AdChangeLogReply, error)
	// 更新AdChangeLog
	UpdateAdChangeLog(ctx context.Context, in *AdChangeLogUpdateReq, opts ...grpc.CallOption) (*AdChangeLogUpdateReply, error)
	// 创建AdChangeLog
	CreateAdChangeLog(ctx context.Context, in *AdChangeLogCreateReq, opts ...grpc.CallOption) (*AdChangeLogCreateReply, error)
	// 删除AdChangeLog
	DeleteAdChangeLog(ctx context.Context, in *AdChangeLogDeleteReq, opts ...grpc.CallOption) (*AdChangeLogDeleteReply, error)
	// 批量删除AdChangeLog
	BatchDeleteAdChangeLog(ctx context.Context, in *AdChangeLogBatchDeleteReq, opts ...grpc.CallOption) (*AdChangeLogDeleteReply, error)
}

type adChangeLogClient struct {
	cc grpc.ClientConnInterface
}

func NewAdChangeLogClient(cc grpc.ClientConnInterface) AdChangeLogClient {
	return &adChangeLogClient{cc}
}

func (c *adChangeLogClient) GetAdChangeLogPage(ctx context.Context, in *AdChangeLogPageReq, opts ...grpc.CallOption) (*AdChangeLogPageReply, error) {
	out := new(AdChangeLogPageReply)
	err := c.cc.Invoke(ctx, "/adchangelog.v1.AdChangeLog/GetAdChangeLogPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adChangeLogClient) GetAdChangeLog(ctx context.Context, in *AdChangeLogReq, opts ...grpc.CallOption) (*AdChangeLogReply, error) {
	out := new(AdChangeLogReply)
	err := c.cc.Invoke(ctx, "/adchangelog.v1.AdChangeLog/GetAdChangeLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adChangeLogClient) UpdateAdChangeLog(ctx context.Context, in *AdChangeLogUpdateReq, opts ...grpc.CallOption) (*AdChangeLogUpdateReply, error) {
	out := new(AdChangeLogUpdateReply)
	err := c.cc.Invoke(ctx, "/adchangelog.v1.AdChangeLog/UpdateAdChangeLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adChangeLogClient) CreateAdChangeLog(ctx context.Context, in *AdChangeLogCreateReq, opts ...grpc.CallOption) (*AdChangeLogCreateReply, error) {
	out := new(AdChangeLogCreateReply)
	err := c.cc.Invoke(ctx, "/adchangelog.v1.AdChangeLog/CreateAdChangeLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adChangeLogClient) DeleteAdChangeLog(ctx context.Context, in *AdChangeLogDeleteReq, opts ...grpc.CallOption) (*AdChangeLogDeleteReply, error) {
	out := new(AdChangeLogDeleteReply)
	err := c.cc.Invoke(ctx, "/adchangelog.v1.AdChangeLog/DeleteAdChangeLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adChangeLogClient) BatchDeleteAdChangeLog(ctx context.Context, in *AdChangeLogBatchDeleteReq, opts ...grpc.CallOption) (*AdChangeLogDeleteReply, error) {
	out := new(AdChangeLogDeleteReply)
	err := c.cc.Invoke(ctx, "/adchangelog.v1.AdChangeLog/BatchDeleteAdChangeLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdChangeLogServer is the server API for AdChangeLog service.
// All implementations must embed UnimplementedAdChangeLogServer
// for forward compatibility
type AdChangeLogServer interface {
	// 分页查询AdChangeLog
	GetAdChangeLogPage(context.Context, *AdChangeLogPageReq) (*AdChangeLogPageReply, error)
	// 获取AdChangeLog
	GetAdChangeLog(context.Context, *AdChangeLogReq) (*AdChangeLogReply, error)
	// 更新AdChangeLog
	UpdateAdChangeLog(context.Context, *AdChangeLogUpdateReq) (*AdChangeLogUpdateReply, error)
	// 创建AdChangeLog
	CreateAdChangeLog(context.Context, *AdChangeLogCreateReq) (*AdChangeLogCreateReply, error)
	// 删除AdChangeLog
	DeleteAdChangeLog(context.Context, *AdChangeLogDeleteReq) (*AdChangeLogDeleteReply, error)
	// 批量删除AdChangeLog
	BatchDeleteAdChangeLog(context.Context, *AdChangeLogBatchDeleteReq) (*AdChangeLogDeleteReply, error)
	mustEmbedUnimplementedAdChangeLogServer()
}

// UnimplementedAdChangeLogServer must be embedded to have forward compatible implementations.
type UnimplementedAdChangeLogServer struct {
}

func (UnimplementedAdChangeLogServer) GetAdChangeLogPage(context.Context, *AdChangeLogPageReq) (*AdChangeLogPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdChangeLogPage not implemented")
}
func (UnimplementedAdChangeLogServer) GetAdChangeLog(context.Context, *AdChangeLogReq) (*AdChangeLogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdChangeLog not implemented")
}
func (UnimplementedAdChangeLogServer) UpdateAdChangeLog(context.Context, *AdChangeLogUpdateReq) (*AdChangeLogUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAdChangeLog not implemented")
}
func (UnimplementedAdChangeLogServer) CreateAdChangeLog(context.Context, *AdChangeLogCreateReq) (*AdChangeLogCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdChangeLog not implemented")
}
func (UnimplementedAdChangeLogServer) DeleteAdChangeLog(context.Context, *AdChangeLogDeleteReq) (*AdChangeLogDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAdChangeLog not implemented")
}
func (UnimplementedAdChangeLogServer) BatchDeleteAdChangeLog(context.Context, *AdChangeLogBatchDeleteReq) (*AdChangeLogDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteAdChangeLog not implemented")
}
func (UnimplementedAdChangeLogServer) mustEmbedUnimplementedAdChangeLogServer() {}

// UnsafeAdChangeLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdChangeLogServer will
// result in compilation errors.
type UnsafeAdChangeLogServer interface {
	mustEmbedUnimplementedAdChangeLogServer()
}

func RegisterAdChangeLogServer(s grpc.ServiceRegistrar, srv AdChangeLogServer) {
	s.RegisterService(&AdChangeLog_ServiceDesc, srv)
}

func _AdChangeLog_GetAdChangeLogPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdChangeLogPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdChangeLogServer).GetAdChangeLogPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adchangelog.v1.AdChangeLog/GetAdChangeLogPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdChangeLogServer).GetAdChangeLogPage(ctx, req.(*AdChangeLogPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdChangeLog_GetAdChangeLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdChangeLogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdChangeLogServer).GetAdChangeLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adchangelog.v1.AdChangeLog/GetAdChangeLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdChangeLogServer).GetAdChangeLog(ctx, req.(*AdChangeLogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdChangeLog_UpdateAdChangeLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdChangeLogUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdChangeLogServer).UpdateAdChangeLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adchangelog.v1.AdChangeLog/UpdateAdChangeLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdChangeLogServer).UpdateAdChangeLog(ctx, req.(*AdChangeLogUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdChangeLog_CreateAdChangeLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdChangeLogCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdChangeLogServer).CreateAdChangeLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adchangelog.v1.AdChangeLog/CreateAdChangeLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdChangeLogServer).CreateAdChangeLog(ctx, req.(*AdChangeLogCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdChangeLog_DeleteAdChangeLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdChangeLogDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdChangeLogServer).DeleteAdChangeLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adchangelog.v1.AdChangeLog/DeleteAdChangeLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdChangeLogServer).DeleteAdChangeLog(ctx, req.(*AdChangeLogDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdChangeLog_BatchDeleteAdChangeLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdChangeLogBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdChangeLogServer).BatchDeleteAdChangeLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adchangelog.v1.AdChangeLog/BatchDeleteAdChangeLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdChangeLogServer).BatchDeleteAdChangeLog(ctx, req.(*AdChangeLogBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AdChangeLog_ServiceDesc is the grpc.ServiceDesc for AdChangeLog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdChangeLog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "adchangelog.v1.AdChangeLog",
	HandlerType: (*AdChangeLogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdChangeLogPage",
			Handler:    _AdChangeLog_GetAdChangeLogPage_Handler,
		},
		{
			MethodName: "GetAdChangeLog",
			Handler:    _AdChangeLog_GetAdChangeLog_Handler,
		},
		{
			MethodName: "UpdateAdChangeLog",
			Handler:    _AdChangeLog_UpdateAdChangeLog_Handler,
		},
		{
			MethodName: "CreateAdChangeLog",
			Handler:    _AdChangeLog_CreateAdChangeLog_Handler,
		},
		{
			MethodName: "DeleteAdChangeLog",
			Handler:    _AdChangeLog_DeleteAdChangeLog_Handler,
		},
		{
			MethodName: "BatchDeleteAdChangeLog",
			Handler:    _AdChangeLog_BatchDeleteAdChangeLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/novel/adchangelog/v1/ad_change_log.proto",
}
