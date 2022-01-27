// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/novel/activity/v1/activity.proto

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

// ActivityClient is the client API for Activity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityClient interface {
	// 分页查询Activity
	GetActivityPage(ctx context.Context, in *ActivityPageReq, opts ...grpc.CallOption) (*ActivityPageReply, error)
	// 获取Activity
	GetActivity(ctx context.Context, in *ActivityReq, opts ...grpc.CallOption) (*ActivityReply, error)
	// 更新Activity
	UpdateActivity(ctx context.Context, in *ActivityUpdateReq, opts ...grpc.CallOption) (*ActivityUpdateReply, error)
	// 创建Activity
	CreateActivity(ctx context.Context, in *ActivityCreateReq, opts ...grpc.CallOption) (*ActivityCreateReply, error)
	// 删除Activity
	DeleteActivity(ctx context.Context, in *ActivityDeleteReq, opts ...grpc.CallOption) (*ActivityDeleteReply, error)
	// 批量删除Activity
	BatchDeleteActivity(ctx context.Context, in *ActivityBatchDeleteReq, opts ...grpc.CallOption) (*ActivityDeleteReply, error)
}

type activityClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityClient(cc grpc.ClientConnInterface) ActivityClient {
	return &activityClient{cc}
}

func (c *activityClient) GetActivityPage(ctx context.Context, in *ActivityPageReq, opts ...grpc.CallOption) (*ActivityPageReply, error) {
	out := new(ActivityPageReply)
	err := c.cc.Invoke(ctx, "/activity.v1.Activity/GetActivityPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) GetActivity(ctx context.Context, in *ActivityReq, opts ...grpc.CallOption) (*ActivityReply, error) {
	out := new(ActivityReply)
	err := c.cc.Invoke(ctx, "/activity.v1.Activity/GetActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) UpdateActivity(ctx context.Context, in *ActivityUpdateReq, opts ...grpc.CallOption) (*ActivityUpdateReply, error) {
	out := new(ActivityUpdateReply)
	err := c.cc.Invoke(ctx, "/activity.v1.Activity/UpdateActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) CreateActivity(ctx context.Context, in *ActivityCreateReq, opts ...grpc.CallOption) (*ActivityCreateReply, error) {
	out := new(ActivityCreateReply)
	err := c.cc.Invoke(ctx, "/activity.v1.Activity/CreateActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) DeleteActivity(ctx context.Context, in *ActivityDeleteReq, opts ...grpc.CallOption) (*ActivityDeleteReply, error) {
	out := new(ActivityDeleteReply)
	err := c.cc.Invoke(ctx, "/activity.v1.Activity/DeleteActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) BatchDeleteActivity(ctx context.Context, in *ActivityBatchDeleteReq, opts ...grpc.CallOption) (*ActivityDeleteReply, error) {
	out := new(ActivityDeleteReply)
	err := c.cc.Invoke(ctx, "/activity.v1.Activity/BatchDeleteActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServer is the server API for Activity service.
// All implementations must embed UnimplementedActivityServer
// for forward compatibility
type ActivityServer interface {
	// 分页查询Activity
	GetActivityPage(context.Context, *ActivityPageReq) (*ActivityPageReply, error)
	// 获取Activity
	GetActivity(context.Context, *ActivityReq) (*ActivityReply, error)
	// 更新Activity
	UpdateActivity(context.Context, *ActivityUpdateReq) (*ActivityUpdateReply, error)
	// 创建Activity
	CreateActivity(context.Context, *ActivityCreateReq) (*ActivityCreateReply, error)
	// 删除Activity
	DeleteActivity(context.Context, *ActivityDeleteReq) (*ActivityDeleteReply, error)
	// 批量删除Activity
	BatchDeleteActivity(context.Context, *ActivityBatchDeleteReq) (*ActivityDeleteReply, error)
	mustEmbedUnimplementedActivityServer()
}

// UnimplementedActivityServer must be embedded to have forward compatible implementations.
type UnimplementedActivityServer struct {
}

func (UnimplementedActivityServer) GetActivityPage(context.Context, *ActivityPageReq) (*ActivityPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivityPage not implemented")
}
func (UnimplementedActivityServer) GetActivity(context.Context, *ActivityReq) (*ActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivity not implemented")
}
func (UnimplementedActivityServer) UpdateActivity(context.Context, *ActivityUpdateReq) (*ActivityUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateActivity not implemented")
}
func (UnimplementedActivityServer) CreateActivity(context.Context, *ActivityCreateReq) (*ActivityCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateActivity not implemented")
}
func (UnimplementedActivityServer) DeleteActivity(context.Context, *ActivityDeleteReq) (*ActivityDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteActivity not implemented")
}
func (UnimplementedActivityServer) BatchDeleteActivity(context.Context, *ActivityBatchDeleteReq) (*ActivityDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteActivity not implemented")
}
func (UnimplementedActivityServer) mustEmbedUnimplementedActivityServer() {}

// UnsafeActivityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityServer will
// result in compilation errors.
type UnsafeActivityServer interface {
	mustEmbedUnimplementedActivityServer()
}

func RegisterActivityServer(s grpc.ServiceRegistrar, srv ActivityServer) {
	s.RegisterService(&Activity_ServiceDesc, srv)
}

func _Activity_GetActivityPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).GetActivityPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.v1.Activity/GetActivityPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).GetActivityPage(ctx, req.(*ActivityPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_GetActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).GetActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.v1.Activity/GetActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).GetActivity(ctx, req.(*ActivityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_UpdateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).UpdateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.v1.Activity/UpdateActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).UpdateActivity(ctx, req.(*ActivityUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_CreateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).CreateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.v1.Activity/CreateActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).CreateActivity(ctx, req.(*ActivityCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_DeleteActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).DeleteActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.v1.Activity/DeleteActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).DeleteActivity(ctx, req.(*ActivityDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_BatchDeleteActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).BatchDeleteActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.v1.Activity/BatchDeleteActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).BatchDeleteActivity(ctx, req.(*ActivityBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Activity_ServiceDesc is the grpc.ServiceDesc for Activity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Activity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "activity.v1.Activity",
	HandlerType: (*ActivityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetActivityPage",
			Handler:    _Activity_GetActivityPage_Handler,
		},
		{
			MethodName: "GetActivity",
			Handler:    _Activity_GetActivity_Handler,
		},
		{
			MethodName: "UpdateActivity",
			Handler:    _Activity_UpdateActivity_Handler,
		},
		{
			MethodName: "CreateActivity",
			Handler:    _Activity_CreateActivity_Handler,
		},
		{
			MethodName: "DeleteActivity",
			Handler:    _Activity_DeleteActivity_Handler,
		},
		{
			MethodName: "BatchDeleteActivity",
			Handler:    _Activity_BatchDeleteActivity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/novel/activity/v1/activity.proto",
}
