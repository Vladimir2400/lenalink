package repository

import (
	"context"
	"github.com/lenalink/backend/internal/domain"
)

// RouteRepository defines operations for route persistence
type RouteRepository interface {
	// FindByID retrieves a route by ID
	FindByID(ctx context.Context, id string) (*domain.Route, error)

	// FindAll retrieves all routes
	FindAll(ctx context.Context) ([]domain.Route, error)

	// Save stores a new route
	Save(ctx context.Context, route *domain.Route) error

	// Update modifies an existing route
	Update(ctx context.Context, route *domain.Route) error

	// Delete removes a route
	Delete(ctx context.Context, id string) error

	// FindByCriteria searches routes by criteria
	FindByCriteria(ctx context.Context, criteria *domain.RouteSearchCriteria) ([]domain.Route, error)
}

// BookingRepository defines operations for booking persistence
type BookingRepository interface {
	// FindByID retrieves a booking by ID
	FindByID(ctx context.Context, id string) (*domain.Booking, error)

	// FindAll retrieves all bookings
	FindAll(ctx context.Context) ([]domain.Booking, error)

	// Save stores a new booking
	Save(ctx context.Context, booking *domain.Booking) error

	// Update modifies an existing booking
	Update(ctx context.Context, booking *domain.Booking) error

	// Delete removes a booking
	Delete(ctx context.Context, id string) error

	// FindByPassenger finds bookings by passenger email
	FindByPassenger(ctx context.Context, email string) ([]domain.Booking, error)

	// FindByStatus finds bookings by status
	FindByStatus(ctx context.Context, status domain.BookingStatus) ([]domain.Booking, error)
}

// Transaction represents a database transaction for ACID guarantees
type Transaction interface {
	// Commit commits the transaction
	Commit() error

	// Rollback rolls back the transaction
	Rollback() error

	// RouteRepository returns a route repository within this transaction
	RouteRepository() RouteRepository

	// BookingRepository returns a booking repository within this transaction
	BookingRepository() BookingRepository
}

// TransactionManager manages database transactions
type TransactionManager interface {
	// Begin starts a new transaction
	Begin(ctx context.Context) (Transaction, error)

	// WithTransaction executes a function within a transaction
	// Automatically commits on success, rolls back on error
	WithTransaction(ctx context.Context, fn func(tx Transaction) error) error
}
