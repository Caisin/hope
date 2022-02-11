package server

import (
	"context"
	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/handlers"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"hope/pkg/auth"
	"hope/pkg/conf"
	"hope/pkg/middleware"
	"strings"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server,
	regFun []func(server *http.Server),
	logger log.Logger,
	hv *prometheus.HistogramVec,
	rdc *redis.Client,
	cv *prometheus.CounterVec) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
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
			//应用授权
			selector.Server(
				jwt.Server(auth.SecretKeyFun),
			).Match(func(ctx context.Context, operation string) bool {
				if strings.HasSuffix(operation, "/Login") || strings.HasSuffix(operation, "Logout") {
					return false
				}
				return true
			}).Build(),
			//权限校验
			middleware.PermissionCheck(rdc),
		),
		//cors
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	for _, f := range regFun {
		f(srv)
	}
	//注册swagger-ui
	h := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", h)
	//注册服务监控
	srv.Handle("/metrics", promhttp.Handler())
	return srv
}
