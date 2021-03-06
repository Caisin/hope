// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/novel/assetitem/v1/asset_item.proto

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

// AssetItemClient is the client API for AssetItem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetItemClient interface {
	// 分页查询AssetItem
	GetAssetItemPage(ctx context.Context, in *AssetItemPageReq, opts ...grpc.CallOption) (*AssetItemPageReply, error)
	// 获取AssetItem
	GetAssetItem(ctx context.Context, in *AssetItemReq, opts ...grpc.CallOption) (*AssetItemReply, error)
	// 更新AssetItem
	UpdateAssetItem(ctx context.Context, in *AssetItemUpdateReq, opts ...grpc.CallOption) (*AssetItemUpdateReply, error)
	// 创建AssetItem
	CreateAssetItem(ctx context.Context, in *AssetItemCreateReq, opts ...grpc.CallOption) (*AssetItemCreateReply, error)
	// 删除AssetItem
	DeleteAssetItem(ctx context.Context, in *AssetItemDeleteReq, opts ...grpc.CallOption) (*AssetItemDeleteReply, error)
	// 批量删除AssetItem
	BatchDeleteAssetItem(ctx context.Context, in *AssetItemBatchDeleteReq, opts ...grpc.CallOption) (*AssetItemDeleteReply, error)
}

type assetItemClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetItemClient(cc grpc.ClientConnInterface) AssetItemClient {
	return &assetItemClient{cc}
}

