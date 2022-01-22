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

type NovelBuyChapterRecordHTTPServer interface {
	BatchDeleteNovelBuyChapterRecord(context.Context, *NovelBuyChapterRecordBatchDeleteReq) (*NovelBuyChapterRecordDeleteReply, error)
	CreateNovelBuyChapterRecord(context.Context, *NovelBuyChapterRecordCreateReq) (*NovelBuyChapterRecordCreateReply, error)
	DeleteNovelBuyChapterRecord(context.Context, *NovelBuyChapterRecordDeleteReq) (*NovelBuyChapterRecordDeleteReply, error)
	GetNovelBuyChapterRecord(context.Context, *NovelBuyChapterRecordReq) (*NovelBuyChapterRecordReply, error)
	GetPageNovelBuyChapterRecord(context.Context, *NovelBuyChapterRecordPageReq) (*NovelBuyChapterRecordPageReply, error)
	UpdateNovelBuyChapterRecord(context.Context, *NovelBuyChapterRecordUpdateReq) (*NovelBuyChapterRecordUpdateReply, error)
}

func RegisterNovelBuyChapterRecordHTTPServer(s *http.Server, srv NovelBuyChapterRecordHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/novel/buy/chapter/record/page", _NovelBuyChapterRecord_GetPageNovelBuyChapterRecord0_HTTP_Handler(srv))
	r.GET("/v1/novel/buy/chapter/record/{id}", _NovelBuyChapterRecord_GetNovelBuyChapterRecord0_HTTP_Handler(srv))
	r.PUT("/v1/novel/buy/chapter/record/{id}", _NovelBuyChapterRecord_UpdateNovelBuyChapterRecord0_HTTP_Handler(srv))
	r.POST("/v1/novel/buy/chapter/record", _NovelBuyChapterRecord_CreateNovelBuyChapterRecord0_HTTP_Handler(srv))
	r.DELETE("/v1/novel/buy/chapter/record/{id}", _NovelBuyChapterRecord_DeleteNovelBuyChapterRecord0_HTTP_Handler(srv))
	r.DELETE("/v1/novel/buy/chapter/record", _NovelBuyChapterRecord_BatchDeleteNovelBuyChapterRecord0_HTTP_Handler(srv))
}

func _NovelBuyChapterRecord_GetPageNovelBuyChapterRecord0_HTTP_Handler(srv NovelBuyChapterRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelBuyChapterRecordPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelbuychapterrecord.v1.NovelBuyChapterRecord/GetPageNovelBuyChapterRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPageNovelBuyChapterRecord(ctx, req.(*NovelBuyChapterRecordPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelBuyChapterRecordPageReply)
		return ctx.Result(200, reply)
	}
}

func _NovelBuyChapterRecord_GetNovelBuyChapterRecord0_HTTP_Handler(srv NovelBuyChapterRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelBuyChapterRecordReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelbuychapterrecord.v1.NovelBuyChapterRecord/GetNovelBuyChapterRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNovelBuyChapterRecord(ctx, req.(*NovelBuyChapterRecordReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelBuyChapterRecordReply)
		return ctx.Result(200, reply)
	}
}

