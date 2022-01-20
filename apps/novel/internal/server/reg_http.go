package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"

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

func RegisterHTTPServer(

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
) []func(*http.Server) {
	list := make([]func(*http.Server), 0)

	list = append(list, func(srv *http.Server) {
		activity.RegisterActivityHTTPServer(srv, activityService)
	})
	list = append(list, func(srv *http.Server) {
		activitycomponent.RegisterActivityComponentHTTPServer(srv, activityComponentService)
	})
	list = append(list, func(srv *http.Server) {
		adchangelog.RegisterAdChangeLogHTTPServer(srv, adChangeLogService)
	})
	list = append(list, func(srv *http.Server) {
		adchannel.RegisterAdChannelHTTPServer(srv, adChannelService)
	})
	list = append(list, func(srv *http.Server) {
		agreementlog.RegisterAgreementLogHTTPServer(srv, agreementLogService)
	})
	list = append(list, func(srv *http.Server) {
		ambalance.RegisterAmBalanceHTTPServer(srv, amBalanceService)
	})
	list = append(list, func(srv *http.Server) {
		appversion.RegisterAppVersionHTTPServer(srv, appVersionService)
	})
	list = append(list, func(srv *http.Server) {
		assetchangelog.RegisterAssetChangeLogHTTPServer(srv, assetChangeLogService)
	})
	list = append(list, func(srv *http.Server) {
		assetitem.RegisterAssetItemHTTPServer(srv, assetItemService)
	})
	list = append(list, func(srv *http.Server) {
		bookpackage.RegisterBookPackageHTTPServer(srv, bookPackageService)
	})
	list = append(list, func(srv *http.Server) {
		clienterror.RegisterClientErrorHTTPServer(srv, clientErrorService)
	})
	list = append(list, func(srv *http.Server) {
		customernovelconfig.RegisterCustomerNovelConfigHTTPServer(srv, customerNovelConfigService)
	})
	list = append(list, func(srv *http.Server) {
		customernovels.RegisterCustomerNovelsHTTPServer(srv, customerNovelsService)
	})
	list = append(list, func(srv *http.Server) {
		datasource.RegisterDataSourceHTTPServer(srv, dataSourceService)
	})
	list = append(list, func(srv *http.Server) {
		listenrecord.RegisterListenRecordHTTPServer(srv, listenRecordService)
	})
	list = append(list, func(srv *http.Server) {
		novel.RegisterNovelHTTPServer(srv, novelService)
	})
	list = append(list, func(srv *http.Server) {
		novelautobuy.RegisterNovelAutoBuyHTTPServer(srv, novelAutoBuyService)
	})
	list = append(list, func(srv *http.Server) {
		novelbookshelf.RegisterNovelBookshelfHTTPServer(srv, novelBookshelfService)
	})
	list = append(list, func(srv *http.Server) {
		novelbuychapterrecord.RegisterNovelBuyChapterRecordHTTPServer(srv, novelBuyChapterRecordService)
	})
	list = append(list, func(srv *http.Server) {
		novelbuyrecord.RegisterNovelBuyRecordHTTPServer(srv, novelBuyRecordService)
	})
	list = append(list, func(srv *http.Server) {
		novelchapter.RegisterNovelChapterHTTPServer(srv, novelChapterService)
	})
	list = append(list, func(srv *http.Server) {
		novelclassify.RegisterNovelClassifyHTTPServer(srv, novelClassifyService)
	})
	list = append(list, func(srv *http.Server) {
		novelcomment.RegisterNovelCommentHTTPServer(srv, novelCommentService)
	})
	list = append(list, func(srv *http.Server) {
		novelconsume.RegisterNovelConsumeHTTPServer(srv, novelConsumeService)
	})
	list = append(list, func(srv *http.Server) {
		novelmsg.RegisterNovelMsgHTTPServer(srv, novelMsgService)
	})
	list = append(list, func(srv *http.Server) {
		payorder.RegisterPayOrderHTTPServer(srv, payOrderService)
	})
	list = append(list, func(srv *http.Server) {
		socialuser.RegisterSocialUserHTTPServer(srv, socialUserService)
	})
	list = append(list, func(srv *http.Server) {
		tasklog.RegisterTaskLogHTTPServer(srv, taskLogService)
	})
	list = append(list, func(srv *http.Server) {
		userevent.RegisterUserEventHTTPServer(srv, userEventService)
	})
	list = append(list, func(srv *http.Server) {
		usermsg.RegisterUserMsgHTTPServer(srv, userMsgService)
	})
	list = append(list, func(srv *http.Server) {
		vipuser.RegisterVipUserHTTPServer(srv, vipUserService)
	})
	return list
}
