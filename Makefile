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

start-postgres:
	@docker-compose up --build

stop-postgres:
	@docker-compose down

reset-db:
	@docker-compose down -v