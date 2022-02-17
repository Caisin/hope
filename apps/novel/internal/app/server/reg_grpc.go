package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	vipuser "hope/api/novel/vipuser/v1"
	"hope/apps/novel/internal/service"
)

func RegisterGRPCServer(

	vipUserService *service.VipUserService,
) []func(*grpc.Server) {
	list := make([]func(*grpc.Server), 0)

	list = append(list, func(srv *grpc.Server) {
		vipuser.RegisterVipUserServer(srv, vipUserService)
	})
	return list
}
