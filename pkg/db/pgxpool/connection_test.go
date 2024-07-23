package pgxpool

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/suite"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/db"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/logging"
)

type connectionSuite struct {
	suite.Suite
	dockerPool       *dockertest.Pool
	container        *dockertest.Resource
	connectionConfig *db.PostgresConnectionConfig
}

func (s *connectionSuite) SetupSuite() {
	pool, err := dockertest.NewPool("")
	if err != nil {
		logging.SendSimpleError(err)
	}
	s.dockerPool = pool

	err = pool.Client.Ping()
	if err != nil {
		logging.SendSimpleError(err)
	}

	connCfg := &db.PostgresConnectionConfig{
		Username: "user",
		Password: "password",
		Host:     "localhost",
		Port:     5431,
		DBName:   "dbname",
		SSLMode:  db.SslModeDisable,
	}
	s.connectionConfig = connCfg

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Name:       "connection-test",
		Repository: "postgres",
		Tag:        "16.3",
		Env: []string{
			fmt.Sprintf("POSTGRES_USER=%s", connCfg.Username),
			fmt.Sprintf("POSTGRES_PASSWORD=%s", connCfg.Password),
			fmt.Sprintf("POSTGRES_DB=%s", connCfg.DBName),
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		logging.SendSimpleError(err)
	}

	err = resource.Expire(180)
	if err != nil {
		logging.SendSimpleError(err)
	}

	connCfg.Port, err = strconv.Atoi(resource.GetPort("5432/tcp"))
	if err != nil {
		logging.SendSimpleError(err)
	}

	s.container = resource
}

func (s *connectionSuite) TearDownSuite() {
	if s.dockerPool == nil || s.container == nil {
		return
	}

	err := s.dockerPool.Purge(s.container)
	if err != nil {
		logging.SendSimpleError(err)
	}
}

func TestConnectionSuite(t *testing.T) {
	suite.Run(t, new(connectionSuite))
}

func (s *connectionSuite) TestCreatePool() {
	expCfg := &PoolConfig{
		PingTimeout:           5 * time.Second,
		PingAttempts:          3,
		MinConnections:        1,
		MaxConnections:        1,
		MaxConnectionLifetime: 30 * time.Minute,
		MaxConnectionIdleTime: 5 * time.Minute,
		HealthCheckPeriod:     1 * time.Minute,
	}

	cs := db.ParsePGConnString(s.connectionConfig)

	pool, err := CreatePool(
		context.Background(),
		cs,
		expCfg,
	)
	s.NoError(err)

	cfg := pool.Config()
	s.Equal(expCfg.MinConnections, cfg.MinConns)
	s.Equal(expCfg.MaxConnections, cfg.MaxConns)
	s.Equal(expCfg.MaxConnectionLifetime, cfg.MaxConnLifetime)
	s.Equal(expCfg.MaxConnectionIdleTime, cfg.MaxConnIdleTime)
	s.Equal(expCfg.HealthCheckPeriod, cfg.HealthCheckPeriod)
}
