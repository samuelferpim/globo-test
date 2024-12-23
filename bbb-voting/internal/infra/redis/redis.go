package redis

import (
	"context"
	"sync"

	"bbb-voting/internal/repository"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var (
	once        sync.Once
	redisClient *redis.Client
)

type RedisClient struct {
	client *redis.Client
	logger *zap.Logger
}

type RedisPipeliner struct {
	pipeline redis.Pipeliner
}

func NewRedisClient(url string, logger *zap.Logger) (repository.RedisClient, error) {
	var err error
	once.Do(func() {
		opts, parseErr := redis.ParseURL(url)
		if parseErr != nil {
			err = parseErr
			return
		}
		redisClient = redis.NewClient(opts)
	})
	if err != nil {
		return nil, err
	}
	return &RedisClient{
		client: redisClient,
		logger: logger,
	}, nil
}

func (r *RedisClient) Close() error {
	return r.client.Close()
}

func (r *RedisClient) Pipeline() repository.RedisPipeliner {
	return &RedisPipeliner{pipeline: r.client.Pipeline()}
}

func (r *RedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

func (r *RedisClient) HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd {
	return r.client.HGetAll(ctx, key)
}

func (r *RedisClient) Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.client.Scan(ctx, cursor, match, count)
}

func (p *RedisPipeliner) Exec(ctx context.Context) error {
	_, err := p.pipeline.Exec(ctx)
	return err
}

func (p *RedisPipeliner) HMSet(ctx context.Context, key string, fields map[string]interface{}) {
	p.pipeline.HMSet(ctx, key, fields)
}

func (p *RedisPipeliner) Incr(ctx context.Context, key string) {
	p.pipeline.Incr(ctx, key)
}