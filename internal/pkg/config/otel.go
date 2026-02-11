package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type MetricConfiguration interface {
	Integration() (string, error)
	Environment() (string, error)
	Host() (string, error)
	Port() (int, error)
}

type OtelConfiguration struct{}

func NewOtelConfiguration() OtelConfiguration {
	return OtelConfiguration{}
}

func (c OtelConfiguration) Environment() (string, error) {
	return c.get("ENV")
}

func (c OtelConfiguration) Integration() (string, error) {
	return c.get("TELEMETRY_INTEGRATION")
}

func (c OtelConfiguration) Host() (string, error) {
	return c.get("TELEMETRY_HOST")
}

func (c OtelConfiguration) Port() (int, error) {
	env, err := c.get("TELEMETRY_PORT")

	if err != nil {
		return 0, err
	}

	port, err := strconv.Atoi(env)

	if err != nil {
		return 0, errors.New("TELEMETRY_PORT must be an interger value")
	}

	return port, nil
}

func (c OtelConfiguration) get(env string) (string, error) {
	if value, exists := os.LookupEnv(env); exists {
		return value, nil
	}

	return "", fmt.Errorf("Missing required environment variable %s", env)
}
