package repository

import (
	"errors"

	"github.com/zeon-code/tiny-url/internal/db"
	"github.com/zeon-code/tiny-url/internal/pkg/config"
	"github.com/zeon-code/tiny-url/internal/pkg/observability"
)

type Repositories struct {
	Url    URLRepository
	Health HealthRepository

	database db.SQLClient
	memory   db.SQLReader
}

// Shutdown gracefully closes all resources associated with the Repositories.
// It attempts to close both the database and in-memory storage if they are initialized.
// Any errors encountered during closure are aggregated and returned.
// If no errors occur, Shutdown returns nil.
func (r Repositories) Shutdown() error {
	var err error

	if r.database != nil {
		err = errors.Join(err, r.database.Close())
	}

	if r.memory != nil {
		err = errors.Join(err, r.memory.Close())
	}

	return err
}

// NewRepositoriesFromConfig builds and wires all repository dependencies using the
// provided application configuration and logger.
//
// It initializes metric, cache, primary database, and replica database clients.
// If the replica database configuration is missing or fails to initialize, the
// primary database client is used as a fallback to ensure read availability.
//
// The function panics if critical dependencies (cache or primary database)
// cannot be created, as the application cannot operate without them.
//
// Returns a fully initialized Repositories instance
func NewRepositoriesFromConfig(conf config.Configuration, observer observability.Observer) Repositories {
	cache, err := db.NewCacheClient(conf.Cache(), observer)

	if err != nil {
		panic("error building cache client: " + err.Error())
	}

	primary, err := db.NewDBClient(conf.PrimaryDatabase(), observer)

	if err != nil {
		cache.Close()
		panic("error building primary database client: " + err.Error())
	}

	replica, err := db.NewDBClient(conf.ReplicaDatabase(), observer)

	if err != nil {
		replica = primary
	}

	memory, err := db.NewMemoryDatabase(replica, cache, observer)

	if err != nil {
		panic("error building memory client" + err.Error())
	}

	return NewRepositories(primary, memory, observer)
}

// NewRepositories constructs a Repositories container using the provided
// database and memory-backed readers.
//
// The database client is used for write operations, while the memory client
// (typically backed by cache and/or replicas) is used for read operations.
func NewRepositories(primary db.SQLClient, memory db.SQLReader, observer observability.Observer) Repositories {
	return Repositories{
		Url:    NewURLRepository(primary, memory, observer),
		Health: NewHealthRepository(primary, memory, observer),

		database: primary,
		memory:   memory,
	}
}
