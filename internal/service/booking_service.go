package service

import (
	"context"
	"fmt"
	"time"

	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/internal/repository"
	"github.com/lenalink/backend/pkg/utils"
)

// BookingService implements business logic for bookings
type BookingService struct {
	bookingRepo   repository.BookingRepository
	ticketRepo    repository.TicketRepository
	routeRepo     repository.RouteRepository
	txManager     repository.TransactionManager
	insuranceServ *InsuranceService
}

// NewBookingService creates a new booking service
func NewBookingService(
	bookingRepo repository.BookingRepository,
	ticketRepo repository.TicketRepository,
	routeRepo repository.RouteRepository,
	txManager repository.TransactionManager,
	insuranceServ *InsuranceService,
) *BookingService {
	return &BookingService{
		bookingRepo:   bookingRepo,
		ticketRepo:    ticketRepo,
		routeRepo:     routeRepo,
		txManager:     txManager,
		insuranceServ: insuranceServ,
	}
}

// CreateBooking creates a new booking with ACID guarantees
func (s *BookingService) CreateBooking(ctx context.Context, req *domain.BookingCreateRequest) (*domain.Booking, error) {
	if req == nil {
		return nil, fmt.Errorf("booking request cannot be nil")
	}

	// Validate request
	if err := s.validateBookingRequest(req); err != nil {
		return nil, err
	}

	// Get route
	route, err := s.routeRepo.FindByID(ctx, req.RouteID)
	if err != nil {
		return nil, fmt.Errorf("route not found: %w", err)
	}

	// Create booking object
	booking := &domain.Booking{
		ID:        utils.GenerateBookingID(),
		RouteID:   req.RouteID,
		Passengers: req.Passengers,
		Status:    domain.BookingStatusPending,
		BookedAt:  time.Now(),
		PaymentDeadline: time.Now().Add(24 * time.Hour),
		CancellationDeadline: time.Now().Add(2 * time.Hour),
	}

	// Calculate prices
	totalPrice := 0.0
	for range req.Passengers {
		totalPrice += route.TotalPrice
	}
	booking.TotalPrice = totalPrice

	// Calculate insurance if requested
	if req.IncludeInsurance {
		insurance, err := s.insuranceServ.CalculateInsurance(ctx, route, booking)
		if err != nil {
			return nil, fmt.Errorf("insurance calculation failed: %w", err)
		}
		booking.InsurancePremium = insurance.Premium
		booking.InsuranceIncluded = true

		// Create guarantee
		booking.Guarantee = &domain.BookingGuarantee{
			GuaranteeID:    utils.GenerateID(),
			BookingID:      booking.ID,
			ValidUntil:     time.Now().Add(365 * 24 * time.Hour),
			InsuranceLevel: insurance.Level,
			CoverageAmount: insurance.CoverageAmount,
			CreatedAt:      time.Now(),
			SOSHotline:     "+7-800-555-0123",
			RefundPolicy:   "Standard refund policy applies",
		}
	}

	// Use transaction for ACID compliance
	if s.txManager != nil {
		return s.createBookingWithTransaction(ctx, booking, req.Passengers, route)
	}

	// Fallback for in-memory repository (no transaction support)
	return s.createBookingWithoutTransaction(ctx, booking, req.Passengers, route)
}

// GetBookingByID retrieves a booking by ID
func (s *BookingService) GetBookingByID(ctx context.Context, id string) (*domain.Booking, error) {
	if id == "" {
		return nil, fmt.Errorf("booking ID cannot be empty")
	}

	return s.bookingRepo.FindByID(ctx, id)
}

// GetBookingsByRouteID retrieves all bookings for a route
func (s *BookingService) GetBookingsByRouteID(ctx context.Context, routeID string) ([]domain.Booking, error) {
	if routeID == "" {
		return nil, fmt.Errorf("route ID cannot be empty")
	}

	return s.bookingRepo.FindByRouteID(ctx, routeID)
}

