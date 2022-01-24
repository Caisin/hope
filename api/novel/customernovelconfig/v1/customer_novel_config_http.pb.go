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

type CustomerNovelConfigHTTPServer interface {
	BatchDeleteCustomerNovelConfig(context.Context, *CustomerNovelConfigBatchDeleteReq) (*CustomerNovelConfigDeleteReply, error)
	CreateCustomerNovelConfig(context.Context, *CustomerNovelConfigCreateReq) (*CustomerNovelConfigCreateReply, error)
	DeleteCustomerNovelConfig(context.Context, *CustomerNovelConfigDeleteReq) (*CustomerNovelConfigDeleteReply, error)
	GetCustomerNovelConfig(context.Context, *CustomerNovelConfigReq) (*CustomerNovelConfigReply, error)
	GetPageCustomerNovelConfig(context.Context, *CustomerNovelConfigPageReq) (*CustomerNovelConfigPageReply, error)
	UpdateCustomerNovelConfig(context.Context, *CustomerNovelConfigUpdateReq) (*CustomerNovelConfigUpdateReply, error)
}

func RegisterCustomerNovelConfigHTTPServer(s *http.Server, srv CustomerNovelConfigHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/customer/novel/config/page", _CustomerNovelConfig_GetPageCustomerNovelConfig0_HTTP_Handler(srv))
	r.GET("/v1/customer/novel/config/{id}", _CustomerNovelConfig_GetCustomerNovelConfig0_HTTP_Handler(srv))
	r.PUT("/v1/customer/novel/config/{id}", _CustomerNovelConfig_UpdateCustomerNovelConfig0_HTTP_Handler(srv))
	r.POST("/v1/customer/novel/config", _CustomerNovelConfig_CreateCustomerNovelConfig0_HTTP_Handler(srv))
	r.DELETE("/v1/customer/novel/config/{id}", _CustomerNovelConfig_DeleteCustomerNovelConfig0_HTTP_Handler(srv))
	r.DELETE("/v1/customer/novel/config", _CustomerNovelConfig_BatchDeleteCustomerNovelConfig0_HTTP_Handler(srv))
}

func _CustomerNovelConfig_GetPageCustomerNovelConfig0_HTTP_Handler(srv CustomerNovelConfigHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CustomerNovelConfigPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/customernovelconfig.v1.CustomerNovelConfig/GetPageCustomerNovelConfig")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPageCustomerNovelConfig(ctx, req.(*CustomerNovelConfigPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CustomerNovelConfigPageReply)
		return ctx.Result(200, reply)
	}
}

func _CustomerNovelConfig_GetCustomerNovelConfig0_HTTP_Handler(srv CustomerNovelConfigHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CustomerNovelConfigReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/customernovelconfig.v1.CustomerNovelConfig/GetCustomerNovelConfig")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCustomerNovelConfig(ctx, req.(*CustomerNovelConfigReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CustomerNovelConfigReply)
		return ctx.Result(200, reply)
	}
}

func _CustomerNovelConfig_UpdateCustomerNovelConfig0_HTTP_Handler(srv CustomerNovelConfigHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CustomerNovelConfigUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/customernovelconfig.v1.CustomerNovelConfig/UpdateCustomerNovelConfig")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCustomerNovelConfig(ctx, req.(*CustomerNovelConfigUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CustomerNovelConfigUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _CustomerNovelConfig_CreateCustomerNovelConfig0_HTTP_Handler(srv CustomerNovelConfigHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CustomerNovelConfigCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/customernovelconfig.v1.CustomerNovelConfig/CreateCustomerNovelConfig")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCustomerNovelConfig(ctx, req.(*CustomerNovelConfigCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CustomerNovelConfigCreateReply)
		return ctx.Result(200, reply)
	}
}

