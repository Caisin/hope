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

type NovelBuyRecordHTTPServer interface {
	BatchDeleteNovelBuyRecord(context.Context, *NovelBuyRecordBatchDeleteReq) (*NovelBuyRecordDeleteReply, error)
	CreateNovelBuyRecord(context.Context, *NovelBuyRecordCreateReq) (*NovelBuyRecordCreateReply, error)
	DeleteNovelBuyRecord(context.Context, *NovelBuyRecordDeleteReq) (*NovelBuyRecordDeleteReply, error)
	GetNovelBuyRecord(context.Context, *NovelBuyRecordReq) (*NovelBuyRecordReply, error)
	GetPageNovelBuyRecord(context.Context, *NovelBuyRecordPageReq) (*NovelBuyRecordPageReply, error)
	UpdateNovelBuyRecord(context.Context, *NovelBuyRecordUpdateReq) (*NovelBuyRecordUpdateReply, error)
}

func RegisterNovelBuyRecordHTTPServer(s *http.Server, srv NovelBuyRecordHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/novel/buy/record/page", _NovelBuyRecord_GetPageNovelBuyRecord0_HTTP_Handler(srv))
	r.GET("/v1/novel/buy/record/{id}", _NovelBuyRecord_GetNovelBuyRecord0_HTTP_Handler(srv))
	r.PUT("/v1/novel/buy/record/{id}", _NovelBuyRecord_UpdateNovelBuyRecord0_HTTP_Handler(srv))
	r.POST("/v1/novel/buy/record", _NovelBuyRecord_CreateNovelBuyRecord0_HTTP_Handler(srv))
	r.DELETE("/v1/novel/buy/record/{id}", _NovelBuyRecord_DeleteNovelBuyRecord0_HTTP_Handler(srv))
	r.DELETE("/v1/novel/buy/record", _NovelBuyRecord_BatchDeleteNovelBuyRecord0_HTTP_Handler(srv))
}

func _NovelBuyRecord_GetPageNovelBuyRecord0_HTTP_Handler(srv NovelBuyRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelBuyRecordPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelbuyrecord.v1.NovelBuyRecord/GetPageNovelBuyRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPageNovelBuyRecord(ctx, req.(*NovelBuyRecordPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelBuyRecordPageReply)
		return ctx.Result(200, reply)
	}
}

func _NovelBuyRecord_GetNovelBuyRecord0_HTTP_Handler(srv NovelBuyRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelBuyRecordReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelbuyrecord.v1.NovelBuyRecord/GetNovelBuyRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNovelBuyRecord(ctx, req.(*NovelBuyRecordReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelBuyRecordReply)
		return ctx.Result(200, reply)
	}
}