// CancelBooking cancels an existing booking
func (s *BookingService) CancelBooking(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("booking ID cannot be empty")
	}

	booking, err := s.bookingRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if booking.Status == domain.BookingStatusCancelled {
		return fmt.Errorf("booking is already cancelled")
	}

	if time.Now().After(booking.CancellationDeadline) {
		return fmt.Errorf("cancellation deadline has passed")
	}

	booking.Status = domain.BookingStatusCancelled
	booking.Notes = "Cancelled by customer"

	return s.bookingRepo.Update(ctx, booking)
}

// Private methods

func (s *BookingService) createBookingWithTransaction(ctx context.Context, booking *domain.Booking, passengers []domain.Passenger, route *domain.Route) (*domain.Booking, error) {
	err := s.txManager.WithTx(ctx, func(tx repository.Transaction) error {
		bookingRepo := tx.GetBookingRepository()
		ticketRepo := tx.GetTicketRepository()

		// Save booking
		if err := bookingRepo.Save(ctx, booking); err != nil {
			return fmt.Errorf("failed to save booking: %w", err)
		}

		// Create tickets for each passenger and segment
		for _, passenger := range passengers {
			for _, segment := range route.Segments {
				ticket := &domain.BookingSegmentTicket{
					ID:          utils.GenerateID(),
					BookingID:   booking.ID,
					SegmentID:   segment.ID,
					PassengerID: passenger.ID,
					Provider:    segment.Provider,
					TicketNumber: utils.GenerateTicketNumber(segment.Provider),
					Price:       segment.Price,
					Status:      "confirmed",
					CreatedAt:   time.Now(),
				}

				if err := ticketRepo.Save(ctx, ticket); err != nil {
					return fmt.Errorf("failed to save ticket: %w", err)
				}

				booking.Tickets = append(booking.Tickets, *ticket)
			}
		}

		// Update booking with all tickets
		if err := bookingRepo.Update(ctx, booking); err != nil {
			return fmt.Errorf("failed to update booking with tickets: %w", err)
		}

		booking.Status = domain.BookingStatusConfirmed
		return nil
	})

	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (s *BookingService) createBookingWithoutTransaction(ctx context.Context, booking *domain.Booking, passengers []domain.Passenger, route *domain.Route) (*domain.Booking, error) {
	// Save booking
	if err := s.bookingRepo.Save(ctx, booking); err != nil {
		return nil, fmt.Errorf("failed to save booking: %w", err)
	}

	// Create tickets for each passenger and segment
	for _, passenger := range passengers {
		for _, segment := range route.Segments {
			ticket := &domain.BookingSegmentTicket{
				ID:          utils.GenerateID(),
				BookingID:   booking.ID,
				SegmentID:   segment.ID,
				PassengerID: passenger.ID,
				Provider:    segment.Provider,
				TicketNumber: utils.GenerateTicketNumber(segment.Provider),
				Price:       segment.Price,
				Status:      "confirmed",
				CreatedAt:   time.Now(),
			}

			if err := s.ticketRepo.Save(ctx, ticket); err != nil {
				// In a real scenario with transaction support, this would be rolled back
				return nil, fmt.Errorf("failed to save ticket: %w", err)
			}

			booking.Tickets = append(booking.Tickets, *ticket)
		}
	}

	booking.Status = domain.BookingStatusConfirmed
	return booking, nil
}

func (s *BookingService) validateBookingRequest(req *domain.BookingCreateRequest) error {
	if req.RouteID == "" {
		return fmt.Errorf("route_id is required")
	}

	if len(req.Passengers) == 0 {
		return fmt.Errorf("at least one passenger is required")
	}

	for i, p := range req.Passengers {
		if p.FirstName == "" {
			return fmt.Errorf("passenger %d first_name is required", i+1)
		}
		if p.LastName == "" {
			return fmt.Errorf("passenger %d last_name is required", i+1)
		}
		if p.PassportNumber == "" {
			return fmt.Errorf("passenger %d passport_number is required", i+1)
		}
	}

	return nil
}
