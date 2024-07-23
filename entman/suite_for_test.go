package entman

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/yaroslav-koval/graphql-psychologists-courses/bunmodels"
)

var cb = context.Background()

type entmanSuite struct {
	suite.Suite
	em   *EntityManager
	bun  *bun.DB
	mock sqlmock.Sqlmock
}

func (s *entmanSuite) SetupSuite() {
	db, mock, err := sqlmock.New()
	if err != nil {
		s.Fail("sqlmock")
		return
	}
	s.NotEmpty(db)
	s.NotEmpty(mock)

	bunDB := bun.NewDB(db, pgdialect.New())
	em := New(bunDB)

	bunmodels.RegisterBunManyToManyModels(bunDB)

	s.em = em
	s.bun = bunDB
	s.mock = mock
}

func TestCreateSuite(t *testing.T) {
	suite.Run(t, new(entmanSuite))
}
