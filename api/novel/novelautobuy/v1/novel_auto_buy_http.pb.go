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

type NovelAutoBuyHTTPServer interface {
	BatchDeleteNovelAutoBuy(context.Context, *NovelAutoBuyBatchDeleteReq) (*NovelAutoBuyDeleteReply, error)
	CreateNovelAutoBuy(context.Context, *NovelAutoBuyCreateReq) (*NovelAutoBuyCreateReply, error)
	DeleteNovelAutoBuy(context.Context, *NovelAutoBuyDeleteReq) (*NovelAutoBuyDeleteReply, error)
	GetNovelAutoBuy(context.Context, *NovelAutoBuyReq) (*NovelAutoBuyReply, error)
	GetPageNovelAutoBuy(context.Context, *NovelAutoBuyPageReq) (*NovelAutoBuyPageReply, error)
	UpdateNovelAutoBuy(context.Context, *NovelAutoBuyUpdateReq) (*NovelAutoBuyUpdateReply, error)
}

func RegisterNovelAutoBuyHTTPServer(s *http.Server, srv NovelAutoBuyHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/novel/auto/buy/page", _NovelAutoBuy_GetPageNovelAutoBuy0_HTTP_Handler(srv))
	r.GET("/v1/novel/auto/buy/{id}", _NovelAutoBuy_GetNovelAutoBuy0_HTTP_Handler(srv))
	r.PUT("/v1/novel/auto/buy/{id}", _NovelAutoBuy_UpdateNovelAutoBuy0_HTTP_Handler(srv))
	r.POST("/v1/novel/auto/buy", _NovelAutoBuy_CreateNovelAutoBuy0_HTTP_Handler(srv))
	r.DELETE("/v1/novel/auto/buy/{id}", _NovelAutoBuy_DeleteNovelAutoBuy0_HTTP_Handler(srv))
	r.DELETE("/v1/novel/auto/buy", _NovelAutoBuy_BatchDeleteNovelAutoBuy0_HTTP_Handler(srv))
}

func _NovelAutoBuy_GetPageNovelAutoBuy0_HTTP_Handler(srv NovelAutoBuyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelAutoBuyPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelautobuy.v1.NovelAutoBuy/GetPageNovelAutoBuy")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPageNovelAutoBuy(ctx, req.(*NovelAutoBuyPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelAutoBuyPageReply)
		return ctx.Result(200, reply)
	}
}

func _NovelAutoBuy_GetNovelAutoBuy0_HTTP_Handler(srv NovelAutoBuyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelAutoBuyReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelautobuy.v1.NovelAutoBuy/GetNovelAutoBuy")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNovelAutoBuy(ctx, req.(*NovelAutoBuyReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelAutoBuyReply)
		return ctx.Result(200, reply)
	}
}

func _NovelAutoBuy_UpdateNovelAutoBuy0_HTTP_Handler(srv NovelAutoBuyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelAutoBuyUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelautobuy.v1.NovelAutoBuy/UpdateNovelAutoBuy")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateNovelAutoBuy(ctx, req.(*NovelAutoBuyUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelAutoBuyUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _NovelAutoBuy_CreateNovelAutoBuy0_HTTP_Handler(srv NovelAutoBuyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelAutoBuyCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelautobuy.v1.NovelAutoBuy/CreateNovelAutoBuy")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateNovelAutoBuy(ctx, req.(*NovelAutoBuyCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelAutoBuyCreateReply)
		return ctx.Result(200, reply)
	}
}

func _NovelAutoBuy_DeleteNovelAutoBuy0_HTTP_Handler(srv NovelAutoBuyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelAutoBuyDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelautobuy.v1.NovelAutoBuy/DeleteNovelAutoBuy")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteNovelAutoBuy(ctx, req.(*NovelAutoBuyDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelAutoBuyDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _NovelAutoBuy_BatchDeleteNovelAutoBuy0_HTTP_Handler(srv NovelAutoBuyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelAutoBuyBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelautobuy.v1.NovelAutoBuy/BatchDeleteNovelAutoBuy")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteNovelAutoBuy(ctx, req.(*NovelAutoBuyBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelAutoBuyDeleteReply)
		return ctx.Result(200, reply)
	}
}

