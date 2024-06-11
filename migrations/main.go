package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/yaroslav-koval/graphql-psychologists-courses/migrations/migrator"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/db"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/db/pgxpool"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/logging"
)

func main() {
	ctx := context.Background()

	connString := os.Getenv("GRAPHQL_PG_MIGRATIONS_CONNECTION_STRING")
	if connString == "" {
		connString = db.ParsePGConnString(
			"postgres",
			"secret",
			"localhost",
			5432,
			"graphql-psychologists-courses",
		)
	}

	pool, err := pgxpool.CreatePool(ctx, connString)
	if err != nil {
		logging.SendSimpleError(err)
	}

	m := migrator.NewMigrator(pool)

	dir := os.Getenv("GRAPHQL_PSYCHOLOGISTS_COURSES_MIGRATIONS_DIRECTORY")
	if dir == "" {
		logging.Send(logging.Error().Err(errors.New("migrations directory is not specified")))
		return
	}

	logging.Send(logging.Info().Str("message", fmt.Sprintf("Directory for migrations:: %s", dir)))

	md := migrator.Mode(os.Getenv("GRAPHQL_PSYCHOLOGISTS_COURSES_MIGRATIONS_MODE"))
	if md != migrator.UP && md != migrator.DOWN {
		panic("migrations mode is not correct")
	}

	err = m.Migrate(dir, md)
	if err != nil {
		logging.Send(logging.Error().Err(err))
	}
}
