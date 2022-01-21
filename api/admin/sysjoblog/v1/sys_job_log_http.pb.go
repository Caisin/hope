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

type SysJobLogHTTPServer interface {
	BatchDeleteSysJobLog(context.Context, *SysJobLogBatchDeleteReq) (*SysJobLogDeleteReply, error)
	CreateSysJobLog(context.Context, *SysJobLogCreateReq) (*SysJobLogCreateReply, error)
	DeleteSysJobLog(context.Context, *SysJobLogDeleteReq) (*SysJobLogDeleteReply, error)
	GetPageSysJobLog(context.Context, *SysJobLogPageReq) (*SysJobLogPageReply, error)
	GetSysJobLog(context.Context, *SysJobLogReq) (*SysJobLogReply, error)
	UpdateSysJobLog(context.Context, *SysJobLogUpdateReq) (*SysJobLogUpdateReply, error)
}

func RegisterSysJobLogHTTPServer(s *http.Server, srv SysJobLogHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/sysjoblog/page", _SysJobLog_GetPageSysJobLog0_HTTP_Handler(srv))
	r.GET("/v1/sysjoblog/{id}", _SysJobLog_GetSysJobLog0_HTTP_Handler(srv))
	r.PUT("/v1/sysjoblog/{id}", _SysJobLog_UpdateSysJobLog0_HTTP_Handler(srv))
	r.POST("/v1/sysjoblog", _SysJobLog_CreateSysJobLog0_HTTP_Handler(srv))
	r.DELETE("/v1/sysjoblog/{id}", _SysJobLog_DeleteSysJobLog0_HTTP_Handler(srv))
	r.DELETE("/v1/sysjoblog", _SysJobLog_BatchDeleteSysJobLog0_HTTP_Handler(srv))
}

func _SysJobLog_GetPageSysJobLog0_HTTP_Handler(srv SysJobLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysJobLogPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysjoblog.v1.SysJobLog/GetPageSysJobLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPageSysJobLog(ctx, req.(*SysJobLogPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysJobLogPageReply)
		return ctx.Result(200, reply)
	}
}

func _SysJobLog_GetSysJobLog0_HTTP_Handler(srv SysJobLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysJobLogReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysjoblog.v1.SysJobLog/GetSysJobLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSysJobLog(ctx, req.(*SysJobLogReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysJobLogReply)
		return ctx.Result(200, reply)
	}
}

func _SysJobLog_UpdateSysJobLog0_HTTP_Handler(srv SysJobLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysJobLogUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysjoblog.v1.SysJobLog/UpdateSysJobLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateSysJobLog(ctx, req.(*SysJobLogUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysJobLogUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _SysJobLog_CreateSysJobLog0_HTTP_Handler(srv SysJobLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysJobLogCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysjoblog.v1.SysJobLog/CreateSysJobLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSysJobLog(ctx, req.(*SysJobLogCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysJobLogCreateReply)
		return ctx.Result(200, reply)
	}
}

func _SysJobLog_DeleteSysJobLog0_HTTP_Handler(srv SysJobLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysJobLogDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysjoblog.v1.SysJobLog/DeleteSysJobLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteSysJobLog(ctx, req.(*SysJobLogDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysJobLogDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _SysJobLog_BatchDeleteSysJobLog0_HTTP_Handler(srv SysJobLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysJobLogBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysjoblog.v1.SysJobLog/BatchDeleteSysJobLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteSysJobLog(ctx, req.(*SysJobLogBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysJobLogDeleteReply)
		return ctx.Result(200, reply)
	}
}

type SysJobLogHTTPClient interface {
	BatchDeleteSysJobLog(ctx context.Context, req *SysJobLogBatchDeleteReq, opts ...http.CallOption) (rsp *SysJobLogDeleteReply, err error)
	CreateSysJobLog(ctx context.Context, req *SysJobLogCreateReq, opts ...http.CallOption) (rsp *SysJobLogCreateReply, err error)
	DeleteSysJobLog(ctx context.Context, req *SysJobLogDeleteReq, opts ...http.CallOption) (rsp *SysJobLogDeleteReply, err error)
	GetPageSysJobLog(ctx context.Context, req *SysJobLogPageReq, opts ...http.CallOption) (rsp *SysJobLogPageReply, err error)
	GetSysJobLog(ctx context.Context, req *SysJobLogReq, opts ...http.CallOption) (rsp *SysJobLogReply, err error)
	UpdateSysJobLog(ctx context.Context, req *SysJobLogUpdateReq, opts ...http.CallOption) (rsp *SysJobLogUpdateReply, err error)
}

type SysJobLogHTTPClientImpl struct {
	cc *http.Client
}

func NewSysJobLogHTTPClient(client *http.Client) SysJobLogHTTPClient {
	return &SysJobLogHTTPClientImpl{client}
}

func (c *SysJobLogHTTPClientImpl) BatchDeleteSysJobLog(ctx context.Context, in *SysJobLogBatchDeleteReq, opts ...http.CallOption) (*SysJobLogDeleteReply, error) {
	var out SysJobLogDeleteReply
	pattern := "/v1/sysjoblog"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/sysjoblog.v1.SysJobLog/BatchDeleteSysJobLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysJobLogHTTPClientImpl) CreateSysJobLog(ctx context.Context, in *SysJobLogCreateReq, opts ...http.CallOption) (*SysJobLogCreateReply, error) {
	var out SysJobLogCreateReply
	pattern := "/v1/sysjoblog"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/sysjoblog.v1.SysJobLog/CreateSysJobLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysJobLogHTTPClientImpl) DeleteSysJobLog(ctx context.Context, in *SysJobLogDeleteReq, opts ...http.CallOption) (*SysJobLogDeleteReply, error) {
	var out SysJobLogDeleteReply
	pattern := "/v1/sysjoblog/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/sysjoblog.v1.SysJobLog/DeleteSysJobLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysJobLogHTTPClientImpl) GetPageSysJobLog(ctx context.Context, in *SysJobLogPageReq, opts ...http.CallOption) (*SysJobLogPageReply, error) {
	var out SysJobLogPageReply
	pattern := "/v1/sysjoblog/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/sysjoblog.v1.SysJobLog/GetPageSysJobLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysJobLogHTTPClientImpl) GetSysJobLog(ctx context.Context, in *SysJobLogReq, opts ...http.CallOption) (*SysJobLogReply, error) {
	var out SysJobLogReply
	pattern := "/v1/sysjoblog/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/sysjoblog.v1.SysJobLog/GetSysJobLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysJobLogHTTPClientImpl) UpdateSysJobLog(ctx context.Context, in *SysJobLogUpdateReq, opts ...http.CallOption) (*SysJobLogUpdateReply, error) {
	var out SysJobLogUpdateReply
	pattern := "/v1/sysjoblog/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/sysjoblog.v1.SysJobLog/UpdateSysJobLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