func _CustomerNovelConfig_DeleteCustomerNovelConfig0_HTTP_Handler(srv CustomerNovelConfigHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CustomerNovelConfigDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/customernovelconfig.v1.CustomerNovelConfig/DeleteCustomerNovelConfig")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCustomerNovelConfig(ctx, req.(*CustomerNovelConfigDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CustomerNovelConfigDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _CustomerNovelConfig_BatchDeleteCustomerNovelConfig0_HTTP_Handler(srv CustomerNovelConfigHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CustomerNovelConfigBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/customernovelconfig.v1.CustomerNovelConfig/BatchDeleteCustomerNovelConfig")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteCustomerNovelConfig(ctx, req.(*CustomerNovelConfigBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CustomerNovelConfigDeleteReply)
		return ctx.Result(200, reply)
	}
}

type CustomerNovelConfigHTTPClient interface {
	BatchDeleteCustomerNovelConfig(ctx context.Context, req *CustomerNovelConfigBatchDeleteReq, opts ...http.CallOption) (rsp *CustomerNovelConfigDeleteReply, err error)
	CreateCustomerNovelConfig(ctx context.Context, req *CustomerNovelConfigCreateReq, opts ...http.CallOption) (rsp *CustomerNovelConfigCreateReply, err error)
	DeleteCustomerNovelConfig(ctx context.Context, req *CustomerNovelConfigDeleteReq, opts ...http.CallOption) (rsp *CustomerNovelConfigDeleteReply, err error)
	GetCustomerNovelConfig(ctx context.Context, req *CustomerNovelConfigReq, opts ...http.CallOption) (rsp *CustomerNovelConfigReply, err error)
	GetPageCustomerNovelConfig(ctx context.Context, req *CustomerNovelConfigPageReq, opts ...http.CallOption) (rsp *CustomerNovelConfigPageReply, err error)
	UpdateCustomerNovelConfig(ctx context.Context, req *CustomerNovelConfigUpdateReq, opts ...http.CallOption) (rsp *CustomerNovelConfigUpdateReply, err error)
}

type CustomerNovelConfigHTTPClientImpl struct {
	cc *http.Client
}

func NewCustomerNovelConfigHTTPClient(client *http.Client) CustomerNovelConfigHTTPClient {
	return &CustomerNovelConfigHTTPClientImpl{client}
}

func (c *CustomerNovelConfigHTTPClientImpl) BatchDeleteCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigBatchDeleteReq, opts ...http.CallOption) (*CustomerNovelConfigDeleteReply, error) {
	var out CustomerNovelConfigDeleteReply
	pattern := "/v1/customer/novel/config"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/customernovelconfig.v1.CustomerNovelConfig/BatchDeleteCustomerNovelConfig"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CustomerNovelConfigHTTPClientImpl) CreateCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigCreateReq, opts ...http.CallOption) (*CustomerNovelConfigCreateReply, error) {
	var out CustomerNovelConfigCreateReply
	pattern := "/v1/customer/novel/config"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/customernovelconfig.v1.CustomerNovelConfig/CreateCustomerNovelConfig"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CustomerNovelConfigHTTPClientImpl) DeleteCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigDeleteReq, opts ...http.CallOption) (*CustomerNovelConfigDeleteReply, error) {
	var out CustomerNovelConfigDeleteReply
	pattern := "/v1/customer/novel/config/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/customernovelconfig.v1.CustomerNovelConfig/DeleteCustomerNovelConfig"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CustomerNovelConfigHTTPClientImpl) GetCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigReq, opts ...http.CallOption) (*CustomerNovelConfigReply, error) {
	var out CustomerNovelConfigReply
	pattern := "/v1/customer/novel/config/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/customernovelconfig.v1.CustomerNovelConfig/GetCustomerNovelConfig"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CustomerNovelConfigHTTPClientImpl) GetPageCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigPageReq, opts ...http.CallOption) (*CustomerNovelConfigPageReply, error) {
	var out CustomerNovelConfigPageReply
	pattern := "/v1/customer/novel/config/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/customernovelconfig.v1.CustomerNovelConfig/GetPageCustomerNovelConfig"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CustomerNovelConfigHTTPClientImpl) UpdateCustomerNovelConfig(ctx context.Context, in *CustomerNovelConfigUpdateReq, opts ...http.CallOption) (*CustomerNovelConfigUpdateReply, error) {
	var out CustomerNovelConfigUpdateReply
	pattern := "/v1/customer/novel/config/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/customernovelconfig.v1.CustomerNovelConfig/UpdateCustomerNovelConfig"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}