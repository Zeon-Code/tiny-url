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

	mux.HandleFunc("GET /api/v1/url/", url.List)
	mux.HandleFunc("POST /api/v1/url/", url.Create)
	mux.HandleFunc("GET /api/v1/url/{id}", url.GetByID)

	return otelhttp.NewHandler(mux, "server")
}
