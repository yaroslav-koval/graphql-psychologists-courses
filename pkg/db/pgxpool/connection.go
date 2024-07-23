package pgxpool

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/logging"
)

type PoolConfig struct {
	PingTimeout           time.Duration
	PingAttempts          int
	MinConnections        int32
	MaxConnections        int32
	MaxConnectionLifetime time.Duration
	MaxConnectionIdleTime time.Duration
	HealthCheckPeriod     time.Duration
}

func CreatePool(ctx context.Context, connString string, cfg *PoolConfig) (*pgxpool.Pool, error) {
	dbConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	dbConfig.MinConns = cfg.MinConnections
	dbConfig.MaxConns = cfg.MaxConnections
	dbConfig.MaxConnLifetime = cfg.MaxConnectionLifetime
	dbConfig.MaxConnIdleTime = cfg.MaxConnectionIdleTime
	dbConfig.HealthCheckPeriod = cfg.HealthCheckPeriod

	pool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		return nil, err
	}

	timeout := cfg.PingTimeout
	attempts := cfg.PingAttempts
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

		time.Sleep(timeout)
	}

	logging.Send(
		logging.Info().Caller().Str("message", fmt.Sprintf(`Successfuly connected to postgres`)),
	)

	return pool, nil
}
