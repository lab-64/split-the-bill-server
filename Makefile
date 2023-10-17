run:
	@make build
	@./split-the-bill-server

build:
	@go build

watch:
	@reflex -s -r '\.go$$' make run

clean:
	@rm split-the-bill-server

test:
	@go test ./...

start-postgres:
	@docker-compose up --build

stop-postgres:
	@docker-compose down

reset-db:
	@docker-compose down -v