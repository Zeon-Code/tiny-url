package observability

import (
	"context"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func TraceError(ctx context.Context, reason string, err error) {
	span := trace.SpanFromContext(ctx)

	span.SetStatus(codes.Error, reason)
	span.RecordError(err, trace.WithStackTrace(true))
}
