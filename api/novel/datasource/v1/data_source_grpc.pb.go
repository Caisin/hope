// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/novel/datasource/v1/data_source.proto

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

// DataSourceClient is the client API for DataSource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataSourceClient interface {
	// 分页查询DataSource
	GetDataSourcePage(ctx context.Context, in *DataSourcePageReq, opts ...grpc.CallOption) (*DataSourcePageReply, error)
	// 获取DataSource
	GetDataSource(ctx context.Context, in *DataSourceReq, opts ...grpc.CallOption) (*DataSourceReply, error)
	// 更新DataSource
	UpdateDataSource(ctx context.Context, in *DataSourceUpdateReq, opts ...grpc.CallOption) (*DataSourceUpdateReply, error)
	// 创建DataSource
	CreateDataSource(ctx context.Context, in *DataSourceCreateReq, opts ...grpc.CallOption) (*DataSourceCreateReply, error)
	// 删除DataSource
	DeleteDataSource(ctx context.Context, in *DataSourceDeleteReq, opts ...grpc.CallOption) (*DataSourceDeleteReply, error)
	// 批量删除DataSource
	BatchDeleteDataSource(ctx context.Context, in *DataSourceBatchDeleteReq, opts ...grpc.CallOption) (*DataSourceDeleteReply, error)
}

type dataSourceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataSourceClient(cc grpc.ClientConnInterface) DataSourceClient {
	return &dataSourceClient{cc}
}

