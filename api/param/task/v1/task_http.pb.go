// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.3

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type TaskHTTPServer interface {
	BatchDeleteTask(context.Context, *TaskBatchDeleteReq) (*TaskDeleteReply, error)
	CreateTask(context.Context, *TaskCreateReq) (*TaskCreateReply, error)
	DeleteTask(context.Context, *TaskDeleteReq) (*TaskDeleteReply, error)
	GetTask(context.Context, *TaskReq) (*TaskReply, error)
	GetTaskPage(context.Context, *TaskPageReq) (*TaskPageReply, error)
	UpdateTask(context.Context, *TaskUpdateReq) (*TaskUpdateReply, error)
}

func RegisterTaskHTTPServer(s *http.Server, srv TaskHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/task/page", _Task_GetTaskPage0_HTTP_Handler(srv))
	r.GET("/v1/task/{id}", _Task_GetTask0_HTTP_Handler(srv))
	r.PUT("/v1/task/{id}", _Task_UpdateTask0_HTTP_Handler(srv))
	r.POST("/v1/task", _Task_CreateTask0_HTTP_Handler(srv))
	r.DELETE("/v1/task/{id}", _Task_DeleteTask0_HTTP_Handler(srv))
	r.DELETE("/v1/task", _Task_BatchDeleteTask0_HTTP_Handler(srv))
}

func _Task_GetTaskPage0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TaskPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/task.v1.Task/GetTaskPage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTaskPage(ctx, req.(*TaskPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TaskPageReply)
		return ctx.Result(200, reply)
	}
}

func _Task_GetTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TaskReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/task.v1.Task/GetTask")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTask(ctx, req.(*TaskReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_UpdateTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TaskUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/task.v1.Task/UpdateTask")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTask(ctx, req.(*TaskUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TaskUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _Task_CreateTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TaskCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/task.v1.Task/CreateTask")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTask(ctx, req.(*TaskCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TaskCreateReply)
		return ctx.Result(200, reply)
	}
}

func _Task_DeleteTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TaskDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/task.v1.Task/DeleteTask")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteTask(ctx, req.(*TaskDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TaskDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _Task_BatchDeleteTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TaskBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/task.v1.Task/BatchDeleteTask")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteTask(ctx, req.(*TaskBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TaskDeleteReply)
		return ctx.Result(200, reply)
	}
}

type TaskHTTPClient interface {
	BatchDeleteTask(ctx context.Context, req *TaskBatchDeleteReq, opts ...http.CallOption) (rsp *TaskDeleteReply, err error)
	CreateTask(ctx context.Context, req *TaskCreateReq, opts ...http.CallOption) (rsp *TaskCreateReply, err error)
	DeleteTask(ctx context.Context, req *TaskDeleteReq, opts ...http.CallOption) (rsp *TaskDeleteReply, err error)
	GetTask(ctx context.Context, req *TaskReq, opts ...http.CallOption) (rsp *TaskReply, err error)
	GetTaskPage(ctx context.Context, req *TaskPageReq, opts ...http.CallOption) (rsp *TaskPageReply, err error)
	UpdateTask(ctx context.Context, req *TaskUpdateReq, opts ...http.CallOption) (rsp *TaskUpdateReply, err error)
}

type TaskHTTPClientImpl struct {
	cc *http.Client
}

func NewTaskHTTPClient(client *http.Client) TaskHTTPClient {
	return &TaskHTTPClientImpl{client}
}

func (c *TaskHTTPClientImpl) BatchDeleteTask(ctx context.Context, in *TaskBatchDeleteReq, opts ...http.CallOption) (*TaskDeleteReply, error) {
	var out TaskDeleteReply
	pattern := "/v1/task"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/task.v1.Task/BatchDeleteTask"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) CreateTask(ctx context.Context, in *TaskCreateReq, opts ...http.CallOption) (*TaskCreateReply, error) {
	var out TaskCreateReply
	pattern := "/v1/task"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/task.v1.Task/CreateTask"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) DeleteTask(ctx context.Context, in *TaskDeleteReq, opts ...http.CallOption) (*TaskDeleteReply, error) {
	var out TaskDeleteReply
	pattern := "/v1/task/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/task.v1.Task/DeleteTask"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) GetTask(ctx context.Context, in *TaskReq, opts ...http.CallOption) (*TaskReply, error) {
	var out TaskReply
	pattern := "/v1/task/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/task.v1.Task/GetTask"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) GetTaskPage(ctx context.Context, in *TaskPageReq, opts ...http.CallOption) (*TaskPageReply, error) {
	var out TaskPageReply
	pattern := "/v1/task/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/task.v1.Task/GetTaskPage"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) UpdateTask(ctx context.Context, in *TaskUpdateReq, opts ...http.CallOption) (*TaskUpdateReply, error) {
	var out TaskUpdateReply
	pattern := "/v1/task/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/task.v1.Task/UpdateTask"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
