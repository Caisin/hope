package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewCasbinRuleUseCase,
	NewSysApiUseCase,
	NewSysConfigUseCase,
	NewSysDeptUseCase,
	NewSysDictDataUseCase,
	NewSysDictTypeUseCase,
	NewSysJobUseCase,
	NewSysJobLogUseCase,
	NewSysLoginLogUseCase,
	NewSysMenuUseCase,
	NewSysOperaLogUseCase,
	NewSysPostUseCase,
	NewSysRoleUseCase,
	NewSysUserUseCase,
)
