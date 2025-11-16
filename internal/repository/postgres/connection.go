package postgres

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/lenalink/backend/internal/config"
)

// Database wraps sql.DB for PostgreSQL operations
type Database struct {
	db *sql.DB
}

// NewDatabase creates a new PostgreSQL database connection
func NewDatabase(cfg config.DatabaseConfig) (*Database, error) {
	// Open connection
	db, err := sql.Open("postgres", cfg.ConnectionString())
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(cfg.MaxConnections)
	db.SetMaxIdleConns(cfg.MinConnections)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping database: %w", err)
	}

	return &Database{db: db}, nil
}

// DB returns the underlying sql.DB connection
func (d *Database) DB() *sql.DB {
	return d.db
}

// Close closes the database connection
func (d *Database) Close() error {
	if d.db != nil {
		return d.db.Close()
	}
	return nil
}

// Ping tests the database connection
func (d *Database) Ping() error {
	return d.db.Ping()
}

// GetStats returns connection pool statistics
func (d *Database) GetStats() sql.DBStats {
	return d.db.Stats()
}
