package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/internal/repository"
)

// BookingRepository implements repository.BookingRepository interface for PostgreSQL
type BookingRepository struct {
	db *Database
}

// NewBookingRepository creates a new booking repository
func NewBookingRepository(db *Database) repository.BookingRepository {
	return &BookingRepository{db: db}
}

// FindByID retrieves a booking with all details
func (r *BookingRepository) FindByID(ctx context.Context, id string) (*domain.Booking, error) {
	const query = `
		SELECT id, route_id, status, total_price, total_commission, grand_total,
		       insurance_premium, include_insurance, created_at, updated_at,
		       confirmed_at, cancelled_at, cancellation_reason,
		       passenger_first_name, passenger_last_name, passenger_middle_name,
		       passenger_date_of_birth, passenger_passport_number,
		       passenger_email, passenger_phone
		FROM bookings
		WHERE id = $1
	`

	var booking domain.Booking
	var confirmedAt, cancelledAt sql.NullTime
	var middleName, cancellationReason sql.NullString

	err := r.db.db.QueryRowContext(ctx, query, id).Scan(
		&booking.ID,
		&booking.RouteID,
		&booking.Status,
		&booking.TotalPrice,
		&booking.TotalCommission,
		&booking.GrandTotal,
		&booking.InsurancePremium,
		&booking.IncludeInsurance,
		&booking.CreatedAt,
		&booking.UpdatedAt,
		&confirmedAt,
		&cancelledAt,
		&cancellationReason,
		&booking.Passenger.FirstName,
		&booking.Passenger.LastName,
		&middleName,
		&booking.Passenger.DateOfBirth,
		&booking.Passenger.PassportNumber,
		&booking.Passenger.Email,
		&booking.Passenger.Phone,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrBookingNotFound
		}
		return nil, fmt.Errorf("error querying booking: %w", err)
	}

	// Set nullable fields
	if middleName.Valid {
		booking.Passenger.MiddleName = middleName.String
	}
	if confirmedAt.Valid {
		booking.ConfirmedAt = &confirmedAt.Time
	}
	if cancelledAt.Valid {
		booking.CancelledAt = &cancelledAt.Time
	}
	if cancellationReason.Valid {
		booking.CancellationReason = cancellationReason.String
	}

	// Fetch booked segments
	if err := r.fetchBookedSegments(ctx, &booking); err != nil {
		return nil, err
	}

	// Fetch payment if exists
	if err := r.fetchPayment(ctx, &booking); err != nil {
		return nil, err
	}

	return &booking, nil
}

// FindAll retrieves all bookings
func (r *BookingRepository) FindAll(ctx context.Context) ([]domain.Booking, error) {
	const query = `
		SELECT id, route_id, status, total_price, total_commission, grand_total,
		       insurance_premium, include_insurance, created_at, updated_at,
		       confirmed_at, cancelled_at, cancellation_reason,
		       passenger_first_name, passenger_last_name, passenger_middle_name,
		       passenger_date_of_birth, passenger_passport_number,
		       passenger_email, passenger_phone
		FROM bookings
		ORDER BY created_at DESC
		LIMIT 100
	`

	rows, err := r.db.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying bookings: %w", err)
	}
	defer rows.Close()

	var bookings []domain.Booking
	for rows.Next() {
		var booking domain.Booking
		var confirmedAt, cancelledAt sql.NullTime
		var middleName, cancellationReason sql.NullString

		if err := rows.Scan(
			&booking.ID,
			&booking.RouteID,
			&booking.Status,
			&booking.TotalPrice,
			&booking.TotalCommission,
			&booking.GrandTotal,
			&booking.InsurancePremium,
			&booking.IncludeInsurance,
			&booking.CreatedAt,
			&booking.UpdatedAt,
			&confirmedAt,
			&cancelledAt,
			&cancellationReason,
			&booking.Passenger.FirstName,
			&booking.Passenger.LastName,
			&middleName,
			&booking.Passenger.DateOfBirth,
			&booking.Passenger.PassportNumber,
			&booking.Passenger.Email,
			&booking.Passenger.Phone,
		); err != nil {
			return nil, fmt.Errorf("error scanning booking: %w", err)
		}

		if middleName.Valid {
			booking.Passenger.MiddleName = middleName.String
		}
		if confirmedAt.Valid {
			booking.ConfirmedAt = &confirmedAt.Time
		}
		if cancelledAt.Valid {
			booking.CancelledAt = &cancelledAt.Time
		}
		if cancellationReason.Valid {
			booking.CancellationReason = cancellationReason.String
		}

		bookings = append(bookings, booking)
	}

	return bookings, rows.Err()
}

