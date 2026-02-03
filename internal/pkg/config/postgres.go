package config

import (
	"fmt"
	"os"
	"strconv"
)

type DatabaseConfiguration interface {
	GetDNS() string
	GetDriver() string
}

type PostgresConfig struct {
	Prefix string
}

func newPostgresConfig(prefix string) PostgresConfig {
	return PostgresConfig{
		Prefix: prefix,
	}
}

func (c PostgresConfig) GetDriver() string {
	return "postgres"
}

func (c PostgresConfig) GetDNS() string {
	sslMode := ""

	if !c.DBTLSMode() {
		sslMode = "sslmode=disable"
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s", c.DBUser(), c.DBPassword(), c.DBHost(), c.DBPort(), c.DBName(), sslMode)
}

func (c PostgresConfig) DBHost() string {
	return c.get("HOST")
}

func (c PostgresConfig) DBPort() int {
	port, err := strconv.Atoi(c.get("PORT"))

	if err != nil {
		panic("DB_PORT must be an interger value")
	}

	return port
}

func (c PostgresConfig) DBUser() string {
	return c.get("USER")
}

func (c PostgresConfig) DBPassword() string {
	return c.get("PASSWORD")
}

func (c PostgresConfig) DBName() string {
	return c.get("NAME")
}

func (c PostgresConfig) DBTLSMode() bool {
	tlsMode, err := strconv.ParseBool(c.get("TLS_MODE"))

	if err != nil {
		panic("DB_TLS_MODE must be a boolean value")
	}

	return tlsMode
}

func (c PostgresConfig) get(suffix string) string {
	env := fmt.Sprintf("%s_%s", c.Prefix, suffix)

	if value, exists := os.LookupEnv(env); exists {
		return value
	}

	panic(fmt.Sprintf("Missing required environment variable %s", env))
}
