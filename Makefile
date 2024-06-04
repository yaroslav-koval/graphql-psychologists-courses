migrate:
	export GRAPHQL_PSYCHOLOGISTS_COURSES_MIGRATIONS_DIRECTORY="$$(pwd)/migrations" && \
	env | grep GRAPHQL_PSYCHOLOGISTS_COURSES_MIGRATIONS_DIRECTORY && \
	go run migrations/main.go

