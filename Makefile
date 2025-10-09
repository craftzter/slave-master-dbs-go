# Default connection string (bisa di override waktu make jalan)
DB_PRIMARY_URL ?= "postgres://postgres:secret@localhost:5432/mydb?sslmode=disable"
DB_REPLICA_URL ?= "postgres://postgres:secret@localhost:5433/mydb?sslmode=disable"

MIGRATIONS_DIR = migrations

.PHONY: migrate-up migrate-down migrate-new migrate-reset migrate-force build run test clean docker-up docker-down docker-build docker-logs docker-exec docker-clean sqlc-gen setup restart-app

## Apply all migrations ke PRIMARY DB
migrate-up-primary:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_PRIMARY_URL) up

## Apply all migrations ke REPLICA DB
migrate-up-replica:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_REPLICA_URL) up

## Rollback last migration di PRIMARY
migrate-down-primary:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_PRIMARY_URL) down

## Rollback last migration di REPLICA
migrate-down-replica:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_REPLICA_URL) down

## Bikin migration baru
migrate-new:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make migrate-new name=create_users_table"; \
		exit 1; \
	fi
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

## Reset database (rollback semua lalu apply semua) di PRIMARY
migrate-reset:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_PRIMARY_URL) down -all
	migrate -path $(MIGRATIONS_DIR) -database $(DB_PRIMARY_URL) up

## Force version tertentu (contoh: make migrate-force version=1)
migrate-force:
	@if [ -z "$(version)" ]; then \
		echo "Usage: make migrate-force version=1"; \
		exit 1; \
	fi
	migrate -path $(MIGRATIONS_DIR) -database $(DB_PRIMARY_URL) force $(version)

# Build the application
build:
	go build -o bin/app ./cmd

# Run the application
run:
	go run ./cmd

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Start Docker containers
docker-up:
	docker-compose up -d

# Stop Docker containers
docker-down:
	docker-compose down

# Build Docker images
docker-build:
	docker-compose build

# Show Docker logs
docker-logs:
	docker-compose logs -f

# Exec into app container
docker-exec:
	docker-compose exec app sh

# Clean up Docker containers and volumes
docker-clean:
	docker-compose down -v --remove-orphans
	docker system prune -f

# Generate sqlc code
sqlc-gen:
	sqlc generate

# Full setup: docker, migrate, generate, build
setup: docker-up migrate-up-primary sqlc-gen build

# Restart app
restart-app:
	docker-compose restart app
