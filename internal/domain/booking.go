package domain

import "time"

// PassengerType defines the type of passenger
type PassengerType string

const (
	PassengerAdult   PassengerType = "adult"
	PassengerChild   PassengerType = "child"
	PassengerInfant  PassengerType = "infant"
)

// BookingStatus represents booking status
type BookingStatus string

const (
	BookingStatusPending     BookingStatus = "pending"
	BookingStatusConfirmed   BookingStatus = "confirmed"
	BookingStatusFailed      BookingStatus = "failed"
	BookingStatusCancelled   BookingStatus = "cancelled"
)

// Passenger represents a person traveling
type Passenger struct {
	ID             string        `json:"id"`
	FirstName      string        `json:"first_name"`
	LastName       string        `json:"last_name"`
	MiddleName     string        `json:"middle_name,omitempty"`
	DateOfBirth    time.Time     `json:"date_of_birth"`
	PassportNumber string        `json:"passport_number"`
	PassportSeries string        `json:"passport_series"`
	Nationality    string        `json:"nationality"`
	Type           PassengerType `json:"type"`
	Email          string        `json:"email,omitempty"`
	Phone          string        `json:"phone,omitempty"`
}

// BookingSegmentTicket represents a ticket for one segment
type BookingSegmentTicket struct {
	ID              string        `json:"id"`
	BookingID       string        `json:"booking_id"`
	SegmentID       string        `json:"segment_id"`
	PassengerID     string        `json:"passenger_id"`
	Provider        string        `json:"provider"`
	TicketNumber    string        `json:"ticket_number"`
	SeatNumber      string        `json:"seat_number"`
	Price           float64       `json:"price"`
	Status          string        `json:"status"`
	CreatedAt       time.Time     `json:"created_at"`
}

// BookingGuarantee contains guarantee and support information
type BookingGuarantee struct {
	GuaranteeID    string    `json:"guarantee_id"`
	BookingID      string    `json:"booking_id"`
	ValidUntil     time.Time `json:"valid_until"`
	RefundPolicy   string    `json:"refund_policy"`
	SOSHotline     string    `json:"sos_hotline"`
	InsuranceLevel string    `json:"insurance_level"`
	CoverageAmount float64   `json:"coverage_amount"`
	CreatedAt      time.Time `json:"created_at"`
}

// Booking represents a complete booking
type Booking struct {
	ID                   string                   `json:"id"`
	RouteID              string                   `json:"route_id"`
	Passengers           []Passenger              `json:"passengers"`
	Tickets              []BookingSegmentTicket   `json:"tickets"`
	TotalPrice           float64                  `json:"total_price"`
	InsurancePremium     float64                  `json:"insurance_premium"`
	InsuranceIncluded    bool                     `json:"insurance_included"`
	Guarantee            *BookingGuarantee        `json:"guarantee,omitempty"`
	Status               BookingStatus            `json:"status"`
	BookedAt             time.Time                `json:"booked_at"`
	PaymentDeadline      time.Time                `json:"payment_deadline"`
	CancellationDeadline time.Time                `json:"cancellation_deadline"`
	Notes                string                   `json:"notes,omitempty"`
}

// BookingCreateRequest represents request to create booking
type BookingCreateRequest struct {
	RouteID          string      `json:"route_id"`
	Passengers       []Passenger `json:"passengers"`
	IncludeInsurance bool        `json:"include_insurance"`
}

// BookingCreateResult represents result of booking creation
type BookingCreateResult struct {
	Booking *Booking
	Error   error
}
