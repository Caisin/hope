package data

import (
	"github.com/google/wire"
	entData "hope/apps/novel/internal/data"
	"hope/pkg/provider"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	provider.NewRedisClient,
	provider.NewNsqProducer,
	NewEntClient,
	NewData,
	entData.NewActivityRepo,
	entData.NewActivityComponentRepo,
	entData.NewAdChangeLogRepo,
	entData.NewAdChannelRepo,
	entData.NewAgreementLogRepo,
	entData.NewAmBalanceRepo,
	entData.NewAppVersionRepo,
	entData.NewAssetChangeLogRepo,
	entData.NewAssetItemRepo,
	entData.NewBookPackageRepo,
	entData.NewClientErrorRepo,
	entData.NewCustomerNovelConfigRepo,
	entData.NewCustomerNovelsRepo,
	entData.NewDataSourceRepo,
	entData.NewListenRecordRepo,
	entData.NewNovelRepo,
	entData.NewNovelAutoBuyRepo,
	entData.NewNovelBookshelfRepo,
	entData.NewNovelBuyChapterRecordRepo,
	entData.NewNovelBuyRecordRepo,
	entData.NewNovelChapterRepo,
	entData.NewNovelClassifyRepo,
	entData.NewNovelCommentRepo,
	entData.NewNovelConsumeRepo,
	entData.NewNovelMsgRepo,
	entData.NewPayOrderRepo,
	entData.NewSocialUserRepo,
	entData.NewTaskLogRepo,
	entData.NewUserEventRepo,
	entData.NewUserMsgRepo,
	entData.NewVipUserRepo,
)
