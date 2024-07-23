package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/yaroslav-koval/graphql-psychologists-courses/migrations/adapter"
	"github.com/yaroslav-koval/graphql-psychologists-courses/migrations/migrator"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/db"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/db/pgxpool"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/logging"
)

func main() {
	ctx := context.Background()

	logging.SetDefaultLogger()

	connString := getPGConnectionFromEnv()

	poolCfg := &pgxpool.PoolConfig{
		PingTimeout:           5 * time.Second,
		PingAttempts:          5,
		MinConnections:        1,
		MaxConnections:        1,
		MaxConnectionLifetime: 30 * time.Minute,
		MaxConnectionIdleTime: 5 * time.Minute,
		HealthCheckPeriod:     time.Minute,
	}

	pool, err := pgxpool.CreatePool(ctx, connString, poolCfg)
	if err != nil {
		logging.SendSimpleError(err)
		return
	}

	dbConnection := adapter.NewDBConnectionFromPGXPool(pool)
	do := adapter.NewDirectoryOperator()

	m := migrator.NewMigrator(do, dbConnection)

	dir, err := getMigrationsDirectoryFromEnv()
	if err != nil {
		return
	}

	logging.Send(logging.Info().Str("message", fmt.Sprintf("Directory for migrations:: %s", dir)))

	md, err := getMigrationModeFromEnv()
	if err != nil {
		return
	}

	err = m.Migrate(ctx, dir, md)
	if err != nil {
		logging.Send(logging.Error().Err(err))
	}
}

func getPGConnectionFromEnv() string {
	connString := os.Getenv("GRAPHQL_PG_MIGRATIONS_CONNECTION_STRING")
	if connString == "" {
		connString = db.ParsePGConnString(
			&db.PostgresConnectionConfig{
				Username: "postgres",
				Password: "secret",
				Host:     "localhost",
				Port:     5432,
				DBName:   "graphql-psychologists-courses",
				SSLMode:  db.SslModeDisable,
			},
		)
	}

	return connString
}

func getMigrationsDirectoryFromEnv() (string, error) {
	dir := os.Getenv("GRAPHQL_PSYCHOLOGISTS_COURSES_MIGRATIONS_DIRECTORY")
	if dir == "" {
		err := errors.New("migrations directory is not specified")
		logging.Send(logging.Error().Err(err))
		return "", err
	}

	return dir, nil
}

func getMigrationModeFromEnv() (migrator.Mode, error) {
	md := migrator.Mode(os.Getenv("GRAPHQL_PSYCHOLOGISTS_COURSES_MIGRATIONS_MODE"))
	if md != migrator.UP && md != migrator.DOWN {
		err := errors.New("migrations mode is not correct")
		logging.Send(logging.Error().Err(err))
		return "", err
	}

	return md, nil
}
