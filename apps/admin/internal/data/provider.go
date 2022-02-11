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
	NewCasbinRuleRepo,
	NewSysApiRepo,
	NewSysConfigRepo,
	NewSysDeptRepo,
	NewSysDictDataRepo,
	NewSysDictTypeRepo,
	NewSysJobRepo,
	NewSysJobLogRepo,
	NewSysLoginLogRepo,
	NewSysMenuRepo,
	NewSysOperaLogRepo,
	NewSysPostRepo,
	NewSysRoleRepo,
	NewSysUserRepo,
	NewAuthRepo,
)
