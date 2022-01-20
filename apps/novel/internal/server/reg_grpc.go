package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"

	activity "hope/api/novel/activity/v1"
	activitycomponent "hope/api/novel/activitycomponent/v1"
	adchangelog "hope/api/novel/adchangelog/v1"
	adchannel "hope/api/novel/adchannel/v1"
	agreementlog "hope/api/novel/agreementlog/v1"
	ambalance "hope/api/novel/ambalance/v1"
	appversion "hope/api/novel/appversion/v1"
	assetchangelog "hope/api/novel/assetchangelog/v1"
	assetitem "hope/api/novel/assetitem/v1"
	bookpackage "hope/api/novel/bookpackage/v1"
	clienterror "hope/api/novel/clienterror/v1"
	customernovelconfig "hope/api/novel/customernovelconfig/v1"
	customernovels "hope/api/novel/customernovels/v1"
	datasource "hope/api/novel/datasource/v1"
	listenrecord "hope/api/novel/listenrecord/v1"
	novel "hope/api/novel/novel/v1"
	novelautobuy "hope/api/novel/novelautobuy/v1"
	novelbookshelf "hope/api/novel/novelbookshelf/v1"
	novelbuychapterrecord "hope/api/novel/novelbuychapterrecord/v1"
	novelbuyrecord "hope/api/novel/novelbuyrecord/v1"
	novelchapter "hope/api/novel/novelchapter/v1"
	novelclassify "hope/api/novel/novelclassify/v1"
	novelcomment "hope/api/novel/novelcomment/v1"
	novelconsume "hope/api/novel/novelconsume/v1"
	novelmsg "hope/api/novel/novelmsg/v1"
	payorder "hope/api/novel/payorder/v1"
	socialuser "hope/api/novel/socialuser/v1"
	tasklog "hope/api/novel/tasklog/v1"
	userevent "hope/api/novel/userevent/v1"
	usermsg "hope/api/novel/usermsg/v1"
	vipuser "hope/api/novel/vipuser/v1"
	"hope/apps/novel/internal/service"
)

