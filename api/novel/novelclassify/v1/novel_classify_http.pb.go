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

type NovelClassifyHTTPServer interface {
	BatchDeleteNovelClassify(context.Context, *NovelClassifyBatchDeleteReq) (*NovelClassifyDeleteReply, error)
	CreateNovelClassify(context.Context, *NovelClassifyCreateReq) (*NovelClassifyCreateReply, error)
	DeleteNovelClassify(context.Context, *NovelClassifyDeleteReq) (*NovelClassifyDeleteReply, error)
	GetNovelClassify(context.Context, *NovelClassifyReq) (*NovelClassifyReply, error)
	GetNovelClassifyPage(context.Context, *NovelClassifyPageReq) (*NovelClassifyPageReply, error)
	UpdateNovelClassify(context.Context, *NovelClassifyUpdateReq) (*NovelClassifyUpdateReply, error)
}

func RegisterNovelClassifyHTTPServer(s *http.Server, srv NovelClassifyHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/novel/classify/page", _NovelClassify_GetNovelClassifyPage0_HTTP_Handler(srv))
	r.GET("/v1/novel/classify/{id}", _NovelClassify_GetNovelClassify0_HTTP_Handler(srv))
	r.PUT("/v1/novel/classify/{id}", _NovelClassify_UpdateNovelClassify0_HTTP_Handler(srv))
	r.POST("/v1/novel/classify", _NovelClassify_CreateNovelClassify0_HTTP_Handler(srv))
	r.DELETE("/v1/novel/classify/{id}", _NovelClassify_DeleteNovelClassify0_HTTP_Handler(srv))
	r.DELETE("/v1/novel/classify", _NovelClassify_BatchDeleteNovelClassify0_HTTP_Handler(srv))
}

func _NovelClassify_GetNovelClassifyPage0_HTTP_Handler(srv NovelClassifyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelClassifyPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelclassify.v1.NovelClassify/GetNovelClassifyPage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNovelClassifyPage(ctx, req.(*NovelClassifyPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelClassifyPageReply)
		return ctx.Result(200, reply)
	}
}

func _NovelClassify_GetNovelClassify0_HTTP_Handler(srv NovelClassifyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelClassifyReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelclassify.v1.NovelClassify/GetNovelClassify")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNovelClassify(ctx, req.(*NovelClassifyReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelClassifyReply)
		return ctx.Result(200, reply)
	}
}

func _NovelClassify_UpdateNovelClassify0_HTTP_Handler(srv NovelClassifyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelClassifyUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelclassify.v1.NovelClassify/UpdateNovelClassify")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateNovelClassify(ctx, req.(*NovelClassifyUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelClassifyUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _NovelClassify_CreateNovelClassify0_HTTP_Handler(srv NovelClassifyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelClassifyCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelclassify.v1.NovelClassify/CreateNovelClassify")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateNovelClassify(ctx, req.(*NovelClassifyCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelClassifyCreateReply)
		return ctx.Result(200, reply)
	}
}

func _NovelClassify_DeleteNovelClassify0_HTTP_Handler(srv NovelClassifyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelClassifyDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelclassify.v1.NovelClassify/DeleteNovelClassify")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteNovelClassify(ctx, req.(*NovelClassifyDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelClassifyDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _NovelClassify_BatchDeleteNovelClassify0_HTTP_Handler(srv NovelClassifyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelClassifyBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelclassify.v1.NovelClassify/BatchDeleteNovelClassify")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteNovelClassify(ctx, req.(*NovelClassifyBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelClassifyDeleteReply)
		return ctx.Result(200, reply)
	}
}

type NovelClassifyHTTPClient interface {
	BatchDeleteNovelClassify(ctx context.Context, req *NovelClassifyBatchDeleteReq, opts ...http.CallOption) (rsp *NovelClassifyDeleteReply, err error)
	CreateNovelClassify(ctx context.Context, req *NovelClassifyCreateReq, opts ...http.CallOption) (rsp *NovelClassifyCreateReply, err error)
	DeleteNovelClassify(ctx context.Context, req *NovelClassifyDeleteReq, opts ...http.CallOption) (rsp *NovelClassifyDeleteReply, err error)
	GetNovelClassify(ctx context.Context, req *NovelClassifyReq, opts ...http.CallOption) (rsp *NovelClassifyReply, err error)
	GetNovelClassifyPage(ctx context.Context, req *NovelClassifyPageReq, opts ...http.CallOption) (rsp *NovelClassifyPageReply, err error)
	UpdateNovelClassify(ctx context.Context, req *NovelClassifyUpdateReq, opts ...http.CallOption) (rsp *NovelClassifyUpdateReply, err error)
}

type NovelClassifyHTTPClientImpl struct {
	cc *http.Client
}

func NewNovelClassifyHTTPClient(client *http.Client) NovelClassifyHTTPClient {
	return &NovelClassifyHTTPClientImpl{client}
}

func (c *NovelClassifyHTTPClientImpl) BatchDeleteNovelClassify(ctx context.Context, in *NovelClassifyBatchDeleteReq, opts ...http.CallOption) (*NovelClassifyDeleteReply, error) {
	var out NovelClassifyDeleteReply
	pattern := "/v1/novel/classify"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelclassify.v1.NovelClassify/BatchDeleteNovelClassify"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelClassifyHTTPClientImpl) CreateNovelClassify(ctx context.Context, in *NovelClassifyCreateReq, opts ...http.CallOption) (*NovelClassifyCreateReply, error) {
	var out NovelClassifyCreateReply
	pattern := "/v1/novel/classify"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/novelclassify.v1.NovelClassify/CreateNovelClassify"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelClassifyHTTPClientImpl) DeleteNovelClassify(ctx context.Context, in *NovelClassifyDeleteReq, opts ...http.CallOption) (*NovelClassifyDeleteReply, error) {
	var out NovelClassifyDeleteReply
	pattern := "/v1/novel/classify/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelclassify.v1.NovelClassify/DeleteNovelClassify"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelClassifyHTTPClientImpl) GetNovelClassify(ctx context.Context, in *NovelClassifyReq, opts ...http.CallOption) (*NovelClassifyReply, error) {
	var out NovelClassifyReply
	pattern := "/v1/novel/classify/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelclassify.v1.NovelClassify/GetNovelClassify"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelClassifyHTTPClientImpl) GetNovelClassifyPage(ctx context.Context, in *NovelClassifyPageReq, opts ...http.CallOption) (*NovelClassifyPageReply, error) {
	var out NovelClassifyPageReply
	pattern := "/v1/novel/classify/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelclassify.v1.NovelClassify/GetNovelClassifyPage"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelClassifyHTTPClientImpl) UpdateNovelClassify(ctx context.Context, in *NovelClassifyUpdateReq, opts ...http.CallOption) (*NovelClassifyUpdateReply, error) {
	var out NovelClassifyUpdateReply
	pattern := "/v1/novel/classify/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/novelclassify.v1.NovelClassify/UpdateNovelClassify"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
