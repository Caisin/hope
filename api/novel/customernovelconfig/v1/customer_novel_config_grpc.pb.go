// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/novel/customernovelconfig/v1/customer_novel_config.proto

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

// CustomerNovelConfigClient is the client API for CustomerNovelConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerNovelConfigClient interface {
	// 分页查询CustomerNovelConfig
	GetPageCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigPageReq, opts ...grpc.CallOption) (*CustomerNovelConfigPageReply, error)
	// 获取CustomerNovelConfig
	GetCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigReq, opts ...grpc.CallOption) (*CustomerNovelConfigReply, error)
	// 更新CustomerNovelConfig
	UpdateCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigUpdateReq, opts ...grpc.CallOption) (*CustomerNovelConfigUpdateReply, error)
	// 创建CustomerNovelConfig
	CreateCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigCreateReq, opts ...grpc.CallOption) (*CustomerNovelConfigCreateReply, error)
	// 删除CustomerNovelConfig
	DeleteCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigDeleteReq, opts ...grpc.CallOption) (*CustomerNovelConfigDeleteReply, error)
	// 批量删除CustomerNovelConfig
	BatchDeleteCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigBatchDeleteReq, opts ...grpc.CallOption) (*CustomerNovelConfigDeleteReply, error)
}

type customerNovelConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerNovelConfigClient(cc grpc.ClientConnInterface) CustomerNovelConfigClient {
	return &customerNovelConfigClient{cc}
}

