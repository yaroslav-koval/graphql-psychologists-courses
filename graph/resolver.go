package graph

import "github.com/yaroslav-koval/graphql-psychologists-courses/entman"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	em *entman.EntityManager
}

func NewResolver(em *entman.EntityManager) *Resolver {
	return &Resolver{
		em: em,
	}
}
