package dto

import "time"

// SeatSelectionRequest represents seat selection for an air segment
type SeatSelectionRequest struct {
	SegmentIndex int    `json:"segment_index" validate:"required,gte=0"` // 0-based index of segment in route
	SeatType     string `json:"seat_type" validate:"required,oneof=random window aisle extra_legroom"`
}

// CreateBookingRequest represents a request to create a booking
type CreateBookingRequest struct {
	RouteID          string                 `json:"route_id" validate:"required"`
	Passenger        PassengerRequest       `json:"passenger" validate:"required"`
	IncludeInsurance bool                   `json:"include_insurance"`
	PaymentMethod    string                 `json:"payment_method" validate:"required,oneof=card yookassa cloudpay sberpay"`
	Tariff           string                 `json:"tariff" validate:"required,oneof=tarif1 tarif2 tarif3 tarif4"`
	SeatSelections   []SeatSelectionRequest `json:"seat_selections,omitempty"` // Only for air segments
}

// PassengerRequest represents passenger information
type PassengerRequest struct {
	FirstName      string `json:"first_name" validate:"required"`
	LastName       string `json:"last_name" validate:"required"`
	MiddleName     string `json:"middle_name"`
	DateOfBirth    string `json:"date_of_birth" validate:"required"` // YYYY-MM-DD
	PassportNumber string `json:"passport_number" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Phone          string `json:"phone" validate:"required"`
}

// BookingResponse represents a booking in API response
type BookingResponse struct {
	ID               string                  `json:"id"`
	RouteID          string                  `json:"route_id"`
	Status           string                  `json:"status"` // pending, confirmed, failed, cancelled, refunded
	Passenger        PassengerResponse       `json:"passenger"`
	Segments         []BookedSegmentResponse `json:"segments"`
	TotalPrice       float64                 `json:"total_price"`
	TotalSeatsPrice  float64                 `json:"total_seats_price"`
	Tariff           string                  `json:"tariff"`
	TariffPrice      float64                 `json:"tariff_price"`
	InsurancePremium float64                 `json:"insurance_premium,omitempty"`
	GrandTotal       float64                 `json:"grand_total"`
	IncludeInsurance bool                    `json:"include_insurance"`
	Payment          *PaymentResponse        `json:"payment,omitempty"`
	CreatedAt        time.Time               `json:"created_at"`
	ConfirmedAt      *time.Time              `json:"confirmed_at,omitempty"`
	CancelledAt      *time.Time              `json:"cancelled_at,omitempty"`
	CancellationReason string                `json:"cancellation_reason,omitempty"`
}

// PassengerResponse represents passenger information in response
type PassengerResponse struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name,omitempty"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

// BookedSegmentResponse represents a booked segment
type BookedSegmentResponse struct {
	ID                 string       `json:"id"`
	SegmentID          string       `json:"segment_id"`
	Provider           string       `json:"provider"`
	TransportType      string       `json:"transport_type"`
	From               StopResponse `json:"from"`
	To                 StopResponse `json:"to"`
	DepartureTime      time.Time    `json:"departure_time"`
	ArrivalTime        time.Time    `json:"arrival_time"`
	TicketNumber       string       `json:"ticket_number"`
	Price              float64      `json:"price"`
	Commission         float64      `json:"commission"`
	TotalPrice         float64      `json:"total_price"`
	BookingStatus      string       `json:"booking_status"` // confirmed, failed, cancelled
	ProviderBookingRef string       `json:"provider_booking_ref,omitempty"`
	SeatType           string       `json:"seat_type,omitempty"`  // Only for air segments
	SeatPrice          float64      `json:"seat_price,omitempty"` // Price of selected seat
}

// PaymentResponse represents payment information
type PaymentResponse struct {
	ID                string     `json:"id"`
	OrderID           string     `json:"order_id"`
	Amount            float64    `json:"amount"`
	Currency          string     `json:"currency"`
	Method            string     `json:"method"` // card, yookassa, cloudpay, sberpay
	Status            string     `json:"status"` // pending, completed, failed, refunded
	ProviderPaymentID string     `json:"provider_payment_id,omitempty"`
	ConfirmationURL   string     `json:"confirmation_url,omitempty"` // URL for redirect to payment provider (YooKassa)
	CreatedAt         time.Time  `json:"created_at"`
	CompletedAt       *time.Time `json:"completed_at,omitempty"`
	FailureReason     string     `json:"failure_reason,omitempty"`
}

// CancelBookingRequest represents a request to cancel a booking
type CancelBookingRequest struct {
	Reason string `json:"reason" validate:"required"`
}

// BookingListResponse represents a list of bookings
type BookingListResponse struct {
	Bookings []BookingSummaryResponse `json:"bookings"`
	Total    int                      `json:"total"`
}

// BookingSummaryResponse represents a summary of a booking (for lists)
type BookingSummaryResponse struct {
	ID             string    `json:"id"`
	RouteID        string    `json:"route_id"`
	Status         string    `json:"status"`
	PassengerEmail string    `json:"passenger_email"`
	GrandTotal     float64   `json:"grand_total"`
	CreatedAt      time.Time `json:"created_at"`
	ConfirmedAt    *time.Time `json:"confirmed_at,omitempty"`
}
