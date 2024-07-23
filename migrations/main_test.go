package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaroslav-koval/graphql-psychologists-courses/migrations/migrator"
)

func TestGetPGConnectionFromEnv(t *testing.T) {
	cs := getPGConnectionFromEnv()
	assert.Equal(t, "postgres://postgres:secret@localhost:5432/graphql-psychologists-courses?sslmode=disable", cs)

	expConnString := "connectionString"
	err := os.Setenv("GRAPHQL_PG_MIGRATIONS_CONNECTION_STRING", expConnString)
	assert.NoError(t, err)
	cs = getPGConnectionFromEnv()
	assert.Equal(t, expConnString, cs)
}

func TestGetMigrationsDirectoryFromEnv(t *testing.T) {
	dir, err := getMigrationsDirectoryFromEnv()
	assert.Error(t, err)
	assert.Empty(t, dir)

	expDir := "directory"
	err = os.Setenv("GRAPHQL_PSYCHOLOGISTS_COURSES_MIGRATIONS_DIRECTORY", expDir)
	assert.NoError(t, err)
	dir, err = getMigrationsDirectoryFromEnv()
	assert.NoError(t, err)
	assert.Equal(t, expDir, dir)
}

func TestGetMigrationModeFromEnv(t *testing.T) {
	m, err := getMigrationModeFromEnv()
	assert.Error(t, err)
	assert.Empty(t, m)

	err = os.Setenv("GRAPHQL_PSYCHOLOGISTS_COURSES_MIGRATIONS_MODE", "randomValue")
	assert.NoError(t, err)
	m, err = getMigrationModeFromEnv()
	assert.Error(t, err)
	assert.Empty(t, m)

	err = os.Setenv("GRAPHQL_PSYCHOLOGISTS_COURSES_MIGRATIONS_MODE", "up")
	assert.NoError(t, err)
	m, err = getMigrationModeFromEnv()
	assert.NoError(t, err)
	assert.Equal(t, migrator.UP, m)

	err = os.Setenv("GRAPHQL_PSYCHOLOGISTS_COURSES_MIGRATIONS_MODE", "down")
	assert.NoError(t, err)
	m, err = getMigrationModeFromEnv()
	assert.NoError(t, err)
	assert.Equal(t, migrator.DOWN, m)
}
