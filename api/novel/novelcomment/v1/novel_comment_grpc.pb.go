// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/novel/novelcomment/v1/novel_comment.proto

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

// NovelCommentClient is the client API for NovelComment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NovelCommentClient interface {
	// 分页查询NovelComment
	GetNovelCommentPage(ctx context.Context, in *NovelCommentPageReq, opts ...grpc.CallOption) (*NovelCommentPageReply, error)
	// 获取NovelComment
	GetNovelComment(ctx context.Context, in *NovelCommentReq, opts ...grpc.CallOption) (*NovelCommentReply, error)
	// 更新NovelComment
	UpdateNovelComment(ctx context.Context, in *NovelCommentUpdateReq, opts ...grpc.CallOption) (*NovelCommentUpdateReply, error)
	// 创建NovelComment
	CreateNovelComment(ctx context.Context, in *NovelCommentCreateReq, opts ...grpc.CallOption) (*NovelCommentCreateReply, error)
	// 删除NovelComment
	DeleteNovelComment(ctx context.Context, in *NovelCommentDeleteReq, opts ...grpc.CallOption) (*NovelCommentDeleteReply, error)
	// 批量删除NovelComment
	BatchDeleteNovelComment(ctx context.Context, in *NovelCommentBatchDeleteReq, opts ...grpc.CallOption) (*NovelCommentDeleteReply, error)
}

type novelCommentClient struct {
	cc grpc.ClientConnInterface
}

func NewNovelCommentClient(cc grpc.ClientConnInterface) NovelCommentClient {
	return &novelCommentClient{cc}
}

