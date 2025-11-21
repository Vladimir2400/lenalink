package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/internal/repository"
)

// StopRepository implements repository.StopRepository interface for PostgreSQL
type StopRepository struct {
	db *Database
}

// NewStopRepository creates a new stop repository
func NewStopRepository(db *Database) repository.StopRepository {
	return &StopRepository{db: db}
}

// Save stores a new stop
func (r *StopRepository) Save(ctx context.Context, stop *domain.Stop) error {
	const query = `
		INSERT INTO stops (id, name, city, latitude, longitude, stop_type)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	stopType := r.inferStopType(stop.Name)

	_, err := r.db.db.ExecContext(ctx, query,
		stop.ID,
		stop.Name,
		stop.City,
		stop.Latitude,
		stop.Longitude,
		stopType,
	)

	if err != nil {
		return fmt.Errorf("error saving stop: %w", err)
	}

	return nil
}

// Upsert inserts or updates a stop (by unique key name+city)
func (r *StopRepository) Upsert(ctx context.Context, stop *domain.Stop) error {
	const query = `
		INSERT INTO stops (id, name, city, latitude, longitude, stop_type)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (name, city)
		DO UPDATE SET
			latitude = EXCLUDED.latitude,
			longitude = EXCLUDED.longitude,
			stop_type = EXCLUDED.stop_type
	`

	stopType := r.inferStopType(stop.Name)

	_, err := r.db.db.ExecContext(ctx, query,
		stop.ID,
		stop.Name,
		stop.City,
		stop.Latitude,
		stop.Longitude,
		stopType,
	)

	if err != nil {
		return fmt.Errorf("error upserting stop: %w", err)
	}

	return nil
}

// FindByID retrieves a stop by ID
func (r *StopRepository) FindByID(ctx context.Context, id string) (*domain.Stop, error) {
	const query = `
		SELECT id, name, city, latitude, longitude
		FROM stops
		WHERE id = $1
	`

	var stop domain.Stop

	err := r.db.db.QueryRowContext(ctx, query, id).Scan(
		&stop.ID,
		&stop.Name,
		&stop.City,
		&stop.Latitude,
		&stop.Longitude,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("stop not found: %s", id)
		}
		return nil, fmt.Errorf("error querying stop: %w", err)
	}

	return &stop, nil
}

// FindByCity retrieves all stops in a city
func (r *StopRepository) FindByCity(ctx context.Context, city string) ([]domain.Stop, error) {
	const query = `
		SELECT id, name, city, latitude, longitude
		FROM stops
		WHERE city = $1
		ORDER BY name
	`

	rows, err := r.db.db.QueryContext(ctx, query, city)
	if err != nil {
		return nil, fmt.Errorf("error querying stops by city: %w", err)
	}
	defer rows.Close()

	var stops []domain.Stop
	for rows.Next() {
		var stop domain.Stop
		if err := rows.Scan(
			&stop.ID,
			&stop.Name,
			&stop.City,
			&stop.Latitude,
			&stop.Longitude,
		); err != nil {
			return nil, fmt.Errorf("error scanning stop: %w", err)
		}
		stops = append(stops, stop)
	}

	return stops, rows.Err()
}

// FindByCoordinates finds stops within radius (km) from given coordinates
func (r *StopRepository) FindByCoordinates(ctx context.Context, lat, lon float64, radiusKm int) ([]domain.Stop, error) {
	// Using Haversine formula for distance calculation
	const query = `
		SELECT id, name, city, latitude, longitude,
		       6371 * acos(
		           cos(radians($1)) * cos(radians(latitude)) *
		           cos(radians(longitude) - radians($2)) +
		           sin(radians($1)) * sin(radians(latitude))
		       ) AS distance
		FROM stops
		WHERE 6371 * acos(
		           cos(radians($1)) * cos(radians(latitude)) *
		           cos(radians(longitude) - radians($2)) +
		           sin(radians($1)) * sin(radians(latitude))
		       ) <= $3
		ORDER BY distance
	`

	rows, err := r.db.db.QueryContext(ctx, query, lat, lon, radiusKm)
	if err != nil {
		return nil, fmt.Errorf("error querying stops by coordinates: %w", err)
	}
	defer rows.Close()

	var stops []domain.Stop
	for rows.Next() {
		var stop domain.Stop
		var distance float64
		if err := rows.Scan(
			&stop.ID,
			&stop.Name,
			&stop.City,
			&stop.Latitude,
			&stop.Longitude,
			&distance,
		); err != nil {
			return nil, fmt.Errorf("error scanning stop: %w", err)
		}
		stops = append(stops, stop)
	}

	return stops, rows.Err()
}

// FindAll retrieves all stops
func (r *StopRepository) FindAll(ctx context.Context) ([]domain.Stop, error) {
	const query = `
		SELECT id, name, city, latitude, longitude
		FROM stops
		ORDER BY city, name
	`

	rows, err := r.db.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying all stops: %w", err)
	}
	defer rows.Close()

	var stops []domain.Stop
	for rows.Next() {
		var stop domain.Stop
		if err := rows.Scan(
			&stop.ID,
			&stop.Name,
			&stop.City,
			&stop.Latitude,
			&stop.Longitude,
		); err != nil {
			return nil, fmt.Errorf("error scanning stop: %w", err)
		}
		stops = append(stops, stop)
	}

	return stops, rows.Err()
}

// inferStopType infers stop type from name
func (r *StopRepository) inferStopType(name string) string {
	// Simple heuristic based on keywords in name
	name = fmt.Sprintf("%s", name) // ensure lowercase for matching

	// Check for airport keywords
	if containsAny(name, []string{"аэропорт", "airport", "авиа"}) {
		return "airport"
	}

	// Check for river port keywords
	if containsAny(name, []string{"речной", "порт", "пристань", "причал"}) {
		return "port"
	}

	// Check for train station keywords
	if containsAny(name, []string{"вокзал", "станция", "жд"}) {
		return "station"
	}

	// Default to terminal (bus station)
	return "terminal"
}

// containsAny checks if string contains any of the substrings
func containsAny(s string, subs []string) bool {
	for _, sub := range subs {
		if len(s) >= len(sub) {
			for i := 0; i <= len(s)-len(sub); i++ {
				match := true
				for j := 0; j < len(sub); j++ {
					if toLower(s[i+j]) != toLower(sub[j]) {
						match = false
						break
					}
				}
				if match {
					return true
				}
			}
		}
	}
	return false
}

// toLower converts byte to lowercase
func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}
