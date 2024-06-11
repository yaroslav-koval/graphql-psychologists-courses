deploy-pg:
	docker compose -f docker-deployments/docker-compose.pg.yml -p graphql-psychologists-courses up -d

migrate:
	GRAPHQL_PSYCHOLOGISTS_COURSES_MIGRATIONS_DIRECTORY="$$(pwd)/migrations" \
	GRAPHQL_PSYCHOLOGISTS_COURSES_MIGRATIONS_MODE="up" \
	go run migrations/main.go

migrate-down:
	GRAPHQL_PSYCHOLOGISTS_COURSES_MIGRATIONS_DIRECTORY="$$(pwd)/migrations" \
	GRAPHQL_PSYCHOLOGISTS_COURSES_MIGRATIONS_MODE="down" \
	go run migrations/main.go

reset-database:
	make migrate-down && make migrate

gql-init:
	go run github.com/99designs/gqlgen init

gql-gen:
	go run github.com/99designs/gqlgen generate

start:
	go run server.go
