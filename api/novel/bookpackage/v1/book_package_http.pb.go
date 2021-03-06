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

type BookPackageHTTPServer interface {
	BatchDeleteBookPackage(context.Context, *BookPackageBatchDeleteReq) (*BookPackageDeleteReply, error)
	CreateBookPackage(context.Context, *BookPackageCreateReq) (*BookPackageCreateReply, error)
	DeleteBookPackage(context.Context, *BookPackageDeleteReq) (*BookPackageDeleteReply, error)
	GetBookPackage(context.Context, *BookPackageReq) (*BookPackageReply, error)
	GetBookPackagePage(context.Context, *BookPackagePageReq) (*BookPackagePageReply, error)
	UpdateBookPackage(context.Context, *BookPackageUpdateReq) (*BookPackageUpdateReply, error)
}

func RegisterBookPackageHTTPServer(s *http.Server, srv BookPackageHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/book/package/page", _BookPackage_GetBookPackagePage0_HTTP_Handler(srv))
	r.GET("/v1/book/package/{id}", _BookPackage_GetBookPackage0_HTTP_Handler(srv))
	r.PUT("/v1/book/package/{id}", _BookPackage_UpdateBookPackage0_HTTP_Handler(srv))
	r.POST("/v1/book/package", _BookPackage_CreateBookPackage0_HTTP_Handler(srv))
	r.DELETE("/v1/book/package/{id}", _BookPackage_DeleteBookPackage0_HTTP_Handler(srv))
	r.DELETE("/v1/book/package", _BookPackage_BatchDeleteBookPackage0_HTTP_Handler(srv))
}

func _BookPackage_GetBookPackagePage0_HTTP_Handler(srv BookPackageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BookPackagePageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bookpackage.v1.BookPackage/GetBookPackagePage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBookPackagePage(ctx, req.(*BookPackagePageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookPackagePageReply)
		return ctx.Result(200, reply)
	}
}

func _BookPackage_GetBookPackage0_HTTP_Handler(srv BookPackageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BookPackageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bookpackage.v1.BookPackage/GetBookPackage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBookPackage(ctx, req.(*BookPackageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookPackageReply)
		return ctx.Result(200, reply)
	}
}

func _BookPackage_UpdateBookPackage0_HTTP_Handler(srv BookPackageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BookPackageUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bookpackage.v1.BookPackage/UpdateBookPackage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateBookPackage(ctx, req.(*BookPackageUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookPackageUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _BookPackage_CreateBookPackage0_HTTP_Handler(srv BookPackageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BookPackageCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bookpackage.v1.BookPackage/CreateBookPackage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateBookPackage(ctx, req.(*BookPackageCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookPackageCreateReply)
		return ctx.Result(200, reply)
	}
}

func _BookPackage_DeleteBookPackage0_HTTP_Handler(srv BookPackageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BookPackageDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bookpackage.v1.BookPackage/DeleteBookPackage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteBookPackage(ctx, req.(*BookPackageDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookPackageDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _BookPackage_BatchDeleteBookPackage0_HTTP_Handler(srv BookPackageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BookPackageBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bookpackage.v1.BookPackage/BatchDeleteBookPackage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteBookPackage(ctx, req.(*BookPackageBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookPackageDeleteReply)
		return ctx.Result(200, reply)
	}
}

type BookPackageHTTPClient interface {
	BatchDeleteBookPackage(ctx context.Context, req *BookPackageBatchDeleteReq, opts ...http.CallOption) (rsp *BookPackageDeleteReply, err error)
	CreateBookPackage(ctx context.Context, req *BookPackageCreateReq, opts ...http.CallOption) (rsp *BookPackageCreateReply, err error)
	DeleteBookPackage(ctx context.Context, req *BookPackageDeleteReq, opts ...http.CallOption) (rsp *BookPackageDeleteReply, err error)
	GetBookPackage(ctx context.Context, req *BookPackageReq, opts ...http.CallOption) (rsp *BookPackageReply, err error)
	GetBookPackagePage(ctx context.Context, req *BookPackagePageReq, opts ...http.CallOption) (rsp *BookPackagePageReply, err error)
	UpdateBookPackage(ctx context.Context, req *BookPackageUpdateReq, opts ...http.CallOption) (rsp *BookPackageUpdateReply, err error)
}

type BookPackageHTTPClientImpl struct {
	cc *http.Client
}

func NewBookPackageHTTPClient(client *http.Client) BookPackageHTTPClient {
	return &BookPackageHTTPClientImpl{client}
}

func (c *BookPackageHTTPClientImpl) BatchDeleteBookPackage(ctx context.Context, in *BookPackageBatchDeleteReq, opts ...http.CallOption) (*BookPackageDeleteReply, error) {
	var out BookPackageDeleteReply
	pattern := "/v1/book/package"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bookpackage.v1.BookPackage/BatchDeleteBookPackage"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BookPackageHTTPClientImpl) CreateBookPackage(ctx context.Context, in *BookPackageCreateReq, opts ...http.CallOption) (*BookPackageCreateReply, error) {
	var out BookPackageCreateReply
	pattern := "/v1/book/package"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bookpackage.v1.BookPackage/CreateBookPackage"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BookPackageHTTPClientImpl) DeleteBookPackage(ctx context.Context, in *BookPackageDeleteReq, opts ...http.CallOption) (*BookPackageDeleteReply, error) {
	var out BookPackageDeleteReply
	pattern := "/v1/book/package/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bookpackage.v1.BookPackage/DeleteBookPackage"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BookPackageHTTPClientImpl) GetBookPackage(ctx context.Context, in *BookPackageReq, opts ...http.CallOption) (*BookPackageReply, error) {
	var out BookPackageReply
	pattern := "/v1/book/package/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bookpackage.v1.BookPackage/GetBookPackage"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BookPackageHTTPClientImpl) GetBookPackagePage(ctx context.Context, in *BookPackagePageReq, opts ...http.CallOption) (*BookPackagePageReply, error) {
	var out BookPackagePageReply
	pattern := "/v1/book/package/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bookpackage.v1.BookPackage/GetBookPackagePage"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BookPackageHTTPClientImpl) UpdateBookPackage(ctx context.Context, in *BookPackageUpdateReq, opts ...http.CallOption) (*BookPackageUpdateReply, error) {
	var out BookPackageUpdateReply
	pattern := "/v1/book/package/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bookpackage.v1.BookPackage/UpdateBookPackage"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