type NovelAutoBuyHTTPClient interface {
	BatchDeleteNovelAutoBuy(ctx context.Context, req *NovelAutoBuyBatchDeleteReq, opts ...http.CallOption) (rsp *NovelAutoBuyDeleteReply, err error)
	CreateNovelAutoBuy(ctx context.Context, req *NovelAutoBuyCreateReq, opts ...http.CallOption) (rsp *NovelAutoBuyCreateReply, err error)
	DeleteNovelAutoBuy(ctx context.Context, req *NovelAutoBuyDeleteReq, opts ...http.CallOption) (rsp *NovelAutoBuyDeleteReply, err error)
	GetNovelAutoBuy(ctx context.Context, req *NovelAutoBuyReq, opts ...http.CallOption) (rsp *NovelAutoBuyReply, err error)
	GetPageNovelAutoBuy(ctx context.Context, req *NovelAutoBuyPageReq, opts ...http.CallOption) (rsp *NovelAutoBuyPageReply, err error)
	UpdateNovelAutoBuy(ctx context.Context, req *NovelAutoBuyUpdateReq, opts ...http.CallOption) (rsp *NovelAutoBuyUpdateReply, err error)
}

type NovelAutoBuyHTTPClientImpl struct {
	cc *http.Client
}

func NewNovelAutoBuyHTTPClient(client *http.Client) NovelAutoBuyHTTPClient {
	return &NovelAutoBuyHTTPClientImpl{client}
}

func (c *NovelAutoBuyHTTPClientImpl) BatchDeleteNovelAutoBuy(ctx context.Context, in *NovelAutoBuyBatchDeleteReq, opts ...http.CallOption) (*NovelAutoBuyDeleteReply, error) {
	var out NovelAutoBuyDeleteReply
	pattern := "/v1/novel/auto/buy"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelautobuy.v1.NovelAutoBuy/BatchDeleteNovelAutoBuy"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelAutoBuyHTTPClientImpl) CreateNovelAutoBuy(ctx context.Context, in *NovelAutoBuyCreateReq, opts ...http.CallOption) (*NovelAutoBuyCreateReply, error) {
	var out NovelAutoBuyCreateReply
	pattern := "/v1/novel/auto/buy"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/novelautobuy.v1.NovelAutoBuy/CreateNovelAutoBuy"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelAutoBuyHTTPClientImpl) DeleteNovelAutoBuy(ctx context.Context, in *NovelAutoBuyDeleteReq, opts ...http.CallOption) (*NovelAutoBuyDeleteReply, error) {
	var out NovelAutoBuyDeleteReply
	pattern := "/v1/novel/auto/buy/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelautobuy.v1.NovelAutoBuy/DeleteNovelAutoBuy"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelAutoBuyHTTPClientImpl) GetNovelAutoBuy(ctx context.Context, in *NovelAutoBuyReq, opts ...http.CallOption) (*NovelAutoBuyReply, error) {
	var out NovelAutoBuyReply
	pattern := "/v1/novel/auto/buy/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelautobuy.v1.NovelAutoBuy/GetNovelAutoBuy"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelAutoBuyHTTPClientImpl) GetPageNovelAutoBuy(ctx context.Context, in *NovelAutoBuyPageReq, opts ...http.CallOption) (*NovelAutoBuyPageReply, error) {
	var out NovelAutoBuyPageReply
	pattern := "/v1/novel/auto/buy/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelautobuy.v1.NovelAutoBuy/GetPageNovelAutoBuy"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelAutoBuyHTTPClientImpl) UpdateNovelAutoBuy(ctx context.Context, in *NovelAutoBuyUpdateReq, opts ...http.CallOption) (*NovelAutoBuyUpdateReply, error) {
	var out NovelAutoBuyUpdateReply
	pattern := "/v1/novel/auto/buy/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/novelautobuy.v1.NovelAutoBuy/UpdateNovelAutoBuy"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}