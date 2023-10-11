run:
	@make build
	@./split-the-bill-server

build:
	@go build

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