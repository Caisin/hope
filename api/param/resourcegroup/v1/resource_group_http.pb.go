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

type ResourceGroupHTTPServer interface {
	BatchDeleteResourceGroup(context.Context, *ResourceGroupBatchDeleteReq) (*ResourceGroupDeleteReply, error)
	CreateResourceGroup(context.Context, *ResourceGroupCreateReq) (*ResourceGroupCreateReply, error)
	DeleteResourceGroup(context.Context, *ResourceGroupDeleteReq) (*ResourceGroupDeleteReply, error)
	GetResourceGroup(context.Context, *ResourceGroupReq) (*ResourceGroupReply, error)
	GetResourceGroupPage(context.Context, *ResourceGroupPageReq) (*ResourceGroupPageReply, error)
	UpdateResourceGroup(context.Context, *ResourceGroupUpdateReq) (*ResourceGroupUpdateReply, error)
}

func RegisterResourceGroupHTTPServer(s *http.Server, srv ResourceGroupHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/resource/group/page", _ResourceGroup_GetResourceGroupPage0_HTTP_Handler(srv))
	r.GET("/v1/resource/group/{id}", _ResourceGroup_GetResourceGroup0_HTTP_Handler(srv))
	r.PUT("/v1/resource/group/{id}", _ResourceGroup_UpdateResourceGroup0_HTTP_Handler(srv))
	r.POST("/v1/resource/group", _ResourceGroup_CreateResourceGroup0_HTTP_Handler(srv))
	r.DELETE("/v1/resource/group/{id}", _ResourceGroup_DeleteResourceGroup0_HTTP_Handler(srv))
	r.DELETE("/v1/resource/group", _ResourceGroup_BatchDeleteResourceGroup0_HTTP_Handler(srv))
}

func _ResourceGroup_GetResourceGroupPage0_HTTP_Handler(srv ResourceGroupHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ResourceGroupPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/resourcegroup.v1.ResourceGroup/GetResourceGroupPage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetResourceGroupPage(ctx, req.(*ResourceGroupPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ResourceGroupPageReply)
		return ctx.Result(200, reply)
	}
}

func _ResourceGroup_GetResourceGroup0_HTTP_Handler(srv ResourceGroupHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ResourceGroupReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/resourcegroup.v1.ResourceGroup/GetResourceGroup")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetResourceGroup(ctx, req.(*ResourceGroupReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ResourceGroupReply)
		return ctx.Result(200, reply)
	}
}

func _ResourceGroup_UpdateResourceGroup0_HTTP_Handler(srv ResourceGroupHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ResourceGroupUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/resourcegroup.v1.ResourceGroup/UpdateResourceGroup")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateResourceGroup(ctx, req.(*ResourceGroupUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ResourceGroupUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _ResourceGroup_CreateResourceGroup0_HTTP_Handler(srv ResourceGroupHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ResourceGroupCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/resourcegroup.v1.ResourceGroup/CreateResourceGroup")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateResourceGroup(ctx, req.(*ResourceGroupCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ResourceGroupCreateReply)
		return ctx.Result(200, reply)
	}
}

func _ResourceGroup_DeleteResourceGroup0_HTTP_Handler(srv ResourceGroupHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ResourceGroupDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/resourcegroup.v1.ResourceGroup/DeleteResourceGroup")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteResourceGroup(ctx, req.(*ResourceGroupDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ResourceGroupDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _ResourceGroup_BatchDeleteResourceGroup0_HTTP_Handler(srv ResourceGroupHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ResourceGroupBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/resourcegroup.v1.ResourceGroup/BatchDeleteResourceGroup")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteResourceGroup(ctx, req.(*ResourceGroupBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ResourceGroupDeleteReply)
		return ctx.Result(200, reply)
	}
}

type ResourceGroupHTTPClient interface {
	BatchDeleteResourceGroup(ctx context.Context, req *ResourceGroupBatchDeleteReq, opts ...http.CallOption) (rsp *ResourceGroupDeleteReply, err error)
	CreateResourceGroup(ctx context.Context, req *ResourceGroupCreateReq, opts ...http.CallOption) (rsp *ResourceGroupCreateReply, err error)
	DeleteResourceGroup(ctx context.Context, req *ResourceGroupDeleteReq, opts ...http.CallOption) (rsp *ResourceGroupDeleteReply, err error)
	GetResourceGroup(ctx context.Context, req *ResourceGroupReq, opts ...http.CallOption) (rsp *ResourceGroupReply, err error)
	GetResourceGroupPage(ctx context.Context, req *ResourceGroupPageReq, opts ...http.CallOption) (rsp *ResourceGroupPageReply, err error)
	UpdateResourceGroup(ctx context.Context, req *ResourceGroupUpdateReq, opts ...http.CallOption) (rsp *ResourceGroupUpdateReply, err error)
}

type ResourceGroupHTTPClientImpl struct {
	cc *http.Client
}

func NewResourceGroupHTTPClient(client *http.Client) ResourceGroupHTTPClient {
	return &ResourceGroupHTTPClientImpl{client}
}

func (c *ResourceGroupHTTPClientImpl) BatchDeleteResourceGroup(ctx context.Context, in *ResourceGroupBatchDeleteReq, opts ...http.CallOption) (*ResourceGroupDeleteReply, error) {
	var out ResourceGroupDeleteReply
	pattern := "/v1/resource/group"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/resourcegroup.v1.ResourceGroup/BatchDeleteResourceGroup"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ResourceGroupHTTPClientImpl) CreateResourceGroup(ctx context.Context, in *ResourceGroupCreateReq, opts ...http.CallOption) (*ResourceGroupCreateReply, error) {
	var out ResourceGroupCreateReply
	pattern := "/v1/resource/group"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/resourcegroup.v1.ResourceGroup/CreateResourceGroup"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ResourceGroupHTTPClientImpl) DeleteResourceGroup(ctx context.Context, in *ResourceGroupDeleteReq, opts ...http.CallOption) (*ResourceGroupDeleteReply, error) {
	var out ResourceGroupDeleteReply
	pattern := "/v1/resource/group/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/resourcegroup.v1.ResourceGroup/DeleteResourceGroup"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ResourceGroupHTTPClientImpl) GetResourceGroup(ctx context.Context, in *ResourceGroupReq, opts ...http.CallOption) (*ResourceGroupReply, error) {
	var out ResourceGroupReply
	pattern := "/v1/resource/group/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/resourcegroup.v1.ResourceGroup/GetResourceGroup"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ResourceGroupHTTPClientImpl) GetResourceGroupPage(ctx context.Context, in *ResourceGroupPageReq, opts ...http.CallOption) (*ResourceGroupPageReply, error) {
	var out ResourceGroupPageReply
	pattern := "/v1/resource/group/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/resourcegroup.v1.ResourceGroup/GetResourceGroupPage"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ResourceGroupHTTPClientImpl) UpdateResourceGroup(ctx context.Context, in *ResourceGroupUpdateReq, opts ...http.CallOption) (*ResourceGroupUpdateReply, error) {
	var out ResourceGroupUpdateReply
	pattern := "/v1/resource/group/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/resourcegroup.v1.ResourceGroup/UpdateResourceGroup"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
