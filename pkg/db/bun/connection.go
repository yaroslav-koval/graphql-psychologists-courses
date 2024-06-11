package bun

import (
	"database/sql"

	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bunzerolog"
)

func CreateConnection(connectionString string, logger *zerolog.Logger) (*bun.DB, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(connectionString),
		pgdriver.WithInsecure(true),
	))

	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bunzerolog.NewQueryHook(
		bunzerolog.WithLogger(logger),
	))

	err := db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
