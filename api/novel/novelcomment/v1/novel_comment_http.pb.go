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

type NovelCommentHTTPServer interface {
	BatchDeleteNovelComment(context.Context, *NovelCommentBatchDeleteReq) (*NovelCommentDeleteReply, error)
	CreateNovelComment(context.Context, *NovelCommentCreateReq) (*NovelCommentCreateReply, error)
	DeleteNovelComment(context.Context, *NovelCommentDeleteReq) (*NovelCommentDeleteReply, error)
	GetNovelComment(context.Context, *NovelCommentReq) (*NovelCommentReply, error)
	GetPageNovelComment(context.Context, *NovelCommentPageReq) (*NovelCommentPageReply, error)
	UpdateNovelComment(context.Context, *NovelCommentUpdateReq) (*NovelCommentUpdateReply, error)
}

func RegisterNovelCommentHTTPServer(s *http.Server, srv NovelCommentHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/novel/comment/page", _NovelComment_GetPageNovelComment0_HTTP_Handler(srv))
	r.GET("/v1/novel/comment/{id}", _NovelComment_GetNovelComment0_HTTP_Handler(srv))
	r.PUT("/v1/novel/comment/{id}", _NovelComment_UpdateNovelComment0_HTTP_Handler(srv))
	r.POST("/v1/novel/comment", _NovelComment_CreateNovelComment0_HTTP_Handler(srv))
	r.DELETE("/v1/novel/comment/{id}", _NovelComment_DeleteNovelComment0_HTTP_Handler(srv))
	r.DELETE("/v1/novel/comment", _NovelComment_BatchDeleteNovelComment0_HTTP_Handler(srv))
}

func _NovelComment_GetPageNovelComment0_HTTP_Handler(srv NovelCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelCommentPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelcomment.v1.NovelComment/GetPageNovelComment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPageNovelComment(ctx, req.(*NovelCommentPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelCommentPageReply)
		return ctx.Result(200, reply)
	}
}

func _NovelComment_GetNovelComment0_HTTP_Handler(srv NovelCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelCommentReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelcomment.v1.NovelComment/GetNovelComment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNovelComment(ctx, req.(*NovelCommentReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelCommentReply)
		return ctx.Result(200, reply)
	}
}

func _NovelComment_UpdateNovelComment0_HTTP_Handler(srv NovelCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelCommentUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelcomment.v1.NovelComment/UpdateNovelComment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateNovelComment(ctx, req.(*NovelCommentUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelCommentUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _NovelComment_CreateNovelComment0_HTTP_Handler(srv NovelCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelCommentCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelcomment.v1.NovelComment/CreateNovelComment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateNovelComment(ctx, req.(*NovelCommentCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelCommentCreateReply)
		return ctx.Result(200, reply)
	}
}

func _NovelComment_DeleteNovelComment0_HTTP_Handler(srv NovelCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelCommentDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelcomment.v1.NovelComment/DeleteNovelComment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteNovelComment(ctx, req.(*NovelCommentDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelCommentDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _NovelComment_BatchDeleteNovelComment0_HTTP_Handler(srv NovelCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelCommentBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelcomment.v1.NovelComment/BatchDeleteNovelComment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteNovelComment(ctx, req.(*NovelCommentBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelCommentDeleteReply)
		return ctx.Result(200, reply)
	}
}

type NovelCommentHTTPClient interface {
	BatchDeleteNovelComment(ctx context.Context, req *NovelCommentBatchDeleteReq, opts ...http.CallOption) (rsp *NovelCommentDeleteReply, err error)
	CreateNovelComment(ctx context.Context, req *NovelCommentCreateReq, opts ...http.CallOption) (rsp *NovelCommentCreateReply, err error)
	DeleteNovelComment(ctx context.Context, req *NovelCommentDeleteReq, opts ...http.CallOption) (rsp *NovelCommentDeleteReply, err error)
	GetNovelComment(ctx context.Context, req *NovelCommentReq, opts ...http.CallOption) (rsp *NovelCommentReply, err error)
	GetPageNovelComment(ctx context.Context, req *NovelCommentPageReq, opts ...http.CallOption) (rsp *NovelCommentPageReply, err error)
	UpdateNovelComment(ctx context.Context, req *NovelCommentUpdateReq, opts ...http.CallOption) (rsp *NovelCommentUpdateReply, err error)
}

type NovelCommentHTTPClientImpl struct {
	cc *http.Client
}

func NewNovelCommentHTTPClient(client *http.Client) NovelCommentHTTPClient {
	return &NovelCommentHTTPClientImpl{client}
}

func (c *NovelCommentHTTPClientImpl) BatchDeleteNovelComment(ctx context.Context, in *NovelCommentBatchDeleteReq, opts ...http.CallOption) (*NovelCommentDeleteReply, error) {
	var out NovelCommentDeleteReply
	pattern := "/v1/novel/comment"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelcomment.v1.NovelComment/BatchDeleteNovelComment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelCommentHTTPClientImpl) CreateNovelComment(ctx context.Context, in *NovelCommentCreateReq, opts ...http.CallOption) (*NovelCommentCreateReply, error) {
	var out NovelCommentCreateReply
	pattern := "/v1/novel/comment"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/novelcomment.v1.NovelComment/CreateNovelComment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelCommentHTTPClientImpl) DeleteNovelComment(ctx context.Context, in *NovelCommentDeleteReq, opts ...http.CallOption) (*NovelCommentDeleteReply, error) {
	var out NovelCommentDeleteReply
	pattern := "/v1/novel/comment/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelcomment.v1.NovelComment/DeleteNovelComment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelCommentHTTPClientImpl) GetNovelComment(ctx context.Context, in *NovelCommentReq, opts ...http.CallOption) (*NovelCommentReply, error) {
	var out NovelCommentReply
	pattern := "/v1/novel/comment/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelcomment.v1.NovelComment/GetNovelComment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelCommentHTTPClientImpl) GetPageNovelComment(ctx context.Context, in *NovelCommentPageReq, opts ...http.CallOption) (*NovelCommentPageReply, error) {
	var out NovelCommentPageReply
	pattern := "/v1/novel/comment/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelcomment.v1.NovelComment/GetPageNovelComment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelCommentHTTPClientImpl) UpdateNovelComment(ctx context.Context, in *NovelCommentUpdateReq, opts ...http.CallOption) (*NovelCommentUpdateReply, error) {
	var out NovelCommentUpdateReply
	pattern := "/v1/novel/comment/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/novelcomment.v1.NovelComment/UpdateNovelComment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
