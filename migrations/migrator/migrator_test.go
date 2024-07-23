package migrator

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

var cb = context.Background()

type suiteConnection struct {
	exec func(ctx context.Context, sql string) error
}

func (s *suiteConnection) Exec(ctx context.Context, sql string) error {
	return s.exec(ctx, sql)
}

type suiteDirectoryOperator struct {
	isFileExist func(name string) bool
	readFile    func(name string) ([]byte, error)
}

func (sdo *suiteDirectoryOperator) IsFileExist(name string) bool {
	return sdo.isFileExist(name)
}

func (sdo *suiteDirectoryOperator) ReadFile(name string) ([]byte, error) {
	return sdo.readFile(name)
}

type migratorSuite struct {
	suite.Suite
	m  *Migrator
	do *suiteDirectoryOperator
	db *suiteConnection
}

func (s *migratorSuite) SetupTest() {
	s.db = &suiteConnection{}
	s.do = &suiteDirectoryOperator{}
	s.m = &Migrator{
		dirOp: s.do,
		db:    s.db,
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(migratorSuite))
}

func (s *migratorSuite) TestNewMigrator() {
	m := NewMigrator(s.do, s.db)
	s.Equal(s.db, m.db)
}

func (s *migratorSuite) TestMigrateModeUp() {
	readFileCounter := 0
	execCounter := 0
	isFileExistCounter := 0
	s.do.readFile = func(name string) ([]byte, error) {
		readFileCounter++
		return []byte{}, nil
	}
	s.db.exec = func(ctx context.Context, sql string) error {
		execCounter++
		return nil
	}
	s.do.isFileExist = func(name string) bool {
		isFileExistCounter++
		return true
	}
	err := s.m.Migrate(cb, "dir", UP)
	s.NoError(err)
	s.Equal(2, readFileCounter)
	s.Equal(2, execCounter)
	s.Equal(1, isFileExistCounter)
}

func (s *migratorSuite) TestMigrateModeDown() {
	readFileCounter := 0
	execCounter := 0
	isFileExistCounter := 0
	s.do.readFile = func(name string) ([]byte, error) {
		readFileCounter++
		return []byte{}, nil
	}
	s.db.exec = func(ctx context.Context, sql string) error {
		execCounter++
		return nil
	}
	s.do.isFileExist = func(name string) bool {
		isFileExistCounter++
		return true
	}
	err := s.m.Migrate(cb, "dir", DOWN)
	s.NoError(err)
	s.Equal(1, readFileCounter)
	s.Equal(1, execCounter)
	s.Equal(1, isFileExistCounter)
}

func (s *migratorSuite) TestMigrateUp() {
	isReadFileCalled := false
	isExecCalled := false
	expDir := "dir"
	expErr := "error text"
	s.do.readFile = func(name string) ([]byte, error) {
		s.Equal(fmt.Sprintf("%v/sql/up.sql", expDir), name)
		isReadFileCalled = true
		return []byte{}, nil
	}
	s.db.exec = func(ctx context.Context, sql string) error {
		isExecCalled = true
		// so migrateData isn't called
		return fmt.Errorf(expErr)
	}
	err := s.m.Migrate(cb, expDir, UP)
	s.ErrorContains(err, expErr)
	s.True(isReadFileCalled)
	s.True(isExecCalled)
}

func (s *migratorSuite) TestMigrateData() {
	isReadFileCalled := false
	isExecCalled := false
	isFileExistsCalled := false
	expDir := "dir"
	s.do.readFile = func(name string) ([]byte, error) {
		if name != "dir/sql/up.sql" {
			s.Equal(fmt.Sprintf("%v/sql/data.sql", expDir), name)
		}

		isReadFileCalled = true
		return []byte{}, nil
	}
	s.db.exec = func(ctx context.Context, sql string) error {
		isExecCalled = true
		return nil
	}
	s.do.isFileExist = func(name string) bool {
		isFileExistsCalled = true
		s.Equal(fmt.Sprintf("%v/sql/data.sql", expDir), name)
		return true
	}
	err := s.m.Migrate(cb, expDir, UP)
	s.NoError(err)
	s.True(isReadFileCalled)
	s.True(isExecCalled)
	s.True(isFileExistsCalled)
}

func (s *migratorSuite) TestMigrateDown() {

	isReadFileCalled := false
	isExecCalled := false
	isFileExistsCalled := false
	expDir := "dir"
	s.do.readFile = func(name string) ([]byte, error) {
		s.Equal(fmt.Sprintf("%v/sql/down.sql", expDir), name)
		isReadFileCalled = true
		return []byte{}, nil
	}
	s.db.exec = func(ctx context.Context, sql string) error {
		isExecCalled = true
		return nil
	}
	s.do.isFileExist = func(name string) bool {
		isFileExistsCalled = true
		s.Equal(fmt.Sprintf("%v/sql/down.sql", expDir), name)
		return true
	}
	err := s.m.Migrate(cb, expDir, DOWN)
	s.NoError(err)
	s.True(isReadFileCalled)
	s.True(isExecCalled)
	s.True(isFileExistsCalled)
}

func (s *migratorSuite) TestMigrateFile() {
	expFileName := "file/name"
	isReadFileCalled := false
	isExecCalled := false
	query := []byte("query text")
	s.do.readFile = func(name string) ([]byte, error) {
		isReadFileCalled = true
		s.Equal(expFileName, name)
		return query, nil
	}
	s.db.exec = func(ctx context.Context, sql string) error {
		isExecCalled = true
		s.EqualValues(query, sql)
		return nil
	}

	err := s.m.migrateFile(cb, expFileName)
	s.NoError(err)
	s.True(isReadFileCalled)
	s.True(isExecCalled)
}
