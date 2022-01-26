package bitop

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// GetBit 获取值
func GetBit(ctx context.Context, client *redis.Client, key string, offset int64) int64 {
	return client.GetBit(ctx, key, offset).Val()
}

func SetBit(ctx context.Context, client *redis.Client, key string, offset int64, value int) int64 {
	return client.SetBit(ctx, key, offset, value).Val()
}

// BitCount 开始结束是按字符来算的,不是位算的,一个偏移为8
func BitCount(ctx context.Context, client *redis.Client, key string, start, end int64) int64 {
	count := redis.BitCount{
		Start: start,
		End:   end,
	}
	return client.BitCount(ctx, key, &count).Val()
}

func Count(ctx context.Context, client *redis.Client, key string) int64 {
	return client.BitCount(ctx, key, nil).Val()
}

func Or(ctx context.Context, client *redis.Client, key string, keys []string) int64 {
	_ = client.BitOpOr(ctx, key, keys...).Val()
	return Count(ctx, client, key)
}
