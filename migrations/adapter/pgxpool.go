package adapter

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yaroslav-koval/graphql-psychologists-courses/migrations/migrator"
)

type PGXPoolAdapter struct {
	pgxPool *pgxpool.Pool
}

func (p PGXPoolAdapter) Exec(ctx context.Context, sql string) error {
	_, err := p.pgxPool.Exec(ctx, sql)
	return err
}

func NewDBConnectionFromPGXPool(pool *pgxpool.Pool) migrator.DBConnection {
	return &PGXPoolAdapter{
		pgxPool: pool,
	}
}
