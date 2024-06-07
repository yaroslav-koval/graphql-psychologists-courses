package main

import (
	"context"
	"github.com/yaroslav-koval/graphql-psychologists-courses/entman"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/db"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/db/pgxpool"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/logging"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/yaroslav-koval/graphql-psychologists-courses/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ctx := context.Background()

	connString := os.Getenv("GRAPHQL_SERVER_CONNECTION_STRING")
	if connString == "" {
		connString = db.ParsePGConnString("postgres", "secret", "localhost", 5432, "graphql-psychologists-courses")
	}

	pool, err := pgxpool.CreatePool(ctx, connString)
	if err != nil {
		logging.Send(logging.Error().Err(err))
	}

	em := entman.New(pool)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(em)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
