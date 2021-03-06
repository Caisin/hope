// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/admin/sysoperalog/v1/sys_opera_log.proto

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

// SysOperaLogClient is the client API for SysOperaLog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysOperaLogClient interface {
	// 分页查询SysOperaLog
	GetSysOperaLogPage(ctx context.Context, in *SysOperaLogPageReq, opts ...grpc.CallOption) (*SysOperaLogPageReply, error)
	// 获取SysOperaLog
	GetSysOperaLog(ctx context.Context, in *SysOperaLogReq, opts ...grpc.CallOption) (*SysOperaLogReply, error)
	// 更新SysOperaLog
	UpdateSysOperaLog(ctx context.Context, in *SysOperaLogUpdateReq, opts ...grpc.CallOption) (*SysOperaLogUpdateReply, error)
	// 创建SysOperaLog
	CreateSysOperaLog(ctx context.Context, in *SysOperaLogCreateReq, opts ...grpc.CallOption) (*SysOperaLogCreateReply, error)
	// 删除SysOperaLog
	DeleteSysOperaLog(ctx context.Context, in *SysOperaLogDeleteReq, opts ...grpc.CallOption) (*SysOperaLogDeleteReply, error)
	// 批量删除SysOperaLog
	BatchDeleteSysOperaLog(ctx context.Context, in *SysOperaLogBatchDeleteReq, opts ...grpc.CallOption) (*SysOperaLogDeleteReply, error)
}

type sysOperaLogClient struct {
	cc grpc.ClientConnInterface
}

func NewSysOperaLogClient(cc grpc.ClientConnInterface) SysOperaLogClient {
	return &sysOperaLogClient{cc}
}

