package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"

	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/internal/repository"
)

// RouteRepository implements repository.RouteRepository interface for PostgreSQL
type RouteRepository struct {
	db *Database
}

// NewRouteRepository creates a new route repository
func NewRouteRepository(db *Database) repository.RouteRepository {
	return &RouteRepository{db: db}
}

// FindByID retrieves a route with all segments and connections
func (r *RouteRepository) FindByID(ctx context.Context, id string) (*domain.Route, error) {
	// First, fetch the route
	const routeQuery = `
		SELECT id, from_city, to_city, departure_time, arrival_time,
		       total_duration, total_price, reliability_score,
		       insurance_premium, insurance_included, transport_types, saved_at
		FROM routes
		WHERE id = $1
	`

	var route domain.Route
	var transportTypes pq.StringArray

	err := r.db.db.QueryRowContext(ctx, routeQuery, id).Scan(
		&route.ID,
		&route.FromCity,
		&route.ToCity,
		&route.DepartureTime,
		&route.ArrivalTime,
		&route.TotalDuration,
		&route.TotalPrice,
		&route.ReliabilityScore,
		&route.InsurancePremium,
		&route.InsuranceIncluded,
		&transportTypes,
		&route.SavedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrRouteNotFound
		}
		return nil, fmt.Errorf("error querying route: %w", err)
	}

	// Convert transport types
	for _, tt := range transportTypes {
		route.TransportTypes = append(route.TransportTypes, domain.TransportType(tt))
	}

	// Fetch segments
	if err := r.fetchSegments(ctx, &route); err != nil {
		return nil, err
	}

	// Fetch connections
	if err := r.fetchConnections(ctx, &route); err != nil {
		return nil, err
	}

	return &route, nil
}

// FindAll retrieves all routes
func (r *RouteRepository) FindAll(ctx context.Context) ([]domain.Route, error) {
	const query = `
		SELECT id, from_city, to_city, departure_time, arrival_time,
		       total_duration, total_price, reliability_score,
		       insurance_premium, insurance_included, transport_types, saved_at
		FROM routes
		ORDER BY saved_at DESC
		LIMIT 100
	`

	rows, err := r.db.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying routes: %w", err)
	}
	defer rows.Close()

	var routes []domain.Route
	for rows.Next() {
		var route domain.Route
		var transportTypes pq.StringArray

		if err := rows.Scan(
			&route.ID,
			&route.FromCity,
			&route.ToCity,
			&route.DepartureTime,
			&route.ArrivalTime,
			&route.TotalDuration,
			&route.TotalPrice,
			&route.ReliabilityScore,
			&route.InsurancePremium,
			&route.InsuranceIncluded,
			&transportTypes,
			&route.SavedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning route: %w", err)
		}

		// Convert transport types
		for _, tt := range transportTypes {
			route.TransportTypes = append(route.TransportTypes, domain.TransportType(tt))
		}

		routes = append(routes, route)
	}

	return routes, rows.Err()
}

// FindByCriteria searches routes by search criteria
func (r *RouteRepository) FindByCriteria(ctx context.Context, criteria *domain.RouteSearchCriteria) ([]domain.Route, error) {
	query := `
		SELECT id, from_city, to_city, departure_time, arrival_time,
		       total_duration, total_price, reliability_score,
		       insurance_premium, insurance_included, transport_types, saved_at
		FROM routes
		WHERE from_city = $1 AND to_city = $2
		AND DATE(departure_time) = $3
	`

	args := []interface{}{
		criteria.FromCity,
		criteria.ToCity,
		criteria.DepartureDate,
	}

	// Add optional filters
	if criteria.BudgetMax > 0 {
		query += ` AND total_price <= $4`
		args = append(args, criteria.BudgetMax)
	}

	if criteria.BudgetMin > 0 {
		query += ` AND total_price >= $5`
		args = append(args, criteria.BudgetMin)
	}

	query += ` ORDER BY reliability_score DESC, total_price ASC`
	query += ` LIMIT 50`

	rows, err := r.db.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying routes by criteria: %w", err)
	}
	defer rows.Close()

	var routes []domain.Route
	for rows.Next() {
		var route domain.Route
		var transportTypes pq.StringArray

		if err := rows.Scan(
			&route.ID,
			&route.FromCity,
			&route.ToCity,
			&route.DepartureTime,
			&route.ArrivalTime,
			&route.TotalDuration,
			&route.TotalPrice,
			&route.ReliabilityScore,
			&route.InsurancePremium,
			&route.InsuranceIncluded,
			&transportTypes,
			&route.SavedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning route: %w", err)
		}

		for _, tt := range transportTypes {
			route.TransportTypes = append(route.TransportTypes, domain.TransportType(tt))
		}

		routes = append(routes, route)
	}

	return routes, rows.Err()
}

