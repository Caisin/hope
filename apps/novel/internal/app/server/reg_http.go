package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"

	vipuser "hope/api/novel/vipuser/v1"
	"hope/apps/novel/internal/service"
)

func RegisterHTTPServer(

	vipUserService *service.VipUserService,
) []func(*http.Server) {
	list := make([]func(*http.Server), 0)

	list = append(list, func(srv *http.Server) {
		vipuser.RegisterVipUserHTTPServer(srv, vipUserService)
	})
	return list
}
