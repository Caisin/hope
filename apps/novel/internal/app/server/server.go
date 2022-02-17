package server

import (
	"github.com/google/wire"
	"hope/pkg/provider"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(provider.NewHTTPServer, provider.NewGRPCServer, RegisterHTTPServer, RegisterGRPCServer)
