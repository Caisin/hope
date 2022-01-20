package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewActivityUseCase,
	NewActivityComponentUseCase,
	NewAdChangeLogUseCase,
	NewAdChannelUseCase,
	NewAgreementLogUseCase,
	NewAmBalanceUseCase,
	NewAppVersionUseCase,
	NewAssetChangeLogUseCase,
	NewAssetItemUseCase,
	NewBookPackageUseCase,
	NewClientErrorUseCase,
	NewCustomerNovelConfigUseCase,
	NewCustomerNovelsUseCase,
	NewDataSourceUseCase,
	NewListenRecordUseCase,
	NewNovelUseCase,
	NewNovelAutoBuyUseCase,
	NewNovelBookshelfUseCase,
	NewNovelBuyChapterRecordUseCase,
	NewNovelBuyRecordUseCase,
	NewNovelChapterUseCase,
	NewNovelClassifyUseCase,
	NewNovelCommentUseCase,
	NewNovelConsumeUseCase,
	NewNovelMsgUseCase,
	NewPayOrderUseCase,
	NewSocialUserUseCase,
	NewTaskLogUseCase,
	NewUserEventUseCase,
	NewUserMsgUseCase,
	NewVipUserUseCase,
)
