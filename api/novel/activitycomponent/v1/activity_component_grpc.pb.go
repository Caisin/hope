// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/novel/activitycomponent/v1/activity_component.proto

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

// ActivityComponentClient is the client API for ActivityComponent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityComponentClient interface {
	// 分页查询ActivityComponent
	GetActivityComponentPage(ctx context.Context, in *ActivityComponentPageReq, opts ...grpc.CallOption) (*ActivityComponentPageReply, error)
	// 获取ActivityComponent
	GetActivityComponent(ctx context.Context, in *ActivityComponentReq, opts ...grpc.CallOption) (*ActivityComponentReply, error)
	// 更新ActivityComponent
	UpdateActivityComponent(ctx context.Context, in *ActivityComponentUpdateReq, opts ...grpc.CallOption) (*ActivityComponentUpdateReply, error)
	// 创建ActivityComponent
	CreateActivityComponent(ctx context.Context, in *ActivityComponentCreateReq, opts ...grpc.CallOption) (*ActivityComponentCreateReply, error)
	// 删除ActivityComponent
	DeleteActivityComponent(ctx context.Context, in *ActivityComponentDeleteReq, opts ...grpc.CallOption) (*ActivityComponentDeleteReply, error)
	// 批量删除ActivityComponent
	BatchDeleteActivityComponent(ctx context.Context, in *ActivityComponentBatchDeleteReq, opts ...grpc.CallOption) (*ActivityComponentDeleteReply, error)
}

type activityComponentClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityComponentClient(cc grpc.ClientConnInterface) ActivityComponentClient {
	return &activityComponentClient{cc}
}

