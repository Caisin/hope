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

type AgreementLogHTTPServer interface {
	BatchDeleteAgreementLog(context.Context, *AgreementLogBatchDeleteReq) (*AgreementLogDeleteReply, error)
	CreateAgreementLog(context.Context, *AgreementLogCreateReq) (*AgreementLogCreateReply, error)
	DeleteAgreementLog(context.Context, *AgreementLogDeleteReq) (*AgreementLogDeleteReply, error)
	GetAgreementLog(context.Context, *AgreementLogReq) (*AgreementLogReply, error)
	GetPageAgreementLog(context.Context, *AgreementLogPageReq) (*AgreementLogPageReply, error)
	UpdateAgreementLog(context.Context, *AgreementLogUpdateReq) (*AgreementLogUpdateReply, error)
}

func RegisterAgreementLogHTTPServer(s *http.Server, srv AgreementLogHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/agreementlog/page", _AgreementLog_GetPageAgreementLog0_HTTP_Handler(srv))
	r.GET("/v1/agreementlog/{id}", _AgreementLog_GetAgreementLog0_HTTP_Handler(srv))
	r.PUT("/v1/agreementlog/{id}", _AgreementLog_UpdateAgreementLog0_HTTP_Handler(srv))
	r.POST("/v1/agreementlog", _AgreementLog_CreateAgreementLog0_HTTP_Handler(srv))
	r.DELETE("/v1/agreementlog/{id}", _AgreementLog_DeleteAgreementLog0_HTTP_Handler(srv))
	r.DELETE("/v1/agreementlog", _AgreementLog_BatchDeleteAgreementLog0_HTTP_Handler(srv))
}

func _AgreementLog_GetPageAgreementLog0_HTTP_Handler(srv AgreementLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AgreementLogPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/agreementlog.v1.AgreementLog/GetPageAgreementLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPageAgreementLog(ctx, req.(*AgreementLogPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AgreementLogPageReply)
		return ctx.Result(200, reply)
	}
}

func _AgreementLog_GetAgreementLog0_HTTP_Handler(srv AgreementLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AgreementLogReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/agreementlog.v1.AgreementLog/GetAgreementLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAgreementLog(ctx, req.(*AgreementLogReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AgreementLogReply)
		return ctx.Result(200, reply)
	}
}

func _AgreementLog_UpdateAgreementLog0_HTTP_Handler(srv AgreementLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AgreementLogUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/agreementlog.v1.AgreementLog/UpdateAgreementLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateAgreementLog(ctx, req.(*AgreementLogUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AgreementLogUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _AgreementLog_CreateAgreementLog0_HTTP_Handler(srv AgreementLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AgreementLogCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/agreementlog.v1.AgreementLog/CreateAgreementLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAgreementLog(ctx, req.(*AgreementLogCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AgreementLogCreateReply)
		return ctx.Result(200, reply)
	}
}

func _AgreementLog_DeleteAgreementLog0_HTTP_Handler(srv AgreementLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AgreementLogDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/agreementlog.v1.AgreementLog/DeleteAgreementLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteAgreementLog(ctx, req.(*AgreementLogDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AgreementLogDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _AgreementLog_BatchDeleteAgreementLog0_HTTP_Handler(srv AgreementLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AgreementLogBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/agreementlog.v1.AgreementLog/BatchDeleteAgreementLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteAgreementLog(ctx, req.(*AgreementLogBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AgreementLogDeleteReply)
		return ctx.Result(200, reply)
	}
}

type AgreementLogHTTPClient interface {
	BatchDeleteAgreementLog(ctx context.Context, req *AgreementLogBatchDeleteReq, opts ...http.CallOption) (rsp *AgreementLogDeleteReply, err error)
	CreateAgreementLog(ctx context.Context, req *AgreementLogCreateReq, opts ...http.CallOption) (rsp *AgreementLogCreateReply, err error)
	DeleteAgreementLog(ctx context.Context, req *AgreementLogDeleteReq, opts ...http.CallOption) (rsp *AgreementLogDeleteReply, err error)
	GetAgreementLog(ctx context.Context, req *AgreementLogReq, opts ...http.CallOption) (rsp *AgreementLogReply, err error)
	GetPageAgreementLog(ctx context.Context, req *AgreementLogPageReq, opts ...http.CallOption) (rsp *AgreementLogPageReply, err error)
	UpdateAgreementLog(ctx context.Context, req *AgreementLogUpdateReq, opts ...http.CallOption) (rsp *AgreementLogUpdateReply, err error)
}

type AgreementLogHTTPClientImpl struct {
	cc *http.Client
}

func NewAgreementLogHTTPClient(client *http.Client) AgreementLogHTTPClient {
	return &AgreementLogHTTPClientImpl{client}
}

func (c *AgreementLogHTTPClientImpl) BatchDeleteAgreementLog(ctx context.Context, in *AgreementLogBatchDeleteReq, opts ...http.CallOption) (*AgreementLogDeleteReply, error) {
	var out AgreementLogDeleteReply
	pattern := "/v1/agreementlog"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/agreementlog.v1.AgreementLog/BatchDeleteAgreementLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AgreementLogHTTPClientImpl) CreateAgreementLog(ctx context.Context, in *AgreementLogCreateReq, opts ...http.CallOption) (*AgreementLogCreateReply, error) {
	var out AgreementLogCreateReply
	pattern := "/v1/agreementlog"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/agreementlog.v1.AgreementLog/CreateAgreementLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AgreementLogHTTPClientImpl) DeleteAgreementLog(ctx context.Context, in *AgreementLogDeleteReq, opts ...http.CallOption) (*AgreementLogDeleteReply, error) {
	var out AgreementLogDeleteReply
	pattern := "/v1/agreementlog/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/agreementlog.v1.AgreementLog/DeleteAgreementLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AgreementLogHTTPClientImpl) GetAgreementLog(ctx context.Context, in *AgreementLogReq, opts ...http.CallOption) (*AgreementLogReply, error) {
	var out AgreementLogReply
	pattern := "/v1/agreementlog/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/agreementlog.v1.AgreementLog/GetAgreementLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AgreementLogHTTPClientImpl) GetPageAgreementLog(ctx context.Context, in *AgreementLogPageReq, opts ...http.CallOption) (*AgreementLogPageReply, error) {
	var out AgreementLogPageReply
	pattern := "/v1/agreementlog/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/agreementlog.v1.AgreementLog/GetPageAgreementLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AgreementLogHTTPClientImpl) UpdateAgreementLog(ctx context.Context, in *AgreementLogUpdateReq, opts ...http.CallOption) (*AgreementLogUpdateReply, error) {
	var out AgreementLogUpdateReply
	pattern := "/v1/agreementlog/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/agreementlog.v1.AgreementLog/UpdateAgreementLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
