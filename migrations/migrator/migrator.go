package migrator

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/logging"
)

type Migrator struct {
	pg *pgxpool.Pool
}

func NewMigrator(pg *pgxpool.Pool) *Migrator {
	return &Migrator{
		pg: pg,
	}
}

func (m *Migrator) Migrate(workingDir string, md Mode) error {
	workingDir = fmt.Sprintf("%s/sql", workingDir)

	if md == UP {
		err := m.migrateTables(workingDir)
		if err != nil {
			return err
		}

		err = m.migrateData(workingDir)
		if err != nil {
			return err
		}
	}

	if md == DOWN {
		err := m.migrateDown(workingDir)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Migrator) migrateTables(workingDir string) error {
	fn := fmt.Sprintf("%s/up.sql", workingDir)

	return m.migrateFile(fn)
}

func (m *Migrator) migrateData(workingDir string) error {
	fn := fmt.Sprintf("%s/data.sql", workingDir)

	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return nil
	}

	return m.migrateFile(fn)
}

func (m *Migrator) migrateDown(workingDir string) error {
	fn := fmt.Sprintf("%s/down.sql", workingDir)

	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return nil
	}

	return m.migrateFile(fn)
}

func (m *Migrator) migrateFile(file string) error {
	sqlBytes, err := os.ReadFile(file)
	if err != nil {
		logInfoF("Error reading file:: %v", err)
		return err
	}

	_, err = m.pg.Exec(context.Background(), string(sqlBytes))
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
