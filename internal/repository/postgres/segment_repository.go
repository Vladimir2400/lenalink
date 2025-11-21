package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/internal/repository"
)

// SegmentRepository implements repository.SegmentRepository interface for PostgreSQL
type SegmentRepository struct {
	db *Database
}

// NewSegmentRepository creates a new segment repository
func NewSegmentRepository(db *Database) repository.SegmentRepository {
	return &SegmentRepository{db: db}
}

// Save stores a new segment
func (r *SegmentRepository) Save(ctx context.Context, segment *domain.Segment) error {
	const query = `
		INSERT INTO segments (
			id, route_id, transport_type, provider,
			start_stop_id, end_stop_id, departure_time, arrival_time,
			price, duration, seat_count, reliability_rate, distance, sequence_order
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	`

	_, err := r.db.db.ExecContext(ctx, query,
		segment.ID,
		nil, // route_id is NULL for standalone segments
		segment.TransportType,
		segment.Provider,
		segment.StartStop.ID,
		segment.EndStop.ID,
		segment.DepartureTime,
		segment.ArrivalTime,
		segment.Price,
		segment.Duration.Nanoseconds(),
		segment.SeatCount,
		segment.ReliabilityRate,
		segment.Distance,
		0, // sequence_order is 0 for standalone segments
	)

	if err != nil {
		return fmt.Errorf("error saving segment: %w", err)
	}

	return nil
}

