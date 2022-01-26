package rediscache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

// SetBytes 设置缓存
func SetBytes(ctx context.Context, client *redis.Client, key string, data []byte, sec int) error {
	return client.Set(ctx, key, data, time.Duration(sec)*time.Second).Err()
}

// SetStr 设置缓存
func SetStr(ctx context.Context, client *redis.Client, key string, data string, sec int) error {
	return client.Set(ctx, key, data, time.Duration(sec)*time.Second).Err()
}

// SetJson 设置缓存
func SetJson(ctx context.Context, client *redis.Client, key string, data interface{}, sec int) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return SetBytes(ctx, client, key, marshal, sec)
}

// Del 删除缓存
func Del(ctx context.Context, client *redis.Client, key string) error {
	return client.Del(ctx, key).Err()
}

func Dels(ctx context.Context, client *redis.Client, keys []string) {
	for _, key := range keys {
		_ = Del(ctx, client, key)
	}
}

// GetJsonData 获取缓存,将json数据反序列化到data
func GetJsonData(ctx context.Context, client *redis.Client, key string, data interface{}) error {
	bytes, err := client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, data)
}

// Expire 失效时间
func Expire(ctx context.Context, client *redis.Client, key string, dur time.Duration) error {
	return client.Expire(ctx, key, dur).Err()
}

// GetStr 获取缓存
func GetStr(ctx context.Context, client *redis.Client, key string) string {
	ret, err := client.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return ret
}

// GetInt64 获取缓存
func GetInt64(ctx context.Context, client *redis.Client, key string) int64 {
	ret, err := client.Get(ctx, key).Int64()
	if err != nil {
		return 0
	}
	return ret
}

func Keys(ctx context.Context, client *redis.Client, key string) []string {
	keys := client.Keys(ctx, key).Val()
	return keys
}

// DelReg 正则删除
func DelReg(ctx context.Context, client *redis.Client, reg string) {
	keys := Keys(ctx, client, reg)
	Dels(ctx, client, keys)
}
