package auth

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"hope/pkg/redis/bitop"
	"hope/pkg/redis/rediscache"
	"sync"
	"time"
)

var (
	operMapping = make(map[string]int64)
	permMapping = make(map[string]int64)
	whiteList   = make(map[int64]bool)
	refreshTime = time.Now()
	mutex       = new(sync.Mutex)
)

const (
	permMappingKey = "permission:mapping:permission_code"
	operMappingKey = "permission:mapping:operation_code"
	whiteListKey   = "permission:white_list"
)

func MatchPer(hasPerms []string, target string) bool {
	if len(hasPerms) == 0 {
		return false
	}
	for _, perm := range hasPerms {
		if perm == target {
			return true
		}
	}
	return false
}

//权限逻辑
// 1.系统启动,加载operation和权限编码的映射
// 2.用户登陆,刷新和设置用户权限编码是否有权限,通过Redis setBit方法实现
// 3.根据operation渠道权限编码,通过Redis getBit 去对应编码的值,判断是否有当前权限
func genPermKeyByUser(userId int64) string {
	return fmt.Sprintf("user_perm:%d", userId)
}

// HasOperationPer 根据operation判断权限
func HasOperationPer(ctx context.Context, rdc *redis.Client, opera string) bool {
	//白名单过滤
	id := GetIdByOper(opera)
	if whiteList[id] {
		return true
	}
	//用户
	claims, err := GetClaims(ctx)
	if err != nil {
		return false
	}
	userId := claims.UserId
	if userId == 1 {
		return true
	}
	return bitop.GetBit(ctx, rdc, genPermKeyByUser(userId), id) == 1
}

// SetPer 设置用权限缓存
func SetPer(ctx context.Context, rdc *redis.Client, userId, perId int64) int64 {
	return bitop.SetBit(ctx, rdc, genPermKeyByUser(userId), perId, 1)
}

func GetIdByOper(oper string) int64 {
	return operMapping[oper]
}

func GetIdByPerm(perm string) int64 {
	return permMapping[perm]
}

func RefreshFromRedis(ctx context.Context, rdc *redis.Client) {
	sub := time.Now().Sub(refreshTime)
	if sub > 10*time.Minute {
		mutex.Lock()
		defer mutex.Unlock()
		rediscache.GetJsonData(ctx, rdc, permMappingKey, &permMapping)
		rediscache.GetJsonData(ctx, rdc, operMappingKey, &operMapping)
		rediscache.GetJsonData(ctx, rdc, whiteListKey, &whiteList)
	}

}

// SetPermissionMappingToRedis 设置权限到缓存
func SetPermissionMappingToRedis(
	ctx context.Context,
	rdc *redis.Client,
	perm,
	oper map[string]int64,
	white map[int64]bool,
) {
	mutex.Lock()
	defer mutex.Unlock()
	permMapping = perm
	operMapping = oper
	whiteList = white
	rediscache.SetJson(ctx, rdc, permMappingKey, permMapping, -1)
	rediscache.SetJson(ctx, rdc, operMappingKey, operMapping, -1)
	rediscache.SetJson(ctx, rdc, whiteListKey, whiteList, -1)
}
