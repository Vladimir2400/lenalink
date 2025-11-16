package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/internal/repository"
)

// BookingRepository is an in-memory implementation of repository.BookingRepository
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

	// Return a copy to prevent external modification
	bookingCopy := *booking
	return &bookingCopy, nil
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
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.bookings[booking.ID]; exists {
		return fmt.Errorf("booking with ID %s already exists", booking.ID)
	}

	// Store a copy to prevent external modification
	bookingCopy := *booking
	r.bookings[booking.ID] = &bookingCopy

	return nil
}

// Update modifies an existing booking
func (r *BookingRepository) Update(ctx context.Context, booking *domain.Booking) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.bookings[booking.ID]; !exists {
		return domain.ErrBookingNotFound
	}

	// Store a copy to prevent external modification
	bookingCopy := *booking
	r.bookings[booking.ID] = &bookingCopy

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

// FindByPassenger finds bookings by passenger email
func (r *BookingRepository) FindByPassenger(ctx context.Context, email string) ([]domain.Booking, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	bookings := make([]domain.Booking, 0)
	for _, booking := range r.bookings {
		if booking.Passenger.Email == email {
			bookings = append(bookings, *booking)
		}
	}

	return bookings, nil
}

// FindByStatus finds bookings by status
func (r *BookingRepository) FindByStatus(ctx context.Context, status domain.BookingStatus) ([]domain.Booking, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	bookings := make([]domain.Booking, 0)
	for _, booking := range r.bookings {
		if booking.Status == status {
			bookings = append(bookings, *booking)
		}
	}

	return bookings, nil
}
