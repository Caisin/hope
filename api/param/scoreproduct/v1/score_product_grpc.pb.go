// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/param/scoreproduct/v1/score_product.proto

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

// ScoreProductClient is the client API for ScoreProduct service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScoreProductClient interface {
	// 分页查询ScoreProduct
	GetScoreProductPage(ctx context.Context, in *ScoreProductPageReq, opts ...grpc.CallOption) (*ScoreProductPageReply, error)
	// 获取ScoreProduct
	GetScoreProduct(ctx context.Context, in *ScoreProductReq, opts ...grpc.CallOption) (*ScoreProductReply, error)
	// 更新ScoreProduct
	UpdateScoreProduct(ctx context.Context, in *ScoreProductUpdateReq, opts ...grpc.CallOption) (*ScoreProductUpdateReply, error)
	// 创建ScoreProduct
	CreateScoreProduct(ctx context.Context, in *ScoreProductCreateReq, opts ...grpc.CallOption) (*ScoreProductCreateReply, error)
	// 删除ScoreProduct
	DeleteScoreProduct(ctx context.Context, in *ScoreProductDeleteReq, opts ...grpc.CallOption) (*ScoreProductDeleteReply, error)
	// 批量删除ScoreProduct
	BatchDeleteScoreProduct(ctx context.Context, in *ScoreProductBatchDeleteReq, opts ...grpc.CallOption) (*ScoreProductDeleteReply, error)
}

type scoreProductClient struct {
	cc grpc.ClientConnInterface
}

func NewScoreProductClient(cc grpc.ClientConnInterface) ScoreProductClient {
	return &scoreProductClient{cc}
}

