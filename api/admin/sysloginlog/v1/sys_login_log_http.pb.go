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

type SysLoginLogHTTPServer interface {
	BatchDeleteSysLoginLog(context.Context, *SysLoginLogBatchDeleteReq) (*SysLoginLogDeleteReply, error)
	CreateSysLoginLog(context.Context, *SysLoginLogCreateReq) (*SysLoginLogCreateReply, error)
	DeleteSysLoginLog(context.Context, *SysLoginLogDeleteReq) (*SysLoginLogDeleteReply, error)
	GetSysLoginLog(context.Context, *SysLoginLogReq) (*SysLoginLogReply, error)
	GetSysLoginLogPage(context.Context, *SysLoginLogPageReq) (*SysLoginLogPageReply, error)
	UpdateSysLoginLog(context.Context, *SysLoginLogUpdateReq) (*SysLoginLogUpdateReply, error)
}

func RegisterSysLoginLogHTTPServer(s *http.Server, srv SysLoginLogHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/sys/login/log/page", _SysLoginLog_GetSysLoginLogPage0_HTTP_Handler(srv))
	r.GET("/v1/sys/login/log/{id}", _SysLoginLog_GetSysLoginLog0_HTTP_Handler(srv))
	r.PUT("/v1/sys/login/log/{id}", _SysLoginLog_UpdateSysLoginLog0_HTTP_Handler(srv))
	r.POST("/v1/sys/login/log", _SysLoginLog_CreateSysLoginLog0_HTTP_Handler(srv))
	r.DELETE("/v1/sys/login/log/{id}", _SysLoginLog_DeleteSysLoginLog0_HTTP_Handler(srv))
	r.DELETE("/v1/sys/login/log", _SysLoginLog_BatchDeleteSysLoginLog0_HTTP_Handler(srv))
}

func _SysLoginLog_GetSysLoginLogPage0_HTTP_Handler(srv SysLoginLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysLoginLogPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysloginlog.v1.SysLoginLog/GetSysLoginLogPage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSysLoginLogPage(ctx, req.(*SysLoginLogPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysLoginLogPageReply)
		return ctx.Result(200, reply)
	}
}

func _SysLoginLog_GetSysLoginLog0_HTTP_Handler(srv SysLoginLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysLoginLogReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysloginlog.v1.SysLoginLog/GetSysLoginLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSysLoginLog(ctx, req.(*SysLoginLogReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysLoginLogReply)
		return ctx.Result(200, reply)
	}
}

func _SysLoginLog_UpdateSysLoginLog0_HTTP_Handler(srv SysLoginLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysLoginLogUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysloginlog.v1.SysLoginLog/UpdateSysLoginLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateSysLoginLog(ctx, req.(*SysLoginLogUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysLoginLogUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _SysLoginLog_CreateSysLoginLog0_HTTP_Handler(srv SysLoginLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysLoginLogCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysloginlog.v1.SysLoginLog/CreateSysLoginLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSysLoginLog(ctx, req.(*SysLoginLogCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysLoginLogCreateReply)
		return ctx.Result(200, reply)
	}
}

func _SysLoginLog_DeleteSysLoginLog0_HTTP_Handler(srv SysLoginLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysLoginLogDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysloginlog.v1.SysLoginLog/DeleteSysLoginLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteSysLoginLog(ctx, req.(*SysLoginLogDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysLoginLogDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _SysLoginLog_BatchDeleteSysLoginLog0_HTTP_Handler(srv SysLoginLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysLoginLogBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysloginlog.v1.SysLoginLog/BatchDeleteSysLoginLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteSysLoginLog(ctx, req.(*SysLoginLogBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysLoginLogDeleteReply)
		return ctx.Result(200, reply)
	}
}

type SysLoginLogHTTPClient interface {
	BatchDeleteSysLoginLog(ctx context.Context, req *SysLoginLogBatchDeleteReq, opts ...http.CallOption) (rsp *SysLoginLogDeleteReply, err error)
	CreateSysLoginLog(ctx context.Context, req *SysLoginLogCreateReq, opts ...http.CallOption) (rsp *SysLoginLogCreateReply, err error)
	DeleteSysLoginLog(ctx context.Context, req *SysLoginLogDeleteReq, opts ...http.CallOption) (rsp *SysLoginLogDeleteReply, err error)
	GetSysLoginLog(ctx context.Context, req *SysLoginLogReq, opts ...http.CallOption) (rsp *SysLoginLogReply, err error)
	GetSysLoginLogPage(ctx context.Context, req *SysLoginLogPageReq, opts ...http.CallOption) (rsp *SysLoginLogPageReply, err error)
	UpdateSysLoginLog(ctx context.Context, req *SysLoginLogUpdateReq, opts ...http.CallOption) (rsp *SysLoginLogUpdateReply, err error)
}

type SysLoginLogHTTPClientImpl struct {
	cc *http.Client
}

func NewSysLoginLogHTTPClient(client *http.Client) SysLoginLogHTTPClient {
	return &SysLoginLogHTTPClientImpl{client}
}

func (c *SysLoginLogHTTPClientImpl) BatchDeleteSysLoginLog(ctx context.Context, in *SysLoginLogBatchDeleteReq, opts ...http.CallOption) (*SysLoginLogDeleteReply, error) {
	var out SysLoginLogDeleteReply
	pattern := "/v1/sys/login/log"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/sysloginlog.v1.SysLoginLog/BatchDeleteSysLoginLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysLoginLogHTTPClientImpl) CreateSysLoginLog(ctx context.Context, in *SysLoginLogCreateReq, opts ...http.CallOption) (*SysLoginLogCreateReply, error) {
	var out SysLoginLogCreateReply
	pattern := "/v1/sys/login/log"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/sysloginlog.v1.SysLoginLog/CreateSysLoginLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysLoginLogHTTPClientImpl) DeleteSysLoginLog(ctx context.Context, in *SysLoginLogDeleteReq, opts ...http.CallOption) (*SysLoginLogDeleteReply, error) {
	var out SysLoginLogDeleteReply
	pattern := "/v1/sys/login/log/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/sysloginlog.v1.SysLoginLog/DeleteSysLoginLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysLoginLogHTTPClientImpl) GetSysLoginLog(ctx context.Context, in *SysLoginLogReq, opts ...http.CallOption) (*SysLoginLogReply, error) {
	var out SysLoginLogReply
	pattern := "/v1/sys/login/log/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/sysloginlog.v1.SysLoginLog/GetSysLoginLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysLoginLogHTTPClientImpl) GetSysLoginLogPage(ctx context.Context, in *SysLoginLogPageReq, opts ...http.CallOption) (*SysLoginLogPageReply, error) {
	var out SysLoginLogPageReply
	pattern := "/v1/sys/login/log/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/sysloginlog.v1.SysLoginLog/GetSysLoginLogPage"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysLoginLogHTTPClientImpl) UpdateSysLoginLog(ctx context.Context, in *SysLoginLogUpdateReq, opts ...http.CallOption) (*SysLoginLogUpdateReply, error) {
	var out SysLoginLogUpdateReply
	pattern := "/v1/sys/login/log/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/sysloginlog.v1.SysLoginLog/UpdateSysLoginLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
