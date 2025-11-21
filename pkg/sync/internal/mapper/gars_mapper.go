package mapper

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/pkg/sync/api/gars"
)

// GarsStopToDomain converts GARS Stop to domain.Stop
func GarsStopToDomain(garsStop gars.Stop) (*domain.Stop, error) {
	lat, lon, err := parseCoordinates(garsStop.Coordinates)
	if err != nil {
		// If coordinates parsing fails, use default values
		lat, lon = 0.0, 0.0
	}

	return &domain.Stop{
		ID:        garsStop.RefKey,
		Name:      garsStop.Description,
		City:      garsStop.Settlement,
		Latitude:  lat,
		Longitude: lon,
	}, nil
}

// GarsScheduleToSegment converts GARS TripSchedule to domain.Segment
func GarsScheduleToSegment(
	schedule gars.TripSchedule,
	stops []gars.TripScheduleStop,
	garsStops map[string]gars.Stop,
	fare *gars.Fare,
	seats *gars.SeatAvailability,
	date time.Time,
) (*domain.Segment, error) {
	if len(stops) < 2 {
		return nil, fmt.Errorf("schedule must have at least 2 stops")
	}

	firstStop := stops[0]
	lastStop := stops[len(stops)-1]

	// Get stop details
	startGarsStop, ok := garsStops[firstStop.StopKey]
	if !ok {
		return nil, fmt.Errorf("start stop not found: %s", firstStop.StopKey)
	}

	endGarsStop, ok := garsStops[lastStop.StopKey]
	if !ok {
		return nil, fmt.Errorf("end stop not found: %s", lastStop.StopKey)
	}

	// Convert stops
	startStop, err := GarsStopToDomain(startGarsStop)
	if err != nil {
		return nil, fmt.Errorf("error converting start stop: %w", err)
	}

	endStop, err := GarsStopToDomain(endGarsStop)
	if err != nil {
		return nil, fmt.Errorf("error converting end stop: %w", err)
	}

	// Parse times
	departureTime, err := parseGarsTime(firstStop.Departure, date)
	if err != nil {
		return nil, fmt.Errorf("error parsing departure time: %w", err)
	}

	arrivalTime, err := parseGarsTime(lastStop.Arrival, date)
	if err != nil {
		return nil, fmt.Errorf("error parsing arrival time: %w", err)
	}

	// If arrival is before departure, assume it's next day
	if arrivalTime.Before(departureTime) {
		arrivalTime = arrivalTime.Add(24 * time.Hour)
	}

	duration := arrivalTime.Sub(departureTime)

	// Extract price and seat count
	price := 0.0
	if fare != nil {
		price = fare.Price
	}

	seatCount := 0
	if seats != nil {
		seatCount = seats.FreeSeats
	}

	return &domain.Segment{
		ID:              schedule.RefKey,
		TransportType:   domain.TransportBus, // GARS is for buses
		Provider:        "АвиБус (ГАРС)",
		StartStop:       *startStop,
		EndStop:         *endStop,
		DepartureTime:   departureTime,
		ArrivalTime:     arrivalTime,
		Price:           price,
		Duration:        duration,
		SeatCount:       seatCount,
		ReliabilityRate: 85.0, // Default reliability rate for GARS
		Distance:        int(lastStop.Distance),
	}, nil
}

// parseCoordinates parses "latitude,longitude" string
func parseCoordinates(coords string) (lat, lon float64, err error) {
	if coords == "" {
		return 0, 0, fmt.Errorf("empty coordinates")
	}

	parts := strings.Split(coords, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid coordinates format: %s", coords)
	}

	lat, err = strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid latitude: %w", err)
	}

	lon, err = strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid longitude: %w", err)
	}

	return lat, lon, nil
}

// parseGarsTime parses time string from GARS (format: "HH:MM:SS") and combines with date
func parseGarsTime(timeStr string, date time.Time) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, fmt.Errorf("empty time string")
	}

	// Parse time in format "HH:MM:SS"
	parts := strings.Split(timeStr, ":")
	if len(parts) != 3 {
		return time.Time{}, fmt.Errorf("invalid time format: %s", timeStr)
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid hour: %w", err)
	}

	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid minute: %w", err)
	}

	second, err := strconv.Atoi(parts[2])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid second: %w", err)
	}

	// Combine date and time
	return time.Date(
		date.Year(), date.Month(), date.Day(),
		hour, minute, second, 0,
		time.UTC,
	), nil
}
