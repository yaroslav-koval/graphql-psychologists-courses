migrate:
	GRAPHQL_PSYCHOLOGISTS_COURSES_MIGRATIONS_DIRECTORY="$$(pwd)/migrations" \
	go run migrations/main.go