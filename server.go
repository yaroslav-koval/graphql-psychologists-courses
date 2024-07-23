package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/zerolog"
	"github.com/yaroslav-koval/graphql-psychologists-courses/bunmodels"
	"github.com/yaroslav-koval/graphql-psychologists-courses/entman"
	"github.com/yaroslav-koval/graphql-psychologists-courses/graph"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/db"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/db/bun"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/logging"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	connString := os.Getenv("GRAPHQL_SERVER_CONNECTION_STRING")
	if connString == "" {
		connString = db.ParsePGConnString(
			&db.PostgresConnectionConfig{
				Username: "postgres",
				Password: "secret",
				Host:     "localhost",
				Port:     5432,
				DBName:   "graphql-psychologists-courses",
				SSLMode:  db.SslModeDisable,
			},
		)
	}

	logging.Send(logging.Info().Str("connectionString", connString))

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime})
	logging.SetLogger(logger)

	b, err := bun.CreateConnection(connString, &logger)
	if err != nil {
		logging.Send(logging.Error().Err(err))
		return
	}

	bunmodels.RegisterBunManyToManyModels(b)

	em := entman.New(b)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(em)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	serverErrorCh := make(chan error)
	go func() {
		serverErrorCh <- http.ListenAndServe(":"+port, nil)
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGQUIT)

	select {
	case err = <-serverErrorCh:
		logging.SendSimpleErrorAsync(err)
	case s := <-sigCh:
		logging.Send(
			logging.Info().Str("message", fmt.Sprintf("Got signar from os: %s", s)),
		)
	}

	logging.WaitAsyncLogs(context.Background(), time.Second*30)
}
