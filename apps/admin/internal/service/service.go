package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewCasbinRuleService,
	NewSysApiService,
	NewSysConfigService,
	NewSysDeptService,
	NewSysDictDataService,
	NewSysDictTypeService,
	NewSysJobService,
	NewSysJobLogService,
	NewSysLoginLogService,
	NewSysMenuService,
	NewSysOperaLogService,
	NewSysPostService,
	NewSysRoleService,
	NewSysUserService,
	NewAuthService,
)