// BatchSave stores multiple segments in a single transaction
func (r *SegmentRepository) BatchSave(ctx context.Context, segments []domain.Segment) error {
	if len(segments) == 0 {
		return nil
	}

	tx, err := r.db.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error beginning transaction: %w", err)
	}
	defer tx.Rollback()

	const query = `
		INSERT INTO segments (
			id, route_id, transport_type, provider,
			start_stop_id, end_stop_id, departure_time, arrival_time,
			price, duration, seat_count, reliability_rate, distance, sequence_order
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		ON CONFLICT (id) DO UPDATE SET
			departure_time = EXCLUDED.departure_time,
			arrival_time = EXCLUDED.arrival_time,
			price = EXCLUDED.price,
			seat_count = EXCLUDED.seat_count
	`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	for _, segment := range segments {
		_, err := stmt.ExecContext(ctx,
			segment.ID,
			nil, // route_id is NULL for standalone segments
			segment.TransportType,
			segment.Provider,
			segment.StartStop.ID,
			segment.EndStop.ID,
			segment.DepartureTime,
			segment.ArrivalTime,
			segment.Price,
			segment.Duration.Nanoseconds(),
			segment.SeatCount,
			segment.ReliabilityRate,
			segment.Distance,
			0, // sequence_order
		)
		if err != nil {
			return fmt.Errorf("error executing batch insert: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

// FindByID retrieves a segment by ID
func (r *SegmentRepository) FindByID(ctx context.Context, id string) (*domain.Segment, error) {
	const query = `
		SELECT
			s.id, s.transport_type, s.provider,
			s.departure_time, s.arrival_time, s.price, s.duration,
			s.seat_count, s.reliability_rate, s.distance,
			start.id, start.name, start.city, start.latitude, start.longitude,
			end_stop.id, end_stop.name, end_stop.city, end_stop.latitude, end_stop.longitude
		FROM segments s
		JOIN stops start ON s.start_stop_id = start.id
		JOIN stops end_stop ON s.end_stop_id = end_stop.id
		WHERE s.id = $1
	`

	var segment domain.Segment
	var durationNs int64

	err := r.db.db.QueryRowContext(ctx, query, id).Scan(
		&segment.ID,
		&segment.TransportType,
		&segment.Provider,
		&segment.DepartureTime,
		&segment.ArrivalTime,
		&segment.Price,
		&durationNs,
		&segment.SeatCount,
		&segment.ReliabilityRate,
		&segment.Distance,
		&segment.StartStop.ID,
		&segment.StartStop.Name,
		&segment.StartStop.City,
		&segment.StartStop.Latitude,
		&segment.StartStop.Longitude,
		&segment.EndStop.ID,
		&segment.EndStop.Name,
		&segment.EndStop.City,
		&segment.EndStop.Latitude,
		&segment.EndStop.Longitude,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("segment not found: %s", id)
		}
		return nil, fmt.Errorf("error querying segment: %w", err)
	}

	segment.Duration = time.Duration(durationNs)

	return &segment, nil
}

// FindByCriteria searches segments by origin, destination, and date range
func (r *SegmentRepository) FindByCriteria(ctx context.Context, fromCity, toCity string, departureStart, departureEnd time.Time) ([]domain.Segment, error) {
	const query = `
		SELECT
			s.id, s.transport_type, s.provider,
			s.departure_time, s.arrival_time, s.price, s.duration,
			s.seat_count, s.reliability_rate, s.distance,
			start.id, start.name, start.city, start.latitude, start.longitude,
			end_stop.id, end_stop.name, end_stop.city, end_stop.latitude, end_stop.longitude
		FROM segments s
		JOIN stops start ON s.start_stop_id = start.id
		JOIN stops end_stop ON s.end_stop_id = end_stop.id
		WHERE start.city = $1
		  AND end_stop.city = $2
		  AND s.departure_time >= $3
		  AND s.departure_time < $4
		ORDER BY s.departure_time
	`

	rows, err := r.db.db.QueryContext(ctx, query, fromCity, toCity, departureStart, departureEnd)
	if err != nil {
		return nil, fmt.Errorf("error querying segments by criteria: %w", err)
	}
	defer rows.Close()

	var segments []domain.Segment
	for rows.Next() {
		var segment domain.Segment
		var durationNs int64

		if err := rows.Scan(
			&segment.ID,
			&segment.TransportType,
			&segment.Provider,
			&segment.DepartureTime,
			&segment.ArrivalTime,
			&segment.Price,
			&durationNs,
			&segment.SeatCount,
			&segment.ReliabilityRate,
			&segment.Distance,
			&segment.StartStop.ID,
			&segment.StartStop.Name,
			&segment.StartStop.City,
			&segment.StartStop.Latitude,
			&segment.StartStop.Longitude,
			&segment.EndStop.ID,
			&segment.EndStop.Name,
			&segment.EndStop.City,
			&segment.EndStop.Latitude,
			&segment.EndStop.Longitude,
		); err != nil {
			return nil, fmt.Errorf("error scanning segment: %w", err)
		}

		segment.Duration = time.Duration(durationNs)
		segments = append(segments, segment)
	}

	return segments, rows.Err()
}

// DeleteOldSegments removes segments older than specified date
func (r *SegmentRepository) DeleteOldSegments(ctx context.Context, beforeDate time.Time) error {
	const query = `
		DELETE FROM segments
		WHERE route_id IS NULL
		  AND departure_time < $1
	`

	result, err := r.db.db.ExecContext(ctx, query, beforeDate)
	if err != nil {
		return fmt.Errorf("error deleting old segments: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	fmt.Printf("Deleted %d old segments\n", rowsAffected)

	return nil
}

// FindAll retrieves all segments
func (r *SegmentRepository) FindAll(ctx context.Context) ([]domain.Segment, error) {
	const query = `
		SELECT
			s.id, s.transport_type, s.provider,
			s.departure_time, s.arrival_time, s.price, s.duration,
			s.seat_count, s.reliability_rate, s.distance,
			start.id, start.name, start.city, start.latitude, start.longitude,
			end_stop.id, end_stop.name, end_stop.city, end_stop.latitude, end_stop.longitude
		FROM segments s
		JOIN stops start ON s.start_stop_id = start.id
		JOIN stops end_stop ON s.end_stop_id = end_stop.id
		ORDER BY s.departure_time DESC
		LIMIT 1000
	`

	rows, err := r.db.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying all segments: %w", err)
	}
	defer rows.Close()

	var segments []domain.Segment
	for rows.Next() {
		var segment domain.Segment
		var durationNs int64

		if err := rows.Scan(
			&segment.ID,
			&segment.TransportType,
			&segment.Provider,
			&segment.DepartureTime,
			&segment.ArrivalTime,
			&segment.Price,
			&durationNs,
			&segment.SeatCount,
			&segment.ReliabilityRate,
			&segment.Distance,
			&segment.StartStop.ID,
			&segment.StartStop.Name,
			&segment.StartStop.City,
			&segment.StartStop.Latitude,
			&segment.StartStop.Longitude,
			&segment.EndStop.ID,
			&segment.EndStop.Name,
			&segment.EndStop.City,
			&segment.EndStop.Latitude,
			&segment.EndStop.Longitude,
		); err != nil {
			return nil, fmt.Errorf("error scanning segment: %w", err)
		}

		segment.Duration = time.Duration(durationNs)
		segments = append(segments, segment)
	}

	return segments, rows.Err()
}
