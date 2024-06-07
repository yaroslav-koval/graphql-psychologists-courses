package db

import (
	"fmt"
)

func ParsePGConnString(username, password, host string, port int, dbName string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s?sslmode=disable",
		username,
		password,
		host,
		port,
		dbName,
	)
}
