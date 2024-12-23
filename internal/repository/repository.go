package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisPipeliner interface {
	Exec(ctx context.Context) error
	HMSet(ctx context.Context, key string, fields map[string]interface{})
	Incr(ctx context.Context, key string)
	HIncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd
}

type RedisClient interface {
	Pipeline() RedisPipeliner
	Get(ctx context.Context, key string) *redis.StringCmd
	HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd
	Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd
	Close() error
}