// FindByPassenger finds bookings by passenger email
func (r *BookingRepository) FindByPassenger(ctx context.Context, email string) ([]domain.Booking, error) {
	const query = `
		SELECT id, route_id, status, total_price, total_commission, grand_total,
		       insurance_premium, include_insurance, created_at, updated_at,
		       confirmed_at, cancelled_at, cancellation_reason,
		       passenger_first_name, passenger_last_name, passenger_middle_name,
		       passenger_date_of_birth, passenger_passport_number,
		       passenger_email, passenger_phone
		FROM bookings
		WHERE LOWER(passenger_email) = LOWER($1)
		ORDER BY created_at DESC
	`

	rows, err := r.db.db.QueryContext(ctx, query, email)
	if err != nil {
		return nil, fmt.Errorf("error querying bookings by passenger: %w", err)
	}
	defer rows.Close()

	var bookings []domain.Booking
	for rows.Next() {
		var booking domain.Booking
		var confirmedAt, cancelledAt sql.NullTime
		var middleName, cancellationReason sql.NullString

		if err := rows.Scan(
			&booking.ID,
			&booking.RouteID,
			&booking.Status,
			&booking.TotalPrice,
			&booking.TotalCommission,
			&booking.GrandTotal,
			&booking.InsurancePremium,
			&booking.IncludeInsurance,
			&booking.CreatedAt,
			&booking.UpdatedAt,
			&confirmedAt,
			&cancelledAt,
			&cancellationReason,
			&booking.Passenger.FirstName,
			&booking.Passenger.LastName,
			&middleName,
			&booking.Passenger.DateOfBirth,
			&booking.Passenger.PassportNumber,
			&booking.Passenger.Email,
			&booking.Passenger.Phone,
		); err != nil {
			return nil, fmt.Errorf("error scanning booking: %w", err)
		}

		if middleName.Valid {
			booking.Passenger.MiddleName = middleName.String
		}
		if confirmedAt.Valid {
			booking.ConfirmedAt = &confirmedAt.Time
		}
		if cancelledAt.Valid {
			booking.CancelledAt = &cancelledAt.Time
		}
		if cancellationReason.Valid {
			booking.CancellationReason = cancellationReason.String
		}

		bookings = append(bookings, booking)
	}

	return bookings, rows.Err()
}

