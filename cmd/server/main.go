package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	httphandler "github.com/lenalink/backend/internal/handler/http"
	"github.com/lenalink/backend/internal/config"
	postgres "github.com/lenalink/backend/internal/repository/postgres"
	"github.com/lenalink/backend/internal/service"
	"github.com/lenalink/backend/pkg/utils"
)

const (
	AppName    = "LenaLink"
	AppVersion = "0.5.0"
)

func main() {
	// Load .env file (ignore error if file doesn't exist)
	_ = godotenv.Load()

	log.Println("🚀 Starting LenaLink Backend...")

	// Load configuration
	cfg := config.Load()

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	// Print banner
	fmt.Printf("\n%s v%s\n", AppName, AppVersion)
	fmt.Println("Multi-modal Transport Aggregator with Unified Booking")
	fmt.Println("========================================\n")

	log.Printf("Database driver: %s", cfg.Database.Driver)
	log.Printf("Server: %s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Logger level: %s", cfg.Logger.Level)

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

	// Initialize cache
	log.Println("💾 Initializing cache...")
	routeCache := utils.NewCache(10*time.Minute, 1000)
	defer routeCache.Stop()
	log.Println("✓ Cache initialized")

	// Initialize PostgreSQL repositories
	log.Println("🗄️  Initializing repositories...")
	routeRepo := postgres.NewRouteRepository(db)
	bookingRepo := postgres.NewBookingRepository(db)
	log.Println("✓ Repositories initialized")

	// Initialize services
	log.Println("⚙️  Initializing services...")
	routeService := service.NewRouteService(routeRepo)
	commissionSvc := service.NewCommissionService(service.DefaultCommissionConfig())
	insuranceSvc := service.NewInsuranceService(service.DefaultInsuranceConfig())
	paymentSvc := service.NewPaymentService(service.NewMockPaymentGateway(0.0))
	providerBooking := service.NewMockProviderBookingService(0.0)
	bookingService := service.NewBookingService(
		routeRepo,
		bookingRepo,
		commissionSvc,
		insuranceSvc,
		paymentSvc,
		providerBooking,
	)
	log.Println("✓ Services initialized")

	// Initialize router with handlers
	log.Println("🛣️  Setting up HTTP routes...")
	router := httphandler.NewRouter(routeService, bookingService)
	log.Println("✓ HTTP routes configured")

	// Server configuration
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.ShutdownTimeout,
	}

	// Channel for shutdown signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		log.Printf("🌐 HTTP server listening on http://%s\n", server.Addr)
		log.Println("📝 Health check: GET /health")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Log service availability
	log.Println("========================================")
	log.Println("✓ Route Service: Ready (graph-based pathfinding)")
	log.Println("✓ Booking Service: Ready (multi-segment + ACID)")
	log.Println("✓ Commission Service: Ready (5-15% markup)")
	log.Println("✓ Insurance Service: Ready (5% base premium)")
	log.Println("✓ Payment Service: Ready (mock gateway)")
	log.Println("✓ HTTP Server: Ready (listening)")
	log.Println("========================================\n")

	// Wait for shutdown signal
	sig := <-sigChan
	log.Printf("\n⛔ Received signal: %v\n", sig)

	// Graceful shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancel()

	log.Println("🛑 Shutting down server gracefully...")
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	// Clean up resources
	log.Println("🧹 Cleaning up resources...")
	routeCache.Stop()
	log.Println("✓ Server stopped gracefully")
}
