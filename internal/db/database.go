package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/zeon-code/tiny-url/internal/pkg/config"
	"github.com/zeon-code/tiny-url/internal/pkg/observability"
)

type SQLReader interface {
	Close() error

	Select(ctx context.Context, value any, query string, args ...any) error
	Get(ctx context.Context, value any, query string, args ...any) error
}

type SQLTX interface {
	Commit() error
	Rollback() error

	Select(ctx context.Context, value any, query string, args ...any) error
	Get(ctx context.Context, value any, query string, args ...any) error
	Exec(ctx context.Context, query string, args ...any) error
}

type SQLClient interface {
	SQLReader

	Exec(ctx context.Context, query string, args ...any) error
	BeginTx(ctx context.Context, opt *sql.TxOptions) (SQLTX, error)
}

func NewDBClient(c config.DatabaseConfiguration, observer observability.Observer) (SQLClient, error) {
	return NewPostgresClientFromConfig(c, observer)
}

type CacheClient interface {
	Del(ctx context.Context, key string) error
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, value any, key string, ttl time.Duration) error
	Incr(ctx context.Context, key string) (int64, error)
	Close() error
}

func NewCacheClient(c config.DatabaseConfiguration, observer observability.Observer) (CacheClient, error) {
	return NewRedisClientFromConfig(c, observer)
}
