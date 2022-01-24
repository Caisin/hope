// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/novel/ambalance/v1/am_balance.proto

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

// AmBalanceClient is the client API for AmBalance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AmBalanceClient interface {
	// 分页查询AmBalance
	GetPageAmBalance(ctx context.Context, in *AmBalancePageReq, opts ...grpc.CallOption) (*AmBalancePageReply, error)
	// 获取AmBalance
	GetAmBalance(ctx context.Context, in *AmBalanceReq, opts ...grpc.CallOption) (*AmBalanceReply, error)
	// 更新AmBalance
	UpdateAmBalance(ctx context.Context, in *AmBalanceUpdateReq, opts ...grpc.CallOption) (*AmBalanceUpdateReply, error)
	// 创建AmBalance
	CreateAmBalance(ctx context.Context, in *AmBalanceCreateReq, opts ...grpc.CallOption) (*AmBalanceCreateReply, error)
	// 删除AmBalance
	DeleteAmBalance(ctx context.Context, in *AmBalanceDeleteReq, opts ...grpc.CallOption) (*AmBalanceDeleteReply, error)
	// 批量删除AmBalance
	BatchDeleteAmBalance(ctx context.Context, in *AmBalanceBatchDeleteReq, opts ...grpc.CallOption) (*AmBalanceDeleteReply, error)
}

type amBalanceClient struct {
	cc grpc.ClientConnInterface
}

func NewAmBalanceClient(cc grpc.ClientConnInterface) AmBalanceClient {
	return &amBalanceClient{cc}
}

