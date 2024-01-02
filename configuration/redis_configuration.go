package configuration

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pithawatnuckong/go-clean/environment"
)

// TODO not support for cluster

func NewRedis(config environment.RedisEnv) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Host, config.Port),
		Password: config.Password,
	})
}
