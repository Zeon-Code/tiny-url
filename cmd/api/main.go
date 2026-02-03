package main

import (
	"net/http"

	h "github.com/zeon-code/tiny-url/internal/http"
	"github.com/zeon-code/tiny-url/internal/pkg/config"
	"github.com/zeon-code/tiny-url/internal/pkg/log"
	"github.com/zeon-code/tiny-url/internal/repository"
	"github.com/zeon-code/tiny-url/internal/service"
)

func main() {
	conf := config.NewConfiguration()
	logger := log.NewLogger(conf)
	repo := repository.NewRepositories(conf, logger.With("package", "repository"))
	svc := service.NewServices(conf, logger.With("package", "service"), repo)

	server := &http.Server{
		Addr:    ":8080",
		Handler: h.NewRouter(svc, logger.With("package", "handler")),
	}

	logger.Info("Starting server")
	server.ListenAndServe()
}
