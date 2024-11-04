package config

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func RedisClient(config *AppConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		DB:   config.Redis.Database,
	})

	return client
}
