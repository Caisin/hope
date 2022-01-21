// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/novel/novelclassify/v1/novel_classify.proto

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

// NovelClassifyClient is the client API for NovelClassify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NovelClassifyClient interface {
	// 分页查询NovelClassify
	GetPageNovelClassify(ctx context.Context, in *NovelClassifyPageReq, opts ...grpc.CallOption) (*NovelClassifyPageReply, error)
	// 获取NovelClassify
	GetNovelClassify(ctx context.Context, in *NovelClassifyReq, opts ...grpc.CallOption) (*NovelClassifyReply, error)
	// 更新NovelClassify
	UpdateNovelClassify(ctx context.Context, in *NovelClassifyUpdateReq, opts ...grpc.CallOption) (*NovelClassifyUpdateReply, error)
	// 创建NovelClassify
	CreateNovelClassify(ctx context.Context, in *NovelClassifyCreateReq, opts ...grpc.CallOption) (*NovelClassifyCreateReply, error)
	// 删除NovelClassify
	DeleteNovelClassify(ctx context.Context, in *NovelClassifyDeleteReq, opts ...grpc.CallOption) (*NovelClassifyDeleteReply, error)
	// 批量删除NovelClassify
	BatchDeleteNovelClassify(ctx context.Context, in *NovelClassifyBatchDeleteReq, opts ...grpc.CallOption) (*NovelClassifyDeleteReply, error)
}

type novelClassifyClient struct {
	cc grpc.ClientConnInterface
}

func NewNovelClassifyClient(cc grpc.ClientConnInterface) NovelClassifyClient {
	return &novelClassifyClient{cc}
}

