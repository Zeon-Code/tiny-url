package observability

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func HTTPError(ctx context.Context, response http.ResponseWriter, statusCode int, err error) {
	span := trace.SpanFromContext(ctx)

	span.RecordError(err, trace.WithStackTrace(true))
	span.SetStatus(codes.Error, http.StatusText(statusCode))

	http.Error(response, http.StatusText(statusCode), statusCode)
}
