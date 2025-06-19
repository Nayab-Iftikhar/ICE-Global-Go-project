package infrastructure

import (
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(config *redis.Options) *redis.Client {
	return redis.NewClient(config)
}
