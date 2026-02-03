package db

import (
	"errors"

	"github.com/redis/go-redis/v9"
)

var (
	ErrDBInvalidBackend = errors.New("error db invalid backend instance")
)

func mapDBError(err error) error {
	return err
}

var (
	ErrCacheNotFound    = errors.New("error cache not found")
	ErrCacheUnavailable = errors.New("error cache unavailable")
)

func mapCacheError(err error) error {
	switch {
	case errors.Is(err, redis.Nil):
		return ErrCacheNotFound
	}

	return ErrCacheUnavailable
}