// FindByStatus finds bookings by status
func (r *BookingRepository) FindByStatus(ctx context.Context, status domain.BookingStatus) ([]domain.Booking, error) {
	const query = `
		SELECT id, route_id, status, total_price, total_commission, grand_total,
		       insurance_premium, include_insurance, created_at, updated_at,
		       confirmed_at, cancelled_at, cancellation_reason,
		       passenger_first_name, passenger_last_name, passenger_middle_name,
		       passenger_date_of_birth, passenger_passport_number,
		       passenger_email, passenger_phone
		FROM bookings
		WHERE status = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.db.QueryContext(ctx, query, string(status))
	if err != nil {
		return nil, fmt.Errorf("error querying bookings by status: %w", err)
	}
	defer rows.Close()

	var bookings []domain.Booking
	for rows.Next() {
		var booking domain.Booking
		var confirmedAt, cancelledAt sql.NullTime
		var middleName, cancellationReason sql.NullString

		if err := rows.Scan(
			&booking.ID,
			&booking.RouteID,
			&booking.Status,
			&booking.TotalPrice,
			&booking.TotalCommission,
			&booking.GrandTotal,
			&booking.InsurancePremium,
			&booking.IncludeInsurance,
			&booking.CreatedAt,
			&booking.UpdatedAt,
			&confirmedAt,
			&cancelledAt,
			&cancellationReason,
			&booking.Passenger.FirstName,
			&booking.Passenger.LastName,
			&middleName,
			&booking.Passenger.DateOfBirth,
			&booking.Passenger.PassportNumber,
			&booking.Passenger.Email,
			&booking.Passenger.Phone,
		); err != nil {
			return nil, fmt.Errorf("error scanning booking: %w", err)
		}

		if middleName.Valid {
			booking.Passenger.MiddleName = middleName.String
		}
		if confirmedAt.Valid {
			booking.ConfirmedAt = &confirmedAt.Time
		}
		if cancelledAt.Valid {
			booking.CancelledAt = &cancelledAt.Time
		}
		if cancellationReason.Valid {
			booking.CancellationReason = cancellationReason.String
		}

		bookings = append(bookings, booking)
	}

	return bookings, rows.Err()
}

// Save stores a new booking
func (r *BookingRepository) Save(ctx context.Context, booking *domain.Booking) error {
	const query = `
		INSERT INTO bookings (
			id, route_id, status, total_price, total_commission, grand_total,
			insurance_premium, include_insurance, created_at, updated_at,
			confirmed_at, cancelled_at, cancellation_reason,
			passenger_first_name, passenger_last_name, passenger_middle_name,
			passenger_date_of_birth, passenger_passport_number,
			passenger_email, passenger_phone
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15, $16, $17, $18, $19, $20
		)
	`

	_, err := r.db.db.ExecContext(ctx, query,
		booking.ID,
		booking.RouteID,
		string(booking.Status),
		booking.TotalPrice,
		booking.TotalCommission,
		booking.GrandTotal,
		booking.InsurancePremium,
		booking.IncludeInsurance,
		booking.CreatedAt,
		booking.UpdatedAt,
		booking.ConfirmedAt,
		booking.CancelledAt,
		booking.CancellationReason,
		booking.Passenger.FirstName,
		booking.Passenger.LastName,
		booking.Passenger.MiddleName,
		booking.Passenger.DateOfBirth,
		booking.Passenger.PassportNumber,
		booking.Passenger.Email,
		booking.Passenger.Phone,
	)

	if err != nil {
		return fmt.Errorf("error saving booking: %w", err)
	}

	// Save booked segments
	if err := r.saveBookedSegments(ctx, booking); err != nil {
		return err
	}

	// Save payment if exists
	if booking.Payment != nil {
		if err := r.savePayment(ctx, booking.Payment); err != nil {
			return err
		}
	}

	return nil
}

// Update modifies an existing booking
func (r *BookingRepository) Update(ctx context.Context, booking *domain.Booking) error {
	const query = `
		UPDATE bookings
		SET status = $2, updated_at = $3, confirmed_at = $4,
		    cancelled_at = $5, cancellation_reason = $6
		WHERE id = $1
	`

	result, err := r.db.db.ExecContext(ctx, query,
		booking.ID,
		string(booking.Status),
		time.Now(),
		booking.ConfirmedAt,
		booking.CancelledAt,
		booking.CancellationReason,
	)

	if err != nil {
		return fmt.Errorf("error updating booking: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rows == 0 {
		return domain.ErrBookingNotFound
	}

	// Update payment if exists
	if booking.Payment != nil {
		if err := r.updatePayment(ctx, booking.Payment); err != nil {
			return err
		}
	}

	return nil
}

// Delete removes a booking
func (r *BookingRepository) Delete(ctx context.Context, id string) error {
	const query = `DELETE FROM bookings WHERE id = $1`

	result, err := r.db.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting booking: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rows == 0 {
		return domain.ErrBookingNotFound
	}

	return nil
}

// Helper functions

func (r *BookingRepository) fetchBookedSegments(ctx context.Context, booking *domain.Booking) error {
	const query = `
		SELECT bs.id, bs.segment_id, bs.provider, bs.transport_type,
		       bs.from_stop_id, fs.name, fs.city, fs.latitude, fs.longitude,
		       bs.to_stop_id, ts.name, ts.city, ts.latitude, ts.longitude,
		       bs.departure_time, bs.arrival_time,
		       bs.ticket_number, bs.price, bs.commission, bs.total_price,
		       bs.booking_status, bs.provider_booking_ref
		FROM booked_segments bs
		JOIN stops fs ON bs.from_stop_id = fs.id
		JOIN stops ts ON bs.to_stop_id = ts.id
		WHERE bs.booking_id = $1
		ORDER BY bs.sequence_order
	`

	rows, err := r.db.db.QueryContext(ctx, query, booking.ID)
	if err != nil {
		return fmt.Errorf("error querying booked segments: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var segment domain.BookedSegment
		var ticketNumber, providerRef sql.NullString

		if err := rows.Scan(
			&segment.ID,
			&segment.SegmentID,
			&segment.Provider,
			&segment.TransportType,
			&segment.From.ID,
			&segment.From.Name,
			&segment.From.City,
			&segment.From.Latitude,
			&segment.From.Longitude,
			&segment.To.ID,
			&segment.To.Name,
			&segment.To.City,
			&segment.To.Latitude,
			&segment.To.Longitude,
			&segment.DepartureTime,
			&segment.ArrivalTime,
			&ticketNumber,
			&segment.Price,
			&segment.Commission,
			&segment.TotalPrice,
			&segment.BookingStatus,
			&providerRef,
		); err != nil {
			return fmt.Errorf("error scanning booked segment: %w", err)
		}

		if ticketNumber.Valid {
			segment.TicketNumber = ticketNumber.String
		}
		if providerRef.Valid {
			segment.ProviderBookingRef = providerRef.String
		}

		booking.Segments = append(booking.Segments, segment)
	}

	return rows.Err()
}

func (r *BookingRepository) fetchPayment(ctx context.Context, booking *domain.Booking) error {
	const query = `
		SELECT id, order_id, amount, currency, method, status,
		       provider_payment_id, confirmation_url, created_at, completed_at, failure_reason
		FROM payments
		WHERE order_id = $1
	`

	var payment domain.Payment
	var completedAt sql.NullTime
	var providerPaymentID, confirmationURL, failureReason sql.NullString

	err := r.db.db.QueryRowContext(ctx, query, booking.ID).Scan(
		&payment.ID,
		&payment.OrderID,
		&payment.Amount,
		&payment.Currency,
		&payment.Method,
		&payment.Status,
		&providerPaymentID,
		&confirmationURL,
		&payment.CreatedAt,
		&completedAt,
		&failureReason,
	)

	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error querying payment: %w", err)
	}

	if err == nil {
		if completedAt.Valid {
			payment.CompletedAt = &completedAt.Time
		}
		if providerPaymentID.Valid {
			payment.ProviderPaymentID = providerPaymentID.String
		}
		if confirmationURL.Valid {
			payment.ConfirmationURL = confirmationURL.String
		}
		if failureReason.Valid {
			payment.FailureReason = failureReason.String
		}
		booking.Payment = &payment
	}

	return nil
}

func (r *BookingRepository) saveBookedSegments(ctx context.Context, booking *domain.Booking) error {
	const query = `
		INSERT INTO booked_segments (
			id, booking_id, segment_id, provider, transport_type,
			from_stop_id, to_stop_id, departure_time, arrival_time,
			ticket_number, price, commission, total_price,
			booking_status, provider_booking_ref, sequence_order
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
		)
	`

	for i, segment := range booking.Segments {
		_, err := r.db.db.ExecContext(ctx, query,
			segment.ID,
			booking.ID,
			segment.SegmentID,
			segment.Provider,
			string(segment.TransportType),
			segment.From.ID,
			segment.To.ID,
			segment.DepartureTime,
			segment.ArrivalTime,
			segment.TicketNumber,
			segment.Price,
			segment.Commission,
			segment.TotalPrice,
			string(segment.BookingStatus),
			segment.ProviderBookingRef,
			i+1,
		)
		if err != nil {
			return fmt.Errorf("error saving booked segment: %w", err)
		}
	}

	return nil
}

func (r *BookingRepository) savePayment(ctx context.Context, payment *domain.Payment) error {
	const query = `
		INSERT INTO payments (
			id, order_id, amount, currency, method, status,
			provider_payment_id, confirmation_url, created_at, completed_at, failure_reason
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
		)
	`

	_, err := r.db.db.ExecContext(ctx, query,
		payment.ID,
		payment.OrderID,
		payment.Amount,
		payment.Currency,
		string(payment.Method),
		string(payment.Status),
		payment.ProviderPaymentID,
		payment.ConfirmationURL,
		payment.CreatedAt,
		payment.CompletedAt,
		payment.FailureReason,
	)

	if err != nil {
		return fmt.Errorf("error saving payment: %w", err)
	}

	return nil
}

func (r *BookingRepository) updatePayment(ctx context.Context, payment *domain.Payment) error {
	const query = `
		UPDATE payments
		SET status = $2, provider_payment_id = $3, completed_at = $4, failure_reason = $5
		WHERE id = $1
	`

	_, err := r.db.db.ExecContext(ctx, query,
		payment.ID,
		string(payment.Status),
		payment.ProviderPaymentID,
		payment.CompletedAt,
		payment.FailureReason,
	)

	if err != nil {
		return fmt.Errorf("error updating payment: %w", err)
	}

	return nil
}
