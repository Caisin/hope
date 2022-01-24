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

type SysJobHTTPServer interface {
	BatchDeleteSysJob(context.Context, *SysJobBatchDeleteReq) (*SysJobDeleteReply, error)
	CreateSysJob(context.Context, *SysJobCreateReq) (*SysJobCreateReply, error)
	DeleteSysJob(context.Context, *SysJobDeleteReq) (*SysJobDeleteReply, error)
	GetPageSysJob(context.Context, *SysJobPageReq) (*SysJobPageReply, error)
	GetSysJob(context.Context, *SysJobReq) (*SysJobReply, error)
	UpdateSysJob(context.Context, *SysJobUpdateReq) (*SysJobUpdateReply, error)
}

func RegisterSysJobHTTPServer(s *http.Server, srv SysJobHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/sys/job/page", _SysJob_GetPageSysJob0_HTTP_Handler(srv))
	r.GET("/v1/sys/job/{id}", _SysJob_GetSysJob0_HTTP_Handler(srv))
	r.PUT("/v1/sys/job/{id}", _SysJob_UpdateSysJob0_HTTP_Handler(srv))
	r.POST("/v1/sys/job", _SysJob_CreateSysJob0_HTTP_Handler(srv))
	r.DELETE("/v1/sys/job/{id}", _SysJob_DeleteSysJob0_HTTP_Handler(srv))
	r.DELETE("/v1/sys/job", _SysJob_BatchDeleteSysJob0_HTTP_Handler(srv))
}

func _SysJob_GetPageSysJob0_HTTP_Handler(srv SysJobHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysJobPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysjob.v1.SysJob/GetPageSysJob")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPageSysJob(ctx, req.(*SysJobPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysJobPageReply)
		return ctx.Result(200, reply)
	}
}

func _SysJob_GetSysJob0_HTTP_Handler(srv SysJobHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysJobReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysjob.v1.SysJob/GetSysJob")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSysJob(ctx, req.(*SysJobReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysJobReply)
		return ctx.Result(200, reply)
	}
}

func _SysJob_UpdateSysJob0_HTTP_Handler(srv SysJobHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysJobUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysjob.v1.SysJob/UpdateSysJob")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateSysJob(ctx, req.(*SysJobUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysJobUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _SysJob_CreateSysJob0_HTTP_Handler(srv SysJobHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysJobCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysjob.v1.SysJob/CreateSysJob")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSysJob(ctx, req.(*SysJobCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysJobCreateReply)
		return ctx.Result(200, reply)
	}
}

func _SysJob_DeleteSysJob0_HTTP_Handler(srv SysJobHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysJobDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysjob.v1.SysJob/DeleteSysJob")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteSysJob(ctx, req.(*SysJobDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysJobDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _SysJob_BatchDeleteSysJob0_HTTP_Handler(srv SysJobHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysJobBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysjob.v1.SysJob/BatchDeleteSysJob")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteSysJob(ctx, req.(*SysJobBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysJobDeleteReply)
		return ctx.Result(200, reply)
	}
}

type SysJobHTTPClient interface {
	BatchDeleteSysJob(ctx context.Context, req *SysJobBatchDeleteReq, opts ...http.CallOption) (rsp *SysJobDeleteReply, err error)
	CreateSysJob(ctx context.Context, req *SysJobCreateReq, opts ...http.CallOption) (rsp *SysJobCreateReply, err error)
	DeleteSysJob(ctx context.Context, req *SysJobDeleteReq, opts ...http.CallOption) (rsp *SysJobDeleteReply, err error)
	GetPageSysJob(ctx context.Context, req *SysJobPageReq, opts ...http.CallOption) (rsp *SysJobPageReply, err error)
	GetSysJob(ctx context.Context, req *SysJobReq, opts ...http.CallOption) (rsp *SysJobReply, err error)
	UpdateSysJob(ctx context.Context, req *SysJobUpdateReq, opts ...http.CallOption) (rsp *SysJobUpdateReply, err error)
}

type SysJobHTTPClientImpl struct {
	cc *http.Client
}

func NewSysJobHTTPClient(client *http.Client) SysJobHTTPClient {
	return &SysJobHTTPClientImpl{client}
}

func (c *SysJobHTTPClientImpl) BatchDeleteSysJob(ctx context.Context, in *SysJobBatchDeleteReq, opts ...http.CallOption) (*SysJobDeleteReply, error) {
	var out SysJobDeleteReply
	pattern := "/v1/sys/job"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/sysjob.v1.SysJob/BatchDeleteSysJob"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysJobHTTPClientImpl) CreateSysJob(ctx context.Context, in *SysJobCreateReq, opts ...http.CallOption) (*SysJobCreateReply, error) {
	var out SysJobCreateReply
	pattern := "/v1/sys/job"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/sysjob.v1.SysJob/CreateSysJob"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysJobHTTPClientImpl) DeleteSysJob(ctx context.Context, in *SysJobDeleteReq, opts ...http.CallOption) (*SysJobDeleteReply, error) {
	var out SysJobDeleteReply
	pattern := "/v1/sys/job/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/sysjob.v1.SysJob/DeleteSysJob"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysJobHTTPClientImpl) GetPageSysJob(ctx context.Context, in *SysJobPageReq, opts ...http.CallOption) (*SysJobPageReply, error) {
	var out SysJobPageReply
	pattern := "/v1/sys/job/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/sysjob.v1.SysJob/GetPageSysJob"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysJobHTTPClientImpl) GetSysJob(ctx context.Context, in *SysJobReq, opts ...http.CallOption) (*SysJobReply, error) {
	var out SysJobReply
	pattern := "/v1/sys/job/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/sysjob.v1.SysJob/GetSysJob"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysJobHTTPClientImpl) UpdateSysJob(ctx context.Context, in *SysJobUpdateReq, opts ...http.CallOption) (*SysJobUpdateReply, error) {
	var out SysJobUpdateReply
	pattern := "/v1/sys/job/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/sysjob.v1.SysJob/UpdateSysJob"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}