func (c *amBalanceClient) GetPageAmBalance(ctx context.Context, in *AmBalancePageReq, opts ...grpc.CallOption) (*AmBalancePageReply, error) {
	out := new(AmBalancePageReply)
	err := c.cc.Invoke(ctx, "/ambalance.v1.AmBalance/GetPageAmBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *amBalanceClient) GetAmBalance(ctx context.Context, in *AmBalanceReq, opts ...grpc.CallOption) (*AmBalanceReply, error) {
	out := new(AmBalanceReply)
	err := c.cc.Invoke(ctx, "/ambalance.v1.AmBalance/GetAmBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *amBalanceClient) UpdateAmBalance(ctx context.Context, in *AmBalanceUpdateReq, opts ...grpc.CallOption) (*AmBalanceUpdateReply, error) {
	out := new(AmBalanceUpdateReply)
	err := c.cc.Invoke(ctx, "/ambalance.v1.AmBalance/UpdateAmBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *amBalanceClient) CreateAmBalance(ctx context.Context, in *AmBalanceCreateReq, opts ...grpc.CallOption) (*AmBalanceCreateReply, error) {
	out := new(AmBalanceCreateReply)
	err := c.cc.Invoke(ctx, "/ambalance.v1.AmBalance/CreateAmBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *amBalanceClient) DeleteAmBalance(ctx context.Context, in *AmBalanceDeleteReq, opts ...grpc.CallOption) (*AmBalanceDeleteReply, error) {
	out := new(AmBalanceDeleteReply)
	err := c.cc.Invoke(ctx, "/ambalance.v1.AmBalance/DeleteAmBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *amBalanceClient) BatchDeleteAmBalance(ctx context.Context, in *AmBalanceBatchDeleteReq, opts ...grpc.CallOption) (*AmBalanceDeleteReply, error) {
	out := new(AmBalanceDeleteReply)
	err := c.cc.Invoke(ctx, "/ambalance.v1.AmBalance/BatchDeleteAmBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AmBalanceServer is the server API for AmBalance service.
// All implementations must embed UnimplementedAmBalanceServer
// for forward compatibility
type AmBalanceServer interface {
	// 分页查询AmBalance
	GetPageAmBalance(context.Context, *AmBalancePageReq) (*AmBalancePageReply, error)
	// 获取AmBalance
	GetAmBalance(context.Context, *AmBalanceReq) (*AmBalanceReply, error)
	// 更新AmBalance
	UpdateAmBalance(context.Context, *AmBalanceUpdateReq) (*AmBalanceUpdateReply, error)
	// 创建AmBalance
	CreateAmBalance(context.Context, *AmBalanceCreateReq) (*AmBalanceCreateReply, error)
	// 删除AmBalance
	DeleteAmBalance(context.Context, *AmBalanceDeleteReq) (*AmBalanceDeleteReply, error)
	// 批量删除AmBalance
	BatchDeleteAmBalance(context.Context, *AmBalanceBatchDeleteReq) (*AmBalanceDeleteReply, error)
	mustEmbedUnimplementedAmBalanceServer()
}

// UnimplementedAmBalanceServer must be embedded to have forward compatible implementations.
type UnimplementedAmBalanceServer struct {
}

func (UnimplementedAmBalanceServer) GetPageAmBalance(context.Context, *AmBalancePageReq) (*AmBalancePageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPageAmBalance not implemented")
}
func (UnimplementedAmBalanceServer) GetAmBalance(context.Context, *AmBalanceReq) (*AmBalanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAmBalance not implemented")
}
func (UnimplementedAmBalanceServer) UpdateAmBalance(context.Context, *AmBalanceUpdateReq) (*AmBalanceUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAmBalance not implemented")
}
func (UnimplementedAmBalanceServer) CreateAmBalance(context.Context, *AmBalanceCreateReq) (*AmBalanceCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAmBalance not implemented")
}
func (UnimplementedAmBalanceServer) DeleteAmBalance(context.Context, *AmBalanceDeleteReq) (*AmBalanceDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAmBalance not implemented")
}
func (UnimplementedAmBalanceServer) BatchDeleteAmBalance(context.Context, *AmBalanceBatchDeleteReq) (*AmBalanceDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteAmBalance not implemented")
}
func (UnimplementedAmBalanceServer) mustEmbedUnimplementedAmBalanceServer() {}

// UnsafeAmBalanceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AmBalanceServer will
// result in compilation errors.
type UnsafeAmBalanceServer interface {
	mustEmbedUnimplementedAmBalanceServer()
}

func RegisterAmBalanceServer(s grpc.ServiceRegistrar, srv AmBalanceServer) {
	s.RegisterService(&AmBalance_ServiceDesc, srv)
}

func _AmBalance_GetPageAmBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AmBalancePageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmBalanceServer).GetPageAmBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ambalance.v1.AmBalance/GetPageAmBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmBalanceServer).GetPageAmBalance(ctx, req.(*AmBalancePageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AmBalance_GetAmBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AmBalanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmBalanceServer).GetAmBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ambalance.v1.AmBalance/GetAmBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmBalanceServer).GetAmBalance(ctx, req.(*AmBalanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AmBalance_UpdateAmBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AmBalanceUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmBalanceServer).UpdateAmBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ambalance.v1.AmBalance/UpdateAmBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmBalanceServer).UpdateAmBalance(ctx, req.(*AmBalanceUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AmBalance_CreateAmBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AmBalanceCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmBalanceServer).CreateAmBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ambalance.v1.AmBalance/CreateAmBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmBalanceServer).CreateAmBalance(ctx, req.(*AmBalanceCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AmBalance_DeleteAmBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AmBalanceDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmBalanceServer).DeleteAmBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ambalance.v1.AmBalance/DeleteAmBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmBalanceServer).DeleteAmBalance(ctx, req.(*AmBalanceDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AmBalance_BatchDeleteAmBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AmBalanceBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmBalanceServer).BatchDeleteAmBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ambalance.v1.AmBalance/BatchDeleteAmBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmBalanceServer).BatchDeleteAmBalance(ctx, req.(*AmBalanceBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AmBalance_ServiceDesc is the grpc.ServiceDesc for AmBalance service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AmBalance_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ambalance.v1.AmBalance",
	HandlerType: (*AmBalanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPageAmBalance",
			Handler:    _AmBalance_GetPageAmBalance_Handler,
		},
		{
			MethodName: "GetAmBalance",
			Handler:    _AmBalance_GetAmBalance_Handler,
		},
		{
			MethodName: "UpdateAmBalance",
			Handler:    _AmBalance_UpdateAmBalance_Handler,
		},
		{
			MethodName: "CreateAmBalance",
			Handler:    _AmBalance_CreateAmBalance_Handler,
		},
		{
			MethodName: "DeleteAmBalance",
			Handler:    _AmBalance_DeleteAmBalance_Handler,
		},
		{
			MethodName: "BatchDeleteAmBalance",
			Handler:    _AmBalance_BatchDeleteAmBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/novel/ambalance/v1/am_balance.proto",
}