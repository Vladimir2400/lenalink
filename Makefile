.PHONY: help migrate-up migrate-down migrate-force migrate-version migrate-create db-reset db-drop docker-up docker-down

# Database configuration
DB_USER=lenalink
DB_PASSWORD=password
DB_NAME=lenalink_db
DB_HOST=localhost
DB_PORT=15432
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

help:
	@echo "Available commands:"
	@echo "  make migrate-up        - Apply all pending migrations"
	@echo "  make migrate-down      - Rollback last migration"
	@echo "  make migrate-force     - Force migration version (use VERSION=N)"
	@echo "  make migrate-version   - Show current migration version"
	@echo "  make migrate-create    - Create new migration (use NAME=migration_name)"
	@echo "  make db-reset          - Drop and recreate database"
	@echo "  make db-drop           - Drop database only"
	@echo "  make docker-up         - Start PostgreSQL with docker-compose"
	@echo "  make docker-down       - Stop PostgreSQL containers"
	@echo "  make docker-logs       - View PostgreSQL logs"

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
migrate-up:
	@echo "Applying migrations..."
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	@echo "Rolling back last migration..."
	migrate -path migrations -database "$(DB_URL)" down 1

migrate-force:
	@if [ -z "$(VERSION)" ]; then \
		echo "Usage: make migrate-force VERSION=N"; \
		exit 1; \
	fi
	@echo "Forcing migration to version $(VERSION)..."
	migrate -path migrations -database "$(DB_URL)" force $(VERSION)

migrate-version:
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
db-reset: db-drop migrate-up
	@echo "✓ Database reset complete"

db-drop:
	@echo "Dropping database $(DB_NAME)..."
	dropdb --if-exists -U $(DB_USER) $(DB_NAME)
	@echo "✓ Database dropped"

db-create:
	@echo "Creating database $(DB_NAME)..."
	createdb -U $(DB_USER) $(DB_NAME)
	@echo "✓ Database created"

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
