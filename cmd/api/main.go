package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zeon-code/tiny-url/internal/http/handler"
	"github.com/zeon-code/tiny-url/internal/pkg/config"
	"github.com/zeon-code/tiny-url/internal/pkg/observability"
	"github.com/zeon-code/tiny-url/internal/repository"
	"github.com/zeon-code/tiny-url/internal/service"
)

var version string = "0.0.1"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	conf := config.NewConfiguration()
	observer := observability.NewObserver(version, conf)

	if err := observer.Startup(ctx); err != nil {
		observer.Logger().Error(ctx, "Error initializing tracer", slog.Any("error", err))
	}

	defer observer.Shutdown(ctx)

	repo := repository.NewRepositoriesFromConfig(conf, observer)
	svc := service.NewServices(repo, observer)

	server := &http.Server{
		Addr:        ":8080",
		BaseContext: func(net.Listener) context.Context { return ctx },
		Handler:     handler.NewRouter(svc, observer),
	}

	go func(ctx context.Context, stop context.CancelFunc, server *http.Server, observer observability.Observer) {
		observer.Logger().Info(ctx, "Starting server", slog.Any("version", version))

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			observer.Logger().Error(ctx, "Error starting server", slog.Any("error", err))
			stop()
		}
	}(ctx, stop, server, observer)

	<-ctx.Done()

	observer.Logger().Info(ctx, "Shutdown initiated")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := repo.Shutdown(); err != nil {
		observer.Logger().Error(ctx, "Error graceful shutdown repositories failed", slog.Any("error", err))
	}

	observer.Logger().Info(ctx, "Repositories shut down gracefully")

	if err := server.Shutdown(ctx); err != nil {
		observer.Logger().Error(ctx, "Error graceful shutdown failed", slog.Any("error", err))
		return
	}

	observer.Logger().Info(ctx, "Server shut down gracefully")
}
