package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeon-code/tiny-url/internal/pkg/test"
	"github.com/zeon-code/tiny-url/internal/service"
)

func TestHealthService(t *testing.T) {
	t.Run("ping", func(t *testing.T) {
		fake := test.NewFakeDependencies()
		svc := service.NewHealthService(fake.Repositories(), fake.Observer())

		reason, err := svc.Ping(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, "", reason)
	})
}
