#!/bin/bash
# Database initialization script for LenaLink project

set -e

# Configuration
DB_NAME="lenalink_db"
DB_USER="lenalink"
DB_PASSWORD="password"
DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== LenaLink Database Initialization ===${NC}\n"

# Check if PostgreSQL is available
if ! command -v psql &> /dev/null; then
    echo -e "${RED}Error: psql not found. Please install PostgreSQL client.${NC}"
    exit 1
fi

# Create database and user
echo -e "${BLUE}Creating database '$DB_NAME' and user '$DB_USER'...${NC}"

# Try to create user and database
# Note: This assumes PostgreSQL superuser can connect without password
PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -p "$DB_PORT" -U postgres <<-EOSQL
    -- Drop existing database and user if they exist
    SELECT pg_terminate_backend(pg_stat_activity.pid)
    FROM pg_stat_activity
    WHERE pg_stat_activity.datname = '$DB_NAME'
      AND pid <> pg_backend_pid();

    DROP DATABASE IF EXISTS "$DB_NAME";
    DROP USER IF EXISTS "$DB_USER";

    -- Create user
    CREATE USER "$DB_USER" WITH PASSWORD '$DB_PASSWORD';

    -- Create database
    CREATE DATABASE "$DB_NAME" OWNER "$DB_USER"
        ENCODING 'UTF8'
        LC_COLLATE 'en_US.UTF-8'
        LC_CTYPE 'en_US.UTF-8';

    -- Grant privileges
    GRANT ALL PRIVILEGES ON DATABASE "$DB_NAME" TO "$DB_USER";

    -- Connect to the new database and set additional permissions
    \c "$DB_NAME"

    GRANT USAGE ON SCHEMA public TO "$DB_USER";
    GRANT CREATE ON SCHEMA public TO "$DB_USER";
    ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO "$DB_USER";
    ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO "$DB_USER";
    ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON FUNCTIONS TO "$DB_USER";
EOSQL

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Database and user created successfully${NC}"
else
    echo -e "${RED}✗ Failed to create database and user${NC}"
    exit 1
fi

# Test connection
echo -e "${BLUE}Testing connection to database...${NC}"
PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -c "SELECT 1;" > /dev/null 2>&1

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Database connection successful${NC}"
else
    echo -e "${RED}✗ Failed to connect to database${NC}"
    exit 1
fi

echo ""
echo -e "${BLUE}Database initialization completed!${NC}"
echo ""
echo -e "${BLUE}Next steps:${NC}"
echo "1. Run migrations: make migrate-up"
echo "2. Access database:  psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME"
echo "3. Access pgAdmin:   http://localhost:5050"
echo "   Login: admin@lenalink.com / admin"
echo ""
