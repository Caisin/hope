package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"

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

func RegisterHTTPServer(

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
) []func(*http.Server) {
	list := make([]func(*http.Server), 0)

	list = append(list, func(srv *http.Server) {
		novelpayconfig.RegisterNovelPayConfigHTTPServer(srv, novelPayConfigService)
	})
	list = append(list, func(srv *http.Server) {
		noveltag.RegisterNovelTagHTTPServer(srv, novelTagService)
	})
	list = append(list, func(srv *http.Server) {
		pageconfig.RegisterPageConfigHTTPServer(srv, pageConfigService)
	})
	list = append(list, func(srv *http.Server) {
		qiniuconfig.RegisterQiniuConfigHTTPServer(srv, qiniuConfigService)
	})
	list = append(list, func(srv *http.Server) {
		resourcegroup.RegisterResourceGroupHTTPServer(srv, resourceGroupService)
	})
	list = append(list, func(srv *http.Server) {
		resourcestorage.RegisterResourceStorageHTTPServer(srv, resourceStorageService)
	})
	list = append(list, func(srv *http.Server) {
		scoreproduct.RegisterScoreProductHTTPServer(srv, scoreProductService)
	})
	list = append(list, func(srv *http.Server) {
		task.RegisterTaskHTTPServer(srv, taskService)
	})
	list = append(list, func(srv *http.Server) {
		useranalysisstatistics.RegisterUserAnalysisStatisticsHTTPServer(srv, userAnalysisStatisticsService)
	})
	list = append(list, func(srv *http.Server) {
		userconsume.RegisterUserConsumeHTTPServer(srv, userConsumeService)
	})
	list = append(list, func(srv *http.Server) {
		userresource.RegisterUserResourceHTTPServer(srv, userResourceService)
	})
	list = append(list, func(srv *http.Server) {
		userresourcerecord.RegisterUserResourceRecordHTTPServer(srv, userResourceRecordService)
	})
	list = append(list, func(srv *http.Server) {
		viptype.RegisterVipTypeHTTPServer(srv, vipTypeService)
	})
	return list
}
