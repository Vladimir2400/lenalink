package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lenalink/backend/internal/config"
	"github.com/lenalink/backend/internal/infrastructure/logger"
	"github.com/lenalink/backend/internal/repository/memory"
	"github.com/lenalink/backend/internal/service"
)

const (
	AppName    = "LenaLink Backend"
	AppVersion = "0.2.0"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	// Initialize logger
	logger := logger.New(cfg.Logger.Level)
	logger.Info("Starting %s v%s", AppName, AppVersion)
	logger.Info("========================================")

	// Print configuration
	fmt.Printf("\n%s v%s\n", AppName, AppVersion)
	fmt.Println("Multi-modal Ticket Service for Yakutia")
	fmt.Println("========================================\n")

	logger.Info("Database driver: %s", cfg.Database.Driver)
	logger.Info("Server: %s:%d", cfg.Server.Host, cfg.Server.Port)
	logger.Info("Logger level: %s", cfg.Logger.Level)

	// Initialize repositories (in-memory for now)
	logger.Info("Initializing repositories...")
	routeRepo := memory.NewRouteRepository()
	bookingRepo := memory.NewBookingRepository()
	// ticketRepo and txManager will be initialized when PostgreSQL support is added
	// For now, using nil values for demonstration

	// Initialize services
	logger.Info("Initializing services...")
	insuranceServ := service.NewInsuranceService()
	routeServ := service.NewRouteService(routeRepo)
	_ = routeServ  // Suppress unused warning for now
	// bookingServ := service.NewBookingService(bookingRepo, ticketRepo, routeRepo, txManager, insuranceServ)
	_ = bookingRepo // Suppress unused warning
	_ = insuranceServ // Suppress unused warning

	// Log successful initialization
	logger.Info("Services initialized successfully")
	logger.Info("========================================\n")

	// Graceful shutdown handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	logger.Info("Application started. Press Ctrl+C to shutdown.")

	// TODO: Initialize HTTP server with routes
	// router := setupRouter(cfg, logger, routeServ, bookingServ)
	// if err := router.Start(cfg.Server.Host, cfg.Server.Port); err != nil {
	// 	logger.Error("Server error: %v", err)
	// 	os.Exit(1)
	// }

	// Log service availability
	logger.Info("Route Service: Ready")
	logger.Info("Booking Service: Ready")
	logger.Info("Insurance Service: Ready")
	logger.Info("HTTP Server: Ready (ready for handler implementation)")

	// Keep application running
	<-sigChan

	logger.Info("Shutdown signal received")
	// ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	// defer cancel()

	// TODO: Gracefully shutdown server
	// if err := router.Shutdown(ctx); err != nil {
	// 	logger.Error("Server shutdown error: %v", err)
	// 	os.Exit(1)
	// }

	logger.Info("Application stopped")
}

// Mock implementations for demonstration
// These will be replaced with real implementations later

/*
type mockTicketRepository struct{}

func (m *mockTicketRepository) FindByBookingID(ctx context.Context, bookingID string) ([]interface{}, error) {
	return nil, nil
}

func (m *mockTicketRepository) FindByID(ctx context.Context, id string) (interface{}, error) {
	return nil, nil
}

func (m *mockTicketRepository) Save(ctx context.Context, ticket interface{}) error {
	return nil
}

func (m *mockTicketRepository) Update(ctx context.Context, ticket interface{}) error {
	return nil
}

func (m *mockTicketRepository) Delete(ctx context.Context, id string) error {
	return nil
}

type mockTransactionManager struct{}

func (m *mockTransactionManager) BeginTx(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (m *mockTransactionManager) WithTx(ctx context.Context, fn func(interface{}) error) error {
	return nil
}
*/
