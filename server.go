package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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
			"postgres",
			"secret",
			"localhost",
			5432,
			"graphql-psychologists-courses",
		)
	}

	logging.Send(logging.Info().Str("connectionString", connString))

	b, err := bun.CreateConnection(connString)
	if err != nil {
		logging.Send(logging.Error().Err(err))
		return
	}

	b.RegisterModel((*bunmodels.CoursePsychologist)(nil))

	em := entman.New(b)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(em)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
