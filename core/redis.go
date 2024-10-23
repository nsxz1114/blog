package core

import (
	"context"
	"github.com/nsxz1114/blog/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedis() *redis.Client {
	redisConf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       redisConf.DB,
		PoolSize: redisConf.PoolSize,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		global.Log.Fatal("init redis fail", zap.Error(err))
	}
	return rdb
}
