CONTAINER_NAME=split-the-bill-postgres-db
BACKUP_DIR=./storage/database/backups

# Get the current date and time in a format suitable for a filename
DATE := $(shell date +%Y%m%d%H%M%S)
# Set the backup file name
BACKUP_NAME := db-backup-$(DATE)

# Load environment variables from .env file
ifneq (,$(wildcard .env))
	include .env
	export $(shell sed 's/=.*//' .env)
endif

# Export the database data to a file inside the container
.PHONY: export-db
export-db:
	docker exec -t $(CONTAINER_NAME) pg_dump -U $(DB_USER) -d $(DB_NAME) --data-only -f /tmp/$(BACKUP_NAME).sql

# Copy the backup file from the container to the host machine
.PHONY: copy-backup
copy-backup:
	docker cp $(CONTAINER_NAME):/tmp/$(BACKUP_NAME).sql $(BACKUP_DIR)/$(BACKUP_NAME).sql

# Full backup process
.PHONY: backup-db
backup-db:
	@make export-db
	@make copy-backup
	@echo "Backup completed and saved to $(BACKUP_DIR)/$(BACKUP_NAME).sql"

# Copy the backup file from the host to the container
.PHONY: copy-backup-to-container
copy-backup-to-container:
	docker cp $(BACKUP_FILE) $(CONTAINER_NAME):/tmp/$(BACKUP_FILE_NAME)

# Import the backup file into the PostgreSQL database
.PHONY: import-db
import-db:
	docker exec -t $(CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME) -f /tmp/$(BACKUP_FILE_NAME)

# Full import process with a specific backup file
.PHONY: load-backup
load-backup:
	@if [ -z "$(file)" ]; then \
		echo "Please provide a backup file using 'make import-backup file=<backup_file>'"; \
		exit 1; \
	fi
	@make copy-backup-to-container BACKUP_FILE=$(file) BACKUP_FILE_NAME=$(notdir $(file))
	@make import-db BACKUP_FILE_NAME=$(notdir $(file))
	@echo "Backup $(BACKUP_DIR)/$(BACKUP_FILE_NAME) imported into the database"

# Open Postgres Docker container shell
docker-db-shell:
	@docker exec -it $(CONTAINER_NAME) bash

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
	@docker exec split-the-bill-server sh -c "go run script/db_seed.go"

stop-db:
	@docker compose down

reset-db:
	@docker compose down -v