package entman

import (
	"github.com/uptrace/bun"
)

type EntityManager struct {
	db *bun.DB
}

func New(pg *bun.DB) *EntityManager {
	return &EntityManager{
		db: pg,
	}
}
