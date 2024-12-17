export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgres://postgres:postgres@localhost:5432/shelter
export GOOSE_MIGRATION_DIR=cmd/migrate/migrations

migrate:
	goose up

reset:
	goose reset

generate:
	sqlc generate