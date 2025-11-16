package domain

import "time"

// BookingStatus represents the status of a booking
type BookingStatus string

const (
	BookingPending   BookingStatus = "pending"    // Awaiting payment
	BookingConfirmed BookingStatus = "confirmed"  // Payment successful, all segments booked
	BookingFailed    BookingStatus = "failed"     // Booking or payment failed
	BookingCancelled BookingStatus = "cancelled"  // User cancelled
	BookingRefunded  BookingStatus = "refunded"   // Refund processed
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
	ID              string        `json:"id"`
	SegmentID       string        `json:"segment_id"`       // Reference to original segment
	Provider        string        `json:"provider"`         // Provider who issued ticket
	TransportType   TransportType `json:"transport_type"`
	From            Stop          `json:"from"`
	To              Stop          `json:"to"`
	DepartureTime   time.Time     `json:"departure_time"`
	ArrivalTime     time.Time     `json:"arrival_time"`
	TicketNumber    string        `json:"ticket_number"`    // Ticket issued by provider
	Price           float64       `json:"price"`            // Provider's price
	Commission      float64       `json:"commission"`       // Our markup
	TotalPrice      float64       `json:"total_price"`      // price + commission
	BookingStatus   BookingStatus `json:"booking_status"`
	ProviderBookingRef string     `json:"provider_booking_ref"` // Provider's booking reference
}

// Payment represents a payment transaction
type Payment struct {
	ID              string        `json:"id"`
	OrderID         string        `json:"order_id"`
	Amount          float64       `json:"amount"`
	Currency        string        `json:"currency"`
	Method          PaymentMethod `json:"method"`
	Status          PaymentStatus `json:"status"`
	ProviderPaymentID string      `json:"provider_payment_id"` // Payment gateway transaction ID
	CreatedAt       time.Time     `json:"created_at"`
	CompletedAt     *time.Time    `json:"completed_at,omitempty"`
	FailureReason   string        `json:"failure_reason,omitempty"`
}

// Booking represents a complete multi-segment booking
type Booking struct {
	ID                string          `json:"id"` // Order ID
	RouteID           string          `json:"route_id"`
	Passenger         Passenger       `json:"passenger"`
	Segments          []BookedSegment `json:"segments"`
	TotalPrice        float64         `json:"total_price"`        // Sum of all segment prices
	TotalCommission   float64         `json:"total_commission"`   // Sum of all commissions
	GrandTotal        float64         `json:"grand_total"`        // totalPrice + totalCommission
	InsurancePremium  float64         `json:"insurance_premium,omitempty"`
	IncludeInsurance  bool            `json:"include_insurance"`
	Status            BookingStatus   `json:"status"`
	Payment           *Payment        `json:"payment,omitempty"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	ConfirmedAt       *time.Time      `json:"confirmed_at,omitempty"`
	CancelledAt       *time.Time      `json:"cancelled_at,omitempty"`
	CancellationReason string         `json:"cancellation_reason,omitempty"`
}

// AddSegment adds a booked segment to the booking
func (b *Booking) AddSegment(segment BookedSegment) {
	b.Segments = append(b.Segments, segment)
	b.TotalPrice += segment.Price
	b.TotalCommission += segment.Commission
	b.GrandTotal = b.TotalPrice + b.TotalCommission
	if b.IncludeInsurance {
		b.GrandTotal += b.InsurancePremium
	}
}

// MarkAsConfirmed marks booking as confirmed
func (b *Booking) MarkAsConfirmed() {
	b.Status = BookingConfirmed
	now := time.Now()
	b.ConfirmedAt = &now
	b.UpdatedAt = now
}

// MarkAsFailed marks booking as failed
func (b *Booking) MarkAsFailed(reason string) {
	b.Status = BookingFailed
	b.CancellationReason = reason
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

// AllSegmentsBooked checks if all segments are successfully booked
func (b *Booking) AllSegmentsBooked() bool {
	for _, segment := range b.Segments {
		if segment.BookingStatus != BookingConfirmed {
			return false
		}
	}
	return len(b.Segments) > 0
}
