package domain

import "time"

// BookingStatus represents the status of a booking
type BookingStatus string

const (
	BookingPending        BookingStatus = "pending"         // Awaiting payment
	BookingPendingPayment BookingStatus = "pending_payment" // Segments booked, awaiting payment confirmation
	BookingInProgress     BookingStatus = "in_progress"     // Confirmed and active (journey ongoing or upcoming)
	BookingCompleted      BookingStatus = "completed"       // Journey finished
	BookingCancelled      BookingStatus = "cancelled"       // User cancelled or failed
	BookingFailed         BookingStatus = "failed"          // Booking failed
	BookingRefunded       BookingStatus = "refunded"        // Payment refunded
)

// PaymentStatus represents the status of a payment
type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "pending"
	PaymentCompleted PaymentStatus = "completed"
	PaymentFailed    PaymentStatus = "failed"
	PaymentRefunded  PaymentStatus = "refunded"
)

// PaymentMethod represents payment method
type PaymentMethod string

const (
	PaymentCard       PaymentMethod = "card"
	PaymentYooKassa   PaymentMethod = "yookassa"
	PaymentCloudPay   PaymentMethod = "cloudpay"
	PaymentSberPay    PaymentMethod = "sberpay"
)

// Tariff represents tariff type
type Tariff string

const (
	Tariff1 Tariff = "tarif1" // 0 ₽
	Tariff2 Tariff = "tarif2" // 2244 ₽
	Tariff3 Tariff = "tarif3" // 4244 ₽
	Tariff4 Tariff = "tarif4" // 6244 ₽
)

// TariffPrices maps tariff types to their prices
var TariffPrices = map[Tariff]float64{
	Tariff1: 0,
	Tariff2: 2244,
	Tariff3: 4244,
	Tariff4: 6244,
}

// SeatType represents seat selection type for air segments
type SeatType string

const (
	SeatRandom       SeatType = "random"        // 0 ₽ (random selection)
	SeatWindow       SeatType = "window"        // 3000 ₽
	SeatAisle        SeatType = "aisle"         // 2300 ₽
	SeatExtraLegroom SeatType = "extra_legroom" // 7900 ₽
)

// SeatPrices maps seat types to their prices
var SeatPrices = map[SeatType]float64{
	SeatRandom:       0,
	SeatWindow:       3000,
	SeatAisle:        2300,
	SeatExtraLegroom: 7900,
}

// Passenger represents a passenger
type Passenger struct {
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	MiddleName     string    `json:"middle_name,omitempty"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	PassportNumber string    `json:"passport_number"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
}

// BookedSegment represents a single booked segment in a multi-segment journey
type BookedSegment struct {
	ID                 string        `json:"id"`
	SegmentID          string        `json:"segment_id"`            // Reference to original segment
	Provider           string        `json:"provider"`              // Provider who issued ticket
	TransportType      TransportType `json:"transport_type"`
	From               Stop          `json:"from"`
	To                 Stop          `json:"to"`
	DepartureTime      time.Time     `json:"departure_time"`
	ArrivalTime        time.Time     `json:"arrival_time"`
	TicketNumber       string        `json:"ticket_number"`         // Ticket issued by provider
	Price              float64       `json:"price"`                 // Provider's price
	Commission         float64       `json:"commission"`            // Our markup
	TotalPrice         float64       `json:"total_price"`           // price + commission
	BookingStatus      BookingStatus `json:"booking_status"`
	ProviderBookingRef string        `json:"provider_booking_ref"`  // Provider's booking reference
	SeatType           *SeatType     `json:"seat_type,omitempty"`   // Seat selection (only for air)
	SeatPrice          float64       `json:"seat_price,omitempty"`  // Seat price (0 for non-air or random)
}

