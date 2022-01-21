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

type CasbinRuleHTTPServer interface {
	BatchDeleteCasbinRule(context.Context, *CasbinRuleBatchDeleteReq) (*CasbinRuleDeleteReply, error)
	CreateCasbinRule(context.Context, *CasbinRuleCreateReq) (*CasbinRuleCreateReply, error)
	DeleteCasbinRule(context.Context, *CasbinRuleDeleteReq) (*CasbinRuleDeleteReply, error)
	GetCasbinRule(context.Context, *CasbinRuleReq) (*CasbinRuleReply, error)
	GetPageCasbinRule(context.Context, *CasbinRulePageReq) (*CasbinRulePageReply, error)
	UpdateCasbinRule(context.Context, *CasbinRuleUpdateReq) (*CasbinRuleUpdateReply, error)
}

func RegisterCasbinRuleHTTPServer(s *http.Server, srv CasbinRuleHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/casbinrule/page", _CasbinRule_GetPageCasbinRule0_HTTP_Handler(srv))
	r.GET("/v1/casbinrule/{id}", _CasbinRule_GetCasbinRule0_HTTP_Handler(srv))
	r.PUT("/v1/casbinrule/{id}", _CasbinRule_UpdateCasbinRule0_HTTP_Handler(srv))
	r.POST("/v1/casbinrule", _CasbinRule_CreateCasbinRule0_HTTP_Handler(srv))
	r.DELETE("/v1/casbinrule/{id}", _CasbinRule_DeleteCasbinRule0_HTTP_Handler(srv))
	r.DELETE("/v1/casbinrule", _CasbinRule_BatchDeleteCasbinRule0_HTTP_Handler(srv))
}

func _CasbinRule_GetPageCasbinRule0_HTTP_Handler(srv CasbinRuleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CasbinRulePageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/casbinrule.v1.CasbinRule/GetPageCasbinRule")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPageCasbinRule(ctx, req.(*CasbinRulePageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CasbinRulePageReply)
		return ctx.Result(200, reply)
	}
}

func _CasbinRule_GetCasbinRule0_HTTP_Handler(srv CasbinRuleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CasbinRuleReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/casbinrule.v1.CasbinRule/GetCasbinRule")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCasbinRule(ctx, req.(*CasbinRuleReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CasbinRuleReply)
		return ctx.Result(200, reply)
	}
}

func _CasbinRule_UpdateCasbinRule0_HTTP_Handler(srv CasbinRuleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CasbinRuleUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/casbinrule.v1.CasbinRule/UpdateCasbinRule")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCasbinRule(ctx, req.(*CasbinRuleUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CasbinRuleUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _CasbinRule_CreateCasbinRule0_HTTP_Handler(srv CasbinRuleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CasbinRuleCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/casbinrule.v1.CasbinRule/CreateCasbinRule")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCasbinRule(ctx, req.(*CasbinRuleCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CasbinRuleCreateReply)
		return ctx.Result(200, reply)
	}
}

func _CasbinRule_DeleteCasbinRule0_HTTP_Handler(srv CasbinRuleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CasbinRuleDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/casbinrule.v1.CasbinRule/DeleteCasbinRule")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCasbinRule(ctx, req.(*CasbinRuleDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CasbinRuleDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _CasbinRule_BatchDeleteCasbinRule0_HTTP_Handler(srv CasbinRuleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CasbinRuleBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/casbinrule.v1.CasbinRule/BatchDeleteCasbinRule")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteCasbinRule(ctx, req.(*CasbinRuleBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CasbinRuleDeleteReply)
		return ctx.Result(200, reply)
	}
}

type CasbinRuleHTTPClient interface {
	BatchDeleteCasbinRule(ctx context.Context, req *CasbinRuleBatchDeleteReq, opts ...http.CallOption) (rsp *CasbinRuleDeleteReply, err error)
	CreateCasbinRule(ctx context.Context, req *CasbinRuleCreateReq, opts ...http.CallOption) (rsp *CasbinRuleCreateReply, err error)
	DeleteCasbinRule(ctx context.Context, req *CasbinRuleDeleteReq, opts ...http.CallOption) (rsp *CasbinRuleDeleteReply, err error)
	GetCasbinRule(ctx context.Context, req *CasbinRuleReq, opts ...http.CallOption) (rsp *CasbinRuleReply, err error)
	GetPageCasbinRule(ctx context.Context, req *CasbinRulePageReq, opts ...http.CallOption) (rsp *CasbinRulePageReply, err error)
	UpdateCasbinRule(ctx context.Context, req *CasbinRuleUpdateReq, opts ...http.CallOption) (rsp *CasbinRuleUpdateReply, err error)
}

type CasbinRuleHTTPClientImpl struct {
	cc *http.Client
}

func NewCasbinRuleHTTPClient(client *http.Client) CasbinRuleHTTPClient {
	return &CasbinRuleHTTPClientImpl{client}
}

func (c *CasbinRuleHTTPClientImpl) BatchDeleteCasbinRule(ctx context.Context, in *CasbinRuleBatchDeleteReq, opts ...http.CallOption) (*CasbinRuleDeleteReply, error) {
	var out CasbinRuleDeleteReply
	pattern := "/v1/casbinrule"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/casbinrule.v1.CasbinRule/BatchDeleteCasbinRule"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CasbinRuleHTTPClientImpl) CreateCasbinRule(ctx context.Context, in *CasbinRuleCreateReq, opts ...http.CallOption) (*CasbinRuleCreateReply, error) {
	var out CasbinRuleCreateReply
	pattern := "/v1/casbinrule"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/casbinrule.v1.CasbinRule/CreateCasbinRule"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CasbinRuleHTTPClientImpl) DeleteCasbinRule(ctx context.Context, in *CasbinRuleDeleteReq, opts ...http.CallOption) (*CasbinRuleDeleteReply, error) {
	var out CasbinRuleDeleteReply
	pattern := "/v1/casbinrule/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/casbinrule.v1.CasbinRule/DeleteCasbinRule"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CasbinRuleHTTPClientImpl) GetCasbinRule(ctx context.Context, in *CasbinRuleReq, opts ...http.CallOption) (*CasbinRuleReply, error) {
	var out CasbinRuleReply
	pattern := "/v1/casbinrule/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/casbinrule.v1.CasbinRule/GetCasbinRule"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CasbinRuleHTTPClientImpl) GetPageCasbinRule(ctx context.Context, in *CasbinRulePageReq, opts ...http.CallOption) (*CasbinRulePageReply, error) {
	var out CasbinRulePageReply
	pattern := "/v1/casbinrule/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/casbinrule.v1.CasbinRule/GetPageCasbinRule"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CasbinRuleHTTPClientImpl) UpdateCasbinRule(ctx context.Context, in *CasbinRuleUpdateReq, opts ...http.CallOption) (*CasbinRuleUpdateReply, error) {
	var out CasbinRuleUpdateReply
	pattern := "/v1/casbinrule/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/casbinrule.v1.CasbinRule/UpdateCasbinRule"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
