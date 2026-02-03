package repository

import (
	"log/slog"

	"github.com/zeon-code/tiny-url/internal/db"
	"github.com/zeon-code/tiny-url/internal/pkg/config"
	"github.com/zeon-code/tiny-url/internal/pkg/metric"
)

type Repositories struct {
	Url URLRepository
}

func NewRepositories(c config.Configuration, l *slog.Logger) Repositories {
	metric := metric.NewMetricClient(c, l.With("client", "metric"))

	cache, err := db.NewCacheClient(c.Cache(), metric, l.With("client", "cache"))

	if err != nil {
		panic("error building cache client: " + err.Error())
	}

	database, err := db.NewDBClient(c.PrimaryDatabase(), metric, l.With("client", "primary-db"))

	if err != nil {
		panic("error building primary database client: " + err.Error())
	}

	replica, err := db.NewDBClient(c.ReplicaDatabase(), metric, l.With("client", "replica-db"))

	if err != nil {
		replica = database
	}

	memory := db.NewMemoryDatabase(replica, cache, metric, l.With("client", "memory"))

	return Repositories{
		Url: NewURLRepository(database, memory, l.With("repository", "url")),
	}
}