func (c *activityComponentClient) GetActivityComponentPage(ctx context.Context, in *ActivityComponentPageReq, opts ...grpc.CallOption) (*ActivityComponentPageReply, error) {
	out := new(ActivityComponentPageReply)
	err := c.cc.Invoke(ctx, "/activitycomponent.v1.ActivityComponent/GetActivityComponentPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityComponentClient) GetActivityComponent(ctx context.Context, in *ActivityComponentReq, opts ...grpc.CallOption) (*ActivityComponentReply, error) {
	out := new(ActivityComponentReply)
	err := c.cc.Invoke(ctx, "/activitycomponent.v1.ActivityComponent/GetActivityComponent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityComponentClient) UpdateActivityComponent(ctx context.Context, in *ActivityComponentUpdateReq, opts ...grpc.CallOption) (*ActivityComponentUpdateReply, error) {
	out := new(ActivityComponentUpdateReply)
	err := c.cc.Invoke(ctx, "/activitycomponent.v1.ActivityComponent/UpdateActivityComponent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityComponentClient) CreateActivityComponent(ctx context.Context, in *ActivityComponentCreateReq, opts ...grpc.CallOption) (*ActivityComponentCreateReply, error) {
	out := new(ActivityComponentCreateReply)
	err := c.cc.Invoke(ctx, "/activitycomponent.v1.ActivityComponent/CreateActivityComponent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityComponentClient) DeleteActivityComponent(ctx context.Context, in *ActivityComponentDeleteReq, opts ...grpc.CallOption) (*ActivityComponentDeleteReply, error) {
	out := new(ActivityComponentDeleteReply)
	err := c.cc.Invoke(ctx, "/activitycomponent.v1.ActivityComponent/DeleteActivityComponent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityComponentClient) BatchDeleteActivityComponent(ctx context.Context, in *ActivityComponentBatchDeleteReq, opts ...grpc.CallOption) (*ActivityComponentDeleteReply, error) {
	out := new(ActivityComponentDeleteReply)
	err := c.cc.Invoke(ctx, "/activitycomponent.v1.ActivityComponent/BatchDeleteActivityComponent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityComponentServer is the server API for ActivityComponent service.
// All implementations must embed UnimplementedActivityComponentServer
// for forward compatibility
type ActivityComponentServer interface {
	// 分页查询ActivityComponent
	GetActivityComponentPage(context.Context, *ActivityComponentPageReq) (*ActivityComponentPageReply, error)
	// 获取ActivityComponent
	GetActivityComponent(context.Context, *ActivityComponentReq) (*ActivityComponentReply, error)
	// 更新ActivityComponent
	UpdateActivityComponent(context.Context, *ActivityComponentUpdateReq) (*ActivityComponentUpdateReply, error)
	// 创建ActivityComponent
	CreateActivityComponent(context.Context, *ActivityComponentCreateReq) (*ActivityComponentCreateReply, error)
	// 删除ActivityComponent
	DeleteActivityComponent(context.Context, *ActivityComponentDeleteReq) (*ActivityComponentDeleteReply, error)
	// 批量删除ActivityComponent
	BatchDeleteActivityComponent(context.Context, *ActivityComponentBatchDeleteReq) (*ActivityComponentDeleteReply, error)
	mustEmbedUnimplementedActivityComponentServer()
}

// UnimplementedActivityComponentServer must be embedded to have forward compatible implementations.
type UnimplementedActivityComponentServer struct {
}

func (UnimplementedActivityComponentServer) GetActivityComponentPage(context.Context, *ActivityComponentPageReq) (*ActivityComponentPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivityComponentPage not implemented")
}
func (UnimplementedActivityComponentServer) GetActivityComponent(context.Context, *ActivityComponentReq) (*ActivityComponentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivityComponent not implemented")
}
func (UnimplementedActivityComponentServer) UpdateActivityComponent(context.Context, *ActivityComponentUpdateReq) (*ActivityComponentUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateActivityComponent not implemented")
}
func (UnimplementedActivityComponentServer) CreateActivityComponent(context.Context, *ActivityComponentCreateReq) (*ActivityComponentCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateActivityComponent not implemented")
}
func (UnimplementedActivityComponentServer) DeleteActivityComponent(context.Context, *ActivityComponentDeleteReq) (*ActivityComponentDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteActivityComponent not implemented")
}
func (UnimplementedActivityComponentServer) BatchDeleteActivityComponent(context.Context, *ActivityComponentBatchDeleteReq) (*ActivityComponentDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteActivityComponent not implemented")
}
func (UnimplementedActivityComponentServer) mustEmbedUnimplementedActivityComponentServer() {}

// UnsafeActivityComponentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityComponentServer will
// result in compilation errors.
type UnsafeActivityComponentServer interface {
	mustEmbedUnimplementedActivityComponentServer()
}

func RegisterActivityComponentServer(s grpc.ServiceRegistrar, srv ActivityComponentServer) {
	s.RegisterService(&ActivityComponent_ServiceDesc, srv)
}

func _ActivityComponent_GetActivityComponentPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityComponentPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityComponentServer).GetActivityComponentPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitycomponent.v1.ActivityComponent/GetActivityComponentPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityComponentServer).GetActivityComponentPage(ctx, req.(*ActivityComponentPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityComponent_GetActivityComponent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityComponentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityComponentServer).GetActivityComponent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitycomponent.v1.ActivityComponent/GetActivityComponent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityComponentServer).GetActivityComponent(ctx, req.(*ActivityComponentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityComponent_UpdateActivityComponent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityComponentUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityComponentServer).UpdateActivityComponent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitycomponent.v1.ActivityComponent/UpdateActivityComponent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityComponentServer).UpdateActivityComponent(ctx, req.(*ActivityComponentUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityComponent_CreateActivityComponent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityComponentCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityComponentServer).CreateActivityComponent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitycomponent.v1.ActivityComponent/CreateActivityComponent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityComponentServer).CreateActivityComponent(ctx, req.(*ActivityComponentCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityComponent_DeleteActivityComponent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityComponentDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityComponentServer).DeleteActivityComponent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitycomponent.v1.ActivityComponent/DeleteActivityComponent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityComponentServer).DeleteActivityComponent(ctx, req.(*ActivityComponentDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityComponent_BatchDeleteActivityComponent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityComponentBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityComponentServer).BatchDeleteActivityComponent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitycomponent.v1.ActivityComponent/BatchDeleteActivityComponent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityComponentServer).BatchDeleteActivityComponent(ctx, req.(*ActivityComponentBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ActivityComponent_ServiceDesc is the grpc.ServiceDesc for ActivityComponent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActivityComponent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "activitycomponent.v1.ActivityComponent",
	HandlerType: (*ActivityComponentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetActivityComponentPage",
			Handler:    _ActivityComponent_GetActivityComponentPage_Handler,
		},
		{
			MethodName: "GetActivityComponent",
			Handler:    _ActivityComponent_GetActivityComponent_Handler,
		},
		{
			MethodName: "UpdateActivityComponent",
			Handler:    _ActivityComponent_UpdateActivityComponent_Handler,
		},
		{
			MethodName: "CreateActivityComponent",
			Handler:    _ActivityComponent_CreateActivityComponent_Handler,
		},
		{
			MethodName: "DeleteActivityComponent",
			Handler:    _ActivityComponent_DeleteActivityComponent_Handler,
		},
		{
			MethodName: "BatchDeleteActivityComponent",
			Handler:    _ActivityComponent_BatchDeleteActivityComponent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/novel/activitycomponent/v1/activity_component.proto",
}
