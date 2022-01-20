package data

import "github.com/google/wire"

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewEntClient,
	NewRedisClient,
	NewData,
	NewCasbinRuleRepo,
	NewSysApiRepo,
	NewSysColumnsRepo,
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
	NewSysTablesRepo,
	NewSysUserRepo,
)