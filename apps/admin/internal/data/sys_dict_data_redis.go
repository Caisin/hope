package data

import (
	"context"
	"fmt"
	"hope/apps/admin/internal/data/ent"
	"hope/pkg/redis/rediscache"
)

const (
	sysDictCacheKey = "sys_dict_data:%s"
	cacheSec        = -1
)

func (r *sysDictDataRepo) GetByTypeCache(ctx context.Context, typeCode string, list *[]*ent.SysDictData) error {
	return rediscache.GetJsonData(ctx, r.data.rdb, r.getCacheKey(typeCode), list)
}

func (r *sysDictDataRepo) SetByTypeCache(ctx context.Context, typeCode string, list *[]*ent.SysDictData) error {
	return rediscache.SetJson(ctx, r.data.rdb, r.getCacheKey(typeCode), list, cacheSec)
}

func (r *sysDictDataRepo) DelByTypeCache(ctx context.Context, typeCode string) error {
	return rediscache.Del(ctx, r.data.rdb, r.getCacheKey(typeCode))
}

func (r *sysDictDataRepo) getCacheKey(typeCode string) string {
	return fmt.Sprintf(sysDictCacheKey, typeCode)
}
