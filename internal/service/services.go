package service

import (
	"log/slog"

	"github.com/zeon-code/tiny-url/internal/pkg/config"
	"github.com/zeon-code/tiny-url/internal/repository"
)

type Services struct {
	Url URLService
}

func NewServices(c config.Configuration, l *slog.Logger, repositories repository.Repositories) Services {
	return Services{
		Url: NewUrlService(repositories, l.With("service", "url")),
	}
}
