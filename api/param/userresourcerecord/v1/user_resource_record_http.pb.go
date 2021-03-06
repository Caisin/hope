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

type UserResourceRecordHTTPServer interface {
	BatchDeleteUserResourceRecord(context.Context, *UserResourceRecordBatchDeleteReq) (*UserResourceRecordDeleteReply, error)
	CreateUserResourceRecord(context.Context, *UserResourceRecordCreateReq) (*UserResourceRecordCreateReply, error)
	DeleteUserResourceRecord(context.Context, *UserResourceRecordDeleteReq) (*UserResourceRecordDeleteReply, error)
	GetUserResourceRecord(context.Context, *UserResourceRecordReq) (*UserResourceRecordReply, error)
	GetUserResourceRecordPage(context.Context, *UserResourceRecordPageReq) (*UserResourceRecordPageReply, error)
	UpdateUserResourceRecord(context.Context, *UserResourceRecordUpdateReq) (*UserResourceRecordUpdateReply, error)
}

func RegisterUserResourceRecordHTTPServer(s *http.Server, srv UserResourceRecordHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/user/resource/record/page", _UserResourceRecord_GetUserResourceRecordPage0_HTTP_Handler(srv))
	r.GET("/v1/user/resource/record/{id}", _UserResourceRecord_GetUserResourceRecord0_HTTP_Handler(srv))
	r.PUT("/v1/user/resource/record/{id}", _UserResourceRecord_UpdateUserResourceRecord0_HTTP_Handler(srv))
	r.POST("/v1/user/resource/record", _UserResourceRecord_CreateUserResourceRecord0_HTTP_Handler(srv))
	r.DELETE("/v1/user/resource/record/{id}", _UserResourceRecord_DeleteUserResourceRecord0_HTTP_Handler(srv))
	r.DELETE("/v1/user/resource/record", _UserResourceRecord_BatchDeleteUserResourceRecord0_HTTP_Handler(srv))
}

func _UserResourceRecord_GetUserResourceRecordPage0_HTTP_Handler(srv UserResourceRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserResourceRecordPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/userresourcerecord.v1.UserResourceRecord/GetUserResourceRecordPage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserResourceRecordPage(ctx, req.(*UserResourceRecordPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserResourceRecordPageReply)
		return ctx.Result(200, reply)
	}
}

func _UserResourceRecord_GetUserResourceRecord0_HTTP_Handler(srv UserResourceRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserResourceRecordReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/userresourcerecord.v1.UserResourceRecord/GetUserResourceRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserResourceRecord(ctx, req.(*UserResourceRecordReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserResourceRecordReply)
		return ctx.Result(200, reply)
	}
}

