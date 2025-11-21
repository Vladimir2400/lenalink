package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lenalink/backend/internal/config"
	"github.com/lenalink/backend/internal/repository/postgres"
	syncpkg "github.com/lenalink/backend/pkg/sync"
	"github.com/lenalink/backend/pkg/sync/api/aviasales"
	"github.com/lenalink/backend/pkg/sync/api/gars"
	"github.com/lenalink/backend/pkg/sync/api/rzd"
)

const (
	AppName    = "LenaLink Data Seeder"
	AppVersion = "0.1.0"
)

func main() {
	// Load .env file (ignore error if file doesn't exist)
	_ = godotenv.Load()

	fmt.Printf("\n%s v%s\n", AppName, AppVersion)
	fmt.Println("Data synchronization tool for LenaLink")
	fmt.Println("========================================\n")

	// Load main application configuration
	cfg := config.Load()

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	log.Printf("Database: %s", cfg.Database.Driver)

	// Connect to PostgreSQL
	log.Println("📦 Connecting to PostgreSQL...")
	db, err := postgres.NewDatabase(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Println("✓ PostgreSQL connected")

	// Run migrations
	log.Println("🔄 Running database migrations...")
	if err := postgres.RunMigrations(db.DB()); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("✓ Migrations completed")

	// Load sync configuration
	log.Println("⚙️  Loading sync configuration...")
	garsConfig := syncpkg.LoadGARSConfig()
	aviasalesConfig := syncpkg.LoadAviasalesConfig()
	rzdConfig := syncpkg.LoadRZDConfig()

	log.Printf("GARS BaseURL: %s", garsConfig.BaseURL)
	log.Printf("GARS Username: %s", garsConfig.Username)
	log.Printf("Aviasales Token: %s", maskString(aviasalesConfig.Token))
	log.Printf("RZD Enabled: %v", rzdConfig.Enabled)

	// Validate sync configuration
	if err := garsConfig.Validate(); err != nil {
		log.Fatalf("GARS configuration error: %v", err)
	}

	// Create provider clients
	log.Println("\n🔌 Initializing provider clients...")

	log.Println("  - Creating GARS client...")
	garsClient, err := gars.NewClient(gars.Config{
		BaseURL:  garsConfig.BaseURL,
		Username: garsConfig.Username,
		Password: garsConfig.Password,
		Timeout:  garsConfig.Timeout,
	})
	if err != nil {
		log.Fatalf("Failed to create GARS client: %v", err)
	}
	log.Println("  ✓ GARS client created")

	log.Println("  - Creating Aviasales client...")
	aviasalesClient, err := aviasales.NewClient(aviasales.Config{
		APIToken: aviasalesConfig.Token,
	})
	if err != nil {
		log.Fatalf("Failed to create Aviasales client: %v", err)
	}
	log.Println("  ✓ Aviasales client created")

	log.Println("  - Creating RZD client (mock)...")
	rzdClient := rzd.NewMockClient()
	log.Println("  ✓ RZD client created")

	// Initialize repositories
	log.Println("\n🗄️  Initializing repositories...")
	stopRepo := postgres.NewStopRepository(db)
	segmentRepo := postgres.NewSegmentRepository(db)
	log.Println("✓ Repositories initialized")

	// Create sync service
	log.Println("\n🔄 Creating sync service...")
	syncer := syncpkg.New(garsClient, aviasalesClient, rzdClient, stopRepo, segmentRepo)
	log.Println("✓ Sync service created")

	// Check current data
	log.Println("\n📊 Checking current data in database...")
	ctx := context.Background()

	stops, err := stopRepo.FindAll(ctx)
	if err != nil {
		log.Printf("Warning: Could not count stops: %v", err)
	} else {
		log.Printf("  Current stops in database: %d", len(stops))
	}

	segments, err := segmentRepo.FindAll(ctx)
	if err != nil {
		log.Printf("Warning: Could not count segments: %v", err)
	} else {
		log.Printf("  Current segments in database: %d", len(segments))
	}

	// Run synchronization
	log.Println("\n🚀 Starting data synchronization...")
	fmt.Println("========================================")

	startTime := time.Now()

	// Option to sync specific provider
	provider := os.Getenv("SYNC_PROVIDER")
	if provider != "" {
		log.Printf("Syncing only provider: %s", provider)
		if err := syncer.SyncProvider(ctx, syncpkg.Provider(provider)); err != nil {
			log.Fatalf("❌ Sync failed: %v", err)
		}
	} else {
		log.Println("Syncing all providers...")
		if err := syncer.SyncAll(ctx); err != nil {
			log.Fatalf("❌ Sync failed: %v", err)
		}
	}

	elapsed := time.Since(startTime)
	fmt.Println("========================================")
	log.Printf("✅ Synchronization completed in %s\n", elapsed.Round(time.Second))

	// Show results
	log.Println("\n📈 Final statistics:")

	stops, err = stopRepo.FindAll(ctx)
	if err != nil {
		log.Printf("Warning: Could not count stops: %v", err)
	} else {
		log.Printf("  Total stops: %d", len(stops))
	}

	segments, err = segmentRepo.FindAll(ctx)
	if err != nil {
		log.Printf("Warning: Could not count segments: %v", err)
	} else {
		log.Printf("  Total segments: %d", len(segments))
	}

	log.Println("\n✓ Seeding completed successfully!")
}

// maskString masks sensitive data, showing only first 4 characters
func maskString(s string) string {
	if s == "" {
		return "(not set)"
	}
	if len(s) <= 4 {
		return "****"
	}
	return s[:4] + "****"
}
