package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewNovelPayConfigService,
	NewNovelTagService,
	NewPageConfigService,
	NewQiniuConfigService,
	NewResourceGroupService,
	NewResourceStorageService,
	NewScoreProductService,
	NewTaskService,
	NewUserAnalysisStatisticsService,
	NewUserConsumeService,
	NewUserResourceService,
	NewUserResourceRecordService,
	NewVipTypeService,
)
