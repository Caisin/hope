// Package sortedset https://www.tizi365.com/archives/304.html
package sortedset

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
)

// ZRank 根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
func ZRank(ctx context.Context, client *redis.Client, key string, member string) (int64, error) {
	// 查询集合元素member的分数
	return client.ZRank(ctx, key, member).Result()
}

// ZRevRank 根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从大到小排序
func ZRevRank(ctx context.Context, client *redis.Client, key string, member string) (int64, error) {
	// 查询集合元素member的分数
	return client.ZRevRank(ctx, key, member).Result()
}

// ZScore 查询元素对应的分数
func ZScore(ctx context.Context, client *redis.Client, key string, member string) (float64, error) {
	// 查询集合元素member的分数
	return client.ZScore(ctx, key, member).Result()
}

// ZRem 删除集合元素
func ZRem(ctx context.Context, client *redis.Client, key string, members ...interface{}) error {
	err := client.ZRem(ctx, key, members...).Err()
	return err
}

// ZRemRangeByRank 根据索引范围删除元素
func ZRemRangeByRank(ctx context.Context, client *redis.Client, key string, start, stop int64) error {
	//// 集合元素按分数排序，从最低分到高分，删除第0个元素到第5个元素。
	//// 这里相当于删除最低分的几个元素
	//// 位置参数写成负数，代表从高分开始删除。
	//// 这个例子，删除最高分数的两个元素，-1代表最高分数的位置，-2第二高分，以此类推。
	//client.ZRemRangeByRank("key", -1, -2)
	err := client.ZRemRangeByRank(ctx, key, start, stop).Err()
	return err
}

// ZRemRangeByScore 根据分数范围删除元素
func ZRemRangeByScore(ctx context.Context, client *redis.Client, key string, start, stop string) error {
	//// 删除范围： 2<=分数<=5 的元素
	//client.ZRemRangeByScore("key", "2", "5")
	//
	//// 删除范围： 2<=分数<5 的元素
	//client.ZRemRangeByScore("key", "2", "(5")
	err := client.ZRemRangeByScore(ctx, key, start, stop).Err()
	return err
}

func ZAddOne(ctx context.Context, client *redis.Client, key string, score float64, member interface{}) error {
	return ZAdd(ctx, client, key, &redis.Z{Score: score, Member: member})
}

// ZAdd 添加一个或者多个元素到集合，如果元素已经存在则更新分数
func ZAdd(ctx context.Context, client *redis.Client, key string, members ...*redis.Z) error {
	err := client.ZAdd(ctx, key, members...).Err()
	return err
}

// ZCard 返回集合元素个数
func ZCard(ctx context.Context, client *redis.Client, key string) (int64, error) {
	return client.ZCard(ctx, key).Result()
}

// ZCount 统计某个分数范围内的元素个数
func ZCount(ctx context.Context, client *redis.Client, key string, min, max float64) (int64, error) {
	//返回： 1<=分数<=5 的元素个数, 注意："1", "5"两个参数是字符串
	return client.ZCount(ctx, key, cast.ToString(min), cast.ToString(max)).Result()
}

// ZIncrBy 增加元素的分数
func ZIncrBy(ctx context.Context, client *redis.Client, key, member string, increment float64) (float64, error) {
	//给元素member，加上increment分
	return client.ZIncrBy(ctx, key, increment, member).Result()
}

// ZRange 返回集合中某个索引范围的元素，根据分数从小到大排序
func ZRange(ctx context.Context, client *redis.Client, key string, start, stop int64) ([]string, error) {
	//// 返回从0到-1位置的集合元素， 元素按分数从小到大排序
	//// 0到-1代表则返回全部数据
	//vals, err := client.ZRange("key", 0,-1).Result()
	//if err != nil {
	//	panic(err)
	//}
	return client.ZRange(ctx, key, start, stop).Result()
}

