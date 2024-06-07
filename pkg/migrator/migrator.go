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

func New(pg *pgxpool.Pool) *Migrator {
	return &Migrator{
		pg: pg,
	}
}

func (m *Migrator) Migrate(workingDir string, migrationLevel int) error {
	for i := 1; i <= migrationLevel; i++ {
		err := m.migrateQueries(workingDir, i)
		if err != nil {
			return err
		}

		err = m.migrateData(workingDir, i)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Migrator) migrateQueries(workingDir string, migrationLevel int) error {
	fn := fmt.Sprintf("%s/%v.up.sql", workingDir, migrationLevel)

	return m.migrateFile(fn)
}

func (m *Migrator) migrateData(workingDir string, migrationLevel int) error {
	fn := fmt.Sprintf("%s/%v.data.sql", workingDir, migrationLevel)

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
