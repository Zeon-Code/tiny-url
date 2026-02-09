package observability

import (
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

func NewInstrumentedRedis(observer Observer, dsn string) (*redis.Client, error) {
	opt, err := redis.ParseURL(dsn)

	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(opt)

	if err := redisotel.InstrumentTracing(rdb); err != nil {
		return nil, err
	}

	if err := redisotel.InstrumentMetrics(rdb); err != nil {
		return nil, err
	}

	return rdb, err
}