func (c *dataSourceClient) GetDataSourcePage(ctx context.Context, in *DataSourcePageReq, opts ...grpc.CallOption) (*DataSourcePageReply, error) {
	out := new(DataSourcePageReply)
	err := c.cc.Invoke(ctx, "/datasource.v1.DataSource/GetDataSourcePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourceClient) GetDataSource(ctx context.Context, in *DataSourceReq, opts ...grpc.CallOption) (*DataSourceReply, error) {
	out := new(DataSourceReply)
	err := c.cc.Invoke(ctx, "/datasource.v1.DataSource/GetDataSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourceClient) UpdateDataSource(ctx context.Context, in *DataSourceUpdateReq, opts ...grpc.CallOption) (*DataSourceUpdateReply, error) {
	out := new(DataSourceUpdateReply)
	err := c.cc.Invoke(ctx, "/datasource.v1.DataSource/UpdateDataSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourceClient) CreateDataSource(ctx context.Context, in *DataSourceCreateReq, opts ...grpc.CallOption) (*DataSourceCreateReply, error) {
	out := new(DataSourceCreateReply)
	err := c.cc.Invoke(ctx, "/datasource.v1.DataSource/CreateDataSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourceClient) DeleteDataSource(ctx context.Context, in *DataSourceDeleteReq, opts ...grpc.CallOption) (*DataSourceDeleteReply, error) {
	out := new(DataSourceDeleteReply)
	err := c.cc.Invoke(ctx, "/datasource.v1.DataSource/DeleteDataSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourceClient) BatchDeleteDataSource(ctx context.Context, in *DataSourceBatchDeleteReq, opts ...grpc.CallOption) (*DataSourceDeleteReply, error) {
	out := new(DataSourceDeleteReply)
	err := c.cc.Invoke(ctx, "/datasource.v1.DataSource/BatchDeleteDataSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataSourceServer is the server API for DataSource service.
// All implementations must embed UnimplementedDataSourceServer
// for forward compatibility
type DataSourceServer interface {
	// 分页查询DataSource
	GetDataSourcePage(context.Context, *DataSourcePageReq) (*DataSourcePageReply, error)
	// 获取DataSource
	GetDataSource(context.Context, *DataSourceReq) (*DataSourceReply, error)
	// 更新DataSource
	UpdateDataSource(context.Context, *DataSourceUpdateReq) (*DataSourceUpdateReply, error)
	// 创建DataSource
	CreateDataSource(context.Context, *DataSourceCreateReq) (*DataSourceCreateReply, error)
	// 删除DataSource
	DeleteDataSource(context.Context, *DataSourceDeleteReq) (*DataSourceDeleteReply, error)
	// 批量删除DataSource
	BatchDeleteDataSource(context.Context, *DataSourceBatchDeleteReq) (*DataSourceDeleteReply, error)
	mustEmbedUnimplementedDataSourceServer()
}

// UnimplementedDataSourceServer must be embedded to have forward compatible implementations.
type UnimplementedDataSourceServer struct {
}

func (UnimplementedDataSourceServer) GetDataSourcePage(context.Context, *DataSourcePageReq) (*DataSourcePageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataSourcePage not implemented")
}
func (UnimplementedDataSourceServer) GetDataSource(context.Context, *DataSourceReq) (*DataSourceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataSource not implemented")
}
func (UnimplementedDataSourceServer) UpdateDataSource(context.Context, *DataSourceUpdateReq) (*DataSourceUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDataSource not implemented")
}
func (UnimplementedDataSourceServer) CreateDataSource(context.Context, *DataSourceCreateReq) (*DataSourceCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDataSource not implemented")
}
func (UnimplementedDataSourceServer) DeleteDataSource(context.Context, *DataSourceDeleteReq) (*DataSourceDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDataSource not implemented")
}
func (UnimplementedDataSourceServer) BatchDeleteDataSource(context.Context, *DataSourceBatchDeleteReq) (*DataSourceDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteDataSource not implemented")
}
func (UnimplementedDataSourceServer) mustEmbedUnimplementedDataSourceServer() {}

// UnsafeDataSourceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataSourceServer will
// result in compilation errors.
type UnsafeDataSourceServer interface {
	mustEmbedUnimplementedDataSourceServer()
}

func RegisterDataSourceServer(s grpc.ServiceRegistrar, srv DataSourceServer) {
	s.RegisterService(&DataSource_ServiceDesc, srv)
}

func _DataSource_GetDataSourcePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSourcePageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).GetDataSourcePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datasource.v1.DataSource/GetDataSourcePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).GetDataSourcePage(ctx, req.(*DataSourcePageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSource_GetDataSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSourceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).GetDataSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datasource.v1.DataSource/GetDataSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).GetDataSource(ctx, req.(*DataSourceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSource_UpdateDataSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSourceUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).UpdateDataSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datasource.v1.DataSource/UpdateDataSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).UpdateDataSource(ctx, req.(*DataSourceUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSource_CreateDataSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSourceCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).CreateDataSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datasource.v1.DataSource/CreateDataSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).CreateDataSource(ctx, req.(*DataSourceCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSource_DeleteDataSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSourceDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).DeleteDataSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datasource.v1.DataSource/DeleteDataSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).DeleteDataSource(ctx, req.(*DataSourceDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSource_BatchDeleteDataSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSourceBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServer).BatchDeleteDataSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datasource.v1.DataSource/BatchDeleteDataSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServer).BatchDeleteDataSource(ctx, req.(*DataSourceBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DataSource_ServiceDesc is the grpc.ServiceDesc for DataSource service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataSource_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "datasource.v1.DataSource",
	HandlerType: (*DataSourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDataSourcePage",
			Handler:    _DataSource_GetDataSourcePage_Handler,
		},
		{
			MethodName: "GetDataSource",
			Handler:    _DataSource_GetDataSource_Handler,
		},
		{
			MethodName: "UpdateDataSource",
			Handler:    _DataSource_UpdateDataSource_Handler,
		},
		{
			MethodName: "CreateDataSource",
			Handler:    _DataSource_CreateDataSource_Handler,
		},
		{
			MethodName: "DeleteDataSource",
			Handler:    _DataSource_DeleteDataSource_Handler,
		},
		{
			MethodName: "BatchDeleteDataSource",
			Handler:    _DataSource_BatchDeleteDataSource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/novel/datasource/v1/data_source.proto",
}
