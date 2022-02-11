package data

import (
	"github.com/google/wire"
	"hope/pkg/provider"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	provider.NewRedisClient,
	provider.NewNsqProducer,
	NewEntClient,
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