func (c *novelCommentClient) GetNovelCommentPage(ctx context.Context, in *NovelCommentPageReq, opts ...grpc.CallOption) (*NovelCommentPageReply, error) {
	out := new(NovelCommentPageReply)
	err := c.cc.Invoke(ctx, "/novelcomment.v1.NovelComment/GetNovelCommentPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *novelCommentClient) GetNovelComment(ctx context.Context, in *NovelCommentReq, opts ...grpc.CallOption) (*NovelCommentReply, error) {
	out := new(NovelCommentReply)
	err := c.cc.Invoke(ctx, "/novelcomment.v1.NovelComment/GetNovelComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *novelCommentClient) UpdateNovelComment(ctx context.Context, in *NovelCommentUpdateReq, opts ...grpc.CallOption) (*NovelCommentUpdateReply, error) {
	out := new(NovelCommentUpdateReply)
	err := c.cc.Invoke(ctx, "/novelcomment.v1.NovelComment/UpdateNovelComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *novelCommentClient) CreateNovelComment(ctx context.Context, in *NovelCommentCreateReq, opts ...grpc.CallOption) (*NovelCommentCreateReply, error) {
	out := new(NovelCommentCreateReply)
	err := c.cc.Invoke(ctx, "/novelcomment.v1.NovelComment/CreateNovelComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *novelCommentClient) DeleteNovelComment(ctx context.Context, in *NovelCommentDeleteReq, opts ...grpc.CallOption) (*NovelCommentDeleteReply, error) {
	out := new(NovelCommentDeleteReply)
	err := c.cc.Invoke(ctx, "/novelcomment.v1.NovelComment/DeleteNovelComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *novelCommentClient) BatchDeleteNovelComment(ctx context.Context, in *NovelCommentBatchDeleteReq, opts ...grpc.CallOption) (*NovelCommentDeleteReply, error) {
	out := new(NovelCommentDeleteReply)
	err := c.cc.Invoke(ctx, "/novelcomment.v1.NovelComment/BatchDeleteNovelComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NovelCommentServer is the server API for NovelComment service.
// All implementations must embed UnimplementedNovelCommentServer
// for forward compatibility
type NovelCommentServer interface {
	// 分页查询NovelComment
	GetNovelCommentPage(context.Context, *NovelCommentPageReq) (*NovelCommentPageReply, error)
	// 获取NovelComment
	GetNovelComment(context.Context, *NovelCommentReq) (*NovelCommentReply, error)
	// 更新NovelComment
	UpdateNovelComment(context.Context, *NovelCommentUpdateReq) (*NovelCommentUpdateReply, error)
	// 创建NovelComment
	CreateNovelComment(context.Context, *NovelCommentCreateReq) (*NovelCommentCreateReply, error)
	// 删除NovelComment
	DeleteNovelComment(context.Context, *NovelCommentDeleteReq) (*NovelCommentDeleteReply, error)
	// 批量删除NovelComment
	BatchDeleteNovelComment(context.Context, *NovelCommentBatchDeleteReq) (*NovelCommentDeleteReply, error)
	mustEmbedUnimplementedNovelCommentServer()
}

// UnimplementedNovelCommentServer must be embedded to have forward compatible implementations.
type UnimplementedNovelCommentServer struct {
}

func (UnimplementedNovelCommentServer) GetNovelCommentPage(context.Context, *NovelCommentPageReq) (*NovelCommentPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNovelCommentPage not implemented")
}
func (UnimplementedNovelCommentServer) GetNovelComment(context.Context, *NovelCommentReq) (*NovelCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNovelComment not implemented")
}
func (UnimplementedNovelCommentServer) UpdateNovelComment(context.Context, *NovelCommentUpdateReq) (*NovelCommentUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNovelComment not implemented")
}
func (UnimplementedNovelCommentServer) CreateNovelComment(context.Context, *NovelCommentCreateReq) (*NovelCommentCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNovelComment not implemented")
}
func (UnimplementedNovelCommentServer) DeleteNovelComment(context.Context, *NovelCommentDeleteReq) (*NovelCommentDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNovelComment not implemented")
}
func (UnimplementedNovelCommentServer) BatchDeleteNovelComment(context.Context, *NovelCommentBatchDeleteReq) (*NovelCommentDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteNovelComment not implemented")
}
func (UnimplementedNovelCommentServer) mustEmbedUnimplementedNovelCommentServer() {}

// UnsafeNovelCommentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NovelCommentServer will
// result in compilation errors.
type UnsafeNovelCommentServer interface {
	mustEmbedUnimplementedNovelCommentServer()
}

func RegisterNovelCommentServer(s grpc.ServiceRegistrar, srv NovelCommentServer) {
	s.RegisterService(&NovelComment_ServiceDesc, srv)
}

func _NovelComment_GetNovelCommentPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelCommentPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelCommentServer).GetNovelCommentPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novelcomment.v1.NovelComment/GetNovelCommentPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelCommentServer).GetNovelCommentPage(ctx, req.(*NovelCommentPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NovelComment_GetNovelComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelCommentServer).GetNovelComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novelcomment.v1.NovelComment/GetNovelComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelCommentServer).GetNovelComment(ctx, req.(*NovelCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NovelComment_UpdateNovelComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelCommentUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelCommentServer).UpdateNovelComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novelcomment.v1.NovelComment/UpdateNovelComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelCommentServer).UpdateNovelComment(ctx, req.(*NovelCommentUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NovelComment_CreateNovelComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelCommentCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelCommentServer).CreateNovelComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novelcomment.v1.NovelComment/CreateNovelComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelCommentServer).CreateNovelComment(ctx, req.(*NovelCommentCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NovelComment_DeleteNovelComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelCommentDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelCommentServer).DeleteNovelComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novelcomment.v1.NovelComment/DeleteNovelComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelCommentServer).DeleteNovelComment(ctx, req.(*NovelCommentDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NovelComment_BatchDeleteNovelComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NovelCommentBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelCommentServer).BatchDeleteNovelComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novelcomment.v1.NovelComment/BatchDeleteNovelComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelCommentServer).BatchDeleteNovelComment(ctx, req.(*NovelCommentBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// NovelComment_ServiceDesc is the grpc.ServiceDesc for NovelComment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NovelComment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "novelcomment.v1.NovelComment",
	HandlerType: (*NovelCommentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNovelCommentPage",
			Handler:    _NovelComment_GetNovelCommentPage_Handler,
		},
		{
			MethodName: "GetNovelComment",
			Handler:    _NovelComment_GetNovelComment_Handler,
		},
		{
			MethodName: "UpdateNovelComment",
			Handler:    _NovelComment_UpdateNovelComment_Handler,
		},
		{
			MethodName: "CreateNovelComment",
			Handler:    _NovelComment_CreateNovelComment_Handler,
		},
		{
			MethodName: "DeleteNovelComment",
			Handler:    _NovelComment_DeleteNovelComment_Handler,
		},
		{
			MethodName: "BatchDeleteNovelComment",
			Handler:    _NovelComment_BatchDeleteNovelComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/novel/novelcomment/v1/novel_comment.proto",
}
