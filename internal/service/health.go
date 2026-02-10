package service

import (
	"context"

	"github.com/zeon-code/tiny-url/internal/pkg/observability"
	"github.com/zeon-code/tiny-url/internal/repository"
)

type HealthService interface {
	Ping(ctx context.Context) (string, error)
}

type HealthSvc struct {
	repo   repository.HealthRepository
	logger observability.Logger
}

func NewHealthService(repositories repository.Repositories, observer observability.Observer) HealthService {
	return HealthSvc{
		repo:   repositories.Health,
		logger: observer.Logger().With("service", "health"),
	}
}

func (s HealthSvc) Ping(ctx context.Context) (string, error) {
	return s.repo.Ping(ctx)
}