func (c *assetItemClient) GetAssetItemPage(ctx context.Context, in *AssetItemPageReq, opts ...grpc.CallOption) (*AssetItemPageReply, error) {
	out := new(AssetItemPageReply)
	err := c.cc.Invoke(ctx, "/assetitem.v1.AssetItem/GetAssetItemPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetItemClient) GetAssetItem(ctx context.Context, in *AssetItemReq, opts ...grpc.CallOption) (*AssetItemReply, error) {
	out := new(AssetItemReply)
	err := c.cc.Invoke(ctx, "/assetitem.v1.AssetItem/GetAssetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetItemClient) UpdateAssetItem(ctx context.Context, in *AssetItemUpdateReq, opts ...grpc.CallOption) (*AssetItemUpdateReply, error) {
	out := new(AssetItemUpdateReply)
	err := c.cc.Invoke(ctx, "/assetitem.v1.AssetItem/UpdateAssetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetItemClient) CreateAssetItem(ctx context.Context, in *AssetItemCreateReq, opts ...grpc.CallOption) (*AssetItemCreateReply, error) {
	out := new(AssetItemCreateReply)
	err := c.cc.Invoke(ctx, "/assetitem.v1.AssetItem/CreateAssetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetItemClient) DeleteAssetItem(ctx context.Context, in *AssetItemDeleteReq, opts ...grpc.CallOption) (*AssetItemDeleteReply, error) {
	out := new(AssetItemDeleteReply)
	err := c.cc.Invoke(ctx, "/assetitem.v1.AssetItem/DeleteAssetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetItemClient) BatchDeleteAssetItem(ctx context.Context, in *AssetItemBatchDeleteReq, opts ...grpc.CallOption) (*AssetItemDeleteReply, error) {
	out := new(AssetItemDeleteReply)
	err := c.cc.Invoke(ctx, "/assetitem.v1.AssetItem/BatchDeleteAssetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetItemServer is the server API for AssetItem service.
// All implementations must embed UnimplementedAssetItemServer
// for forward compatibility
type AssetItemServer interface {
	// 分页查询AssetItem
	GetAssetItemPage(context.Context, *AssetItemPageReq) (*AssetItemPageReply, error)
	// 获取AssetItem
	GetAssetItem(context.Context, *AssetItemReq) (*AssetItemReply, error)
	// 更新AssetItem
	UpdateAssetItem(context.Context, *AssetItemUpdateReq) (*AssetItemUpdateReply, error)
	// 创建AssetItem
	CreateAssetItem(context.Context, *AssetItemCreateReq) (*AssetItemCreateReply, error)
	// 删除AssetItem
	DeleteAssetItem(context.Context, *AssetItemDeleteReq) (*AssetItemDeleteReply, error)
	// 批量删除AssetItem
	BatchDeleteAssetItem(context.Context, *AssetItemBatchDeleteReq) (*AssetItemDeleteReply, error)
	mustEmbedUnimplementedAssetItemServer()
}

// UnimplementedAssetItemServer must be embedded to have forward compatible implementations.
type UnimplementedAssetItemServer struct {
}

func (UnimplementedAssetItemServer) GetAssetItemPage(context.Context, *AssetItemPageReq) (*AssetItemPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssetItemPage not implemented")
}
func (UnimplementedAssetItemServer) GetAssetItem(context.Context, *AssetItemReq) (*AssetItemReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssetItem not implemented")
}
func (UnimplementedAssetItemServer) UpdateAssetItem(context.Context, *AssetItemUpdateReq) (*AssetItemUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAssetItem not implemented")
}
func (UnimplementedAssetItemServer) CreateAssetItem(context.Context, *AssetItemCreateReq) (*AssetItemCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAssetItem not implemented")
}
func (UnimplementedAssetItemServer) DeleteAssetItem(context.Context, *AssetItemDeleteReq) (*AssetItemDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAssetItem not implemented")
}
func (UnimplementedAssetItemServer) BatchDeleteAssetItem(context.Context, *AssetItemBatchDeleteReq) (*AssetItemDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteAssetItem not implemented")
}
func (UnimplementedAssetItemServer) mustEmbedUnimplementedAssetItemServer() {}

// UnsafeAssetItemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetItemServer will
// result in compilation errors.
type UnsafeAssetItemServer interface {
	mustEmbedUnimplementedAssetItemServer()
}

func RegisterAssetItemServer(s grpc.ServiceRegistrar, srv AssetItemServer) {
	s.RegisterService(&AssetItem_ServiceDesc, srv)
}

func _AssetItem_GetAssetItemPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetItemPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetItemServer).GetAssetItemPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetitem.v1.AssetItem/GetAssetItemPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetItemServer).GetAssetItemPage(ctx, req.(*AssetItemPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetItem_GetAssetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetItemServer).GetAssetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetitem.v1.AssetItem/GetAssetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetItemServer).GetAssetItem(ctx, req.(*AssetItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetItem_UpdateAssetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetItemUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetItemServer).UpdateAssetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetitem.v1.AssetItem/UpdateAssetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetItemServer).UpdateAssetItem(ctx, req.(*AssetItemUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetItem_CreateAssetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetItemCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetItemServer).CreateAssetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetitem.v1.AssetItem/CreateAssetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetItemServer).CreateAssetItem(ctx, req.(*AssetItemCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetItem_DeleteAssetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetItemDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetItemServer).DeleteAssetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetitem.v1.AssetItem/DeleteAssetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetItemServer).DeleteAssetItem(ctx, req.(*AssetItemDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetItem_BatchDeleteAssetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetItemBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetItemServer).BatchDeleteAssetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetitem.v1.AssetItem/BatchDeleteAssetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetItemServer).BatchDeleteAssetItem(ctx, req.(*AssetItemBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetItem_ServiceDesc is the grpc.ServiceDesc for AssetItem service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetItem_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "assetitem.v1.AssetItem",
	HandlerType: (*AssetItemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAssetItemPage",
			Handler:    _AssetItem_GetAssetItemPage_Handler,
		},
		{
			MethodName: "GetAssetItem",
			Handler:    _AssetItem_GetAssetItem_Handler,
		},
		{
			MethodName: "UpdateAssetItem",
			Handler:    _AssetItem_UpdateAssetItem_Handler,
		},
		{
			MethodName: "CreateAssetItem",
			Handler:    _AssetItem_CreateAssetItem_Handler,
		},
		{
			MethodName: "DeleteAssetItem",
			Handler:    _AssetItem_DeleteAssetItem_Handler,
		},
		{
			MethodName: "BatchDeleteAssetItem",
			Handler:    _AssetItem_BatchDeleteAssetItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/novel/assetitem/v1/asset_item.proto",
}