// Payment represents a payment transaction
type Payment struct {
	ID                string        `json:"id"`
	OrderID           string        `json:"order_id"`
	Amount            float64       `json:"amount"`
	Currency          string        `json:"currency"`
	Method            PaymentMethod `json:"method"`
	Status            PaymentStatus `json:"status"`
	ProviderPaymentID string        `json:"provider_payment_id,omitempty"` // Payment gateway transaction ID
	ConfirmationURL   string        `json:"confirmation_url,omitempty"`    // URL for redirect to payment provider (YooKassa)
	CreatedAt         time.Time     `json:"created_at"`
	CompletedAt       *time.Time    `json:"completed_at,omitempty"`
	FailureReason     string        `json:"failure_reason,omitempty"`
}

// Booking represents a complete multi-segment booking
type Booking struct {
	ID                 string          `json:"id"` // Order ID
	UserID             string          `json:"user_id,omitempty"`
	RouteID            string          `json:"route_id"`
	Passenger          Passenger       `json:"passenger"`
	Segments           []BookedSegment `json:"segments"`
	TotalPrice         float64         `json:"total_price"`         // Sum of all segment prices
	TotalSeatsPrice    float64         `json:"total_seats_price"`   // Sum of all seat prices
	Tariff             Tariff          `json:"tariff"`              // Selected tariff
	TariffPrice        float64         `json:"tariff_price"`        // Tariff price
	GrandTotal         float64         `json:"grand_total"`         // totalPrice + tariff + seats + insurance
	InsurancePremium   float64         `json:"insurance_premium,omitempty"`
	IncludeInsurance   bool            `json:"include_insurance"`
	Status             BookingStatus   `json:"status"`
	Payment            *Payment        `json:"payment,omitempty"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
	ConfirmedAt        *time.Time      `json:"confirmed_at,omitempty"`
	CancelledAt        *time.Time      `json:"cancelled_at,omitempty"`
	CancellationReason string          `json:"cancellation_reason,omitempty"`
}

// AddSegment adds a booked segment to the booking
func (b *Booking) AddSegment(segment BookedSegment) {
	b.Segments = append(b.Segments, segment)
	// Use segment.TotalPrice (base price + commission) so booking total already includes commission
	b.TotalPrice += segment.TotalPrice
	b.TotalSeatsPrice += segment.SeatPrice
	b.RecalculateGrandTotal()
}

// RecalculateGrandTotal recalculates the grand total
func (b *Booking) RecalculateGrandTotal() {
	b.GrandTotal = b.TotalPrice + b.TariffPrice + b.TotalSeatsPrice
	if b.IncludeInsurance {
		b.GrandTotal += b.InsurancePremium
	}
}

// MarkAsInProgress marks booking as in progress (confirmed and active)
func (b *Booking) MarkAsInProgress() {
	b.Status = BookingInProgress
	now := time.Now()
	b.ConfirmedAt = &now
	b.UpdatedAt = now
}

// MarkAsConfirmed marks booking as confirmed (alias for MarkAsInProgress)
func (b *Booking) MarkAsConfirmed() {
	b.MarkAsInProgress()
}

// MarkAsCompleted marks booking as completed
func (b *Booking) MarkAsCompleted() {
	b.Status = BookingCompleted
	b.UpdatedAt = time.Now()
}

// MarkAsCancelled marks booking as cancelled
func (b *Booking) MarkAsCancelled(reason string) {
	b.Status = BookingCancelled
	now := time.Now()
	b.CancelledAt = &now
	b.CancellationReason = reason
	b.UpdatedAt = now
}

// MarkAsFailed marks booking as failed
func (b *Booking) MarkAsFailed(reason string) {
	b.Status = BookingFailed
	now := time.Now()
	b.CancelledAt = &now
	b.CancellationReason = reason
	b.UpdatedAt = now
}

// AllSegmentsBooked checks if all segments are successfully booked
func (b *Booking) AllSegmentsBooked() bool {
	for _, segment := range b.Segments {
		if segment.BookingStatus != BookingInProgress {
			return false
		}
	}
	return len(b.Segments) > 0
}
