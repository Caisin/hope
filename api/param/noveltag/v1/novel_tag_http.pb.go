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

type NovelTagHTTPServer interface {
	BatchDeleteNovelTag(context.Context, *NovelTagBatchDeleteReq) (*NovelTagDeleteReply, error)
	CreateNovelTag(context.Context, *NovelTagCreateReq) (*NovelTagCreateReply, error)
	DeleteNovelTag(context.Context, *NovelTagDeleteReq) (*NovelTagDeleteReply, error)
	GetNovelTag(context.Context, *NovelTagReq) (*NovelTagReply, error)
	GetNovelTagPage(context.Context, *NovelTagPageReq) (*NovelTagPageReply, error)
	UpdateNovelTag(context.Context, *NovelTagUpdateReq) (*NovelTagUpdateReply, error)
}

func RegisterNovelTagHTTPServer(s *http.Server, srv NovelTagHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/novel/tag/page", _NovelTag_GetNovelTagPage0_HTTP_Handler(srv))
	r.GET("/v1/novel/tag/{id}", _NovelTag_GetNovelTag0_HTTP_Handler(srv))
	r.PUT("/v1/novel/tag/{id}", _NovelTag_UpdateNovelTag0_HTTP_Handler(srv))
	r.POST("/v1/novel/tag", _NovelTag_CreateNovelTag0_HTTP_Handler(srv))
	r.DELETE("/v1/novel/tag/{id}", _NovelTag_DeleteNovelTag0_HTTP_Handler(srv))
	r.DELETE("/v1/novel/tag", _NovelTag_BatchDeleteNovelTag0_HTTP_Handler(srv))
}

func _NovelTag_GetNovelTagPage0_HTTP_Handler(srv NovelTagHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelTagPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/noveltag.v1.NovelTag/GetNovelTagPage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNovelTagPage(ctx, req.(*NovelTagPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelTagPageReply)
		return ctx.Result(200, reply)
	}
}

func _NovelTag_GetNovelTag0_HTTP_Handler(srv NovelTagHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelTagReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/noveltag.v1.NovelTag/GetNovelTag")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNovelTag(ctx, req.(*NovelTagReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelTagReply)
		return ctx.Result(200, reply)
	}
}

func _NovelTag_UpdateNovelTag0_HTTP_Handler(srv NovelTagHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelTagUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/noveltag.v1.NovelTag/UpdateNovelTag")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateNovelTag(ctx, req.(*NovelTagUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelTagUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _NovelTag_CreateNovelTag0_HTTP_Handler(srv NovelTagHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelTagCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/noveltag.v1.NovelTag/CreateNovelTag")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateNovelTag(ctx, req.(*NovelTagCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelTagCreateReply)
		return ctx.Result(200, reply)
	}
}

func _NovelTag_DeleteNovelTag0_HTTP_Handler(srv NovelTagHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelTagDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/noveltag.v1.NovelTag/DeleteNovelTag")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteNovelTag(ctx, req.(*NovelTagDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelTagDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _NovelTag_BatchDeleteNovelTag0_HTTP_Handler(srv NovelTagHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelTagBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/noveltag.v1.NovelTag/BatchDeleteNovelTag")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteNovelTag(ctx, req.(*NovelTagBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelTagDeleteReply)
		return ctx.Result(200, reply)
	}
}

type NovelTagHTTPClient interface {
	BatchDeleteNovelTag(ctx context.Context, req *NovelTagBatchDeleteReq, opts ...http.CallOption) (rsp *NovelTagDeleteReply, err error)
	CreateNovelTag(ctx context.Context, req *NovelTagCreateReq, opts ...http.CallOption) (rsp *NovelTagCreateReply, err error)
	DeleteNovelTag(ctx context.Context, req *NovelTagDeleteReq, opts ...http.CallOption) (rsp *NovelTagDeleteReply, err error)
	GetNovelTag(ctx context.Context, req *NovelTagReq, opts ...http.CallOption) (rsp *NovelTagReply, err error)
	GetNovelTagPage(ctx context.Context, req *NovelTagPageReq, opts ...http.CallOption) (rsp *NovelTagPageReply, err error)
	UpdateNovelTag(ctx context.Context, req *NovelTagUpdateReq, opts ...http.CallOption) (rsp *NovelTagUpdateReply, err error)
}

type NovelTagHTTPClientImpl struct {
	cc *http.Client
}

func NewNovelTagHTTPClient(client *http.Client) NovelTagHTTPClient {
	return &NovelTagHTTPClientImpl{client}
}

func (c *NovelTagHTTPClientImpl) BatchDeleteNovelTag(ctx context.Context, in *NovelTagBatchDeleteReq, opts ...http.CallOption) (*NovelTagDeleteReply, error) {
	var out NovelTagDeleteReply
	pattern := "/v1/novel/tag"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/noveltag.v1.NovelTag/BatchDeleteNovelTag"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelTagHTTPClientImpl) CreateNovelTag(ctx context.Context, in *NovelTagCreateReq, opts ...http.CallOption) (*NovelTagCreateReply, error) {
	var out NovelTagCreateReply
	pattern := "/v1/novel/tag"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/noveltag.v1.NovelTag/CreateNovelTag"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelTagHTTPClientImpl) DeleteNovelTag(ctx context.Context, in *NovelTagDeleteReq, opts ...http.CallOption) (*NovelTagDeleteReply, error) {
	var out NovelTagDeleteReply
	pattern := "/v1/novel/tag/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/noveltag.v1.NovelTag/DeleteNovelTag"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelTagHTTPClientImpl) GetNovelTag(ctx context.Context, in *NovelTagReq, opts ...http.CallOption) (*NovelTagReply, error) {
	var out NovelTagReply
	pattern := "/v1/novel/tag/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/noveltag.v1.NovelTag/GetNovelTag"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelTagHTTPClientImpl) GetNovelTagPage(ctx context.Context, in *NovelTagPageReq, opts ...http.CallOption) (*NovelTagPageReply, error) {
	var out NovelTagPageReply
	pattern := "/v1/novel/tag/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/noveltag.v1.NovelTag/GetNovelTagPage"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelTagHTTPClientImpl) UpdateNovelTag(ctx context.Context, in *NovelTagUpdateReq, opts ...http.CallOption) (*NovelTagUpdateReply, error) {
	var out NovelTagUpdateReply
	pattern := "/v1/novel/tag/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/noveltag.v1.NovelTag/UpdateNovelTag"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
