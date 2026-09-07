-include .env
export

migrate-up:
	@migrate -path internal/core/migrations -database "$(DATABASE_URI)" up

migrate-down:
	@migrate -path internal/core/migrations -database "$(DATABASE_URI)" down

run-server:
	@go run cmd/gophermart/main.go

run-accrual:
	@cmd/accrual/accrual_darwin_arm64 -d="$(ACCRUAL_DATABASE_URI)" -a=localhost:8081
