package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type DatabaseConfiguration interface {
	GetDNS() (string, error)
	GetDriver() string
}

type PostgresConfig struct {
	Prefix string
}

func NewPostgresConfig(prefix string) PostgresConfig {
	return PostgresConfig{
		Prefix: prefix,
	}
}

func (c PostgresConfig) GetDriver() string {
	return "postgres"
}

func (c PostgresConfig) GetDNS() (string, error) {
	isTLSMode, err := c.DBTLSMode()

	if err != nil {
		return "", err
	}

	user, err := c.DBUser()

	if err != nil {
		return "", err
	}

	password, err := c.DBPassword()

	if err != nil {
		return "", err
	}

	host, err := c.DBHost()

	if err != nil {
		return "", err
	}

	port, err := c.DBPort()

	if err != nil {
		return "", err
	}

	name, err := c.DBName()

	if err != nil {
		return "", err
	}

	sslMode := ""

	if !isTLSMode {
		sslMode = "sslmode=disable"
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s", user, password, host, port, name, sslMode), nil
}

func (c PostgresConfig) DBHost() (string, error) {
	return c.get("HOST")
}

func (c PostgresConfig) DBPort() (int, error) {
	env, err := c.get("PORT")

	if err != nil {
		return 0, err
	}

	port, err := strconv.Atoi(env)

	if err != nil {
		return 0, errors.New("DB_PORT must be an interger value")
	}

	return port, nil
}

func (c PostgresConfig) DBUser() (string, error) {
	return c.get("USER")
}

func (c PostgresConfig) DBPassword() (string, error) {
	return c.get("PASSWORD")
}

func (c PostgresConfig) DBName() (string, error) {
	return c.get("NAME")
}

func (c PostgresConfig) DBTLSMode() (bool, error) {
	env, err := c.get("TLS_MODE")

	if err != nil {
		return false, err
	}

	tlsMode, err := strconv.ParseBool(env)

	if err != nil {
		return false, errors.New("DB_TLS_MODE must be a boolean value")
	}

	return tlsMode, nil
}

func (c PostgresConfig) get(suffix string) (string, error) {
	env := fmt.Sprintf("%s_%s", c.Prefix, suffix)

	if value, exists := os.LookupEnv(env); exists {
		return value, nil
	}

	return "", fmt.Errorf("Missing required environment variable %s", env)
}