// ZRevRange 返回集合中某个索引范围的元素，根据分数从大到小排序
func ZRevRange(ctx context.Context, client *redis.Client, key string, start, stop int64) ([]string, error) {
	//// 返回从0到-1位置的集合元素， 元素按分数从大到小排序
	//// 0到-1代表则返回全部数据
	//vals, err := client.ZRevRange("key", 0,-1).Result()
	//if err != nil {
	//	panic(err)
	//}
	return client.ZRevRange(ctx, key, start, stop).Result()
}

// ZRangeWithScores 返回集合中某个索引范围的元素和分数，根据分数从小到大排序
func ZRangeWithScores(ctx context.Context, client *redis.Client, key string, start, stop int64) ([]redis.Z, error) {
	//// 返回从0到-1位置的集合元素， 元素按分数从小到大排序
	//// 0到-1代表则返回全部数据
	//vals, err := client.ZRange("key", 0,-1).Result()
	//if err != nil {
	//	panic(err)
	//}
	return client.ZRangeWithScores(ctx, key, start, stop).Result()
}

// ZRevRangeWithScores 返回集合中某个索引范围的元素和分数，根据分数从大到小排序
func ZRevRangeWithScores(ctx context.Context, client *redis.Client, key string, start, stop int64) ([]redis.Z, error) {
	//// 返回从0到-1位置的集合元素， 元素按分数从大到小排序
	//// 0到-1代表则返回全部数据
	//vals, err := client.ZRevRange("key", 0,-1).Result()
	//if err != nil {
	//	panic(err)
	//}
	return client.ZRevRangeWithScores(ctx, key, start, stop).Result()
}

// ZRangeByScore 返回集合中某个分数范围的元素，根据分数从小到大排序
func ZRangeByScore(ctx context.Context, client *redis.Client, key string, min, max float64, offset, count int64) ([]string, error) {
	op := redis.ZRangeBy{
		Min:    cast.ToString(min), // 最小分数
		Max:    cast.ToString(max), // 最大分数
		Offset: offset,             // 类似sql的limit, 表示开始偏移量
		Count:  count,              // 一次返回多少数据
	}
	return client.ZRangeByScore(ctx, key, &op).Result()
}

// ZRevRangeByScore 返回集合中某个分数范围的元素，根据分数从大到小排序
func ZRevRangeByScore(ctx context.Context, client *redis.Client, key string, min, max float64, offset, count int64) ([]string, error) {
	op := redis.ZRangeBy{
		Min:    cast.ToString(min), // 最小分数
		Max:    cast.ToString(max), // 最大分数
		Offset: offset,             // 类似sql的limit, 表示开始偏移量
		Count:  count,              // 一次返回多少数据
	}
	return client.ZRevRangeByScore(ctx, key, &op).Result()
}

// ZRangeByScoreWithScores 返回集合中某个分数范围的元素和分数，根据分数从小到大排序
func ZRangeByScoreWithScores(ctx context.Context, client *redis.Client, key string, min, max float64, offset, count int64) ([]redis.Z, error) {
	op := redis.ZRangeBy{
		Min:    cast.ToString(min), // 最小分数
		Max:    cast.ToString(max), // 最大分数
		Offset: offset,             // 类似sql的limit, 表示开始偏移量
		Count:  count,              // 一次返回多少数据
	}
	return client.ZRangeByScoreWithScores(ctx, key, &op).Result()
}

// ZRevRangeByScoreWithScores 返回集合中某个分数范围的元素和分数，根据分数从大到小排序
func ZRevRangeByScoreWithScores(ctx context.Context, client *redis.Client, key string, min, max float64, offset, count int64) ([]redis.Z, error) {
	op := redis.ZRangeBy{
		Min:    cast.ToString(min), // 最小分数
		Max:    cast.ToString(max), // 最大分数
		Offset: offset,             // 类似sql的limit, 表示开始偏移量
		Count:  count,              // 一次返回多少数据
	}
	return client.ZRevRangeByScoreWithScores(ctx, key, &op).Result()
}
