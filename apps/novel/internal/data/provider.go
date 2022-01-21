package data

import "github.com/google/wire"

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewEntClient,
	NewRedisClient,
	NewData,
	NewActivityRepo,
	NewActivityComponentRepo,
	NewAdChangeLogRepo,
	NewAdChannelRepo,
	NewAgreementLogRepo,
	NewAmBalanceRepo,
	NewAppVersionRepo,
	NewAssetChangeLogRepo,
	NewAssetItemRepo,
	NewBookPackageRepo,
	NewClientErrorRepo,
	NewCustomerNovelConfigRepo,
	NewCustomerNovelsRepo,
	NewDataSourceRepo,
	NewListenRecordRepo,
	NewNovelRepo,
	NewNovelAutoBuyRepo,
	NewNovelBookshelfRepo,
	NewNovelBuyChapterRecordRepo,
	NewNovelBuyRecordRepo,
	NewNovelChapterRepo,
	NewNovelClassifyRepo,
	NewNovelCommentRepo,
	NewNovelConsumeRepo,
	NewNovelMsgRepo,
	NewPayOrderRepo,
	NewSocialUserRepo,
	NewTaskLogRepo,
	NewUserEventRepo,
	NewUserMsgRepo,
	NewVipUserRepo,

)
