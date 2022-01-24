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

type AuthHTTPServer interface {
	GetPermCode(context.Context, *GetPermReq) (*GetPermReply, error)
	GetUserInfo(context.Context, *GetUserInfoReq) (*LoginReply, error)
	LogOut(context.Context, *LogOutReq) (*LogOutReply, error)
	Login(context.Context, *LoginReq) (*LoginReply, error)
}

func RegisterAuthHTTPServer(s *http.Server, srv AuthHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/sys/auth/login", _Auth_Login0_HTTP_Handler(srv))
	r.POST("/v1/sys/auth/logout", _Auth_LogOut0_HTTP_Handler(srv))
	r.GET("/v1/sys/auth/getUserInfo", _Auth_GetUserInfo0_HTTP_Handler(srv))
	r.GET("/v1/sys/auth/getPermCode", _Auth_GetPermCode0_HTTP_Handler(srv))
}

func _Auth_Login0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysuser.v1.Auth/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_LogOut0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogOutReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysuser.v1.Auth/LogOut")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LogOut(ctx, req.(*LogOutReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogOutReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_GetUserInfo0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserInfoReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysuser.v1.Auth/GetUserInfo")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserInfo(ctx, req.(*GetUserInfoReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_GetPermCode0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPermReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/sysuser.v1.Auth/GetPermCode")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPermCode(ctx, req.(*GetPermReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetPermReply)
		return ctx.Result(200, reply)
	}
}

type AuthHTTPClient interface {
	GetPermCode(ctx context.Context, req *GetPermReq, opts ...http.CallOption) (rsp *GetPermReply, err error)
	GetUserInfo(ctx context.Context, req *GetUserInfoReq, opts ...http.CallOption) (rsp *LoginReply, err error)
	LogOut(ctx context.Context, req *LogOutReq, opts ...http.CallOption) (rsp *LogOutReply, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginReply, err error)
}

type AuthHTTPClientImpl struct {
	cc *http.Client
}

func NewAuthHTTPClient(client *http.Client) AuthHTTPClient {
	return &AuthHTTPClientImpl{client}
}

func (c *AuthHTTPClientImpl) GetPermCode(ctx context.Context, in *GetPermReq, opts ...http.CallOption) (*GetPermReply, error) {
	var out GetPermReply
	pattern := "/v1/sys/auth/getPermCode"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/sysuser.v1.Auth/GetPermCode"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/sys/auth/getUserInfo"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/sysuser.v1.Auth/GetUserInfo"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) LogOut(ctx context.Context, in *LogOutReq, opts ...http.CallOption) (*LogOutReply, error) {
	var out LogOutReply
	pattern := "/v1/sys/auth/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/sysuser.v1.Auth/LogOut"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/sys/auth/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/sysuser.v1.Auth/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
