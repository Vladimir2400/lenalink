package memory

import (
	"context"
	"fmt"
	"sync"
	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/internal/repository"
)

// BookingRepository implements BookingRepository interface using in-memory storage
type BookingRepository struct {
	mu       sync.RWMutex
	bookings map[string]*domain.Booking
}

// NewBookingRepository creates a new in-memory booking repository
func NewBookingRepository() repository.BookingRepository {
	return &BookingRepository{
		bookings: make(map[string]*domain.Booking),
	}
}

// FindByID retrieves a booking by ID
func (r *BookingRepository) FindByID(ctx context.Context, id string) (*domain.Booking, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	booking, exists := r.bookings[id]
	if !exists {
		return nil, domain.ErrBookingNotFound
	}

	return booking, nil
}

// FindAll retrieves all bookings
func (r *BookingRepository) FindAll(ctx context.Context) ([]domain.Booking, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	bookings := make([]domain.Booking, 0, len(r.bookings))
	for _, booking := range r.bookings {
		bookings = append(bookings, *booking)
	}

	return bookings, nil
}

// Save stores a new booking
func (r *BookingRepository) Save(ctx context.Context, booking *domain.Booking) error {
	if booking == nil {
		return domain.ErrInvalidBooking
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.bookings[booking.ID]; exists {
		return fmt.Errorf("booking with ID %s already exists", booking.ID)
	}

	r.bookings[booking.ID] = booking

	return nil
}

// Update modifies an existing booking
func (r *BookingRepository) Update(ctx context.Context, booking *domain.Booking) error {
	if booking == nil {
		return domain.ErrInvalidBooking
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.bookings[booking.ID]; !exists {
		return domain.ErrBookingNotFound
	}

	r.bookings[booking.ID] = booking

	return nil
}

// Delete removes a booking
func (r *BookingRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.bookings[id]; !exists {
		return domain.ErrBookingNotFound
	}

	delete(r.bookings, id)

	return nil
}

// FindByRouteID retrieves bookings for a specific route
func (r *BookingRepository) FindByRouteID(ctx context.Context, routeID string) ([]domain.Booking, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var results []domain.Booking

	for _, booking := range r.bookings {
		if booking.RouteID == routeID {
			results = append(results, *booking)
		}
	}

	return results, nil
}