func (c *sysOperaLogClient) GetSysOperaLogPage(ctx context.Context, in *SysOperaLogPageReq, opts ...grpc.CallOption) (*SysOperaLogPageReply, error) {
	out := new(SysOperaLogPageReply)
	err := c.cc.Invoke(ctx, "/sysoperalog.v1.SysOperaLog/GetSysOperaLogPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysOperaLogClient) GetSysOperaLog(ctx context.Context, in *SysOperaLogReq, opts ...grpc.CallOption) (*SysOperaLogReply, error) {
	out := new(SysOperaLogReply)
	err := c.cc.Invoke(ctx, "/sysoperalog.v1.SysOperaLog/GetSysOperaLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysOperaLogClient) UpdateSysOperaLog(ctx context.Context, in *SysOperaLogUpdateReq, opts ...grpc.CallOption) (*SysOperaLogUpdateReply, error) {
	out := new(SysOperaLogUpdateReply)
	err := c.cc.Invoke(ctx, "/sysoperalog.v1.SysOperaLog/UpdateSysOperaLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysOperaLogClient) CreateSysOperaLog(ctx context.Context, in *SysOperaLogCreateReq, opts ...grpc.CallOption) (*SysOperaLogCreateReply, error) {
	out := new(SysOperaLogCreateReply)
	err := c.cc.Invoke(ctx, "/sysoperalog.v1.SysOperaLog/CreateSysOperaLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysOperaLogClient) DeleteSysOperaLog(ctx context.Context, in *SysOperaLogDeleteReq, opts ...grpc.CallOption) (*SysOperaLogDeleteReply, error) {
	out := new(SysOperaLogDeleteReply)
	err := c.cc.Invoke(ctx, "/sysoperalog.v1.SysOperaLog/DeleteSysOperaLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysOperaLogClient) BatchDeleteSysOperaLog(ctx context.Context, in *SysOperaLogBatchDeleteReq, opts ...grpc.CallOption) (*SysOperaLogDeleteReply, error) {
	out := new(SysOperaLogDeleteReply)
	err := c.cc.Invoke(ctx, "/sysoperalog.v1.SysOperaLog/BatchDeleteSysOperaLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysOperaLogServer is the server API for SysOperaLog service.
// All implementations must embed UnimplementedSysOperaLogServer
// for forward compatibility
type SysOperaLogServer interface {
	// 分页查询SysOperaLog
	GetSysOperaLogPage(context.Context, *SysOperaLogPageReq) (*SysOperaLogPageReply, error)
	// 获取SysOperaLog
	GetSysOperaLog(context.Context, *SysOperaLogReq) (*SysOperaLogReply, error)
	// 更新SysOperaLog
	UpdateSysOperaLog(context.Context, *SysOperaLogUpdateReq) (*SysOperaLogUpdateReply, error)
	// 创建SysOperaLog
	CreateSysOperaLog(context.Context, *SysOperaLogCreateReq) (*SysOperaLogCreateReply, error)
	// 删除SysOperaLog
	DeleteSysOperaLog(context.Context, *SysOperaLogDeleteReq) (*SysOperaLogDeleteReply, error)
	// 批量删除SysOperaLog
	BatchDeleteSysOperaLog(context.Context, *SysOperaLogBatchDeleteReq) (*SysOperaLogDeleteReply, error)
	mustEmbedUnimplementedSysOperaLogServer()
}

// UnimplementedSysOperaLogServer must be embedded to have forward compatible implementations.
type UnimplementedSysOperaLogServer struct {
}

func (UnimplementedSysOperaLogServer) GetSysOperaLogPage(context.Context, *SysOperaLogPageReq) (*SysOperaLogPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSysOperaLogPage not implemented")
}
func (UnimplementedSysOperaLogServer) GetSysOperaLog(context.Context, *SysOperaLogReq) (*SysOperaLogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSysOperaLog not implemented")
}
func (UnimplementedSysOperaLogServer) UpdateSysOperaLog(context.Context, *SysOperaLogUpdateReq) (*SysOperaLogUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSysOperaLog not implemented")
}
func (UnimplementedSysOperaLogServer) CreateSysOperaLog(context.Context, *SysOperaLogCreateReq) (*SysOperaLogCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSysOperaLog not implemented")
}
func (UnimplementedSysOperaLogServer) DeleteSysOperaLog(context.Context, *SysOperaLogDeleteReq) (*SysOperaLogDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSysOperaLog not implemented")
}
func (UnimplementedSysOperaLogServer) BatchDeleteSysOperaLog(context.Context, *SysOperaLogBatchDeleteReq) (*SysOperaLogDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteSysOperaLog not implemented")
}
func (UnimplementedSysOperaLogServer) mustEmbedUnimplementedSysOperaLogServer() {}

// UnsafeSysOperaLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysOperaLogServer will
// result in compilation errors.
type UnsafeSysOperaLogServer interface {
	mustEmbedUnimplementedSysOperaLogServer()
}

func RegisterSysOperaLogServer(s grpc.ServiceRegistrar, srv SysOperaLogServer) {
	s.RegisterService(&SysOperaLog_ServiceDesc, srv)
}

func _SysOperaLog_GetSysOperaLogPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysOperaLogPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysOperaLogServer).GetSysOperaLogPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysoperalog.v1.SysOperaLog/GetSysOperaLogPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysOperaLogServer).GetSysOperaLogPage(ctx, req.(*SysOperaLogPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysOperaLog_GetSysOperaLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysOperaLogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysOperaLogServer).GetSysOperaLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysoperalog.v1.SysOperaLog/GetSysOperaLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysOperaLogServer).GetSysOperaLog(ctx, req.(*SysOperaLogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysOperaLog_UpdateSysOperaLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysOperaLogUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysOperaLogServer).UpdateSysOperaLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysoperalog.v1.SysOperaLog/UpdateSysOperaLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysOperaLogServer).UpdateSysOperaLog(ctx, req.(*SysOperaLogUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysOperaLog_CreateSysOperaLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysOperaLogCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysOperaLogServer).CreateSysOperaLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysoperalog.v1.SysOperaLog/CreateSysOperaLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysOperaLogServer).CreateSysOperaLog(ctx, req.(*SysOperaLogCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysOperaLog_DeleteSysOperaLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysOperaLogDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysOperaLogServer).DeleteSysOperaLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysoperalog.v1.SysOperaLog/DeleteSysOperaLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysOperaLogServer).DeleteSysOperaLog(ctx, req.(*SysOperaLogDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysOperaLog_BatchDeleteSysOperaLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysOperaLogBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysOperaLogServer).BatchDeleteSysOperaLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysoperalog.v1.SysOperaLog/BatchDeleteSysOperaLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysOperaLogServer).BatchDeleteSysOperaLog(ctx, req.(*SysOperaLogBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SysOperaLog_ServiceDesc is the grpc.ServiceDesc for SysOperaLog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SysOperaLog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sysoperalog.v1.SysOperaLog",
	HandlerType: (*SysOperaLogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSysOperaLogPage",
			Handler:    _SysOperaLog_GetSysOperaLogPage_Handler,
		},
		{
			MethodName: "GetSysOperaLog",
			Handler:    _SysOperaLog_GetSysOperaLog_Handler,
		},
		{
			MethodName: "UpdateSysOperaLog",
			Handler:    _SysOperaLog_UpdateSysOperaLog_Handler,
		},
		{
			MethodName: "CreateSysOperaLog",
			Handler:    _SysOperaLog_CreateSysOperaLog_Handler,
		},
		{
			MethodName: "DeleteSysOperaLog",
			Handler:    _SysOperaLog_DeleteSysOperaLog_Handler,
		},
		{
			MethodName: "BatchDeleteSysOperaLog",
			Handler:    _SysOperaLog_BatchDeleteSysOperaLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/admin/sysoperalog/v1/sys_opera_log.proto",
}
