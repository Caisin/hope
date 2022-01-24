// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/param/noveltag/v1/novel_tag.proto

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

// NovelTagClient is the client API for NovelTag service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NovelTagClient interface {
	// 分页查询NovelTag
	GetPageNovelTag(ctx context.Context, in *NovelTagPageReq, opts ...grpc.CallOption) (*NovelTagPageReply, error)
	// 获取NovelTag
	GetNovelTag(ctx context.Context, in *NovelTagReq, opts ...grpc.CallOption) (*NovelTagReply, error)
	// 更新NovelTag
	UpdateNovelTag(ctx context.Context, in *NovelTagUpdateReq, opts ...grpc.CallOption) (*NovelTagUpdateReply, error)
	// 创建NovelTag
	CreateNovelTag(ctx context.Context, in *NovelTagCreateReq, opts ...grpc.CallOption) (*NovelTagCreateReply, error)
	// 删除NovelTag
	DeleteNovelTag(ctx context.Context, in *NovelTagDeleteReq, opts ...grpc.CallOption) (*NovelTagDeleteReply, error)
	// 批量删除NovelTag
	BatchDeleteNovelTag(ctx context.Context, in *NovelTagBatchDeleteReq, opts ...grpc.CallOption) (*NovelTagDeleteReply, error)
}

type novelTagClient struct {
	cc grpc.ClientConnInterface
}

func NewNovelTagClient(cc grpc.ClientConnInterface) NovelTagClient {
	return &novelTagClient{cc}
}

