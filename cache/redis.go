package cache

import (
	"github.com/Rayato159/nevers-kube/config"
	"github.com/redis/go-redis/v9"
)

func ExampleClient(conf *config.RedisConfig) *redis.Client {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		Password:      conf.Password,
		MasterName:    conf.MasterName,
		SentinelAddrs: []string{conf.SentinelAddrs},
	})

	return rdb
}
