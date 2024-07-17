package initialize

import (
	"context"
	"github.com/redis/go-redis/v9"
	"erp_api/gin_admin/app/global"
)

func Redis() {
	redisCfg := global.CONFIG.Redis
	var client redis.UniversalClient
	// 使用集群模式
	if redisCfg.UseCluster {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    redisCfg.ClusterAddrs,
			Password: redisCfg.Password,
		})
	} else {
		// 使用单例模式
		client = redis.NewClient(&redis.Options{
			Addr:     redisCfg.Addr,
			Password: redisCfg.Password,
			DB:       redisCfg.DB,
		})
	}
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		// global.GVA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
		println("redis connect ping failed, err:", err)
		panic(err)
	} else {
		println("redis connect ping response:", pong)
		// global.GVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.REDIS = client
	}
}
