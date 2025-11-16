package main

import (
	"fmt"
	"log"

	"github.com/lenalink/backend/config"
	postgres "github.com/lenalink/backend/internal/repository/postgres"
)

func main() {
	// Load database configuration
	dbConfig := config.LoadDatabaseConfig()

	// Connect to PostgreSQL
	db, err := postgres.NewDatabase(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Run migrations
	if err := postgres.RunMigrations(db.DB()); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	fmt.Println("✓ All migrations applied successfully!")
}
