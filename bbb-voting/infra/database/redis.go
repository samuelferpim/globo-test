package database

import (
	"context"

	interfaces "bbb-voting/internal/core/ports"

	"github.com/go-redis/redis/v8"
)

type RedisDB struct {
	client *redis.Client
}

type RedisPipeline struct {
	pipeline redis.Pipeliner
}

func NewRedisDB(url string) (interfaces.RedisRepository, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	return &RedisDB{client: client}, nil
}

func (r *RedisDB) Close() error {
	return r.client.Close()
}

func (r *RedisDB) Pipeline() interfaces.RedisPipeliner {
	return &RedisPipeline{pipeline: r.client.Pipeline()}
}

func (r *RedisDB) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

func (r *RedisDB) HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd {
	return r.client.HGetAll(ctx, key)
}

func (r *RedisDB) Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.client.Scan(ctx, cursor, match, count)
}

func (p *RedisPipeline) Exec(ctx context.Context) error {
	_, err := p.pipeline.Exec(ctx)
	return err
}

func (p *RedisPipeline) HMSet(ctx context.Context, key string, fields map[string]interface{}) {
	p.pipeline.HMSet(ctx, key, fields)
}

func (p *RedisPipeline) Incr(ctx context.Context, key string) {
	p.pipeline.Incr(ctx, key)
}