func (c *novelClassifyClient) GetPageNovelClassify(ctx context.Context, in *NovelClassifyPageReq, opts ...grpc.CallOption) (*NovelClassifyPageReply, error) {
	out := new(NovelClassifyPageReply)
	err := c.cc.Invoke(ctx, "/novelclassify.v1.NovelClassify/GetPageNovelClassify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *novelClassifyClient) GetNovelClassify(ctx context.Context, in *NovelClassifyReq, opts ...grpc.CallOption) (*NovelClassifyReply, error) {
	out := new(NovelClassifyReply)
	err := c.cc.Invoke(ctx, "/novelclassify.v1.NovelClassify/GetNovelClassify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *novelClassifyClient) UpdateNovelClassify(ctx context.Context, in *NovelClassifyUpdateReq, opts ...grpc.CallOption) (*NovelClassifyUpdateReply, error) {
	out := new(NovelClassifyUpdateReply)
	err := c.cc.Invoke(ctx, "/novelclassify.v1.NovelClassify/UpdateNovelClassify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *novelClassifyClient) CreateNovelClassify(ctx context.Context, in *NovelClassifyCreateReq, opts ...grpc.CallOption) (*NovelClassifyCreateReply, error) {
	out := new(NovelClassifyCreateReply)
	err := c.cc.Invoke(ctx, "/novelclassify.v1.NovelClassify/CreateNovelClassify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *novelClassifyClient) DeleteNovelClassify(ctx context.Context, in *NovelClassifyDeleteReq, opts ...grpc.CallOption) (*NovelClassifyDeleteReply, error) {
	out := new(NovelClassifyDeleteReply)
	err := c.cc.Invoke(ctx, "/novelclassify.v1.NovelClassify/DeleteNovelClassify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *novelClassifyClient) BatchDeleteNovelClassify(ctx context.Context, in *NovelClassifyBatchDeleteReq, opts ...grpc.CallOption) (*NovelClassifyDeleteReply, error) {
	out := new(NovelClassifyDeleteReply)
	err := c.cc.Invoke(ctx, "/novelclassify.v1.NovelClassify/BatchDeleteNovelClassify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NovelClassifyServer is the server API for NovelClassify service.
// All implementations must embed UnimplementedNovelClassifyServer
// for forward compatibility
type NovelClassifyServer interface {
	// 分页查询NovelClassify
	GetPageNovelClassify(context.Context, *NovelClassifyPageReq) (*NovelClassifyPageReply, error)
	// 获取NovelClassify
	GetNovelClassify(context.Context, *NovelClassifyReq) (*NovelClassifyReply, error)
	// 更新NovelClassify
	UpdateNovelClassify(context.Context, *NovelClassifyUpdateReq) (*NovelClassifyUpdateReply, error)
	// 创建NovelClassify
	CreateNovelClassify(context.Context, *NovelClassifyCreateReq) (*NovelClassifyCreateReply, error)
	// 删除NovelClassify
	DeleteNovelClassify(context.Context, *NovelClassifyDeleteReq) (*NovelClassifyDeleteReply, error)
	// 批量删除NovelClassify
	BatchDeleteNovelClassify(context.Context, *NovelClassifyBatchDeleteReq) (*NovelClassifyDeleteReply, error)
	mustEmbedUnimplementedNovelClassifyServer()
}

// UnimplementedNovelClassifyServer must be embedded to have forward compatible implementations.
type UnimplementedNovelClassifyServer struct {
}

func (UnimplementedNovelClassifyServer) GetPageNovelClassify(context.Context, *NovelClassifyPageReq) (*NovelClassifyPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPageNovelClassify not implemented")
}
func (UnimplementedNovelClassifyServer) GetNovelClassify(context.Context, *NovelClassifyReq) (*NovelClassifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNovelClassify not implemented")
}
func (UnimplementedNovelClassifyServer) UpdateNovelClassify(context.Context, *NovelClassifyUpdateReq) (*NovelClassifyUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNovelClassify not implemented")
}
func (UnimplementedNovelClassifyServer) CreateNovelClassify(context.Context, *NovelClassifyCreateReq) (*NovelClassifyCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNovelClassify not implemented")
}
func (UnimplementedNovelClassifyServer) DeleteNovelClassify(context.Context, *NovelClassifyDeleteReq) (*NovelClassifyDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNovelClassify not implemented")
}
func (UnimplementedNovelClassifyServer) BatchDeleteNovelClassify(context.Context, *NovelClassifyBatchDeleteReq) (*NovelClassifyDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteNovelClassify not implemented")
}
func (UnimplementedNovelClassifyServer) mustEmbedUnimplementedNovelClassifyServer() {}

// UnsafeNovelClassifyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NovelClassifyServer will
// result in compilation errors.
type UnsafeNovelClassifyServer interface {
	mustEmbedUnimplementedNovelClassifyServer()
}

func RegisterNovelClassifyServer(s grpc.ServiceRegistrar, srv NovelClassifyServer) {
	s.RegisterService(&NovelClassify_ServiceDesc, srv)
}

func _NovelClassify_GetPageNovelClassify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelClassifyPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelClassifyServer).GetPageNovelClassify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novelclassify.v1.NovelClassify/GetPageNovelClassify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelClassifyServer).GetPageNovelClassify(ctx, req.(*NovelClassifyPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NovelClassify_GetNovelClassify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelClassifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelClassifyServer).GetNovelClassify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novelclassify.v1.NovelClassify/GetNovelClassify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelClassifyServer).GetNovelClassify(ctx, req.(*NovelClassifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NovelClassify_UpdateNovelClassify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelClassifyUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelClassifyServer).UpdateNovelClassify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novelclassify.v1.NovelClassify/UpdateNovelClassify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelClassifyServer).UpdateNovelClassify(ctx, req.(*NovelClassifyUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NovelClassify_CreateNovelClassify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelClassifyCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelClassifyServer).CreateNovelClassify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novelclassify.v1.NovelClassify/CreateNovelClassify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelClassifyServer).CreateNovelClassify(ctx, req.(*NovelClassifyCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NovelClassify_DeleteNovelClassify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelClassifyDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelClassifyServer).DeleteNovelClassify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novelclassify.v1.NovelClassify/DeleteNovelClassify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelClassifyServer).DeleteNovelClassify(ctx, req.(*NovelClassifyDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NovelClassify_BatchDeleteNovelClassify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelClassifyBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelClassifyServer).BatchDeleteNovelClassify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novelclassify.v1.NovelClassify/BatchDeleteNovelClassify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelClassifyServer).BatchDeleteNovelClassify(ctx, req.(*NovelClassifyBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// NovelClassify_ServiceDesc is the grpc.ServiceDesc for NovelClassify service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NovelClassify_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "novelclassify.v1.NovelClassify",
	HandlerType: (*NovelClassifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPageNovelClassify",
			Handler:    _NovelClassify_GetPageNovelClassify_Handler,
		},
		{
			MethodName: "GetNovelClassify",
			Handler:    _NovelClassify_GetNovelClassify_Handler,
		},
		{
			MethodName: "UpdateNovelClassify",
			Handler:    _NovelClassify_UpdateNovelClassify_Handler,
		},
		{
			MethodName: "CreateNovelClassify",
			Handler:    _NovelClassify_CreateNovelClassify_Handler,
		},
		{
			MethodName: "DeleteNovelClassify",
			Handler:    _NovelClassify_DeleteNovelClassify_Handler,
		},
		{
			MethodName: "BatchDeleteNovelClassify",
			Handler:    _NovelClassify_BatchDeleteNovelClassify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/novel/novelclassify/v1/novel_classify.proto",
}
