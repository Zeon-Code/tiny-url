package handler

import (
	"net/http"

	"github.com/zeon-code/tiny-url/internal/pkg/observability"
	"github.com/zeon-code/tiny-url/internal/service"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func NewRouter(svc service.Services, observer observability.Observer) http.Handler {
	mux := http.NewServeMux()

	url := NewUrlHandler(svc, observer)
	health := NewHealthHandler(svc, observer)

	mux.HandleFunc("GET /r/{code}", url.Redirect)

	mux.HandleFunc("GET /api/v1/url/", url.List)
	mux.HandleFunc("POST /api/v1/url/", url.Create)
	mux.HandleFunc("GET /api/v1/url/{id}", url.GetByID)

	mux.HandleFunc("GET /health/ready", health.Ready)
	mux.HandleFunc("GET /health/live", health.Live)

	return otelhttp.NewHandler(mux, "server")
}
