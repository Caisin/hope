// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/novel/agreementlog/v1/agreement_log.proto

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

// AgreementLogClient is the client API for AgreementLog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgreementLogClient interface {
	// 分页查询AgreementLog
	GetAgreementLogPage(ctx context.Context, in *AgreementLogPageReq, opts ...grpc.CallOption) (*AgreementLogPageReply, error)
	// 获取AgreementLog
	GetAgreementLog(ctx context.Context, in *AgreementLogReq, opts ...grpc.CallOption) (*AgreementLogReply, error)
	// 更新AgreementLog
	UpdateAgreementLog(ctx context.Context, in *AgreementLogUpdateReq, opts ...grpc.CallOption) (*AgreementLogUpdateReply, error)
	// 创建AgreementLog
	CreateAgreementLog(ctx context.Context, in *AgreementLogCreateReq, opts ...grpc.CallOption) (*AgreementLogCreateReply, error)
	// 删除AgreementLog
	DeleteAgreementLog(ctx context.Context, in *AgreementLogDeleteReq, opts ...grpc.CallOption) (*AgreementLogDeleteReply, error)
	// 批量删除AgreementLog
	BatchDeleteAgreementLog(ctx context.Context, in *AgreementLogBatchDeleteReq, opts ...grpc.CallOption) (*AgreementLogDeleteReply, error)
}

type agreementLogClient struct {
	cc grpc.ClientConnInterface
}

func NewAgreementLogClient(cc grpc.ClientConnInterface) AgreementLogClient {
	return &agreementLogClient{cc}
}

