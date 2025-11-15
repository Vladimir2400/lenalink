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

	// FindByRouteID retrieves bookings for a specific route
	FindByRouteID(ctx context.Context, routeID string) ([]domain.Booking, error)
}

// TicketRepository defines operations for ticket persistence
type TicketRepository interface {
	// FindByBookingID retrieves all tickets for a booking
	FindByBookingID(ctx context.Context, bookingID string) ([]domain.BookingSegmentTicket, error)

	// FindByID retrieves a ticket by ID
	FindByID(ctx context.Context, id string) (*domain.BookingSegmentTicket, error)

	// Save stores a new ticket
	Save(ctx context.Context, ticket *domain.BookingSegmentTicket) error

	// Update modifies an existing ticket
	Update(ctx context.Context, ticket *domain.BookingSegmentTicket) error

	// Delete removes a ticket
	Delete(ctx context.Context, id string) error
}

// Transaction defines ACID transaction operations
type Transaction interface {
	// Commit commits the transaction
	Commit(ctx context.Context) error

	// Rollback rolls back the transaction
	Rollback(ctx context.Context) error

	// GetBookingRepository returns booking repository for this transaction
	GetBookingRepository() BookingRepository

	// GetTicketRepository returns ticket repository for this transaction
	GetTicketRepository() TicketRepository

	// GetRouteRepository returns route repository for this transaction
	GetRouteRepository() RouteRepository
}

// TransactionManager defines transaction lifecycle management
type TransactionManager interface {
	// BeginTx starts a new transaction
	BeginTx(ctx context.Context) (Transaction, error)

	// WithTx executes a function within a transaction
	WithTx(ctx context.Context, fn func(tx Transaction) error) error
}
