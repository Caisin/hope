//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/conf"
	"hope/apps/novel/internal/data"
	"hope/apps/novel/internal/server"
	"hope/apps/novel/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
