package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	task "hope/api/param/task/v1"
	"hope/apps/param/internal/service"
)

func RegisterGRPCServer(
	taskService *service.TaskService,
) []func(*grpc.Server) {
	list := make([]func(*grpc.Server), 0)
	list = append(list, func(srv *grpc.Server) {
		task.RegisterTaskServer(srv, taskService)
	})
	return list
}
