package redis

import "github.com/go-redis/redis/v8"

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "172.17.0.3:6379",
	})

	return rdb
}
