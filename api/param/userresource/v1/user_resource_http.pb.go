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

type UserResourceHTTPServer interface {
	BatchDeleteUserResource(context.Context, *UserResourceBatchDeleteReq) (*UserResourceDeleteReply, error)
	CreateUserResource(context.Context, *UserResourceCreateReq) (*UserResourceCreateReply, error)
	DeleteUserResource(context.Context, *UserResourceDeleteReq) (*UserResourceDeleteReply, error)
	GetPageUserResource(context.Context, *UserResourcePageReq) (*UserResourcePageReply, error)
	GetUserResource(context.Context, *UserResourceReq) (*UserResourceReply, error)
	UpdateUserResource(context.Context, *UserResourceUpdateReq) (*UserResourceUpdateReply, error)
}

func RegisterUserResourceHTTPServer(s *http.Server, srv UserResourceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/user/resource/page", _UserResource_GetPageUserResource0_HTTP_Handler(srv))
	r.GET("/v1/user/resource/{id}", _UserResource_GetUserResource0_HTTP_Handler(srv))
	r.PUT("/v1/user/resource/{id}", _UserResource_UpdateUserResource0_HTTP_Handler(srv))
	r.POST("/v1/user/resource", _UserResource_CreateUserResource0_HTTP_Handler(srv))
	r.DELETE("/v1/user/resource/{id}", _UserResource_DeleteUserResource0_HTTP_Handler(srv))
	r.DELETE("/v1/user/resource", _UserResource_BatchDeleteUserResource0_HTTP_Handler(srv))
}

func _UserResource_GetPageUserResource0_HTTP_Handler(srv UserResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserResourcePageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/userresource.v1.UserResource/GetPageUserResource")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPageUserResource(ctx, req.(*UserResourcePageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserResourcePageReply)
		return ctx.Result(200, reply)
	}
}

func _UserResource_GetUserResource0_HTTP_Handler(srv UserResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserResourceReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/userresource.v1.UserResource/GetUserResource")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserResource(ctx, req.(*UserResourceReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserResourceReply)
		return ctx.Result(200, reply)
	}
}

func _UserResource_UpdateUserResource0_HTTP_Handler(srv UserResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserResourceUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/userresource.v1.UserResource/UpdateUserResource")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserResource(ctx, req.(*UserResourceUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserResourceUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _UserResource_CreateUserResource0_HTTP_Handler(srv UserResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserResourceCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/userresource.v1.UserResource/CreateUserResource")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUserResource(ctx, req.(*UserResourceCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserResourceCreateReply)
		return ctx.Result(200, reply)
	}
}

func _UserResource_DeleteUserResource0_HTTP_Handler(srv UserResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserResourceDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/userresource.v1.UserResource/DeleteUserResource")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUserResource(ctx, req.(*UserResourceDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserResourceDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _UserResource_BatchDeleteUserResource0_HTTP_Handler(srv UserResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserResourceBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/userresource.v1.UserResource/BatchDeleteUserResource")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteUserResource(ctx, req.(*UserResourceBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserResourceDeleteReply)
		return ctx.Result(200, reply)
	}
}

type UserResourceHTTPClient interface {
	BatchDeleteUserResource(ctx context.Context, req *UserResourceBatchDeleteReq, opts ...http.CallOption) (rsp *UserResourceDeleteReply, err error)
	CreateUserResource(ctx context.Context, req *UserResourceCreateReq, opts ...http.CallOption) (rsp *UserResourceCreateReply, err error)
	DeleteUserResource(ctx context.Context, req *UserResourceDeleteReq, opts ...http.CallOption) (rsp *UserResourceDeleteReply, err error)
	GetPageUserResource(ctx context.Context, req *UserResourcePageReq, opts ...http.CallOption) (rsp *UserResourcePageReply, err error)
	GetUserResource(ctx context.Context, req *UserResourceReq, opts ...http.CallOption) (rsp *UserResourceReply, err error)
	UpdateUserResource(ctx context.Context, req *UserResourceUpdateReq, opts ...http.CallOption) (rsp *UserResourceUpdateReply, err error)
}

type UserResourceHTTPClientImpl struct {
	cc *http.Client
}

func NewUserResourceHTTPClient(client *http.Client) UserResourceHTTPClient {
	return &UserResourceHTTPClientImpl{client}
}

func (c *UserResourceHTTPClientImpl) BatchDeleteUserResource(ctx context.Context, in *UserResourceBatchDeleteReq, opts ...http.CallOption) (*UserResourceDeleteReply, error) {
	var out UserResourceDeleteReply
	pattern := "/v1/user/resource"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/userresource.v1.UserResource/BatchDeleteUserResource"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserResourceHTTPClientImpl) CreateUserResource(ctx context.Context, in *UserResourceCreateReq, opts ...http.CallOption) (*UserResourceCreateReply, error) {
	var out UserResourceCreateReply
	pattern := "/v1/user/resource"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/userresource.v1.UserResource/CreateUserResource"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserResourceHTTPClientImpl) DeleteUserResource(ctx context.Context, in *UserResourceDeleteReq, opts ...http.CallOption) (*UserResourceDeleteReply, error) {
	var out UserResourceDeleteReply
	pattern := "/v1/user/resource/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/userresource.v1.UserResource/DeleteUserResource"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserResourceHTTPClientImpl) GetPageUserResource(ctx context.Context, in *UserResourcePageReq, opts ...http.CallOption) (*UserResourcePageReply, error) {
	var out UserResourcePageReply
	pattern := "/v1/user/resource/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/userresource.v1.UserResource/GetPageUserResource"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserResourceHTTPClientImpl) GetUserResource(ctx context.Context, in *UserResourceReq, opts ...http.CallOption) (*UserResourceReply, error) {
	var out UserResourceReply
	pattern := "/v1/user/resource/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/userresource.v1.UserResource/GetUserResource"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserResourceHTTPClientImpl) UpdateUserResource(ctx context.Context, in *UserResourceUpdateReq, opts ...http.CallOption) (*UserResourceUpdateReply, error) {
	var out UserResourceUpdateReply
	pattern := "/v1/user/resource/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/userresource.v1.UserResource/UpdateUserResource"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}