func _UserResourceRecord_UpdateUserResourceRecord0_HTTP_Handler(srv UserResourceRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserResourceRecordUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/userresourcerecord.v1.UserResourceRecord/UpdateUserResourceRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserResourceRecord(ctx, req.(*UserResourceRecordUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserResourceRecordUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _UserResourceRecord_CreateUserResourceRecord0_HTTP_Handler(srv UserResourceRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserResourceRecordCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/userresourcerecord.v1.UserResourceRecord/CreateUserResourceRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUserResourceRecord(ctx, req.(*UserResourceRecordCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserResourceRecordCreateReply)
		return ctx.Result(200, reply)
	}
}

func _UserResourceRecord_DeleteUserResourceRecord0_HTTP_Handler(srv UserResourceRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserResourceRecordDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/userresourcerecord.v1.UserResourceRecord/DeleteUserResourceRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUserResourceRecord(ctx, req.(*UserResourceRecordDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserResourceRecordDeleteReply)
		return ctx.Result(200, reply)
	}
}

func _UserResourceRecord_BatchDeleteUserResourceRecord0_HTTP_Handler(srv UserResourceRecordHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserResourceRecordBatchDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/userresourcerecord.v1.UserResourceRecord/BatchDeleteUserResourceRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchDeleteUserResourceRecord(ctx, req.(*UserResourceRecordBatchDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserResourceRecordDeleteReply)
		return ctx.Result(200, reply)
	}
}

type UserResourceRecordHTTPClient interface {
	BatchDeleteUserResourceRecord(ctx context.Context, req *UserResourceRecordBatchDeleteReq, opts ...http.CallOption) (rsp *UserResourceRecordDeleteReply, err error)
	CreateUserResourceRecord(ctx context.Context, req *UserResourceRecordCreateReq, opts ...http.CallOption) (rsp *UserResourceRecordCreateReply, err error)
	DeleteUserResourceRecord(ctx context.Context, req *UserResourceRecordDeleteReq, opts ...http.CallOption) (rsp *UserResourceRecordDeleteReply, err error)
	GetUserResourceRecord(ctx context.Context, req *UserResourceRecordReq, opts ...http.CallOption) (rsp *UserResourceRecordReply, err error)
	GetUserResourceRecordPage(ctx context.Context, req *UserResourceRecordPageReq, opts ...http.CallOption) (rsp *UserResourceRecordPageReply, err error)
	UpdateUserResourceRecord(ctx context.Context, req *UserResourceRecordUpdateReq, opts ...http.CallOption) (rsp *UserResourceRecordUpdateReply, err error)
}

type UserResourceRecordHTTPClientImpl struct {
	cc *http.Client
}

func NewUserResourceRecordHTTPClient(client *http.Client) UserResourceRecordHTTPClient {
	return &UserResourceRecordHTTPClientImpl{client}
}

func (c *UserResourceRecordHTTPClientImpl) BatchDeleteUserResourceRecord(ctx context.Context, in *UserResourceRecordBatchDeleteReq, opts ...http.CallOption) (*UserResourceRecordDeleteReply, error) {
	var out UserResourceRecordDeleteReply
	pattern := "/v1/user/resource/record"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/userresourcerecord.v1.UserResourceRecord/BatchDeleteUserResourceRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserResourceRecordHTTPClientImpl) CreateUserResourceRecord(ctx context.Context, in *UserResourceRecordCreateReq, opts ...http.CallOption) (*UserResourceRecordCreateReply, error) {
	var out UserResourceRecordCreateReply
	pattern := "/v1/user/resource/record"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/userresourcerecord.v1.UserResourceRecord/CreateUserResourceRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserResourceRecordHTTPClientImpl) DeleteUserResourceRecord(ctx context.Context, in *UserResourceRecordDeleteReq, opts ...http.CallOption) (*UserResourceRecordDeleteReply, error) {
	var out UserResourceRecordDeleteReply
	pattern := "/v1/user/resource/record/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/userresourcerecord.v1.UserResourceRecord/DeleteUserResourceRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserResourceRecordHTTPClientImpl) GetUserResourceRecord(ctx context.Context, in *UserResourceRecordReq, opts ...http.CallOption) (*UserResourceRecordReply, error) {
	var out UserResourceRecordReply
	pattern := "/v1/user/resource/record/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/userresourcerecord.v1.UserResourceRecord/GetUserResourceRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserResourceRecordHTTPClientImpl) GetUserResourceRecordPage(ctx context.Context, in *UserResourceRecordPageReq, opts ...http.CallOption) (*UserResourceRecordPageReply, error) {
	var out UserResourceRecordPageReply
	pattern := "/v1/user/resource/record/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/userresourcerecord.v1.UserResourceRecord/GetUserResourceRecordPage"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserResourceRecordHTTPClientImpl) UpdateUserResourceRecord(ctx context.Context, in *UserResourceRecordUpdateReq, opts ...http.CallOption) (*UserResourceRecordUpdateReply, error) {
	var out UserResourceRecordUpdateReply
	pattern := "/v1/user/resource/record/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/userresourcerecord.v1.UserResourceRecord/UpdateUserResourceRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
