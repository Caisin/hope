package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"

	novelpayconfig "hope/api/param/novelpayconfig/v1"
	noveltag "hope/api/param/noveltag/v1"
	pageconfig "hope/api/param/pageconfig/v1"
	qiniuconfig "hope/api/param/qiniuconfig/v1"
	resourcegroup "hope/api/param/resourcegroup/v1"
	resourcestorage "hope/api/param/resourcestorage/v1"
	scoreproduct "hope/api/param/scoreproduct/v1"
	task "hope/api/param/task/v1"
	useranalysisstatistics "hope/api/param/useranalysisstatistics/v1"
	userconsume "hope/api/param/userconsume/v1"
	userresource "hope/api/param/userresource/v1"
	userresourcerecord "hope/api/param/userresourcerecord/v1"
	viptype "hope/api/param/viptype/v1"
	"hope/apps/param/internal/service"
)

func RegisterGRPCServer(

	novelPayConfigService *service.NovelPayConfigService,
	novelTagService *service.NovelTagService,
	pageConfigService *service.PageConfigService,
	qiniuConfigService *service.QiniuConfigService,
	resourceGroupService *service.ResourceGroupService,
	resourceStorageService *service.ResourceStorageService,
	scoreProductService *service.ScoreProductService,
	taskService *service.TaskService,
	userAnalysisStatisticsService *service.UserAnalysisStatisticsService,
	userConsumeService *service.UserConsumeService,
	userResourceService *service.UserResourceService,
	userResourceRecordService *service.UserResourceRecordService,
	vipTypeService *service.VipTypeService,
) []func(*grpc.Server) {
	list := make([]func(*grpc.Server), 0)

	list = append(list, func(srv *grpc.Server) {
		novelpayconfig.RegisterNovelPayConfigServer(srv, novelPayConfigService)
	})
	list = append(list, func(srv *grpc.Server) {
		noveltag.RegisterNovelTagServer(srv, novelTagService)
	})
	list = append(list, func(srv *grpc.Server) {
		pageconfig.RegisterPageConfigServer(srv, pageConfigService)
	})
	list = append(list, func(srv *grpc.Server) {
		qiniuconfig.RegisterQiniuConfigServer(srv, qiniuConfigService)
	})
	list = append(list, func(srv *grpc.Server) {
		resourcegroup.RegisterResourceGroupServer(srv, resourceGroupService)
	})
	list = append(list, func(srv *grpc.Server) {
		resourcestorage.RegisterResourceStorageServer(srv, resourceStorageService)
	})
	list = append(list, func(srv *grpc.Server) {
		scoreproduct.RegisterScoreProductServer(srv, scoreProductService)
	})
	list = append(list, func(srv *grpc.Server) {
		task.RegisterTaskServer(srv, taskService)
	})
	list = append(list, func(srv *grpc.Server) {
		useranalysisstatistics.RegisterUserAnalysisStatisticsServer(srv, userAnalysisStatisticsService)
	})
	list = append(list, func(srv *grpc.Server) {
		userconsume.RegisterUserConsumeServer(srv, userConsumeService)
	})
	list = append(list, func(srv *grpc.Server) {
		userresource.RegisterUserResourceServer(srv, userResourceService)
	})
	list = append(list, func(srv *grpc.Server) {
		userresourcerecord.RegisterUserResourceRecordServer(srv, userResourceRecordService)
	})
	list = append(list, func(srv *grpc.Server) {
		viptype.RegisterVipTypeServer(srv, vipTypeService)
	})
	return list
}