func (c *scoreProductClient) GetScoreProductPage(ctx context.Context, in *ScoreProductPageReq, opts ...grpc.CallOption) (*ScoreProductPageReply, error) {
	out := new(ScoreProductPageReply)
	err := c.cc.Invoke(ctx, "/scoreproduct.v1.ScoreProduct/GetScoreProductPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreProductClient) GetScoreProduct(ctx context.Context, in *ScoreProductReq, opts ...grpc.CallOption) (*ScoreProductReply, error) {
	out := new(ScoreProductReply)
	err := c.cc.Invoke(ctx, "/scoreproduct.v1.ScoreProduct/GetScoreProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreProductClient) UpdateScoreProduct(ctx context.Context, in *ScoreProductUpdateReq, opts ...grpc.CallOption) (*ScoreProductUpdateReply, error) {
	out := new(ScoreProductUpdateReply)
	err := c.cc.Invoke(ctx, "/scoreproduct.v1.ScoreProduct/UpdateScoreProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreProductClient) CreateScoreProduct(ctx context.Context, in *ScoreProductCreateReq, opts ...grpc.CallOption) (*ScoreProductCreateReply, error) {
	out := new(ScoreProductCreateReply)
	err := c.cc.Invoke(ctx, "/scoreproduct.v1.ScoreProduct/CreateScoreProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreProductClient) DeleteScoreProduct(ctx context.Context, in *ScoreProductDeleteReq, opts ...grpc.CallOption) (*ScoreProductDeleteReply, error) {
	out := new(ScoreProductDeleteReply)
	err := c.cc.Invoke(ctx, "/scoreproduct.v1.ScoreProduct/DeleteScoreProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreProductClient) BatchDeleteScoreProduct(ctx context.Context, in *ScoreProductBatchDeleteReq, opts ...grpc.CallOption) (*ScoreProductDeleteReply, error) {
	out := new(ScoreProductDeleteReply)
	err := c.cc.Invoke(ctx, "/scoreproduct.v1.ScoreProduct/BatchDeleteScoreProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScoreProductServer is the server API for ScoreProduct service.
// All implementations must embed UnimplementedScoreProductServer
// for forward compatibility
type ScoreProductServer interface {
	// 分页查询ScoreProduct
	GetScoreProductPage(context.Context, *ScoreProductPageReq) (*ScoreProductPageReply, error)
	// 获取ScoreProduct
	GetScoreProduct(context.Context, *ScoreProductReq) (*ScoreProductReply, error)
	// 更新ScoreProduct
	UpdateScoreProduct(context.Context, *ScoreProductUpdateReq) (*ScoreProductUpdateReply, error)
	// 创建ScoreProduct
	CreateScoreProduct(context.Context, *ScoreProductCreateReq) (*ScoreProductCreateReply, error)
	// 删除ScoreProduct
	DeleteScoreProduct(context.Context, *ScoreProductDeleteReq) (*ScoreProductDeleteReply, error)
	// 批量删除ScoreProduct
	BatchDeleteScoreProduct(context.Context, *ScoreProductBatchDeleteReq) (*ScoreProductDeleteReply, error)
	mustEmbedUnimplementedScoreProductServer()
}

// UnimplementedScoreProductServer must be embedded to have forward compatible implementations.
type UnimplementedScoreProductServer struct {
}

func (UnimplementedScoreProductServer) GetScoreProductPage(context.Context, *ScoreProductPageReq) (*ScoreProductPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScoreProductPage not implemented")
}
func (UnimplementedScoreProductServer) GetScoreProduct(context.Context, *ScoreProductReq) (*ScoreProductReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScoreProduct not implemented")
}
func (UnimplementedScoreProductServer) UpdateScoreProduct(context.Context, *ScoreProductUpdateReq) (*ScoreProductUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScoreProduct not implemented")
}
func (UnimplementedScoreProductServer) CreateScoreProduct(context.Context, *ScoreProductCreateReq) (*ScoreProductCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateScoreProduct not implemented")
}
func (UnimplementedScoreProductServer) DeleteScoreProduct(context.Context, *ScoreProductDeleteReq) (*ScoreProductDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScoreProduct not implemented")
}
func (UnimplementedScoreProductServer) BatchDeleteScoreProduct(context.Context, *ScoreProductBatchDeleteReq) (*ScoreProductDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteScoreProduct not implemented")
}
func (UnimplementedScoreProductServer) mustEmbedUnimplementedScoreProductServer() {}

// UnsafeScoreProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScoreProductServer will
// result in compilation errors.
type UnsafeScoreProductServer interface {
	mustEmbedUnimplementedScoreProductServer()
}

func RegisterScoreProductServer(s grpc.ServiceRegistrar, srv ScoreProductServer) {
	s.RegisterService(&ScoreProduct_ServiceDesc, srv)
}

func _ScoreProduct_GetScoreProductPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScoreProductPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreProductServer).GetScoreProductPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoreproduct.v1.ScoreProduct/GetScoreProductPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreProductServer).GetScoreProductPage(ctx, req.(*ScoreProductPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoreProduct_GetScoreProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScoreProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreProductServer).GetScoreProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoreproduct.v1.ScoreProduct/GetScoreProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreProductServer).GetScoreProduct(ctx, req.(*ScoreProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoreProduct_UpdateScoreProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScoreProductUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreProductServer).UpdateScoreProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoreproduct.v1.ScoreProduct/UpdateScoreProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreProductServer).UpdateScoreProduct(ctx, req.(*ScoreProductUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoreProduct_CreateScoreProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScoreProductCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreProductServer).CreateScoreProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoreproduct.v1.ScoreProduct/CreateScoreProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreProductServer).CreateScoreProduct(ctx, req.(*ScoreProductCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoreProduct_DeleteScoreProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScoreProductDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreProductServer).DeleteScoreProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoreproduct.v1.ScoreProduct/DeleteScoreProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreProductServer).DeleteScoreProduct(ctx, req.(*ScoreProductDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoreProduct_BatchDeleteScoreProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScoreProductBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreProductServer).BatchDeleteScoreProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoreproduct.v1.ScoreProduct/BatchDeleteScoreProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreProductServer).BatchDeleteScoreProduct(ctx, req.(*ScoreProductBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ScoreProduct_ServiceDesc is the grpc.ServiceDesc for ScoreProduct service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScoreProduct_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scoreproduct.v1.ScoreProduct",
	HandlerType: (*ScoreProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetScoreProductPage",
			Handler:    _ScoreProduct_GetScoreProductPage_Handler,
		},
		{
			MethodName: "GetScoreProduct",
			Handler:    _ScoreProduct_GetScoreProduct_Handler,
		},
		{
			MethodName: "UpdateScoreProduct",
			Handler:    _ScoreProduct_UpdateScoreProduct_Handler,
		},
		{
			MethodName: "CreateScoreProduct",
			Handler:    _ScoreProduct_CreateScoreProduct_Handler,
		},
		{
			MethodName: "DeleteScoreProduct",
			Handler:    _ScoreProduct_DeleteScoreProduct_Handler,
		},
		{
			MethodName: "BatchDeleteScoreProduct",
			Handler:    _ScoreProduct_BatchDeleteScoreProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/param/scoreproduct/v1/score_product.proto",
}
