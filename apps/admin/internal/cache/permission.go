package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/sysmenu"
	"hope/pkg/auth"
	"hope/pkg/util/str"
)

func InitPermission(ctx context.Context, rdb *redis.Client, entc *ent.Client) error {
	menus, err := entc.
		SysMenu.
		Query().
		Select(
			sysmenu.FieldID,
			sysmenu.FieldOperation,
			sysmenu.FieldPermission,
			sysmenu.FieldCheckPermission,
		).All(ctx)
	if err != nil {
		return err
	}
	opMap := make(map[string]int64)
	permMap := make(map[string]int64)
	whiteListMap := make(map[int64]bool)
	for _, menu := range menus {
		if str.IsNotBlank(menu.Operation) {
			opMap[menu.Operation] = menu.ID
		}
		if str.IsNotBlank(menu.Permission) {
			permMap[menu.Permission] = menu.ID
		}
		if !menu.CheckPermission {
			whiteListMap[menu.ID] = true
		}
	}
	auth.SetPermissionMappingToRedis(ctx, rdb, permMap, opMap, whiteListMap)
	return nil
}
