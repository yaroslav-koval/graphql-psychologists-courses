package pgxpool

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/logging"
)

func CreatePool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	dbConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	dbConfig.MinConns = 5
	dbConfig.MaxConns = 20
	dbConfig.MaxConnLifetime = 30 * time.Minute
	dbConfig.MaxConnIdleTime = 5 * time.Minute
	dbConfig.HealthCheckPeriod = time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		return nil, err
	}

	timeout := 5
	attempts := 10
	for i := 1; i <= attempts; i++ {
		logging.Send(
			logging.Info().Caller().Str("message", fmt.Sprintf(`Connection to postgres, attempt #%v`, i)),
		)

		err = pool.Ping(ctx)
		if err == nil {
			break
		}

		logging.Send(
			logging.Info().Caller().Str("message",
				fmt.Sprintf(
					"Connection unsuccessful, err::%s\nNext connection in %v seconds",
					err.Error(),
					timeout)),
		)

		if i == attempts {
			return nil, err
		}

		time.Sleep(time.Duration(timeout) * time.Second)
	}

	logging.Send(
		logging.Info().Caller().Str("message", fmt.Sprintf(`Successfuly connected to postgres`)),
	)

	return pool, nil
}
