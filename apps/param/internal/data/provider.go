package data

import "github.com/google/wire"

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewEntClient,
	NewRedisClient,
	NewData,
	NewNovelPayConfigRepo,
	NewNovelTagRepo,
	NewPageConfigRepo,
	NewQiniuConfigRepo,
	NewResourceGroupRepo,
	NewResourceStorageRepo,
	NewScoreProductRepo,
	NewTaskRepo,
	NewUserAnalysisStatisticsRepo,
	NewUserConsumeRepo,
	NewUserResourceRepo,
	NewUserResourceRecordRepo,
	NewVipTypeRepo,
)
