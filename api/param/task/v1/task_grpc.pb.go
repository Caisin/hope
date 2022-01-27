// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/param/task/v1/task.proto

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

// TaskClient is the client API for Task service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskClient interface {
	// 分页查询Task
	GetTaskPage(ctx context.Context, in *TaskPageReq, opts ...grpc.CallOption) (*TaskPageReply, error)
	// 获取Task
	GetTask(ctx context.Context, in *TaskReq, opts ...grpc.CallOption) (*TaskReply, error)
	// 更新Task
	UpdateTask(ctx context.Context, in *TaskUpdateReq, opts ...grpc.CallOption) (*TaskUpdateReply, error)
	// 创建Task
	CreateTask(ctx context.Context, in *TaskCreateReq, opts ...grpc.CallOption) (*TaskCreateReply, error)
	// 删除Task
	DeleteTask(ctx context.Context, in *TaskDeleteReq, opts ...grpc.CallOption) (*TaskDeleteReply, error)
	// 批量删除Task
	BatchDeleteTask(ctx context.Context, in *TaskBatchDeleteReq, opts ...grpc.CallOption) (*TaskDeleteReply, error)
}

type taskClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskClient(cc grpc.ClientConnInterface) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) GetTaskPage(ctx context.Context, in *TaskPageReq, opts ...grpc.CallOption) (*TaskPageReply, error) {
	out := new(TaskPageReply)
	err := c.cc.Invoke(ctx, "/task.v1.Task/GetTaskPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) GetTask(ctx context.Context, in *TaskReq, opts ...grpc.CallOption) (*TaskReply, error) {
	out := new(TaskReply)
	err := c.cc.Invoke(ctx, "/task.v1.Task/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) UpdateTask(ctx context.Context, in *TaskUpdateReq, opts ...grpc.CallOption) (*TaskUpdateReply, error) {
	out := new(TaskUpdateReply)
	err := c.cc.Invoke(ctx, "/task.v1.Task/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) CreateTask(ctx context.Context, in *TaskCreateReq, opts ...grpc.CallOption) (*TaskCreateReply, error) {
	out := new(TaskCreateReply)
	err := c.cc.Invoke(ctx, "/task.v1.Task/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) DeleteTask(ctx context.Context, in *TaskDeleteReq, opts ...grpc.CallOption) (*TaskDeleteReply, error) {
	out := new(TaskDeleteReply)
	err := c.cc.Invoke(ctx, "/task.v1.Task/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) BatchDeleteTask(ctx context.Context, in *TaskBatchDeleteReq, opts ...grpc.CallOption) (*TaskDeleteReply, error) {
	out := new(TaskDeleteReply)
	err := c.cc.Invoke(ctx, "/task.v1.Task/BatchDeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServer is the server API for Task service.
// All implementations must embed UnimplementedTaskServer
// for forward compatibility
type TaskServer interface {
	// 分页查询Task
	GetTaskPage(context.Context, *TaskPageReq) (*TaskPageReply, error)
	// 获取Task
	GetTask(context.Context, *TaskReq) (*TaskReply, error)
	// 更新Task
	UpdateTask(context.Context, *TaskUpdateReq) (*TaskUpdateReply, error)
	// 创建Task
	CreateTask(context.Context, *TaskCreateReq) (*TaskCreateReply, error)
	// 删除Task
	DeleteTask(context.Context, *TaskDeleteReq) (*TaskDeleteReply, error)
	// 批量删除Task
	BatchDeleteTask(context.Context, *TaskBatchDeleteReq) (*TaskDeleteReply, error)
	mustEmbedUnimplementedTaskServer()
}

// UnimplementedTaskServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServer struct {
}

func (UnimplementedTaskServer) GetTaskPage(context.Context, *TaskPageReq) (*TaskPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskPage not implemented")
}
func (UnimplementedTaskServer) GetTask(context.Context, *TaskReq) (*TaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedTaskServer) UpdateTask(context.Context, *TaskUpdateReq) (*TaskUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedTaskServer) CreateTask(context.Context, *TaskCreateReq) (*TaskCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTaskServer) DeleteTask(context.Context, *TaskDeleteReq) (*TaskDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedTaskServer) BatchDeleteTask(context.Context, *TaskBatchDeleteReq) (*TaskDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteTask not implemented")
}
func (UnimplementedTaskServer) mustEmbedUnimplementedTaskServer() {}

// UnsafeTaskServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServer will
// result in compilation errors.
type UnsafeTaskServer interface {
	mustEmbedUnimplementedTaskServer()
}

func RegisterTaskServer(s grpc.ServiceRegistrar, srv TaskServer) {
	s.RegisterService(&Task_ServiceDesc, srv)
}

func _Task_GetTaskPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).GetTaskPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.v1.Task/GetTaskPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).GetTaskPage(ctx, req.(*TaskPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.v1.Task/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).GetTask(ctx, req.(*TaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.v1.Task/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).UpdateTask(ctx, req.(*TaskUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.v1.Task/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).CreateTask(ctx, req.(*TaskCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.v1.Task/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).DeleteTask(ctx, req.(*TaskDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_BatchDeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskBatchDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).BatchDeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.v1.Task/BatchDeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).BatchDeleteTask(ctx, req.(*TaskBatchDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Task_ServiceDesc is the grpc.ServiceDesc for Task service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Task_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.v1.Task",
	HandlerType: (*TaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTaskPage",
			Handler:    _Task_GetTaskPage_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _Task_GetTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _Task_UpdateTask_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _Task_CreateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _Task_DeleteTask_Handler,
		},
		{
			MethodName: "BatchDeleteTask",
			Handler:    _Task_BatchDeleteTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/param/task/v1/task.proto",
}
