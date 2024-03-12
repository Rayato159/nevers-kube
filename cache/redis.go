package cache

import "github.com/redis/go-redis/v9"

func ExampleClient() *redis.Client {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		Password:      "admin",
		MasterName:    "mymaster",
		SentinelAddrs: []string{"sentinel-0.sentinel:5000", "sentinel-1.sentinel:5000", "sentinel-2.sentinel:5000"},
	})

	return rdb
}
