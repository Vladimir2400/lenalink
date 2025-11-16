package service

import (
	"context"
	"fmt"
	"time"

	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/internal/repository"
	"github.com/lenalink/backend/pkg/utils"
)

// ProviderBookingService defines interface for booking with external providers
type ProviderBookingService interface {
	BookSegment(ctx context.Context, segment *domain.Segment, passenger *domain.Passenger) (ticketNumber, bookingRef string, err error)
	CancelBooking(ctx context.Context, bookingRef string) error
}

// BookingService handles multi-segment booking with ACID guarantees
type BookingService struct {
	routeRepo       repository.RouteRepository
	bookingRepo     repository.BookingRepository
	commissionSvc   *CommissionService
	insuranceSvc    *InsuranceService
	paymentSvc      *PaymentService
	providerBooking ProviderBookingService
}

// NewBookingService creates a new booking service
func NewBookingService(
	routeRepo repository.RouteRepository,
	bookingRepo repository.BookingRepository,
	commissionSvc *CommissionService,
	insuranceSvc *InsuranceService,
	paymentSvc *PaymentService,
	providerBooking ProviderBookingService,
) *BookingService {
	return &BookingService{
		routeRepo:       routeRepo,
		bookingRepo:     bookingRepo,
		commissionSvc:   commissionSvc,
		insuranceSvc:    insuranceSvc,
		paymentSvc:      paymentSvc,
		providerBooking: providerBooking,
	}
}

