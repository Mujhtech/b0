package cache

import (
	"context"
	"time"

	"github.com/mujhtech/b0/config"
	"github.com/mujhtech/b0/internal/redis"
)

type Cache interface {
	Get(ctx context.Context, key string, value interface{}) error
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
}

func NewCache(cfg *config.Config, redis *redis.Redis) (Cache, error) {
	switch cfg.Cache.Provider {
	case config.CacheProviderRedis:
		return NewRedisCache(cfg, redis)
	default:
		return nil, nil
	}
}
