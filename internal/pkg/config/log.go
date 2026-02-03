package config

import (
	"os"
)

type Log interface {
	Level() string
}

type LogConfig struct{}

func newLogConfig() Log {
	return LogConfig{}
}

func (c LogConfig) Level() string {
	value, exists := os.LookupEnv("LOG_LEVEL")

	if exists {
		return value
	}

	return "info"
}
