.PHONY: help migrate-up migrate-down migrate-force migrate-version migrate-create db-reset db-drop docker-up docker-down

# Database configuration
DB_USER=lenalink
DB_PASSWORD=password
DB_NAME=lenalink_db
DB_HOST=localhost
DB_PORT=15432
# For local use (from host machine)
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
# For Docker use (from within containers)
DB_URL_DOCKER=postgres://$(DB_USER):$(DB_PASSWORD)@postgres:5432/$(DB_NAME)?sslmode=disable

help:
	@echo "Available commands:"
	@echo ""
	@echo "🐳 Docker Migration Commands (recommended for production):"
	@echo "  make migrate-up            - Apply all pending migrations via Docker"
	@echo "  make migrate-down          - Rollback last migration via Docker"
	@echo "  make migrate-force         - Force migration version via Docker (use VERSION=N)"
	@echo "  make migrate-version       - Show current migration version via Docker"
	@echo ""
	@echo "💻 Local Migration Commands (requires golang-migrate CLI):"
	@echo "  make migrate-up-local      - Apply all pending migrations locally"
	@echo "  make migrate-down-local    - Rollback last migration locally"
	@echo "  make migrate-force-local   - Force migration version locally (use VERSION=N)"
	@echo "  make migrate-version-local - Show current migration version locally"
	@echo ""
	@echo "📦 Database Management:"
	@echo "  make migrate-create        - Create new migration (use NAME=migration_name)"
	@echo "  make db-reset              - Drop, recreate database and apply migrations"
	@echo "  make db-reset-with-data    - db-reset + load test data from scripts/"
	@echo "  make db-drop               - Drop database only"
	@echo "  make db-create             - Create database only"
	@echo ""
	@echo "🌱 Data Seeding:"
	@echo "  make seed                  - Sync data from external providers (GARS, Aviasales, RZD)"
	@echo "  make seed-gars             - Sync only GARS data"
	@echo "  make seed-aviasales        - Sync only Aviasales data"
	@echo "  make seed-rzd              - Sync only RZD data"
	@echo ""
	@echo "🐳 Docker Commands:"
	@echo "  make docker-up             - Start all services with docker-compose"
	@echo "  make docker-down           - Stop all containers"
	@echo "  make docker-logs           - View PostgreSQL logs"

# Docker commands
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f postgres

docker-ps:
	docker-compose ps

# Database migration commands
# These commands use Docker container for migrations
migrate-up:
	@echo "Applying migrations via Docker..."
	docker compose run --rm migrate -path=/migrations -database "$(DB_URL_DOCKER)" up

migrate-down:
	@echo "Rolling back last migration via Docker..."
	docker compose run --rm migrate -path=/migrations -database "$(DB_URL_DOCKER)" down 1

migrate-force:
	@if [ -z "$(VERSION)" ]; then \
		echo "Usage: make migrate-force VERSION=N"; \
		exit 1; \
	fi
	@echo "Forcing migration to version $(VERSION) via Docker..."
	docker compose run --rm migrate -path=/migrations -database "$(DB_URL_DOCKER)" force $(VERSION)

migrate-version:
	@echo "Current migration version:"
	docker compose run --rm migrate -path=/migrations -database "$(DB_URL_DOCKER)" version

# Alternative: Local migration commands (requires golang-migrate CLI installed)
migrate-up-local:
	@echo "Applying migrations locally..."
	migrate -path migrations -database "$(DB_URL)" up

migrate-down-local:
	@echo "Rolling back last migration locally..."
	migrate -path migrations -database "$(DB_URL)" down 1

migrate-force-local:
	@if [ -z "$(VERSION)" ]; then \
		echo "Usage: make migrate-force VERSION=N"; \
		exit 1; \
	fi
	@echo "Forcing migration to version $(VERSION) locally..."
	migrate -path migrations -database "$(DB_URL)" force $(VERSION)

migrate-version-local:
	@echo "Current migration version:"
	migrate -path migrations -database "$(DB_URL)" version

migrate-create:
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make migrate-create NAME=migration_name"; \
		exit 1; \
	fi
	@echo "Creating migration: $(NAME)"
	migrate create -ext sql -dir migrations -seq $(NAME)

# Database setup commands
db-reset:
	@echo "🔄 Resetting database..."
	@docker-compose exec -T postgres psql -U $(DB_USER) -d postgres -c "DROP DATABASE IF EXISTS $(DB_NAME);" || true
	@docker-compose exec -T postgres psql -U $(DB_USER) -d postgres -c "CREATE DATABASE $(DB_NAME) OWNER $(DB_USER);"
	@echo "✓ Database dropped and recreated"
	@sleep 1
	@make migrate-up
	@echo "✓ Database reset complete"

db-drop:
	@echo "Dropping database $(DB_NAME)..."
	@docker-compose exec -T postgres psql -U $(DB_USER) -d postgres -c "DROP DATABASE IF EXISTS $(DB_NAME);" || true
	@echo "✓ Database dropped"

db-create:
	@echo "Creating database $(DB_NAME)..."
	@docker-compose exec -T postgres psql -U $(DB_USER) -d postgres -c "CREATE DATABASE $(DB_NAME) OWNER $(DB_USER);"
	@echo "✓ Database created"

# Full reset with test data
db-reset-with-data: db-reset
	@echo "📥 Loading test data..."
	@docker-compose exec -T postgres psql -U $(DB_USER) -d $(DB_NAME) < scripts/seed_yakutia_routes.sql
	@echo "✓ Test data loaded successfully"

# Development commands
dev-setup: docker-up
	@sleep 3
	@echo "Waiting for PostgreSQL to be ready..."
	@while ! pg_isready -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) > /dev/null 2>&1; do \
		echo "Waiting for PostgreSQL..."; \
		sleep 1; \
	done
	@echo "✓ PostgreSQL is ready"
	@make migrate-up

dev-clean: docker-down
	@echo "✓ Development environment cleaned"

# Testing commands
test-migrations:
	@echo "Testing migrations..."
	@make db-drop
	@make db-create
	@make migrate-up
	@echo "✓ All migrations applied successfully"
	@make db-drop

# Utility commands
psql:
	psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME)

pgadmin:
	@echo "pgAdmin is available at http://localhost:15050"
	@echo "Default credentials: admin@lenalink.com / admin"

# Seed commands - sync data from external providers
seed:
	@echo "🌱 Seeding database from all providers..."
	go run ./cmd/seed

seed-gars:
	@echo "🌱 Seeding database from GARS only..."
	SYNC_PROVIDER=gars go run ./cmd/seed

seed-aviasales:
	@echo "🌱 Seeding database from Aviasales only..."
	SYNC_PROVIDER=aviasales go run ./cmd/seed

seed-rzd:
	@echo "🌱 Seeding database from RZD only..."
	SYNC_PROVIDER=rzd go run ./cmd/seed

docker-seed:
	docker compose run --rm seed
