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
	@go test --cover ./...

start-db:
	@docker compose up --no-attach pgadmin

seed-db:
	@docker exec split-the-bill-server sh -c "go run data/db_seed.go"

stop-db:
	@docker compose down

reset-db:
	@docker compose down -v