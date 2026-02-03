package db

import (
	"context"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zeon-code/tiny-url/internal/pkg/config"
	"github.com/zeon-code/tiny-url/internal/pkg/metric"
)

// CacheProxy provides a thin abstraction over redis.Client,
// centralizing cache operations and normalizing cache-related
// error handling. It delegates commands to the underlying Redis
// backend while mapping low-level errors to domain-level errors.
type RedisProxy struct {
	backend *redis.Client
	metric  metric.MetricClient
	logger  *slog.Logger
}

func newRedisClient(c config.DatabaseConfiguration, m metric.MetricClient, l *slog.Logger) (*RedisProxy, error) {
	opt, err := redis.ParseURL(c.GetDNS())

	if err != nil {
		return nil, err
	}

	return &RedisProxy{backend: redis.NewClient(opt), metric: m}, err
}

// Get retrieves the cached value associated with the given key.
// If the key exists, the raw cached bytes are returned. If the key
// does not exist returns error.
//
// Returns a mapped cache error for consistent error handling.
func (p RedisProxy) Get(ctx context.Context, key string) ([]byte, error) {
	data, err := p.backend.Get(ctx, key).Bytes()

	if err == redis.Nil {
		p.metric.CacheMiss(key)
		return []byte{}, mapCacheError(err)
	} else if err != nil {
		p.metric.CacheError(key, "failed to read redis key: "+err.Error())
		return []byte{}, mapCacheError(err)
	}

	p.metric.CacheHit(key)
	return data, nil
}

// Set stores the given value in the cache under the provided key with
// the specified TTL. The operation is performed using a set-if-not-exists
// strategy to avoid overwriting existing entries.
//
// Returns a mapped cache error for consistent error handling.
func (p RedisProxy) Set(ctx context.Context, value any, key string, ttl time.Duration) error {
	err := p.backend.SetNX(ctx, key, value, ttl).Err()

	if err != nil {
		p.metric.CacheError(key, "failed to write value into key: "+err.Error())
		return mapCacheError(err)
	}

	return nil
}

// Set stores the given value in the cache under the provided key with
// the specified TTL. The operation is performed using a set-if-not-exists
// strategy to avoid overwriting existing entries.
//
// Returns a mapped cache error for consistent error handling.
func (p RedisProxy) Del(ctx context.Context, key string) error {
	err := p.backend.Del(ctx, key).Err()

	if err != nil {
		p.metric.CacheError(key, "failed to delete redis key: "+err.Error())
		return mapCacheError(err)
	}

	return nil
}

// Incr atomically increments the integer value stored at the given key
// and returns the updated value. If the key does not exist, it is
// initialized before being incremented.
//
// Returns a mapped cache error for consistent error handling.
func (p RedisProxy) Incr(ctx context.Context, key string) (int64, error) {
	current, err := p.backend.Incr(ctx, key).Result()

	if err != nil {
		p.metric.CacheError(key, "failed to increment redis key: "+err.Error())
		return 0, mapCacheError(err)
	}

	return current, nil
}
