package db

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	"github.com/zeon-code/tiny-url/internal/pkg/config"
	"github.com/zeon-code/tiny-url/internal/pkg/metric"
)

type SQLReader interface {
	Select(ctx context.Context, value any, query string, args ...any) error
	Get(ctx context.Context, value any, query string, args ...any) error
}

type SQLWrite interface {
	Exec(ctx context.Context, query string, args ...any) error
}

type SQLTX interface {
	SQLReader
	SQLWrite

	Commit() error
	Rollback() error
}

type SQLClient interface {
	SQLReader
	SQLWrite

	BeginTx(ctx context.Context, opt *sql.TxOptions) (SQLTX, error)
}

func NewDBClient(c config.DatabaseConfiguration, m metric.MetricClient, l *slog.Logger) (SQLClient, error) {
	return newPostgresClient(c, m, l)
}

type CacheClient interface {
	Del(ctx context.Context, key string) error
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, value any, key string, ttl time.Duration) error
	Incr(ctx context.Context, key string) (int64, error)
}

func NewCacheClient(c config.DatabaseConfiguration, m metric.MetricClient, l *slog.Logger) (CacheClient, error) {
	return newRedisClient(c, m, l)
}
