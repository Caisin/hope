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

type AssetChangeLogHTTPServer interface {
	BatchDeleteAssetChangeLog(context.Context, *AssetChangeLogBatchDeleteReq) (*AssetChangeLogDeleteReply, error)
	CreateAssetChangeLog(context.Context, *AssetChangeLogCreateReq) (*AssetChangeLogCreateReply, error)
	DeleteAssetChangeLog(context.Context, *AssetChangeLogDeleteReq) (*AssetChangeLogDeleteReply, error)
	GetAssetChangeLog(context.Context, *AssetChangeLogReq) (*AssetChangeLogReply, error)
	GetPageAssetChangeLog(context.Context, *AssetChangeLogPageReq) (*AssetChangeLogPageReply, error)
	UpdateAssetChangeLog(context.Context, *AssetChangeLogUpdateReq) (*AssetChangeLogUpdateReply, error)
}

func RegisterAssetChangeLogHTTPServer(s *http.Server, srv AssetChangeLogHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/asset/change/log/page", _AssetChangeLog_GetPageAssetChangeLog0_HTTP_Handler(srv))
	r.GET("/v1/asset/change/log/{id}", _AssetChangeLog_GetAssetChangeLog0_HTTP_Handler(srv))
	r.PUT("/v1/asset/change/log/{id}", _AssetChangeLog_UpdateAssetChangeLog0_HTTP_Handler(srv))
	r.POST("/v1/asset/change/log", _AssetChangeLog_CreateAssetChangeLog0_HTTP_Handler(srv))
	r.DELETE("/v1/asset/change/log/{id}", _AssetChangeLog_DeleteAssetChangeLog0_HTTP_Handler(srv))
	r.DELETE("/v1/asset/change/log", _AssetChangeLog_BatchDeleteAssetChangeLog0_HTTP_Handler(srv))
}

func _AssetChangeLog_GetPageAssetChangeLog0_HTTP_Handler(srv AssetChangeLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AssetChangeLogPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/assetchangelog.v1.AssetChangeLog/GetPageAssetChangeLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPageAssetChangeLog(ctx, req.(*AssetChangeLogPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AssetChangeLogPageReply)
		return ctx.Result(200, reply)
	}
}

func _AssetChangeLog_GetAssetChangeLog0_HTTP_Handler(srv AssetChangeLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AssetChangeLogReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/assetchangelog.v1.AssetChangeLog/GetAssetChangeLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAssetChangeLog(ctx, req.(*AssetChangeLogReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AssetChangeLogReply)
		return ctx.Result(200, reply)
	}
}

func _AssetChangeLog_UpdateAssetChangeLog0_HTTP_Handler(srv AssetChangeLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AssetChangeLogUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/assetchangelog.v1.AssetChangeLog/UpdateAssetChangeLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateAssetChangeLog(ctx, req.(*AssetChangeLogUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AssetChangeLogUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _AssetChangeLog_CreateAssetChangeLog0_HTTP_Handler(srv AssetChangeLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AssetChangeLogCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/assetchangelog.v1.AssetChangeLog/CreateAssetChangeLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAssetChangeLog(ctx, req.(*AssetChangeLogCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AssetChangeLogCreateReply)
		return ctx.Result(200, reply)
	}
}

func _AssetChangeLog_DeleteAssetChangeLog0_HTTP_Handler(srv AssetChangeLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AssetChangeLogDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/assetchangelog.v1.AssetChangeLog/DeleteAssetChangeLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteAssetChangeLog(ctx, req.(*AssetChangeLogDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AssetChangeLogDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _AssetChangeLog_BatchDeleteAssetChangeLog0_HTTP_Handler(srv AssetChangeLogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AssetChangeLogBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/assetchangelog.v1.AssetChangeLog/BatchDeleteAssetChangeLog")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteAssetChangeLog(ctx, req.(*AssetChangeLogBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AssetChangeLogDeleteReply)
		return ctx.Result(200, reply)
	}
}

type AssetChangeLogHTTPClient interface {
	BatchDeleteAssetChangeLog(ctx context.Context, req *AssetChangeLogBatchDeleteReq, opts ...http.CallOption) (rsp *AssetChangeLogDeleteReply, err error)
	CreateAssetChangeLog(ctx context.Context, req *AssetChangeLogCreateReq, opts ...http.CallOption) (rsp *AssetChangeLogCreateReply, err error)
	DeleteAssetChangeLog(ctx context.Context, req *AssetChangeLogDeleteReq, opts ...http.CallOption) (rsp *AssetChangeLogDeleteReply, err error)
	GetAssetChangeLog(ctx context.Context, req *AssetChangeLogReq, opts ...http.CallOption) (rsp *AssetChangeLogReply, err error)
	GetPageAssetChangeLog(ctx context.Context, req *AssetChangeLogPageReq, opts ...http.CallOption) (rsp *AssetChangeLogPageReply, err error)
	UpdateAssetChangeLog(ctx context.Context, req *AssetChangeLogUpdateReq, opts ...http.CallOption) (rsp *AssetChangeLogUpdateReply, err error)
}

type AssetChangeLogHTTPClientImpl struct {
	cc *http.Client
}

func NewAssetChangeLogHTTPClient(client *http.Client) AssetChangeLogHTTPClient {
	return &AssetChangeLogHTTPClientImpl{client}
}

func (c *AssetChangeLogHTTPClientImpl) BatchDeleteAssetChangeLog(ctx context.Context, in *AssetChangeLogBatchDeleteReq, opts ...http.CallOption) (*AssetChangeLogDeleteReply, error) {
	var out AssetChangeLogDeleteReply
	pattern := "/v1/asset/change/log"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/assetchangelog.v1.AssetChangeLog/BatchDeleteAssetChangeLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AssetChangeLogHTTPClientImpl) CreateAssetChangeLog(ctx context.Context, in *AssetChangeLogCreateReq, opts ...http.CallOption) (*AssetChangeLogCreateReply, error) {
	var out AssetChangeLogCreateReply
	pattern := "/v1/asset/change/log"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/assetchangelog.v1.AssetChangeLog/CreateAssetChangeLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AssetChangeLogHTTPClientImpl) DeleteAssetChangeLog(ctx context.Context, in *AssetChangeLogDeleteReq, opts ...http.CallOption) (*AssetChangeLogDeleteReply, error) {
	var out AssetChangeLogDeleteReply
	pattern := "/v1/asset/change/log/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/assetchangelog.v1.AssetChangeLog/DeleteAssetChangeLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AssetChangeLogHTTPClientImpl) GetAssetChangeLog(ctx context.Context, in *AssetChangeLogReq, opts ...http.CallOption) (*AssetChangeLogReply, error) {
	var out AssetChangeLogReply
	pattern := "/v1/asset/change/log/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/assetchangelog.v1.AssetChangeLog/GetAssetChangeLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AssetChangeLogHTTPClientImpl) GetPageAssetChangeLog(ctx context.Context, in *AssetChangeLogPageReq, opts ...http.CallOption) (*AssetChangeLogPageReply, error) {
	var out AssetChangeLogPageReply
	pattern := "/v1/asset/change/log/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/assetchangelog.v1.AssetChangeLog/GetPageAssetChangeLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AssetChangeLogHTTPClientImpl) UpdateAssetChangeLog(ctx context.Context, in *AssetChangeLogUpdateReq, opts ...http.CallOption) (*AssetChangeLogUpdateReply, error) {
	var out AssetChangeLogUpdateReply
	pattern := "/v1/asset/change/log/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/assetchangelog.v1.AssetChangeLog/UpdateAssetChangeLog"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}