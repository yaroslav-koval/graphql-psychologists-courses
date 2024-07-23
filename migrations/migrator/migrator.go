package migrator

import (
	"context"
	"fmt"

	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/logging"
)

type DBConnection interface {
	Exec(ctx context.Context, sql string) error
}

type DirectoryOperator interface {
	IsFileExist(name string) bool
	ReadFile(name string) ([]byte, error)
}

type Migrator struct {
	dirOp DirectoryOperator
	db    DBConnection
}

func NewMigrator(do DirectoryOperator, db DBConnection) *Migrator {
	return &Migrator{
		dirOp: do,
		db:    db,
	}
}

func (m *Migrator) Migrate(ctx context.Context, workingDir string, md Mode) error {
	workingDir = fmt.Sprintf("%s/sql", workingDir)

	if md == UP {
		err := m.migrateUp(ctx, workingDir)
		if err != nil {
			return err
		}

		err = m.migrateData(ctx, workingDir)
		if err != nil {
			return err
		}
	}

	if md == DOWN {
		err := m.migrateDown(ctx, workingDir)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Migrator) migrateUp(ctx context.Context, workingDir string) error {
	fn := fmt.Sprintf("%s/up.sql", workingDir)

	return m.migrateFile(ctx, fn)
}

func (m *Migrator) migrateData(ctx context.Context, workingDir string) error {
	fn := fmt.Sprintf("%s/data.sql", workingDir)

	if !m.dirOp.IsFileExist(fn) {
		return nil
	}

	return m.migrateFile(ctx, fn)
}

func (m *Migrator) migrateDown(ctx context.Context, workingDir string) error {
	fn := fmt.Sprintf("%s/down.sql", workingDir)

	if !m.dirOp.IsFileExist(fn) {
		return nil
	}

	return m.migrateFile(ctx, fn)
}

func (m *Migrator) migrateFile(ctx context.Context, file string) error {
	sqlBytes, err := m.dirOp.ReadFile(file)
	if err != nil {
		logInfoF("Error reading file:: %v", err)
		return err
	}

	err = m.db.Exec(ctx, string(sqlBytes))
	if err != nil {
		logInfoF("Error executing SQL:: %v", err)
		return err
	}

	logInfoF("Executed query migration:: %s", file)

	return nil
}

func logInfoF(format string, a ...any) {
	logging.Send(
		logging.Info().Str("message", fmt.Sprintf(format, a...)),
		1,
	)
}