func (c *novelTagClient) GetPageNovelTag(ctx context.Context, in *NovelTagPageReq, opts ...grpc.CallOption) (*NovelTagPageReply, error) {
	out := new(NovelTagPageReply)
	err := c.cc.Invoke(ctx, "/noveltag.v1.NovelTag/GetPageNovelTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *novelTagClient) GetNovelTag(ctx context.Context, in *NovelTagReq, opts ...grpc.CallOption) (*NovelTagReply, error) {
	out := new(NovelTagReply)
	err := c.cc.Invoke(ctx, "/noveltag.v1.NovelTag/GetNovelTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *novelTagClient) UpdateNovelTag(ctx context.Context, in *NovelTagUpdateReq, opts ...grpc.CallOption) (*NovelTagUpdateReply, error) {
	out := new(NovelTagUpdateReply)
	err := c.cc.Invoke(ctx, "/noveltag.v1.NovelTag/UpdateNovelTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *novelTagClient) CreateNovelTag(ctx context.Context, in *NovelTagCreateReq, opts ...grpc.CallOption) (*NovelTagCreateReply, error) {
	out := new(NovelTagCreateReply)
	err := c.cc.Invoke(ctx, "/noveltag.v1.NovelTag/CreateNovelTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *novelTagClient) DeleteNovelTag(ctx context.Context, in *NovelTagDeleteReq, opts ...grpc.CallOption) (*NovelTagDeleteReply, error) {
	out := new(NovelTagDeleteReply)
	err := c.cc.Invoke(ctx, "/noveltag.v1.NovelTag/DeleteNovelTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *novelTagClient) BatchDeleteNovelTag(ctx context.Context, in *NovelTagBatchDeleteReq, opts ...grpc.CallOption) (*NovelTagDeleteReply, error) {
	out := new(NovelTagDeleteReply)
	err := c.cc.Invoke(ctx, "/noveltag.v1.NovelTag/BatchDeleteNovelTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NovelTagServer is the server API for NovelTag service.
// All implementations must embed UnimplementedNovelTagServer
// for forward compatibility
type NovelTagServer interface {
	// 分页查询NovelTag
	GetPageNovelTag(context.Context, *NovelTagPageReq) (*NovelTagPageReply, error)
	// 获取NovelTag
	GetNovelTag(context.Context, *NovelTagReq) (*NovelTagReply, error)
	// 更新NovelTag
	UpdateNovelTag(context.Context, *NovelTagUpdateReq) (*NovelTagUpdateReply, error)
	// 创建NovelTag
	CreateNovelTag(context.Context, *NovelTagCreateReq) (*NovelTagCreateReply, error)
	// 删除NovelTag
	DeleteNovelTag(context.Context, *NovelTagDeleteReq) (*NovelTagDeleteReply, error)
	// 批量删除NovelTag
	BatchDeleteNovelTag(context.Context, *NovelTagBatchDeleteReq) (*NovelTagDeleteReply, error)
	mustEmbedUnimplementedNovelTagServer()
}

// UnimplementedNovelTagServer must be embedded to have forward compatible implementations.
type UnimplementedNovelTagServer struct {
}

func (UnimplementedNovelTagServer) GetPageNovelTag(context.Context, *NovelTagPageReq) (*NovelTagPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPageNovelTag not implemented")
}
func (UnimplementedNovelTagServer) GetNovelTag(context.Context, *NovelTagReq) (*NovelTagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNovelTag not implemented")
}
func (UnimplementedNovelTagServer) UpdateNovelTag(context.Context, *NovelTagUpdateReq) (*NovelTagUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNovelTag not implemented")
}
func (UnimplementedNovelTagServer) CreateNovelTag(context.Context, *NovelTagCreateReq) (*NovelTagCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNovelTag not implemented")
}
func (UnimplementedNovelTagServer) DeleteNovelTag(context.Context, *NovelTagDeleteReq) (*NovelTagDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNovelTag not implemented")
}
func (UnimplementedNovelTagServer) BatchDeleteNovelTag(context.Context, *NovelTagBatchDeleteReq) (*NovelTagDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteNovelTag not implemented")
}
func (UnimplementedNovelTagServer) mustEmbedUnimplementedNovelTagServer() {}

// UnsafeNovelTagServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NovelTagServer will
// result in compilation errors.
type UnsafeNovelTagServer interface {
	mustEmbedUnimplementedNovelTagServer()
}

func RegisterNovelTagServer(s grpc.ServiceRegistrar, srv NovelTagServer) {
	s.RegisterService(&NovelTag_ServiceDesc, srv)
}

func _NovelTag_GetPageNovelTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelTagPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelTagServer).GetPageNovelTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noveltag.v1.NovelTag/GetPageNovelTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelTagServer).GetPageNovelTag(ctx, req.(*NovelTagPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NovelTag_GetNovelTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelTagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelTagServer).GetNovelTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noveltag.v1.NovelTag/GetNovelTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelTagServer).GetNovelTag(ctx, req.(*NovelTagReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NovelTag_UpdateNovelTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelTagUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelTagServer).UpdateNovelTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noveltag.v1.NovelTag/UpdateNovelTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelTagServer).UpdateNovelTag(ctx, req.(*NovelTagUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NovelTag_CreateNovelTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelTagCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelTagServer).CreateNovelTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noveltag.v1.NovelTag/CreateNovelTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelTagServer).CreateNovelTag(ctx, req.(*NovelTagCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NovelTag_DeleteNovelTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelTagDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelTagServer).DeleteNovelTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noveltag.v1.NovelTag/DeleteNovelTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelTagServer).DeleteNovelTag(ctx, req.(*NovelTagDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NovelTag_BatchDeleteNovelTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelTagBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelTagServer).BatchDeleteNovelTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noveltag.v1.NovelTag/BatchDeleteNovelTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelTagServer).BatchDeleteNovelTag(ctx, req.(*NovelTagBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// NovelTag_ServiceDesc is the grpc.ServiceDesc for NovelTag service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NovelTag_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "noveltag.v1.NovelTag",
	HandlerType: (*NovelTagServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPageNovelTag",
			Handler:    _NovelTag_GetPageNovelTag_Handler,
		},
		{
			MethodName: "GetNovelTag",
			Handler:    _NovelTag_GetNovelTag_Handler,
		},
		{
			MethodName: "UpdateNovelTag",
			Handler:    _NovelTag_UpdateNovelTag_Handler,
		},
		{
			MethodName: "CreateNovelTag",
			Handler:    _NovelTag_CreateNovelTag_Handler,
		},
		{
			MethodName: "DeleteNovelTag",
			Handler:    _NovelTag_DeleteNovelTag_Handler,
		},
		{
			MethodName: "BatchDeleteNovelTag",
			Handler:    _NovelTag_BatchDeleteNovelTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/param/noveltag/v1/novel_tag.proto",
}