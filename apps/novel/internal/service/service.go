package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewActivityService,
	NewActivityComponentService,
	NewAdChangeLogService,
	NewAdChannelService,
	NewAgreementLogService,
	NewAmBalanceService,
	NewAppVersionService,
	NewAssetChangeLogService,
	NewAssetItemService,
	NewBookPackageService,
	NewClientErrorService,
	NewCustomerNovelConfigService,
	NewCustomerNovelsService,
	NewDataSourceService,
	NewListenRecordService,
	NewNovelService,
	NewNovelAutoBuyService,
	NewNovelBookshelfService,
	NewNovelBuyChapterRecordService,
	NewNovelBuyRecordService,
	NewNovelChapterService,
	NewNovelClassifyService,
	NewNovelCommentService,
	NewNovelConsumeService,
	NewNovelMsgService,
	NewPayOrderService,
	NewSocialUserService,
	NewTaskLogService,
	NewUserEventService,
	NewUserMsgService,
	NewVipUserService,
)