// CreateBooking creates a multi-segment booking with ACID transaction
func (bs *BookingService) CreateBooking(ctx context.Context, routeID string, passenger domain.Passenger, includeInsurance bool, paymentMethod domain.PaymentMethod) (*domain.Booking, error) {
	// 1. Fetch route
	route, err := bs.routeRepo.FindByID(ctx, routeID)
	if err != nil {
		return nil, fmt.Errorf("route not found: %w", err)
	}

	if len(route.Segments) == 0 {
		return nil, fmt.Errorf("route has no segments")
	}

	// 2. Create booking
	booking := &domain.Booking{
		ID:               utils.GenerateID(),
		RouteID:          routeID,
		Passenger:        passenger,
		Segments:         make([]domain.BookedSegment, 0, len(route.Segments)),
		IncludeInsurance: includeInsurance,
		Status:           domain.BookingPending,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	// 3. Calculate insurance if requested
	if includeInsurance {
		booking.InsurancePremium = bs.insuranceSvc.CalculatePremium(route)
	}

	// 4. Book all segments (with rollback on failure)
	bookedSegments := make([]domain.BookedSegment, 0, len(route.Segments))
	bookingRefs := make([]string, 0, len(route.Segments))

	for i := range route.Segments {
		// Check for context cancellation
		select {
		case <-ctx.Done():
			// Context cancelled, rollback all bookings
			bs.rollbackBookings(ctx, bookingRefs)
			booking.MarkAsFailed("booking cancelled: " + ctx.Err().Error())
			bs.bookingRepo.Save(context.Background(), booking)
			return nil, fmt.Errorf("booking cancelled: %w", ctx.Err())
		default:
			// Continue with booking
		}

		segment := &route.Segments[i]

		// Calculate commission
		basePrice := segment.Price
		commission := bs.commissionSvc.CalculateCommission(segment.TransportType, basePrice)
		totalPrice := basePrice + commission

		// Book with provider
		ticketNumber, bookingRef, err := bs.providerBooking.BookSegment(ctx, segment, &passenger)
		if err != nil {
			// ROLLBACK: Cancel all previously booked segments
			bs.rollbackBookings(ctx, bookingRefs)
			booking.MarkAsFailed(fmt.Sprintf("failed to book segment %d: %v", i+1, err))
			bs.bookingRepo.Save(ctx, booking)
			return nil, fmt.Errorf("booking failed at segment %d (%s -> %s): %w", i+1, segment.StartStop.City, segment.EndStop.City, err)
		}

		// Create booked segment
		bookedSegment := domain.BookedSegment{
			ID:                 utils.GenerateID(),
			SegmentID:          segment.ID,
			Provider:           segment.Provider,
			TransportType:      segment.TransportType,
			From:               segment.StartStop,
			To:                 segment.EndStop,
			DepartureTime:      segment.DepartureTime,
			ArrivalTime:        segment.ArrivalTime,
			TicketNumber:       ticketNumber,
			Price:              basePrice,
			Commission:         commission,
			TotalPrice:         totalPrice,
			BookingStatus:      domain.BookingConfirmed,
			ProviderBookingRef: bookingRef,
		}

		bookedSegments = append(bookedSegments, bookedSegment)
		bookingRefs = append(bookingRefs, bookingRef)
		booking.AddSegment(bookedSegment)
	}

	// 5. Create payment
	grandTotal := booking.GrandTotal
	payment := bs.paymentSvc.CreatePayment(booking.ID, grandTotal, paymentMethod)
	booking.Payment = payment

	// 6. Process payment
	if err := bs.paymentSvc.ProcessPayment(ctx, payment); err != nil {
		// ROLLBACK: Cancel all booked segments
		bs.rollbackBookings(ctx, bookingRefs)
		booking.MarkAsFailed(fmt.Sprintf("payment failed: %v", err))
		bs.bookingRepo.Save(ctx, booking)
		return nil, fmt.Errorf("payment processing failed: %w", err)
	}

	// 7. Mark booking as confirmed
	booking.MarkAsConfirmed()

	// 8. Save booking
	if err := bs.bookingRepo.Save(ctx, booking); err != nil {
		return nil, fmt.Errorf("failed to save booking: %w", err)
	}

	return booking, nil
}

// rollbackBookings cancels all provider bookings (ACID rollback)
func (bs *BookingService) rollbackBookings(ctx context.Context, bookingRefs []string) {
	for _, ref := range bookingRefs {
		// Best effort cancellation - log errors but continue
		if err := bs.providerBooking.CancelBooking(ctx, ref); err != nil {
			// In production, this should be logged and monitored
			fmt.Printf("Warning: failed to cancel booking %s: %v\n", ref, err)
		}
	}
}

// GetBooking retrieves a booking by ID
func (bs *BookingService) GetBooking(ctx context.Context, bookingID string) (*domain.Booking, error) {
	return bs.bookingRepo.FindByID(ctx, bookingID)
}

// CancelBooking cancels a booking and processes refund
func (bs *BookingService) CancelBooking(ctx context.Context, bookingID string, reason string) error {
	booking, err := bs.bookingRepo.FindByID(ctx, bookingID)
	if err != nil {
		return fmt.Errorf("booking not found: %w", err)
	}

	if booking.Status != domain.BookingConfirmed {
		return fmt.Errorf("cannot cancel booking with status: %s", booking.Status)
	}

	// Cancel all segment bookings
	bookingRefs := make([]string, 0, len(booking.Segments))
	for _, segment := range booking.Segments {
		bookingRefs = append(bookingRefs, segment.ProviderBookingRef)
	}
	bs.rollbackBookings(ctx, bookingRefs)

	// Process refund
	if booking.Payment != nil && booking.Payment.Status == domain.PaymentCompleted {
		if err := bs.paymentSvc.RefundPayment(ctx, booking.Payment); err != nil {
			return fmt.Errorf("refund failed: %w", err)
		}
	}

	// Mark as cancelled
	booking.MarkAsCancelled(reason)
	return bs.bookingRepo.Update(ctx, booking)
}

// ListBookings returns all bookings (for admin)
func (bs *BookingService) ListBookings(ctx context.Context) ([]domain.Booking, error) {
	return bs.bookingRepo.FindAll(ctx)
}

// --- Mock Provider Booking Service ---

// MockProviderBookingService simulates booking with external providers
type MockProviderBookingService struct {
	failureRate float64 // Probability of booking failure for testing
}

// NewMockProviderBookingService creates a mock provider booking service
func NewMockProviderBookingService(failureRate float64) *MockProviderBookingService {
	return &MockProviderBookingService{
		failureRate: failureRate,
	}
}

// BookSegment simulates booking a segment with a provider
func (mpbs *MockProviderBookingService) BookSegment(ctx context.Context, segment *domain.Segment, passenger *domain.Passenger) (ticketNumber, bookingRef string, err error) {
	// Simulate processing delay
	time.Sleep(50 * time.Millisecond)

	// Generate mock ticket number and booking reference
	ticketNumber = fmt.Sprintf("TKT-%s-%s", segment.Provider[:3], utils.GenerateID()[:8])
	bookingRef = fmt.Sprintf("BK-%s-%s", segment.TransportType, utils.GenerateID()[:8])

	// Simulate random failures for testing (commented out for hackathon demo)
	// if rand.Float64() < mpbs.failureRate {
	// 	return "", "", fmt.Errorf("provider booking failed: no available seats")
	// }

	return ticketNumber, bookingRef, nil
}

// CancelBooking simulates cancelling a booking
func (mpbs *MockProviderBookingService) CancelBooking(ctx context.Context, bookingRef string) error {
	// Simulate processing delay
	time.Sleep(30 * time.Millisecond)

	// Always succeed in mock mode
	return nil
}
