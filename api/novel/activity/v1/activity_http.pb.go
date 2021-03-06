// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.3

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	result "hope/pkg/result"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type ActivityHTTPServer interface {
	BatchDeleteActivity(context.Context, *ActivityBatchDeleteReq) (*result.Reply, error)
	CreateActivity(context.Context, *ActivityCreateReq) (*result.Reply, error)
	DeleteActivity(context.Context, *ActivityDeleteReq) (*result.Reply, error)
	GetActivity(context.Context, *ActivityReq) (*result.Reply, error)
	GetActivityPage(context.Context, *ActivityPageReq) (*result.Reply, error)
	UpdateActivity(context.Context, *ActivityUpdateReq) (*result.Reply, error)
}

func RegisterActivityHTTPServer(s *http.Server, srv ActivityHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/activity/page", _Activity_GetActivityPage0_HTTP_Handler(srv))
	r.GET("/v1/activity/{id}", _Activity_GetActivity0_HTTP_Handler(srv))
	r.PUT("/v1/activity/{id}", _Activity_UpdateActivity0_HTTP_Handler(srv))
	r.POST("/v1/activity", _Activity_CreateActivity0_HTTP_Handler(srv))
	r.DELETE("/v1/activity/{id}", _Activity_DeleteActivity0_HTTP_Handler(srv))
	r.DELETE("/v1/activity", _Activity_BatchDeleteActivity0_HTTP_Handler(srv))
}

func _Activity_GetActivityPage0_HTTP_Handler(srv ActivityHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ActivityPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/activity.v1.Activity/GetActivityPage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetActivityPage(ctx, req.(*ActivityPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*result.Reply)
		return ctx.Result(200, reply)
	}
}

func _Activity_GetActivity0_HTTP_Handler(srv ActivityHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ActivityReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/activity.v1.Activity/GetActivity")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetActivity(ctx, req.(*ActivityReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*result.Reply)
		return ctx.Result(200, reply)
	}
}

func _Activity_UpdateActivity0_HTTP_Handler(srv ActivityHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ActivityUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/activity.v1.Activity/UpdateActivity")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateActivity(ctx, req.(*ActivityUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*result.Reply)
		return ctx.Result(200, reply)
	}
}

func _Activity_CreateActivity0_HTTP_Handler(srv ActivityHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ActivityCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/activity.v1.Activity/CreateActivity")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateActivity(ctx, req.(*ActivityCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*result.Reply)
		return ctx.Result(200, reply)
	}
}

func _Activity_DeleteActivity0_HTTP_Handler(srv ActivityHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ActivityDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/activity.v1.Activity/DeleteActivity")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteActivity(ctx, req.(*ActivityDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*result.Reply)
		return ctx.Result(200, reply)
	}
}

func _Activity_BatchDeleteActivity0_HTTP_Handler(srv ActivityHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ActivityBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/activity.v1.Activity/BatchDeleteActivity")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteActivity(ctx, req.(*ActivityBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*result.Reply)
		return ctx.Result(200, reply)
	}
}

type ActivityHTTPClient interface {
	BatchDeleteActivity(ctx context.Context, req *ActivityBatchDeleteReq, opts ...http.CallOption) (rsp *result.Reply, err error)
	CreateActivity(ctx context.Context, req *ActivityCreateReq, opts ...http.CallOption) (rsp *result.Reply, err error)
	DeleteActivity(ctx context.Context, req *ActivityDeleteReq, opts ...http.CallOption) (rsp *result.Reply, err error)
	GetActivity(ctx context.Context, req *ActivityReq, opts ...http.CallOption) (rsp *result.Reply, err error)
	GetActivityPage(ctx context.Context, req *ActivityPageReq, opts ...http.CallOption) (rsp *result.Reply, err error)
	UpdateActivity(ctx context.Context, req *ActivityUpdateReq, opts ...http.CallOption) (rsp *result.Reply, err error)
}

type ActivityHTTPClientImpl struct {
	cc *http.Client
}

func NewActivityHTTPClient(client *http.Client) ActivityHTTPClient {
	return &ActivityHTTPClientImpl{client}
}

func (c *ActivityHTTPClientImpl) BatchDeleteActivity(ctx context.Context, in *ActivityBatchDeleteReq, opts ...http.CallOption) (*result.Reply, error) {
	var out result.Reply
	pattern := "/v1/activity"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/activity.v1.Activity/BatchDeleteActivity"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ActivityHTTPClientImpl) CreateActivity(ctx context.Context, in *ActivityCreateReq, opts ...http.CallOption) (*result.Reply, error) {
	var out result.Reply
	pattern := "/v1/activity"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/activity.v1.Activity/CreateActivity"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ActivityHTTPClientImpl) DeleteActivity(ctx context.Context, in *ActivityDeleteReq, opts ...http.CallOption) (*result.Reply, error) {
	var out result.Reply
	pattern := "/v1/activity/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/activity.v1.Activity/DeleteActivity"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ActivityHTTPClientImpl) GetActivity(ctx context.Context, in *ActivityReq, opts ...http.CallOption) (*result.Reply, error) {
	var out result.Reply
	pattern := "/v1/activity/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/activity.v1.Activity/GetActivity"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ActivityHTTPClientImpl) GetActivityPage(ctx context.Context, in *ActivityPageReq, opts ...http.CallOption) (*result.Reply, error) {
	var out result.Reply
	pattern := "/v1/activity/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/activity.v1.Activity/GetActivityPage"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ActivityHTTPClientImpl) UpdateActivity(ctx context.Context, in *ActivityUpdateReq, opts ...http.CallOption) (*result.Reply, error) {
	var out result.Reply
	pattern := "/v1/activity/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/activity.v1.Activity/UpdateActivity"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
