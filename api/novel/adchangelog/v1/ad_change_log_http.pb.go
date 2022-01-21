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

type AdChangeLogHTTPServer interface {
	BatchDeleteAdChangeLog(context.Context, *AdChangeLogBatchDeleteReq) (*AdChangeLogDeleteReply, error)
	CreateAdChangeLog(context.Context, *AdChangeLogCreateReq) (*AdChangeLogCreateReply, error)
	DeleteAdChangeLog(context.Context, *AdChangeLogDeleteReq) (*AdChangeLogDeleteReply, error)
	GetAdChangeLog(context.Context, *AdChangeLogReq) (*AdChangeLogReply, error)
	GetPageAdChangeLog(context.Context, *AdChangeLogPageReq) (*AdChangeLogPageReply, error)
	UpdateAdChangeLog(context.Context, *AdChangeLogUpdateReq) (*AdChangeLogUpdateReply, error)
}

func RegisterAdChangeLogHTTPServer(s *http.Server, srv AdChangeLogHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/adchangelog/page", _AdChangeLog_GetPageAdChangeLog0_HTTP_Handler(srv))
	r.GET("/v1/adchangelog/{id}", _AdChangeLog_GetAdChangeLog0_HTTP_Handler(srv))
	r.PUT("/v1/adchangelog/{id}", _AdChangeLog_UpdateAdChangeLog0_HTTP_Handler(srv))
	r.POST("/v1/adchangelog", _AdChangeLog_CreateAdChangeLog0_HTTP_Handler(srv))
	r.DELETE("/v1/adchangelog/{id}", _AdChangeLog_DeleteAdChangeLog0_HTTP_Handler(srv))
	r.DELETE("/v1/adchangelog", _AdChangeLog_BatchDeleteAdChangeLog0_HTTP_Handler(srv))
}

func _AdChangeLog_GetPageAdChangeLog0_HTTP_Handler(srv AdChangeLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdChangeLogPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/adchangelog.v1.AdChangeLog/GetPageAdChangeLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPageAdChangeLog(ctx, req.(*AdChangeLogPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdChangeLogPageReply)
		return ctx.Result(200, reply)
	}
}

func _AdChangeLog_GetAdChangeLog0_HTTP_Handler(srv AdChangeLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdChangeLogReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/adchangelog.v1.AdChangeLog/GetAdChangeLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAdChangeLog(ctx, req.(*AdChangeLogReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdChangeLogReply)
		return ctx.Result(200, reply)
	}
}

func _AdChangeLog_UpdateAdChangeLog0_HTTP_Handler(srv AdChangeLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdChangeLogUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/adchangelog.v1.AdChangeLog/UpdateAdChangeLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateAdChangeLog(ctx, req.(*AdChangeLogUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdChangeLogUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _AdChangeLog_CreateAdChangeLog0_HTTP_Handler(srv AdChangeLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdChangeLogCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/adchangelog.v1.AdChangeLog/CreateAdChangeLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAdChangeLog(ctx, req.(*AdChangeLogCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdChangeLogCreateReply)
		return ctx.Result(200, reply)
	}
}

func _AdChangeLog_DeleteAdChangeLog0_HTTP_Handler(srv AdChangeLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdChangeLogDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/adchangelog.v1.AdChangeLog/DeleteAdChangeLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteAdChangeLog(ctx, req.(*AdChangeLogDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdChangeLogDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _AdChangeLog_BatchDeleteAdChangeLog0_HTTP_Handler(srv AdChangeLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdChangeLogBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/adchangelog.v1.AdChangeLog/BatchDeleteAdChangeLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteAdChangeLog(ctx, req.(*AdChangeLogBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdChangeLogDeleteReply)
		return ctx.Result(200, reply)
	}
}

type AdChangeLogHTTPClient interface {
	BatchDeleteAdChangeLog(ctx context.Context, req *AdChangeLogBatchDeleteReq, opts ...http.CallOption) (rsp *AdChangeLogDeleteReply, err error)
	CreateAdChangeLog(ctx context.Context, req *AdChangeLogCreateReq, opts ...http.CallOption) (rsp *AdChangeLogCreateReply, err error)
	DeleteAdChangeLog(ctx context.Context, req *AdChangeLogDeleteReq, opts ...http.CallOption) (rsp *AdChangeLogDeleteReply, err error)
	GetAdChangeLog(ctx context.Context, req *AdChangeLogReq, opts ...http.CallOption) (rsp *AdChangeLogReply, err error)
	GetPageAdChangeLog(ctx context.Context, req *AdChangeLogPageReq, opts ...http.CallOption) (rsp *AdChangeLogPageReply, err error)
	UpdateAdChangeLog(ctx context.Context, req *AdChangeLogUpdateReq, opts ...http.CallOption) (rsp *AdChangeLogUpdateReply, err error)
}

type AdChangeLogHTTPClientImpl struct {
	cc *http.Client
}

func NewAdChangeLogHTTPClient(client *http.Client) AdChangeLogHTTPClient {
	return &AdChangeLogHTTPClientImpl{client}
}

func (c *AdChangeLogHTTPClientImpl) BatchDeleteAdChangeLog(ctx context.Context, in *AdChangeLogBatchDeleteReq, opts ...http.CallOption) (*AdChangeLogDeleteReply, error) {
	var out AdChangeLogDeleteReply
	pattern := "/v1/adchangelog"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/adchangelog.v1.AdChangeLog/BatchDeleteAdChangeLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdChangeLogHTTPClientImpl) CreateAdChangeLog(ctx context.Context, in *AdChangeLogCreateReq, opts ...http.CallOption) (*AdChangeLogCreateReply, error) {
	var out AdChangeLogCreateReply
	pattern := "/v1/adchangelog"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/adchangelog.v1.AdChangeLog/CreateAdChangeLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdChangeLogHTTPClientImpl) DeleteAdChangeLog(ctx context.Context, in *AdChangeLogDeleteReq, opts ...http.CallOption) (*AdChangeLogDeleteReply, error) {
	var out AdChangeLogDeleteReply
	pattern := "/v1/adchangelog/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/adchangelog.v1.AdChangeLog/DeleteAdChangeLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdChangeLogHTTPClientImpl) GetAdChangeLog(ctx context.Context, in *AdChangeLogReq, opts ...http.CallOption) (*AdChangeLogReply, error) {
	var out AdChangeLogReply
	pattern := "/v1/adchangelog/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/adchangelog.v1.AdChangeLog/GetAdChangeLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdChangeLogHTTPClientImpl) GetPageAdChangeLog(ctx context.Context, in *AdChangeLogPageReq, opts ...http.CallOption) (*AdChangeLogPageReply, error) {
	var out AdChangeLogPageReply
	pattern := "/v1/adchangelog/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/adchangelog.v1.AdChangeLog/GetPageAdChangeLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdChangeLogHTTPClientImpl) UpdateAdChangeLog(ctx context.Context, in *AdChangeLogUpdateReq, opts ...http.CallOption) (*AdChangeLogUpdateReply, error) {
	var out AdChangeLogUpdateReply
	pattern := "/v1/adchangelog/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/adchangelog.v1.AdChangeLog/UpdateAdChangeLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
