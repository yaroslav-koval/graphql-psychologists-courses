package entman

import "github.com/jackc/pgx/v5/pgxpool"

type EntityManager struct {
	pg *pgxpool.Pool
}

func New(pg *pgxpool.Pool) *EntityManager {
	return &EntityManager{
		pg: pg,
	}
}
