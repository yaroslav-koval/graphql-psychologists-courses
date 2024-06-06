package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/db"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/db/pgxpool"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/logging"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/migrator"
)

func main() {
	ctx := context.Background()

	connString := db.ParsePGConnString("postgres", "secret", "localhost", 5432, "graphql-psychologists-courses")
	pool, err := pgxpool.CreatePool(ctx, connString)
	if err != nil {
		logging.Send(logging.Error().Err(err))
	}

	m := migrator.New(pool)

	dir := os.Getenv("GRAPHQL_PSYCHOLOGISTS_COURSES_MIGRATIONS_DIRECTORY")
	if dir == "" {
		logging.Send(logging.Error().Err(errors.New("migrations directory is not specified")))
		return
	}

	logging.Send(logging.Info().Str("message", fmt.Sprintf("Directory for migrations:: %s", dir)))

	err = m.Migrate(dir, 1)
	if err != nil {
		logging.Send(logging.Error().Err(err))
	}
}
