package service

import (
	"github.com/zeon-code/tiny-url/internal/pkg/observability"
	"github.com/zeon-code/tiny-url/internal/repository"
)

type Services struct {
	Url    URLService
	Health HealthService
}

func NewServices(repo repository.Repositories, observer observability.Observer) Services {
	return Services{
		Url:    NewUrlService(repo, observer),
		Health: NewHealthService(repo, observer),
	}
}
