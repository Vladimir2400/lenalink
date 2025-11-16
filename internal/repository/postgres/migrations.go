package postgres

import (
	"database/sql"
	"fmt"
)

// RunMigrations checks and reports migration status
// Note: Manual migration management is required via golang-migrate CLI or psql
// All migrations should be applied before running the application
func RunMigrations(db *sql.DB) error {
	// Check if schema_migrations table exists
	var exists bool
	err := db.QueryRow(`
		SELECT EXISTS (
			SELECT FROM information_schema.tables
			WHERE table_schema = 'public'
			AND table_name = 'schema_migrations'
		);
	`).Scan(&exists)

	if err != nil {
		return fmt.Errorf("could not check schema_migrations table: %w", err)
	}

	if !exists {
		return fmt.Errorf("schema_migrations table does not exist - run migrations first")
	}

	// Check current migration version
	var version int
	var dirty bool
	err = db.QueryRow(`
		SELECT COALESCE(MAX(version), 0), false
		FROM schema_migrations
		LIMIT 1
	`).Scan(&version, &dirty)

	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("could not get migration version: %w", err)
	}

	if version == 0 {
		fmt.Println("✓ Database initialized (no migrations applied yet)")
	} else {
		if dirty {
			fmt.Printf("⚠ Database is in dirty state at version %v\n", version)
		} else {
			fmt.Printf("✓ Database migrated to version %v\n", version)
		}
	}

	return nil
}

// MigrateDown is a placeholder for rollback functionality
// Manual rollback should be done using golang-migrate CLI or psql
func MigrateDown(db *sql.DB) error {
	return fmt.Errorf("migrate down not supported - use golang-migrate CLI or psql directly")
}

// GetMigrationVersion returns current migration version from schema_migrations table
func GetMigrationVersion(db *sql.DB) (uint, bool, error) {
	var version uint
	var dirty bool

	err := db.QueryRow(`
		SELECT COALESCE(MAX(version), 0), false
		FROM schema_migrations
		LIMIT 1
	`).Scan(&version, &dirty)

	if err != nil && err != sql.ErrNoRows {
		return 0, false, err
	}

	return version, dirty, nil
}
