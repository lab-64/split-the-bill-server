run:
	@make swag
	@make build
	@./split-the-bill-server

build:
	@go build

watch:
	@reflex -s -r '\.go$$' -R 'docs.go' make run

swag:
	@swag init && swag fmt

clean:
	@rm split-the-bill-server

test-all:
	@go test ./...

seed:
	@docker exec split-the-bill-server sh -c "go run scripts/seed.go"

start-postgres:
	@docker compose up --no-attach pgadmin

stop-postgres:
	@docker compose down

reset-db:
	@docker compose down -v