package config

import (
	"fmt"
	"os"
	"strconv"
)

type RedisConfig struct{}

func (c RedisConfig) GetDriver() string {
	return "redis"
}

func (c RedisConfig) GetDNS() string {

	if c.DBPassword() != "" {
		return fmt.Sprintf("redis://:%s@%s:%d/%d", c.DBPassword(), c.DBHost(), c.DBPort(), c.DBName())
	}

	return fmt.Sprintf("redis://%s:%d/%d", c.DBHost(), c.DBPort(), c.DBName())
}

func (c RedisConfig) DBHost() string {
	return c.get("CACHE_HOST")
}

func (c RedisConfig) DBPort() int {
	port, err := strconv.Atoi(c.get("CACHE_PORT"))

	if err != nil {
		panic("CACHE_PORT must be an interger value")
	}

	return port
}

func (c RedisConfig) DBPassword() string {
	return c.get("CACHE_PASSWORD")
}

func (c RedisConfig) DBName() int {
	name, err := strconv.Atoi(c.get("CACHE_NAME"))

	if err != nil {
		panic("CACHE_NAME must be an interger value")
	}

	return name
}

func (c RedisConfig) get(env string) string {
	if value, exists := os.LookupEnv(env); exists {
		return value
	}

	panic(fmt.Sprintf("Missing required environment variable %s", env))
}
