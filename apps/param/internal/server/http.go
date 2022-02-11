package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/handlers"
	"go.opentelemetry.io/otel"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	task "hope/api/param/task/v1"
	"hope/apps/param/internal/service"
	"hope/pkg/auth"
	"hope/pkg/conf"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, regFun []func(*http.Server), tp *tracesdk.TracerProvider, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			jwt.Server(auth.SecretKeyFun),
		),
		//跨域
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}

	otel.SetTracerProvider(tp)

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
	return srv
}

func RegHttpFun(
	taskService *service.TaskService,
) []func(*http.Server) {
	list := make([]func(*http.Server), 0)
	list = append(list, func(srv *http.Server) {
		task.RegisterTaskHTTPServer(srv, taskService)
	})
	return list
}
