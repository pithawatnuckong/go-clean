package configuration

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pithawatnuckong/go-clean/environment"
	"github.com/pithawatnuckong/go-clean/exception"
)

// TODO not support for cluster

func NewRedis(config environment.RedisEnv) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Host, config.Port),
		Password: config.Password,
	})
}

func FindByIdAndSetCache[T any, ID int | string](redisClient *redis.Client, ctx context.Context, prefix string, id ID, findByIdFn func(ctx context.Context, id ID) *T) *T {
	key := fmt.Sprintf("%v:%v", prefix, id)
	var data []byte
	var response T
	if err := redisClient.Get(ctx, key).Scan(&data); err == nil {
		err = json.Unmarshal(data, &response)
		exception.PanicLogging(err)

		return &response
	}

	value := findByIdFn(ctx, id)
	if value == nil {
		panic(exception.ValidationError{
			Message: fmt.Sprintf("%v ID %v not found.", prefix, id),
		})
	}

	valueBytes, err := json.Marshal(value)
	exception.PanicLogging(err)

	err = redisClient.Set(ctx, key, valueBytes, redis.KeepTTL).Err()
	exception.PanicLogging(err)

	return value
}
