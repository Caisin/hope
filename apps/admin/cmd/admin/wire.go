//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/prometheus/client_golang/prometheus"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/conf"
	"hope/apps/admin/internal/data"
	"hope/apps/admin/internal/server"
	"hope/apps/admin/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger, *prometheus.HistogramVec, *prometheus.CounterVec) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
