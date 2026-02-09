package test

import (
	"context"

	"github.com/zeon-code/tiny-url/internal/pkg/observability"
	"go.opentelemetry.io/otel/metric"
)

type FakeObserver struct {
	metrics observability.MetricClient
}

func NewFakeObserver(metrics observability.MetricClient) *FakeObserver {
	return &FakeObserver{
		metrics: metrics,
	}
}

func (o *FakeObserver) Startup(ctx context.Context) error {
	return nil
}

func (o *FakeObserver) Shutdown(ctx context.Context) error {
	return nil
}

func (o *FakeObserver) Logger() observability.Logger {
	return FakeLogger{}
}

func (o *FakeObserver) Metric() (observability.MetricClient, error) {
	return o.metrics, nil
}

func (o *FakeObserver) RegisterDB(dbStats metric.Registration) {}
