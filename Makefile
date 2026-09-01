include .env
export

migrate-up:
	@migrate -path internal/core/migrations -database "$(DATABASE_URI)" up

migrate-down:
	@migrate -path internal/core/migrations -database "$(DATABASE_URI)" down

run:
	@go run cmd/gophermart/main.go