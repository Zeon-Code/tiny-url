package service

import (
	"github.com/zeon-code/tiny-url/internal/pkg/observability"
	"github.com/zeon-code/tiny-url/internal/repository"
)

type Services struct {
	Url URLService
}

func NewServices(r repository.Repositories, observer observability.Observer) Services {
	return Services{
		Url: NewUrlService(r, observer),
	}
}
