package server

import (
	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/prometheus/client_golang/prometheus"
	"hope/pkg/auth"
	"hope/pkg/conf"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, regFun []func(*grpc.Server), logger log.Logger, hv *prometheus.HistogramVec, cv *prometheus.CounterVec) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			//限流
			ratelimit.Server(),
			//异常恢复
			recovery.Recovery(recovery.WithLogger(logger)),
			//打印服务收到或发起的请求详情。
			logging.Server(logger),
			//链路追踪
			tracing.Server(),
			//服务监控
			metrics.Server(
				metrics.WithSeconds(prom.NewHistogram(hv)),
				metrics.WithRequests(prom.NewCounter(cv)),
			),
			jwt.Server(auth.SecretKeyFun),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	for _, f := range regFun {
		f(srv)
	}
	return srv
}