func RegisterGRPCServer(

	activityService *service.ActivityService,
	activityComponentService *service.ActivityComponentService,
	adChangeLogService *service.AdChangeLogService,
	adChannelService *service.AdChannelService,
	agreementLogService *service.AgreementLogService,
	amBalanceService *service.AmBalanceService,
	appVersionService *service.AppVersionService,
	assetChangeLogService *service.AssetChangeLogService,
	assetItemService *service.AssetItemService,
	bookPackageService *service.BookPackageService,
	clientErrorService *service.ClientErrorService,
	customerNovelConfigService *service.CustomerNovelConfigService,
	customerNovelsService *service.CustomerNovelsService,
	dataSourceService *service.DataSourceService,
	listenRecordService *service.ListenRecordService,
	novelService *service.NovelService,
	novelAutoBuyService *service.NovelAutoBuyService,
	novelBookshelfService *service.NovelBookshelfService,
	novelBuyChapterRecordService *service.NovelBuyChapterRecordService,
	novelBuyRecordService *service.NovelBuyRecordService,
	novelChapterService *service.NovelChapterService,
	novelClassifyService *service.NovelClassifyService,
	novelCommentService *service.NovelCommentService,
	novelConsumeService *service.NovelConsumeService,
	novelMsgService *service.NovelMsgService,
	payOrderService *service.PayOrderService,
	socialUserService *service.SocialUserService,
	taskLogService *service.TaskLogService,
	userEventService *service.UserEventService,
	userMsgService *service.UserMsgService,
	vipUserService *service.VipUserService,
) []func(*grpc.Server) {
	list := make([]func(*grpc.Server), 0)

	list = append(list, func(srv *grpc.Server) {
		activity.RegisterActivityServer(srv, activityService)
	})
	list = append(list, func(srv *grpc.Server) {
		activitycomponent.RegisterActivityComponentServer(srv, activityComponentService)
	})
	list = append(list, func(srv *grpc.Server) {
		adchangelog.RegisterAdChangeLogServer(srv, adChangeLogService)
	})
	list = append(list, func(srv *grpc.Server) {
		adchannel.RegisterAdChannelServer(srv, adChannelService)
	})
	list = append(list, func(srv *grpc.Server) {
		agreementlog.RegisterAgreementLogServer(srv, agreementLogService)
	})
	list = append(list, func(srv *grpc.Server) {
		ambalance.RegisterAmBalanceServer(srv, amBalanceService)
	})
	list = append(list, func(srv *grpc.Server) {
		appversion.RegisterAppVersionServer(srv, appVersionService)
	})
	list = append(list, func(srv *grpc.Server) {
		assetchangelog.RegisterAssetChangeLogServer(srv, assetChangeLogService)
	})
	list = append(list, func(srv *grpc.Server) {
		assetitem.RegisterAssetItemServer(srv, assetItemService)
	})
	list = append(list, func(srv *grpc.Server) {
		bookpackage.RegisterBookPackageServer(srv, bookPackageService)
	})
	list = append(list, func(srv *grpc.Server) {
		clienterror.RegisterClientErrorServer(srv, clientErrorService)
	})
	list = append(list, func(srv *grpc.Server) {
		customernovelconfig.RegisterCustomerNovelConfigServer(srv, customerNovelConfigService)
	})
	list = append(list, func(srv *grpc.Server) {
		customernovels.RegisterCustomerNovelsServer(srv, customerNovelsService)
	})
	list = append(list, func(srv *grpc.Server) {
		datasource.RegisterDataSourceServer(srv, dataSourceService)
	})
	list = append(list, func(srv *grpc.Server) {
		listenrecord.RegisterListenRecordServer(srv, listenRecordService)
	})
	list = append(list, func(srv *grpc.Server) {
		novel.RegisterNovelServer(srv, novelService)
	})
	list = append(list, func(srv *grpc.Server) {
		novelautobuy.RegisterNovelAutoBuyServer(srv, novelAutoBuyService)
	})
	list = append(list, func(srv *grpc.Server) {
		novelbookshelf.RegisterNovelBookshelfServer(srv, novelBookshelfService)
	})
	list = append(list, func(srv *grpc.Server) {
		novelbuychapterrecord.RegisterNovelBuyChapterRecordServer(srv, novelBuyChapterRecordService)
	})
	list = append(list, func(srv *grpc.Server) {
		novelbuyrecord.RegisterNovelBuyRecordServer(srv, novelBuyRecordService)
	})
	list = append(list, func(srv *grpc.Server) {
		novelchapter.RegisterNovelChapterServer(srv, novelChapterService)
	})
	list = append(list, func(srv *grpc.Server) {
		novelclassify.RegisterNovelClassifyServer(srv, novelClassifyService)
	})
	list = append(list, func(srv *grpc.Server) {
		novelcomment.RegisterNovelCommentServer(srv, novelCommentService)
	})
	list = append(list, func(srv *grpc.Server) {
		novelconsume.RegisterNovelConsumeServer(srv, novelConsumeService)
	})
	list = append(list, func(srv *grpc.Server) {
		novelmsg.RegisterNovelMsgServer(srv, novelMsgService)
	})
	list = append(list, func(srv *grpc.Server) {
		payorder.RegisterPayOrderServer(srv, payOrderService)
	})
	list = append(list, func(srv *grpc.Server) {
		socialuser.RegisterSocialUserServer(srv, socialUserService)
	})
	list = append(list, func(srv *grpc.Server) {
		tasklog.RegisterTaskLogServer(srv, taskLogService)
	})
	list = append(list, func(srv *grpc.Server) {
		userevent.RegisterUserEventServer(srv, userEventService)
	})
	list = append(list, func(srv *grpc.Server) {
		usermsg.RegisterUserMsgServer(srv, userMsgService)
	})
	list = append(list, func(srv *grpc.Server) {
		vipuser.RegisterVipUserServer(srv, vipUserService)
	})
	return list
}
