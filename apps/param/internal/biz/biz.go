package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewNovelPayConfigUseCase,
	NewNovelTagUseCase,
	NewPageConfigUseCase,
	NewQiniuConfigUseCase,
	NewResourceGroupUseCase,
	NewResourceStorageUseCase,
	NewScoreProductUseCase,
	NewTaskUseCase,
	NewUserAnalysisStatisticsUseCase,
	NewUserConsumeUseCase,
	NewUserResourceUseCase,
	NewUserResourceRecordUseCase,
	NewVipTypeUseCase,
)
