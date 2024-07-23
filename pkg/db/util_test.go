package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgresSSLModeStringer(t *testing.T) {
	var s fmt.Stringer

	s = SslModeDisable
	assert.Equal(t, string(SslModeDisable), s.String())
}

func TestParsePGConnString(t *testing.T) {
	cfg := &PostgresConnectionConfig{
		Username: "login",
		Password: "password",
		Host:     "localhost",
		Port:     5432,
		DBName:   "dbName",
		SSLMode:  SslModeDisable,
	}

	connStr := ParsePGConnString(cfg)

	assert.Equal(t, "postgres://login:password@localhost:5432/dbName?sslmode=disable", connStr)
}
