package db

import (
	"fmt"
)

type PostgresSSLMode string

const (
	SslModeDisable    PostgresSSLMode = "disable"
	SslModeAllow      PostgresSSLMode = "allow"
	SslModePrefer     PostgresSSLMode = "prefer"
	SslModeRequire    PostgresSSLMode = "require"
	SslModeVerifyCa   PostgresSSLMode = "verify-ca"
	SslModeVerifyFull PostgresSSLMode = "verify-full"
)

func (m PostgresSSLMode) String() string {
	return string(m)
}

type PostgresConnectionConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
	SSLMode  PostgresSSLMode
}

func ParsePGConnString(cfg *PostgresConnectionConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s?sslmode=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
	)
}
