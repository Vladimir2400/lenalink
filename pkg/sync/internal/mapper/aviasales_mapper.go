package mapper

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/pkg/sync/api/aviasales"
)

// AviasalesAirportToDomain converts Aviasales Airport to domain.Stop
func AviasalesAirportToDomain(airport aviasales.Airport) (*domain.Stop, error) {
	return &domain.Stop{
		ID:        airport.Code, // Use IATA code as ID
		Name:      airport.Name,
		City:      airport.CityName,
		Latitude:  airport.Latitude,
		Longitude: airport.Longitude,
	}, nil
}

// AviasalesFlightToSegment converts Aviasales Flight to domain.Segment
func AviasalesFlightToSegment(flight aviasales.Flight, airports map[string]aviasales.Airport) (*domain.Segment, error) {
	// Get origin and destination airports
	originAirport, ok := airports[flight.OriginAirport]
	if !ok {
		return nil, fmt.Errorf("origin airport not found: %s", flight.OriginAirport)
	}

	destAirport, ok := airports[flight.DestinationAirport]
	if !ok {
		return nil, fmt.Errorf("destination airport not found: %s", flight.DestinationAirport)
	}

	// Convert airports to stops
	startStop, err := AviasalesAirportToDomain(originAirport)
	if err != nil {
		return nil, fmt.Errorf("error converting origin airport: %w", err)
	}

	endStop, err := AviasalesAirportToDomain(destAirport)
	if err != nil {
		return nil, fmt.Errorf("error converting destination airport: %w", err)
	}

	// Parse departure time
	departureTime, err := time.Parse(time.RFC3339, flight.DepartureAt)
	if err != nil {
		return nil, fmt.Errorf("error parsing departure time: %w", err)
	}

	// Calculate arrival time from duration
	arrivalTime := departureTime.Add(time.Duration(flight.Duration) * time.Minute)

	// Generate segment ID
	segmentID := uuid.New().String()

	// Calculate distance (approximate - not provided by Aviasales)
	distance := estimateDistance(startStop.Latitude, startStop.Longitude, endStop.Latitude, endStop.Longitude)

	return &domain.Segment{
		ID:              segmentID,
		TransportType:   domain.TransportAir,
		Provider:        fmt.Sprintf("Aviasales (%s)", flight.Airline),
		StartStop:       *startStop,
		EndStop:         *endStop,
		DepartureTime:   departureTime,
		ArrivalTime:     arrivalTime,
		Price:           flight.Price,
		Duration:        time.Duration(flight.Duration) * time.Minute,
		SeatCount:       flight.AvailableSeats,
		ReliabilityRate: 90.0, // Default reliability rate for airlines
		Distance:        distance,
	}, nil
}

// estimateDistance estimates distance between two coordinates using Haversine formula
func estimateDistance(lat1, lon1, lat2, lon2 float64) int {
	const earthRadius = 6371.0 // Earth radius in km

	// Convert to radians
	lat1Rad := lat1 * 3.14159265359 / 180
	lat2Rad := lat2 * 3.14159265359 / 180
	deltaLat := (lat2 - lat1) * 3.14159265359 / 180
	deltaLon := (lon2 - lon1) * 3.14159265359 / 180

	// Haversine formula
	a := sin(deltaLat/2)*sin(deltaLat/2) +
		cos(lat1Rad)*cos(lat2Rad)*sin(deltaLon/2)*sin(deltaLon/2)
	c := 2 * atan2(sqrt(a), sqrt(1-a))

	return int(earthRadius * c)
}

// Math helper functions
func sin(x float64) float64 {
	// Simple approximation for small angles
	return x - (x*x*x)/6 + (x*x*x*x*x)/120
}

func cos(x float64) float64 {
	return sqrt(1 - sin(x)*sin(x))
}

func sqrt(x float64) float64 {
	if x == 0 {
		return 0
	}
	// Newton's method
	z := x
	for i := 0; i < 10; i++ {
		z = (z + x/z) / 2
	}
	return z
}

func atan2(y, x float64) float64 {
	// Simple approximation
	if x > 0 {
		return atan(y / x)
	}
	if x < 0 && y >= 0 {
		return atan(y/x) + 3.14159265359
	}
	if x < 0 && y < 0 {
		return atan(y/x) - 3.14159265359
	}
	if x == 0 && y > 0 {
		return 3.14159265359 / 2
	}
	if x == 0 && y < 0 {
		return -3.14159265359 / 2
	}
	return 0
}

func atan(x float64) float64 {
	// Taylor series approximation
	return x - (x*x*x)/3 + (x*x*x*x*x)/5
}
