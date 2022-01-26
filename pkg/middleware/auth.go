package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-redis/redis/v8"
	"hope/pkg/auth"
)

// PermissionCheck 权限校验中间件
func PermissionCheck(rdc *redis.Client) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if info, ok := transport.FromServerContext(ctx); ok {
				operation := info.Operation()
				id := auth.GetIdByOper(operation)
				if id > 0 {
					if auth.HasOperationPer(ctx, rdc, operation) {
						return handler(ctx, req)
					}
				}
			}
			return nil, errors.Unauthorized("UNAUTHORIZED", "JWT token is missing")
		}
	}
}
