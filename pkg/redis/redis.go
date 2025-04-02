package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v3"
	"os"
	"time"
	"wrblog-api-go/config"
)

// redisConfig 配置映射
type redisConfig struct {
	Redis struct {
		Db           int    `yaml:"db"`
		Addr         string `yaml:"addr"`
		Port         int    `yaml:"port"`
		Password     string `yaml:"password"`
		PoolSize     int    `yaml:"pool-size"`
		MaxRetries   int    `yaml:"max-retries"`
		DialTimeout  int    `yaml:"dial-timeout"`
		ReadTimeout  int    `yaml:"read-timeout"`
		WriteTimeout int    `yaml:"write-timeout"`
	} `yaml:"redis"`
}

var redisClient *redis.Client
var ctx = context.Background()

// 初始化
func init() {
	analysisRedis().initRedis()
	fmt.Println("Redis initialization connection successful......")
}

func analysisRedis() (redisConf *redisConfig) {
	// 打开 YAML 文件
	file, err := os.ReadFile(config.ConfPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse redis configuration:%s", err.Error()))
	}
	err = yaml.Unmarshal(file, &redisConf)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse redis configuration:%s", err.Error()))
	}
	return
}

func (redisConf *redisConfig) initRedis() {
	redisClient = redisConf.getRedisDb()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis initialization connection failed:%s", err.Error()))
	}
}

func (redisConf *redisConfig) getRedisDb() *redis.Client {
	conf := redisConf.Redis
	return redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", conf.Addr, conf.Port), // 连接地址
		Password:     conf.Password,                              // 密码
		DB:           conf.Db,                                    // 数据库编号
		PoolSize:     conf.PoolSize,
		MaxRetries:   conf.MaxRetries,
		DialTimeout:  time.Duration(conf.DialTimeout) * time.Millisecond, // 链接超时
		ReadTimeout:  time.Duration(conf.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(conf.WriteTimeout) * time.Millisecond,
	})
}
