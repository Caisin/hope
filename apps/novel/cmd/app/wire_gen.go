// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/data"
	"hope/apps/novel/internal/server"
	"hope/apps/novel/internal/service"
	"hope/pkg/conf"
	"hope/pkg/provider"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	client := data.NewEntClient(confData, logger)
	redisClient := provider.NewRedisClient(confData, logger)
	producer := provider.NewNsqProducer(confData)
	dataData, cleanup, err := data.NewData(client, redisClient, producer, logger)
	if err != nil {
		return nil, nil, err
	}
	activityRepo := data.NewActivityRepo(dataData, logger)
	activityUseCase := biz.NewActivityUseCase(activityRepo, logger)
	activityService := service.NewActivityService(activityUseCase, logger)
	activityComponentRepo := data.NewActivityComponentRepo(dataData, logger)
	activityComponentUseCase := biz.NewActivityComponentUseCase(activityComponentRepo, logger)
	activityComponentService := service.NewActivityComponentService(activityComponentUseCase, logger)
	adChangeLogRepo := data.NewAdChangeLogRepo(dataData, logger)
	adChangeLogUseCase := biz.NewAdChangeLogUseCase(adChangeLogRepo, logger)
	adChangeLogService := service.NewAdChangeLogService(adChangeLogUseCase, logger)
	adChannelRepo := data.NewAdChannelRepo(dataData, logger)
	adChannelUseCase := biz.NewAdChannelUseCase(adChannelRepo, logger)
	adChannelService := service.NewAdChannelService(adChannelUseCase, logger)
	agreementLogRepo := data.NewAgreementLogRepo(dataData, logger)
	agreementLogUseCase := biz.NewAgreementLogUseCase(agreementLogRepo, logger)
	agreementLogService := service.NewAgreementLogService(agreementLogUseCase, logger)
	amBalanceRepo := data.NewAmBalanceRepo(dataData, logger)
	amBalanceUseCase := biz.NewAmBalanceUseCase(amBalanceRepo, logger)
	amBalanceService := service.NewAmBalanceService(amBalanceUseCase, logger)
	appVersionRepo := data.NewAppVersionRepo(dataData, logger)
	appVersionUseCase := biz.NewAppVersionUseCase(appVersionRepo, logger)
	appVersionService := service.NewAppVersionService(appVersionUseCase, logger)
	assetChangeLogRepo := data.NewAssetChangeLogRepo(dataData, logger)
	assetChangeLogUseCase := biz.NewAssetChangeLogUseCase(assetChangeLogRepo, logger)
	assetChangeLogService := service.NewAssetChangeLogService(assetChangeLogUseCase, logger)
	assetItemRepo := data.NewAssetItemRepo(dataData, logger)
	assetItemUseCase := biz.NewAssetItemUseCase(assetItemRepo, logger)
	assetItemService := service.NewAssetItemService(assetItemUseCase, logger)
	bookPackageRepo := data.NewBookPackageRepo(dataData, logger)
	bookPackageUseCase := biz.NewBookPackageUseCase(bookPackageRepo, logger)
	bookPackageService := service.NewBookPackageService(bookPackageUseCase, logger)
	clientErrorRepo := data.NewClientErrorRepo(dataData, logger)
	clientErrorUseCase := biz.NewClientErrorUseCase(clientErrorRepo, logger)
	clientErrorService := service.NewClientErrorService(clientErrorUseCase, logger)
	customerNovelConfigRepo := data.NewCustomerNovelConfigRepo(dataData, logger)
	customerNovelConfigUseCase := biz.NewCustomerNovelConfigUseCase(customerNovelConfigRepo, logger)
	customerNovelConfigService := service.NewCustomerNovelConfigService(customerNovelConfigUseCase, logger)
	customerNovelsRepo := data.NewCustomerNovelsRepo(dataData, logger)
	customerNovelsUseCase := biz.NewCustomerNovelsUseCase(customerNovelsRepo, logger)
	customerNovelsService := service.NewCustomerNovelsService(customerNovelsUseCase, logger)
	dataSourceRepo := data.NewDataSourceRepo(dataData, logger)
	dataSourceUseCase := biz.NewDataSourceUseCase(dataSourceRepo, logger)
	dataSourceService := service.NewDataSourceService(dataSourceUseCase, logger)
	listenRecordRepo := data.NewListenRecordRepo(dataData, logger)
	listenRecordUseCase := biz.NewListenRecordUseCase(listenRecordRepo, logger)
	listenRecordService := service.NewListenRecordService(listenRecordUseCase, logger)
	novelRepo := data.NewNovelRepo(dataData, logger)
	novelUseCase := biz.NewNovelUseCase(novelRepo, logger)
	novelService := service.NewNovelService(novelUseCase, logger)
	novelAutoBuyRepo := data.NewNovelAutoBuyRepo(dataData, logger)
	novelAutoBuyUseCase := biz.NewNovelAutoBuyUseCase(novelAutoBuyRepo, logger)
	novelAutoBuyService := service.NewNovelAutoBuyService(novelAutoBuyUseCase, logger)
	novelBookshelfRepo := data.NewNovelBookshelfRepo(dataData, logger)
	novelBookshelfUseCase := biz.NewNovelBookshelfUseCase(novelBookshelfRepo, logger)
	novelBookshelfService := service.NewNovelBookshelfService(novelBookshelfUseCase, logger)
	novelBuyChapterRecordRepo := data.NewNovelBuyChapterRecordRepo(dataData, logger)
	novelBuyChapterRecordUseCase := biz.NewNovelBuyChapterRecordUseCase(novelBuyChapterRecordRepo, logger)
	novelBuyChapterRecordService := service.NewNovelBuyChapterRecordService(novelBuyChapterRecordUseCase, logger)
	novelBuyRecordRepo := data.NewNovelBuyRecordRepo(dataData, logger)
	novelBuyRecordUseCase := biz.NewNovelBuyRecordUseCase(novelBuyRecordRepo, logger)
	novelBuyRecordService := service.NewNovelBuyRecordService(novelBuyRecordUseCase, logger)
	novelChapterRepo := data.NewNovelChapterRepo(dataData, logger)
	novelChapterUseCase := biz.NewNovelChapterUseCase(novelChapterRepo, logger)
	novelChapterService := service.NewNovelChapterService(novelChapterUseCase, logger)
	novelClassifyRepo := data.NewNovelClassifyRepo(dataData, logger)
	novelClassifyUseCase := biz.NewNovelClassifyUseCase(novelClassifyRepo, logger)
	novelClassifyService := service.NewNovelClassifyService(novelClassifyUseCase, logger)
	novelCommentRepo := data.NewNovelCommentRepo(dataData, logger)
	novelCommentUseCase := biz.NewNovelCommentUseCase(novelCommentRepo, logger)
	novelCommentService := service.NewNovelCommentService(novelCommentUseCase, logger)
	novelConsumeRepo := data.NewNovelConsumeRepo(dataData, logger)
	novelConsumeUseCase := biz.NewNovelConsumeUseCase(novelConsumeRepo, logger)
	novelConsumeService := service.NewNovelConsumeService(novelConsumeUseCase, logger)
	novelMsgRepo := data.NewNovelMsgRepo(dataData, logger)
	novelMsgUseCase := biz.NewNovelMsgUseCase(novelMsgRepo, logger)
	novelMsgService := service.NewNovelMsgService(novelMsgUseCase, logger)
	payOrderRepo := data.NewPayOrderRepo(dataData, logger)
	payOrderUseCase := biz.NewPayOrderUseCase(payOrderRepo, logger)
	payOrderService := service.NewPayOrderService(payOrderUseCase, logger)
	socialUserRepo := data.NewSocialUserRepo(dataData, logger)
	socialUserUseCase := biz.NewSocialUserUseCase(socialUserRepo, logger)
	socialUserService := service.NewSocialUserService(socialUserUseCase, logger)
	taskLogRepo := data.NewTaskLogRepo(dataData, logger)
	taskLogUseCase := biz.NewTaskLogUseCase(taskLogRepo, logger)
	taskLogService := service.NewTaskLogService(taskLogUseCase, logger)
	userEventRepo := data.NewUserEventRepo(dataData, logger)
	userEventUseCase := biz.NewUserEventUseCase(userEventRepo, logger)
	userEventService := service.NewUserEventService(userEventUseCase, logger)
	userMsgRepo := data.NewUserMsgRepo(dataData, logger)
	userMsgUseCase := biz.NewUserMsgUseCase(userMsgRepo, logger)
	userMsgService := service.NewUserMsgService(userMsgUseCase, logger)
	vipUserRepo := data.NewVipUserRepo(dataData, logger)
	vipUserUseCase := biz.NewVipUserUseCase(vipUserRepo, logger)
	vipUserService := service.NewVipUserService(vipUserUseCase, logger)
	v := server.RegisterHTTPServer(activityService, activityComponentService, adChangeLogService, adChannelService, agreementLogService, amBalanceService, appVersionService, assetChangeLogService, assetItemService, bookPackageService, clientErrorService, customerNovelConfigService, customerNovelsService, dataSourceService, listenRecordService, novelService, novelAutoBuyService, novelBookshelfService, novelBuyChapterRecordService, novelBuyRecordService, novelChapterService, novelClassifyService, novelCommentService, novelConsumeService, novelMsgService, payOrderService, socialUserService, taskLogService, userEventService, userMsgService, vipUserService)
	httpServer := provider.NewHTTPServer(confServer, v, redisClient, logger)
	v2 := server.RegisterGRPCServer(activityService, activityComponentService, adChangeLogService, adChannelService, agreementLogService, amBalanceService, appVersionService, assetChangeLogService, assetItemService, bookPackageService, clientErrorService, customerNovelConfigService, customerNovelsService, dataSourceService, listenRecordService, novelService, novelAutoBuyService, novelBookshelfService, novelBuyChapterRecordService, novelBuyRecordService, novelChapterService, novelClassifyService, novelCommentService, novelConsumeService, novelMsgService, payOrderService, socialUserService, taskLogService, userEventService, userMsgService, vipUserService)
	grpcServer := provider.NewGRPCServer(confServer, v2, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
