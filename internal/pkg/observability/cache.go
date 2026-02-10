package observability

import (
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

func NewInstrumentedRedis(opt *redis.Options, observer Observer) (*redis.Client, error) {
	rdb := redis.NewClient(opt)

	if err := redisotel.InstrumentTracing(rdb); err != nil {
		return nil, err
	}

	if err := redisotel.InstrumentMetrics(rdb); err != nil {
		return nil, err
	}

	return rdb, nil
}
