package http

import (
	"log/slog"
	"net/http"

	"github.com/zeon-code/tiny-url/internal/http/handler"
	"github.com/zeon-code/tiny-url/internal/service"
)

func NewRouter(svc service.Services, l *slog.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	url := handler.NewUrlHandler(svc, l.With("handler", "url"))

	mux.HandleFunc("GET /api/v1/url/", url.List)
	mux.HandleFunc("POST /api/v1/url/", url.Create)
	mux.HandleFunc("GET /api/v1/url/{id}", url.GetByID)

	return mux
}