func (c *agreementLogClient) GetAgreementLogPage(ctx context.Context, in *AgreementLogPageReq, opts ...grpc.CallOption) (*AgreementLogPageReply, error) {
	out := new(AgreementLogPageReply)
	err := c.cc.Invoke(ctx, "/agreementlog.v1.AgreementLog/GetAgreementLogPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agreementLogClient) GetAgreementLog(ctx context.Context, in *AgreementLogReq, opts ...grpc.CallOption) (*AgreementLogReply, error) {
	out := new(AgreementLogReply)
	err := c.cc.Invoke(ctx, "/agreementlog.v1.AgreementLog/GetAgreementLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agreementLogClient) UpdateAgreementLog(ctx context.Context, in *AgreementLogUpdateReq, opts ...grpc.CallOption) (*AgreementLogUpdateReply, error) {
	out := new(AgreementLogUpdateReply)
	err := c.cc.Invoke(ctx, "/agreementlog.v1.AgreementLog/UpdateAgreementLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agreementLogClient) CreateAgreementLog(ctx context.Context, in *AgreementLogCreateReq, opts ...grpc.CallOption) (*AgreementLogCreateReply, error) {
	out := new(AgreementLogCreateReply)
	err := c.cc.Invoke(ctx, "/agreementlog.v1.AgreementLog/CreateAgreementLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agreementLogClient) DeleteAgreementLog(ctx context.Context, in *AgreementLogDeleteReq, opts ...grpc.CallOption) (*AgreementLogDeleteReply, error) {
	out := new(AgreementLogDeleteReply)
	err := c.cc.Invoke(ctx, "/agreementlog.v1.AgreementLog/DeleteAgreementLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agreementLogClient) BatchDeleteAgreementLog(ctx context.Context, in *AgreementLogBatchDeleteReq, opts ...grpc.CallOption) (*AgreementLogDeleteReply, error) {
	out := new(AgreementLogDeleteReply)
	err := c.cc.Invoke(ctx, "/agreementlog.v1.AgreementLog/BatchDeleteAgreementLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgreementLogServer is the server API for AgreementLog service.
// All implementations must embed UnimplementedAgreementLogServer
// for forward compatibility
type AgreementLogServer interface {
	// 分页查询AgreementLog
	GetAgreementLogPage(context.Context, *AgreementLogPageReq) (*AgreementLogPageReply, error)
	// 获取AgreementLog
	GetAgreementLog(context.Context, *AgreementLogReq) (*AgreementLogReply, error)
	// 更新AgreementLog
	UpdateAgreementLog(context.Context, *AgreementLogUpdateReq) (*AgreementLogUpdateReply, error)
	// 创建AgreementLog
	CreateAgreementLog(context.Context, *AgreementLogCreateReq) (*AgreementLogCreateReply, error)
	// 删除AgreementLog
	DeleteAgreementLog(context.Context, *AgreementLogDeleteReq) (*AgreementLogDeleteReply, error)
	// 批量删除AgreementLog
	BatchDeleteAgreementLog(context.Context, *AgreementLogBatchDeleteReq) (*AgreementLogDeleteReply, error)
	mustEmbedUnimplementedAgreementLogServer()
}

// UnimplementedAgreementLogServer must be embedded to have forward compatible implementations.
type UnimplementedAgreementLogServer struct {
}

func (UnimplementedAgreementLogServer) GetAgreementLogPage(context.Context, *AgreementLogPageReq) (*AgreementLogPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgreementLogPage not implemented")
}
func (UnimplementedAgreementLogServer) GetAgreementLog(context.Context, *AgreementLogReq) (*AgreementLogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgreementLog not implemented")
}
func (UnimplementedAgreementLogServer) UpdateAgreementLog(context.Context, *AgreementLogUpdateReq) (*AgreementLogUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAgreementLog not implemented")
}
func (UnimplementedAgreementLogServer) CreateAgreementLog(context.Context, *AgreementLogCreateReq) (*AgreementLogCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAgreementLog not implemented")
}
func (UnimplementedAgreementLogServer) DeleteAgreementLog(context.Context, *AgreementLogDeleteReq) (*AgreementLogDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAgreementLog not implemented")
}
func (UnimplementedAgreementLogServer) BatchDeleteAgreementLog(context.Context, *AgreementLogBatchDeleteReq) (*AgreementLogDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteAgreementLog not implemented")
}
func (UnimplementedAgreementLogServer) mustEmbedUnimplementedAgreementLogServer() {}

// UnsafeAgreementLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgreementLogServer will
// result in compilation errors.
type UnsafeAgreementLogServer interface {
	mustEmbedUnimplementedAgreementLogServer()
}

func RegisterAgreementLogServer(s grpc.ServiceRegistrar, srv AgreementLogServer) {
	s.RegisterService(&AgreementLog_ServiceDesc, srv)
}

func _AgreementLog_GetAgreementLogPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgreementLogPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgreementLogServer).GetAgreementLogPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agreementlog.v1.AgreementLog/GetAgreementLogPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgreementLogServer).GetAgreementLogPage(ctx, req.(*AgreementLogPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgreementLog_GetAgreementLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgreementLogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgreementLogServer).GetAgreementLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agreementlog.v1.AgreementLog/GetAgreementLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgreementLogServer).GetAgreementLog(ctx, req.(*AgreementLogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgreementLog_UpdateAgreementLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgreementLogUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgreementLogServer).UpdateAgreementLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agreementlog.v1.AgreementLog/UpdateAgreementLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgreementLogServer).UpdateAgreementLog(ctx, req.(*AgreementLogUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgreementLog_CreateAgreementLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgreementLogCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgreementLogServer).CreateAgreementLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agreementlog.v1.AgreementLog/CreateAgreementLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgreementLogServer).CreateAgreementLog(ctx, req.(*AgreementLogCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgreementLog_DeleteAgreementLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgreementLogDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgreementLogServer).DeleteAgreementLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agreementlog.v1.AgreementLog/DeleteAgreementLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgreementLogServer).DeleteAgreementLog(ctx, req.(*AgreementLogDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgreementLog_BatchDeleteAgreementLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgreementLogBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgreementLogServer).BatchDeleteAgreementLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agreementlog.v1.AgreementLog/BatchDeleteAgreementLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgreementLogServer).BatchDeleteAgreementLog(ctx, req.(*AgreementLogBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AgreementLog_ServiceDesc is the grpc.ServiceDesc for AgreementLog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgreementLog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agreementlog.v1.AgreementLog",
	HandlerType: (*AgreementLogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAgreementLogPage",
			Handler:    _AgreementLog_GetAgreementLogPage_Handler,
		},
		{
			MethodName: "GetAgreementLog",
			Handler:    _AgreementLog_GetAgreementLog_Handler,
		},
		{
			MethodName: "UpdateAgreementLog",
			Handler:    _AgreementLog_UpdateAgreementLog_Handler,
		},
		{
			MethodName: "CreateAgreementLog",
			Handler:    _AgreementLog_CreateAgreementLog_Handler,
		},
		{
			MethodName: "DeleteAgreementLog",
			Handler:    _AgreementLog_DeleteAgreementLog_Handler,
		},
		{
			MethodName: "BatchDeleteAgreementLog",
			Handler:    _AgreementLog_BatchDeleteAgreementLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/novel/agreementlog/v1/agreement_log.proto",
}
