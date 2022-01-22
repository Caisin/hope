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

type NovelMsgHTTPServer interface {
	BatchDeleteNovelMsg(context.Context, *NovelMsgBatchDeleteReq) (*NovelMsgDeleteReply, error)
	CreateNovelMsg(context.Context, *NovelMsgCreateReq) (*NovelMsgCreateReply, error)
	DeleteNovelMsg(context.Context, *NovelMsgDeleteReq) (*NovelMsgDeleteReply, error)
	GetNovelMsg(context.Context, *NovelMsgReq) (*NovelMsgReply, error)
	GetPageNovelMsg(context.Context, *NovelMsgPageReq) (*NovelMsgPageReply, error)
	UpdateNovelMsg(context.Context, *NovelMsgUpdateReq) (*NovelMsgUpdateReply, error)
}

func RegisterNovelMsgHTTPServer(s *http.Server, srv NovelMsgHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/novel/msg/page", _NovelMsg_GetPageNovelMsg0_HTTP_Handler(srv))
	r.GET("/v1/novel/msg/{id}", _NovelMsg_GetNovelMsg0_HTTP_Handler(srv))
	r.PUT("/v1/novel/msg/{id}", _NovelMsg_UpdateNovelMsg0_HTTP_Handler(srv))
	r.POST("/v1/novel/msg", _NovelMsg_CreateNovelMsg0_HTTP_Handler(srv))
	r.DELETE("/v1/novel/msg/{id}", _NovelMsg_DeleteNovelMsg0_HTTP_Handler(srv))
	r.DELETE("/v1/novel/msg", _NovelMsg_BatchDeleteNovelMsg0_HTTP_Handler(srv))
}

func _NovelMsg_GetPageNovelMsg0_HTTP_Handler(srv NovelMsgHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelMsgPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelmsg.v1.NovelMsg/GetPageNovelMsg")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPageNovelMsg(ctx, req.(*NovelMsgPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelMsgPageReply)
		return ctx.Result(200, reply)
	}
}

func _NovelMsg_GetNovelMsg0_HTTP_Handler(srv NovelMsgHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelMsgReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelmsg.v1.NovelMsg/GetNovelMsg")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNovelMsg(ctx, req.(*NovelMsgReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelMsgReply)
		return ctx.Result(200, reply)
	}
}

func _NovelMsg_UpdateNovelMsg0_HTTP_Handler(srv NovelMsgHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelMsgUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelmsg.v1.NovelMsg/UpdateNovelMsg")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateNovelMsg(ctx, req.(*NovelMsgUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelMsgUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _NovelMsg_CreateNovelMsg0_HTTP_Handler(srv NovelMsgHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelMsgCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelmsg.v1.NovelMsg/CreateNovelMsg")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateNovelMsg(ctx, req.(*NovelMsgCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelMsgCreateReply)
		return ctx.Result(200, reply)
	}
}

func _NovelMsg_DeleteNovelMsg0_HTTP_Handler(srv NovelMsgHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelMsgDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelmsg.v1.NovelMsg/DeleteNovelMsg")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteNovelMsg(ctx, req.(*NovelMsgDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelMsgDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _NovelMsg_BatchDeleteNovelMsg0_HTTP_Handler(srv NovelMsgHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelMsgBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelmsg.v1.NovelMsg/BatchDeleteNovelMsg")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteNovelMsg(ctx, req.(*NovelMsgBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelMsgDeleteReply)
		return ctx.Result(200, reply)
	}
}

type NovelMsgHTTPClient interface {
	BatchDeleteNovelMsg(ctx context.Context, req *NovelMsgBatchDeleteReq, opts ...http.CallOption) (rsp *NovelMsgDeleteReply, err error)
	CreateNovelMsg(ctx context.Context, req *NovelMsgCreateReq, opts ...http.CallOption) (rsp *NovelMsgCreateReply, err error)
	DeleteNovelMsg(ctx context.Context, req *NovelMsgDeleteReq, opts ...http.CallOption) (rsp *NovelMsgDeleteReply, err error)
	GetNovelMsg(ctx context.Context, req *NovelMsgReq, opts ...http.CallOption) (rsp *NovelMsgReply, err error)
	GetPageNovelMsg(ctx context.Context, req *NovelMsgPageReq, opts ...http.CallOption) (rsp *NovelMsgPageReply, err error)
	UpdateNovelMsg(ctx context.Context, req *NovelMsgUpdateReq, opts ...http.CallOption) (rsp *NovelMsgUpdateReply, err error)
}

type NovelMsgHTTPClientImpl struct {
	cc *http.Client
}

func NewNovelMsgHTTPClient(client *http.Client) NovelMsgHTTPClient {
	return &NovelMsgHTTPClientImpl{client}
}

func (c *NovelMsgHTTPClientImpl) BatchDeleteNovelMsg(ctx context.Context, in *NovelMsgBatchDeleteReq, opts ...http.CallOption) (*NovelMsgDeleteReply, error) {
	var out NovelMsgDeleteReply
	pattern := "/v1/novel/msg"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelmsg.v1.NovelMsg/BatchDeleteNovelMsg"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelMsgHTTPClientImpl) CreateNovelMsg(ctx context.Context, in *NovelMsgCreateReq, opts ...http.CallOption) (*NovelMsgCreateReply, error) {
	var out NovelMsgCreateReply
	pattern := "/v1/novel/msg"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/novelmsg.v1.NovelMsg/CreateNovelMsg"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelMsgHTTPClientImpl) DeleteNovelMsg(ctx context.Context, in *NovelMsgDeleteReq, opts ...http.CallOption) (*NovelMsgDeleteReply, error) {
	var out NovelMsgDeleteReply
	pattern := "/v1/novel/msg/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelmsg.v1.NovelMsg/DeleteNovelMsg"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelMsgHTTPClientImpl) GetNovelMsg(ctx context.Context, in *NovelMsgReq, opts ...http.CallOption) (*NovelMsgReply, error) {
	var out NovelMsgReply
	pattern := "/v1/novel/msg/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelmsg.v1.NovelMsg/GetNovelMsg"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelMsgHTTPClientImpl) GetPageNovelMsg(ctx context.Context, in *NovelMsgPageReq, opts ...http.CallOption) (*NovelMsgPageReply, error) {
	var out NovelMsgPageReply
	pattern := "/v1/novel/msg/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelmsg.v1.NovelMsg/GetPageNovelMsg"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelMsgHTTPClientImpl) UpdateNovelMsg(ctx context.Context, in *NovelMsgUpdateReq, opts ...http.CallOption) (*NovelMsgUpdateReply, error) {
	var out NovelMsgUpdateReply
	pattern := "/v1/novel/msg/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/novelmsg.v1.NovelMsg/UpdateNovelMsg"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
