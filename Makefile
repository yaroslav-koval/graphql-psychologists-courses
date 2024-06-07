deploy-pg:
	docker compose -f docker-deployments/docker-compose.pg.yml -p graphql-psychologists-courses up -d

migrate:
	GRAPHQL_PSYCHOLOGISTS_COURSES_MIGRATIONS_DIRECTORY="$$(pwd)/migrations" \
	go run migrations/main.go

gql-init:
	go run github.com/99designs/gqlgen init

gql-gen:
	go run github.com/99designs/gqlgen generate

start:
	go run server.go