func (c *customerNovelConfigClient) GetPageCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigPageReq, opts ...grpc.CallOption) (*CustomerNovelConfigPageReply, error) {
	out := new(CustomerNovelConfigPageReply)
	err := c.cc.Invoke(ctx, "/customernovelconfig.v1.CustomerNovelConfig/GetPageCustomerNovelConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerNovelConfigClient) GetCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigReq, opts ...grpc.CallOption) (*CustomerNovelConfigReply, error) {
	out := new(CustomerNovelConfigReply)
	err := c.cc.Invoke(ctx, "/customernovelconfig.v1.CustomerNovelConfig/GetCustomerNovelConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerNovelConfigClient) UpdateCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigUpdateReq, opts ...grpc.CallOption) (*CustomerNovelConfigUpdateReply, error) {
	out := new(CustomerNovelConfigUpdateReply)
	err := c.cc.Invoke(ctx, "/customernovelconfig.v1.CustomerNovelConfig/UpdateCustomerNovelConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerNovelConfigClient) CreateCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigCreateReq, opts ...grpc.CallOption) (*CustomerNovelConfigCreateReply, error) {
	out := new(CustomerNovelConfigCreateReply)
	err := c.cc.Invoke(ctx, "/customernovelconfig.v1.CustomerNovelConfig/CreateCustomerNovelConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerNovelConfigClient) DeleteCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigDeleteReq, opts ...grpc.CallOption) (*CustomerNovelConfigDeleteReply, error) {
	out := new(CustomerNovelConfigDeleteReply)
	err := c.cc.Invoke(ctx, "/customernovelconfig.v1.CustomerNovelConfig/DeleteCustomerNovelConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerNovelConfigClient) BatchDeleteCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigBatchDeleteReq, opts ...grpc.CallOption) (*CustomerNovelConfigDeleteReply, error) {
	out := new(CustomerNovelConfigDeleteReply)
	err := c.cc.Invoke(ctx, "/customernovelconfig.v1.CustomerNovelConfig/BatchDeleteCustomerNovelConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerNovelConfigServer is the server API for CustomerNovelConfig service.
// All implementations must embed UnimplementedCustomerNovelConfigServer
// for forward compatibility
type CustomerNovelConfigServer interface {
	// 分页查询CustomerNovelConfig
	GetPageCustomerNovelConfig(context.Context, *CustomerNovelConfigPageReq) (*CustomerNovelConfigPageReply, error)
	// 获取CustomerNovelConfig
	GetCustomerNovelConfig(context.Context, *CustomerNovelConfigReq) (*CustomerNovelConfigReply, error)
	// 更新CustomerNovelConfig
	UpdateCustomerNovelConfig(context.Context, *CustomerNovelConfigUpdateReq) (*CustomerNovelConfigUpdateReply, error)
	// 创建CustomerNovelConfig
	CreateCustomerNovelConfig(context.Context, *CustomerNovelConfigCreateReq) (*CustomerNovelConfigCreateReply, error)
	// 删除CustomerNovelConfig
	DeleteCustomerNovelConfig(context.Context, *CustomerNovelConfigDeleteReq) (*CustomerNovelConfigDeleteReply, error)
	// 批量删除CustomerNovelConfig
	BatchDeleteCustomerNovelConfig(context.Context, *CustomerNovelConfigBatchDeleteReq) (*CustomerNovelConfigDeleteReply, error)
	mustEmbedUnimplementedCustomerNovelConfigServer()
}

// UnimplementedCustomerNovelConfigServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerNovelConfigServer struct {
}

func (UnimplementedCustomerNovelConfigServer) GetPageCustomerNovelConfig(context.Context, *CustomerNovelConfigPageReq) (*CustomerNovelConfigPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPageCustomerNovelConfig not implemented")
}
func (UnimplementedCustomerNovelConfigServer) GetCustomerNovelConfig(context.Context, *CustomerNovelConfigReq) (*CustomerNovelConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerNovelConfig not implemented")
}
func (UnimplementedCustomerNovelConfigServer) UpdateCustomerNovelConfig(context.Context, *CustomerNovelConfigUpdateReq) (*CustomerNovelConfigUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomerNovelConfig not implemented")
}
func (UnimplementedCustomerNovelConfigServer) CreateCustomerNovelConfig(context.Context, *CustomerNovelConfigCreateReq) (*CustomerNovelConfigCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomerNovelConfig not implemented")
}
func (UnimplementedCustomerNovelConfigServer) DeleteCustomerNovelConfig(context.Context, *CustomerNovelConfigDeleteReq) (*CustomerNovelConfigDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomerNovelConfig not implemented")
}
func (UnimplementedCustomerNovelConfigServer) BatchDeleteCustomerNovelConfig(context.Context, *CustomerNovelConfigBatchDeleteReq) (*CustomerNovelConfigDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteCustomerNovelConfig not implemented")
}
func (UnimplementedCustomerNovelConfigServer) mustEmbedUnimplementedCustomerNovelConfigServer() {}

// UnsafeCustomerNovelConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerNovelConfigServer will
// result in compilation errors.
type UnsafeCustomerNovelConfigServer interface {
	mustEmbedUnimplementedCustomerNovelConfigServer()
}

func RegisterCustomerNovelConfigServer(s grpc.ServiceRegistrar, srv CustomerNovelConfigServer) {
	s.RegisterService(&CustomerNovelConfig_ServiceDesc, srv)
}

func _CustomerNovelConfig_GetPageCustomerNovelConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerNovelConfigPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerNovelConfigServer).GetPageCustomerNovelConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customernovelconfig.v1.CustomerNovelConfig/GetPageCustomerNovelConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerNovelConfigServer).GetPageCustomerNovelConfig(ctx, req.(*CustomerNovelConfigPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerNovelConfig_GetCustomerNovelConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerNovelConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerNovelConfigServer).GetCustomerNovelConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customernovelconfig.v1.CustomerNovelConfig/GetCustomerNovelConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerNovelConfigServer).GetCustomerNovelConfig(ctx, req.(*CustomerNovelConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerNovelConfig_UpdateCustomerNovelConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerNovelConfigUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerNovelConfigServer).UpdateCustomerNovelConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customernovelconfig.v1.CustomerNovelConfig/UpdateCustomerNovelConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerNovelConfigServer).UpdateCustomerNovelConfig(ctx, req.(*CustomerNovelConfigUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerNovelConfig_CreateCustomerNovelConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerNovelConfigCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerNovelConfigServer).CreateCustomerNovelConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customernovelconfig.v1.CustomerNovelConfig/CreateCustomerNovelConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerNovelConfigServer).CreateCustomerNovelConfig(ctx, req.(*CustomerNovelConfigCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerNovelConfig_DeleteCustomerNovelConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerNovelConfigDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerNovelConfigServer).DeleteCustomerNovelConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customernovelconfig.v1.CustomerNovelConfig/DeleteCustomerNovelConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerNovelConfigServer).DeleteCustomerNovelConfig(ctx, req.(*CustomerNovelConfigDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerNovelConfig_BatchDeleteCustomerNovelConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerNovelConfigBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerNovelConfigServer).BatchDeleteCustomerNovelConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customernovelconfig.v1.CustomerNovelConfig/BatchDeleteCustomerNovelConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerNovelConfigServer).BatchDeleteCustomerNovelConfig(ctx, req.(*CustomerNovelConfigBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerNovelConfig_ServiceDesc is the grpc.ServiceDesc for CustomerNovelConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerNovelConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "customernovelconfig.v1.CustomerNovelConfig",
	HandlerType: (*CustomerNovelConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPageCustomerNovelConfig",
			Handler:    _CustomerNovelConfig_GetPageCustomerNovelConfig_Handler,
		},
		{
			MethodName: "GetCustomerNovelConfig",
			Handler:    _CustomerNovelConfig_GetCustomerNovelConfig_Handler,
		},
		{
			MethodName: "UpdateCustomerNovelConfig",
			Handler:    _CustomerNovelConfig_UpdateCustomerNovelConfig_Handler,
		},
		{
			MethodName: "CreateCustomerNovelConfig",
			Handler:    _CustomerNovelConfig_CreateCustomerNovelConfig_Handler,
		},
		{
			MethodName: "DeleteCustomerNovelConfig",
			Handler:    _CustomerNovelConfig_DeleteCustomerNovelConfig_Handler,
		},
		{
			MethodName: "BatchDeleteCustomerNovelConfig",
			Handler:    _CustomerNovelConfig_BatchDeleteCustomerNovelConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/novel/customernovelconfig/v1/customer_novel_config.proto",
}
