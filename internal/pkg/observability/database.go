package observability

import (
	"database/sql"

	"github.com/XSAM/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func NewInstrumentedDB(observer Observer, driver string, dsn string) (*sql.DB, error) {
	db, err := otelsql.Open(
		driver,
		dsn,
		otelsql.WithAttributes(
			semconv.DBSystemPostgreSQL,
			semconv.DBName("tiny_url"),
		),
	)

	if err != nil {
		return nil, err
	}

	reg, err := otelsql.RegisterDBStatsMetrics(
		db,
		otelsql.WithAttributes(
			semconv.DBSystemPostgreSQL,
		),
	)

	if err != nil {
		return nil, err
	}

	observer.RegisterDB(reg)
	return db, nil
}
