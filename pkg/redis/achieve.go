package redis

import (
	"time"
	"wrblog-api-go/pkg/mylog"
)

// Set 插入
func Set(key string, value any) bool {
	err := redisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		mylog.MyLog.Panic("redis缓存失败：%s", err)
		return false
	}
	return true
}

func GetSet(key string, value any) bool {
	err := redisClient.GetSet(ctx, key, value).Err()
	if err != nil {
		mylog.MyLog.Panic("redis缓存失败：%s", err)
		return false
	}
	return true
}

// SetTime 插入（带过期时间）
func SetTime(key string, value any, expiration time.Duration) bool {
	err := redisClient.SetEX(ctx, key, value, expiration).Err()
	if err != nil {
		mylog.MyLog.Panic("redis缓存失败：%s", err)
		return false
	}
	if err != nil {
		return false
	}
	return true
}

// RefTime 刷新过期时间
func RefTime(key string, expiration time.Duration) bool {
	_, err := redisClient.Expire(ctx, key, expiration).Result()
	if err != nil {
		return false
	}
	return true
}

// Get 获取
func Get(key string) (val []byte, err error) {
	val, err = redisClient.Get(ctx, key).Bytes()
	return val, err
}

// Get 获取(通配符)
func ScanKeys(key string) (keys []string, err error) {
	cursor := uint64(0)
	count := int64(10000)
	keys, cursor, err = redisClient.Scan(ctx, cursor, key, count).Result()
	if err != nil {
		return nil, err
	}
	return keys, err
}

// Get 获取
func ScanVals(keyM string) (vals [][]byte, err error) {
	keys, err := ScanKeys(keyM)
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		val, errIn := Get(key)
		if errIn != nil {
			return nil, errIn
		} else {
			vals = append(vals, val)
		}
	}
	return vals, err
}

// Del 删除
func Del(key ...string) bool {
	_, err := redisClient.Del(ctx, key...).Result()
	if err != nil {
		return false
	}
	return true
}