func _NovelBuyChapterRecord_UpdateNovelBuyChapterRecord0_HTTP_Handler(srv NovelBuyChapterRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelBuyChapterRecordUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelbuychapterrecord.v1.NovelBuyChapterRecord/UpdateNovelBuyChapterRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateNovelBuyChapterRecord(ctx, req.(*NovelBuyChapterRecordUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelBuyChapterRecordUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _NovelBuyChapterRecord_CreateNovelBuyChapterRecord0_HTTP_Handler(srv NovelBuyChapterRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelBuyChapterRecordCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelbuychapterrecord.v1.NovelBuyChapterRecord/CreateNovelBuyChapterRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateNovelBuyChapterRecord(ctx, req.(*NovelBuyChapterRecordCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelBuyChapterRecordCreateReply)
		return ctx.Result(200, reply)
	}
}

func _NovelBuyChapterRecord_DeleteNovelBuyChapterRecord0_HTTP_Handler(srv NovelBuyChapterRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelBuyChapterRecordDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelbuychapterrecord.v1.NovelBuyChapterRecord/DeleteNovelBuyChapterRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteNovelBuyChapterRecord(ctx, req.(*NovelBuyChapterRecordDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelBuyChapterRecordDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _NovelBuyChapterRecord_BatchDeleteNovelBuyChapterRecord0_HTTP_Handler(srv NovelBuyChapterRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NovelBuyChapterRecordBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/novelbuychapterrecord.v1.NovelBuyChapterRecord/BatchDeleteNovelBuyChapterRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteNovelBuyChapterRecord(ctx, req.(*NovelBuyChapterRecordBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NovelBuyChapterRecordDeleteReply)
		return ctx.Result(200, reply)
	}
}

type NovelBuyChapterRecordHTTPClient interface {
	BatchDeleteNovelBuyChapterRecord(ctx context.Context, req *NovelBuyChapterRecordBatchDeleteReq, opts ...http.CallOption) (rsp *NovelBuyChapterRecordDeleteReply, err error)
	CreateNovelBuyChapterRecord(ctx context.Context, req *NovelBuyChapterRecordCreateReq, opts ...http.CallOption) (rsp *NovelBuyChapterRecordCreateReply, err error)
	DeleteNovelBuyChapterRecord(ctx context.Context, req *NovelBuyChapterRecordDeleteReq, opts ...http.CallOption) (rsp *NovelBuyChapterRecordDeleteReply, err error)
	GetNovelBuyChapterRecord(ctx context.Context, req *NovelBuyChapterRecordReq, opts ...http.CallOption) (rsp *NovelBuyChapterRecordReply, err error)
	GetPageNovelBuyChapterRecord(ctx context.Context, req *NovelBuyChapterRecordPageReq, opts ...http.CallOption) (rsp *NovelBuyChapterRecordPageReply, err error)
	UpdateNovelBuyChapterRecord(ctx context.Context, req *NovelBuyChapterRecordUpdateReq, opts ...http.CallOption) (rsp *NovelBuyChapterRecordUpdateReply, err error)
}

type NovelBuyChapterRecordHTTPClientImpl struct {
	cc *http.Client
}

func NewNovelBuyChapterRecordHTTPClient(client *http.Client) NovelBuyChapterRecordHTTPClient {
	return &NovelBuyChapterRecordHTTPClientImpl{client}
}

func (c *NovelBuyChapterRecordHTTPClientImpl) BatchDeleteNovelBuyChapterRecord(ctx context.Context, in *NovelBuyChapterRecordBatchDeleteReq, opts ...http.CallOption) (*NovelBuyChapterRecordDeleteReply, error) {
	var out NovelBuyChapterRecordDeleteReply
	pattern := "/v1/novel/buy/chapter/record"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelbuychapterrecord.v1.NovelBuyChapterRecord/BatchDeleteNovelBuyChapterRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelBuyChapterRecordHTTPClientImpl) CreateNovelBuyChapterRecord(ctx context.Context, in *NovelBuyChapterRecordCreateReq, opts ...http.CallOption) (*NovelBuyChapterRecordCreateReply, error) {
	var out NovelBuyChapterRecordCreateReply
	pattern := "/v1/novel/buy/chapter/record"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/novelbuychapterrecord.v1.NovelBuyChapterRecord/CreateNovelBuyChapterRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelBuyChapterRecordHTTPClientImpl) DeleteNovelBuyChapterRecord(ctx context.Context, in *NovelBuyChapterRecordDeleteReq, opts ...http.CallOption) (*NovelBuyChapterRecordDeleteReply, error) {
	var out NovelBuyChapterRecordDeleteReply
	pattern := "/v1/novel/buy/chapter/record/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelbuychapterrecord.v1.NovelBuyChapterRecord/DeleteNovelBuyChapterRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelBuyChapterRecordHTTPClientImpl) GetNovelBuyChapterRecord(ctx context.Context, in *NovelBuyChapterRecordReq, opts ...http.CallOption) (*NovelBuyChapterRecordReply, error) {
	var out NovelBuyChapterRecordReply
	pattern := "/v1/novel/buy/chapter/record/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelbuychapterrecord.v1.NovelBuyChapterRecord/GetNovelBuyChapterRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelBuyChapterRecordHTTPClientImpl) GetPageNovelBuyChapterRecord(ctx context.Context, in *NovelBuyChapterRecordPageReq, opts ...http.CallOption) (*NovelBuyChapterRecordPageReply, error) {
	var out NovelBuyChapterRecordPageReply
	pattern := "/v1/novel/buy/chapter/record/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/novelbuychapterrecord.v1.NovelBuyChapterRecord/GetPageNovelBuyChapterRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NovelBuyChapterRecordHTTPClientImpl) UpdateNovelBuyChapterRecord(ctx context.Context, in *NovelBuyChapterRecordUpdateReq, opts ...http.CallOption) (*NovelBuyChapterRecordUpdateReply, error) {
	var out NovelBuyChapterRecordUpdateReply
	pattern := "/v1/novel/buy/chapter/record/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/novelbuychapterrecord.v1.NovelBuyChapterRecord/UpdateNovelBuyChapterRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