// Save stores a new route
func (r *RouteRepository) Save(ctx context.Context, route *domain.Route) error {
	const query = `
		INSERT INTO routes (
			id, from_city, to_city, departure_time, arrival_time,
			total_duration, total_price, reliability_score,
			insurance_premium, insurance_included, transport_types, saved_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	// Convert transport types to string array
	var transportTypes []string
	for _, tt := range route.TransportTypes {
		transportTypes = append(transportTypes, string(tt))
	}

	_, err := r.db.db.ExecContext(ctx, query,
		route.ID,
		route.FromCity,
		route.ToCity,
		route.DepartureTime,
		route.ArrivalTime,
		route.TotalDuration,
		route.TotalPrice,
		route.ReliabilityScore,
		route.InsurancePremium,
		route.InsuranceIncluded,
		pq.Array(transportTypes),
		time.Now(),
	)

	if err != nil {
		return fmt.Errorf("error saving route: %w", err)
	}

	return nil
}

// Update modifies an existing route
func (r *RouteRepository) Update(ctx context.Context, route *domain.Route) error {
	const query = `
		UPDATE routes
		SET from_city = $2, to_city = $3, total_price = $4,
		    reliability_score = $5, insurance_premium = $6
		WHERE id = $1
	`

	result, err := r.db.db.ExecContext(ctx, query,
		route.ID,
		route.FromCity,
		route.ToCity,
		route.TotalPrice,
		route.ReliabilityScore,
		route.InsurancePremium,
	)

	if err != nil {
		return fmt.Errorf("error updating route: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rows == 0 {
		return domain.ErrRouteNotFound
	}

	return nil
}

// Delete removes a route and cascading segments/connections
func (r *RouteRepository) Delete(ctx context.Context, id string) error {
	const query = `DELETE FROM routes WHERE id = $1`

	result, err := r.db.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting route: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rows == 0 {
		return domain.ErrRouteNotFound
	}

	return nil
}

// Helper functions

func (r *RouteRepository) fetchSegments(ctx context.Context, route *domain.Route) error {
	const query = `
		SELECT s.id, s.transport_type, s.provider,
		       s.start_stop_id, ss.name, ss.city, ss.latitude, ss.longitude,
		       s.end_stop_id, es.name, es.city, es.latitude, es.longitude,
		       s.departure_time, s.arrival_time,
		       s.price, s.duration, s.seat_count,
		       s.reliability_rate, s.distance
		FROM segments s
		JOIN stops ss ON s.start_stop_id = ss.id
		JOIN stops es ON s.end_stop_id = es.id
		WHERE s.route_id = $1
		ORDER BY s.sequence_order
	`

	rows, err := r.db.db.QueryContext(ctx, query, route.ID)
	if err != nil {
		return fmt.Errorf("error querying segments: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var segment domain.Segment

		if err := rows.Scan(
			&segment.ID,
			&segment.TransportType,
			&segment.Provider,
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
			&segment.DepartureTime,
			&segment.ArrivalTime,
			&segment.Price,
			&segment.Duration,
			&segment.SeatCount,
			&segment.ReliabilityRate,
			&segment.Distance,
		); err != nil {
			return fmt.Errorf("error scanning segment: %w", err)
		}

		route.Segments = append(route.Segments, segment)
	}

	return rows.Err()
}

func (r *RouteRepository) fetchConnections(ctx context.Context, route *domain.Route) error {
	const query = `
		SELECT id, from_segment_id, to_segment_id,
		       transfer_duration, transfer_distance,
		       requires_transport, is_valid, gap
		FROM connections
		WHERE route_id = $1
		ORDER BY sequence_order
	`

	rows, err := r.db.db.QueryContext(ctx, query, route.ID)
	if err != nil {
		return fmt.Errorf("error querying connections: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var connection domain.Connection
		var fromSegmentID, toSegmentID sql.NullString

		if err := rows.Scan(
			nil, // id (not used)
			&fromSegmentID,
			&toSegmentID,
			&connection.TransferDuration,
			&connection.TransferDistance,
			&connection.RequiresTransport,
			&connection.IsValid,
			&connection.Gap,
		); err != nil {
			return fmt.Errorf("error scanning connection: %w", err)
		}

		// Link segments if IDs exist
		if fromSegmentID.Valid {
			for i := range route.Segments {
				if route.Segments[i].ID == fromSegmentID.String {
					connection.From = &route.Segments[i]
					break
				}
			}
		}

		if toSegmentID.Valid {
			for i := range route.Segments {
				if route.Segments[i].ID == toSegmentID.String {
					connection.To = &route.Segments[i]
					break
				}
			}
		}

		route.Connections = append(route.Connections, connection)
	}

	return rows.Err()
}
