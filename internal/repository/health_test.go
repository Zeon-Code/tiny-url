package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeon-code/tiny-url/internal/pkg/test"
	"github.com/zeon-code/tiny-url/internal/repository"
)

func TestHealthService(t *testing.T) {
	t.Run("ping", func(t *testing.T) {
		fake := test.NewFakeDependencies()
		repo := repository.NewHealthRepository(fake.DB(), fake.Memory(), fake.Observer())

		reason, err := repo.Ping(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, "", reason)
	})
}