func _NovelBuyRecord_UpdateNovelBuyRecord0_HTTP_Handler(srv NovelBuyRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelBuyRecordUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelbuyrecord.v1.NovelBuyRecord/UpdateNovelBuyRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateNovelBuyRecord(ctx, req.(*NovelBuyRecordUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelBuyRecordUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _NovelBuyRecord_CreateNovelBuyRecord0_HTTP_Handler(srv NovelBuyRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelBuyRecordCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelbuyrecord.v1.NovelBuyRecord/CreateNovelBuyRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateNovelBuyRecord(ctx, req.(*NovelBuyRecordCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelBuyRecordCreateReply)
		return ctx.Result(200, reply)
	}
}

func _NovelBuyRecord_DeleteNovelBuyRecord0_HTTP_Handler(srv NovelBuyRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelBuyRecordDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelbuyrecord.v1.NovelBuyRecord/DeleteNovelBuyRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteNovelBuyRecord(ctx, req.(*NovelBuyRecordDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelBuyRecordDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _NovelBuyRecord_BatchDeleteNovelBuyRecord0_HTTP_Handler(srv NovelBuyRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelBuyRecordBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelbuyrecord.v1.NovelBuyRecord/BatchDeleteNovelBuyRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteNovelBuyRecord(ctx, req.(*NovelBuyRecordBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelBuyRecordDeleteReply)
		return ctx.Result(200, reply)
	}
}

type NovelBuyRecordHTTPClient interface {
	BatchDeleteNovelBuyRecord(ctx context.Context, req *NovelBuyRecordBatchDeleteReq, opts ...http.CallOption) (rsp *NovelBuyRecordDeleteReply, err error)
	CreateNovelBuyRecord(ctx context.Context, req *NovelBuyRecordCreateReq, opts ...http.CallOption) (rsp *NovelBuyRecordCreateReply, err error)
	DeleteNovelBuyRecord(ctx context.Context, req *NovelBuyRecordDeleteReq, opts ...http.CallOption) (rsp *NovelBuyRecordDeleteReply, err error)
	GetNovelBuyRecord(ctx context.Context, req *NovelBuyRecordReq, opts ...http.CallOption) (rsp *NovelBuyRecordReply, err error)
	GetPageNovelBuyRecord(ctx context.Context, req *NovelBuyRecordPageReq, opts ...http.CallOption) (rsp *NovelBuyRecordPageReply, err error)
	UpdateNovelBuyRecord(ctx context.Context, req *NovelBuyRecordUpdateReq, opts ...http.CallOption) (rsp *NovelBuyRecordUpdateReply, err error)
}

type NovelBuyRecordHTTPClientImpl struct {
	cc *http.Client
}

func NewNovelBuyRecordHTTPClient(client *http.Client) NovelBuyRecordHTTPClient {
	return &NovelBuyRecordHTTPClientImpl{client}
}

func (c *NovelBuyRecordHTTPClientImpl) BatchDeleteNovelBuyRecord(ctx context.Context, in *NovelBuyRecordBatchDeleteReq, opts ...http.CallOption) (*NovelBuyRecordDeleteReply, error) {
	var out NovelBuyRecordDeleteReply
	pattern := "/v1/novel/buy/record"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelbuyrecord.v1.NovelBuyRecord/BatchDeleteNovelBuyRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelBuyRecordHTTPClientImpl) CreateNovelBuyRecord(ctx context.Context, in *NovelBuyRecordCreateReq, opts ...http.CallOption) (*NovelBuyRecordCreateReply, error) {
	var out NovelBuyRecordCreateReply
	pattern := "/v1/novel/buy/record"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/novelbuyrecord.v1.NovelBuyRecord/CreateNovelBuyRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelBuyRecordHTTPClientImpl) DeleteNovelBuyRecord(ctx context.Context, in *NovelBuyRecordDeleteReq, opts ...http.CallOption) (*NovelBuyRecordDeleteReply, error) {
	var out NovelBuyRecordDeleteReply
	pattern := "/v1/novel/buy/record/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelbuyrecord.v1.NovelBuyRecord/DeleteNovelBuyRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelBuyRecordHTTPClientImpl) GetNovelBuyRecord(ctx context.Context, in *NovelBuyRecordReq, opts ...http.CallOption) (*NovelBuyRecordReply, error) {
	var out NovelBuyRecordReply
	pattern := "/v1/novel/buy/record/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelbuyrecord.v1.NovelBuyRecord/GetNovelBuyRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelBuyRecordHTTPClientImpl) GetPageNovelBuyRecord(ctx context.Context, in *NovelBuyRecordPageReq, opts ...http.CallOption) (*NovelBuyRecordPageReply, error) {
	var out NovelBuyRecordPageReply
	pattern := "/v1/novel/buy/record/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelbuyrecord.v1.NovelBuyRecord/GetPageNovelBuyRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelBuyRecordHTTPClientImpl) UpdateNovelBuyRecord(ctx context.Context, in *NovelBuyRecordUpdateReq, opts ...http.CallOption) (*NovelBuyRecordUpdateReply, error) {
	var out NovelBuyRecordUpdateReply
	pattern := "/v1/novel/buy/record/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/novelbuyrecord.v1.NovelBuyRecord/UpdateNovelBuyRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}