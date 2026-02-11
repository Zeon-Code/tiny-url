package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeon-code/tiny-url/internal/pkg/config"
)

func TestOtelConfiguration(t *testing.T) {
	conf := config.NewOtelConfiguration()

	t.Run("should return metric environment", func(t *testing.T) {
		os.Setenv("ENV", "dev")
		defer os.Unsetenv("ENV")

		host, err := conf.Environment()
		assert.NoError(t, err)
		assert.Equal(t, "dev", host)
	})

	t.Run("should return error when environment is not set", func(t *testing.T) {
		_, err := conf.Environment()
		assert.Error(t, err)
	})

	t.Run("should return metric integration", func(t *testing.T) {
		os.Setenv("TELEMETRY_INTEGRATION", "integration")
		defer os.Unsetenv("TELEMETRY_INTEGRATION")

		host, err := conf.Integration()
		assert.NoError(t, err)
		assert.Equal(t, "integration", host)
	})

	t.Run("should return error when integration is not set", func(t *testing.T) {
		_, err := conf.Integration()
		assert.Error(t, err)
	})

	t.Run("should return metric host", func(t *testing.T) {
		os.Setenv("TELEMETRY_HOST", "host")
		defer os.Unsetenv("TELEMETRY_HOST")

		host, err := conf.Host()
		assert.NoError(t, err)
		assert.Equal(t, "host", host)
	})

	t.Run("should return error when host is not set", func(t *testing.T) {
		_, err := conf.Host()
		assert.Error(t, err)
	})

	t.Run("should return metric port", func(t *testing.T) {
		os.Setenv("TELEMETRY_PORT", "1")
		defer os.Unsetenv("TELEMETRY_PORT")

		port, err := conf.Port()
		assert.NoError(t, err)
		assert.Equal(t, 1, port)
	})

	t.Run("should return error when port is not an integer", func(t *testing.T) {
		os.Setenv("TELEMETRY_PORT", "port")
		defer os.Unsetenv("TELEMETRY_PORT")

		_, err := conf.Port()
		assert.Error(t, err)
	})

	t.Run("should return error when port is not set", func(t *testing.T) {
		_, err := conf.Port()
		assert.Error(t, err)
	})
}
