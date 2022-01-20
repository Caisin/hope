package server

import "github.com/go-kratos/kratos/v2/transport/http"

func RegisterHTTPServer() []func(*http.Server) {
	list := make([]func(*http.Server), 0)
	return list
}
