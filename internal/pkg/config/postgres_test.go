package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeon-code/tiny-url/internal/pkg/config"
)

func TestPostgresConfiguration(t *testing.T) {
	conf := config.NewPostgresConfig("DB_TEST")

	t.Run("should return database driver", func(t *testing.T) {
		assert.Equal(t, "postgres", conf.GetDriver())
	})

	t.Run("should return database name", func(t *testing.T) {
		os.Setenv("DB_TEST_NAME", "name")
		defer os.Unsetenv("DB_TEST_NAME")

		name, err := conf.DBName()
		assert.NoError(t, err)
		assert.Equal(t, "name", name)
	})

	t.Run("should return error when name is not set", func(t *testing.T) {
		_, err := conf.DBName()
		assert.Error(t, err)
	})

	t.Run("should return database password", func(t *testing.T) {
		os.Setenv("DB_TEST_PASSWORD", "password")
		defer os.Unsetenv("DB_TEST_PASSWORD")

		password, err := conf.DBPassword()
		assert.NoError(t, err)
		assert.Equal(t, "password", password)
	})

	t.Run("should return error when password is not set", func(t *testing.T) {
		_, err := conf.DBPassword()
		assert.Error(t, err)
	})

	t.Run("should return database user", func(t *testing.T) {
		os.Setenv("DB_TEST_USER", "user")
		defer os.Unsetenv("DB_TEST_USER")

		user, err := conf.DBUser()
		assert.NoError(t, err)
		assert.Equal(t, "user", user)
	})

	t.Run("should return error when user is not set", func(t *testing.T) {
		_, err := conf.DBUser()
		assert.Error(t, err)
	})

	t.Run("should return database port", func(t *testing.T) {
		os.Setenv("DB_TEST_PORT", "5555")
		defer os.Unsetenv("DB_TEST_PORT")

		port, err := conf.DBPort()
		assert.NoError(t, err)
		assert.Equal(t, 5555, port)
	})

	t.Run("should return error when port is not an integer", func(t *testing.T) {
		os.Setenv("DB_TEST_PORT", "port")
		defer os.Unsetenv("DB_TEST_PORT")

		_, err := conf.DBPort()
		assert.Error(t, err)
	})

	t.Run("should return error when port is not set", func(t *testing.T) {
		_, err := conf.DBPort()
		assert.Error(t, err)
	})

	t.Run("should return database host", func(t *testing.T) {
		os.Setenv("DB_TEST_HOST", "host")
		defer os.Unsetenv("DB_TEST_HOST")

		host, err := conf.DBHost()
		assert.NoError(t, err)
		assert.Equal(t, "host", host)
	})

	t.Run("should return error when host is not set", func(t *testing.T) {
		_, err := conf.DBHost()
		assert.Error(t, err)
	})

	t.Run("should return database tls mode", func(t *testing.T) {
		os.Setenv("DB_TEST_TLS_MODE", "true")
		defer os.Unsetenv("DB_TEST_TLS_MODE")

		isTLSMode, err := conf.DBTLSMode()
		assert.NoError(t, err)
		assert.Equal(t, true, isTLSMode)
	})

	t.Run("should return error when tls mode is not a boolean", func(t *testing.T) {
		os.Setenv("DB_TEST_TLS_MODE", "tls_mode")
		defer os.Unsetenv("DB_TEST_TLS_MODE")

		_, err := conf.DBHost()
		assert.Error(t, err)
	})

	t.Run("should return error when tls mode is not set", func(t *testing.T) {
		_, err := conf.DBHost()
		assert.Error(t, err)
	})
}
