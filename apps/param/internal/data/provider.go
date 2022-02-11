package data

import (
	"github.com/google/wire"
	"hope/pkg/provider"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	provider.NewRedisClient,
	NewEntClient,
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
