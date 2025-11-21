package mapper

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/pkg/sync/api/rzd"
)

// RzdStationToDomain converts RZD Station to domain.Stop
func RzdStationToDomain(station rzd.Station) (*domain.Stop, error) {
	return &domain.Stop{
		ID:        station.Code,
		Name:      station.Name,
		City:      station.City,
		Latitude:  station.Latitude,
		Longitude: station.Longitude,
	}, nil
}

// RzdTrainToSegment converts RZD Train to domain.Segment
func RzdTrainToSegment(train rzd.Train, stations map[string]rzd.Station, ticket *rzd.Ticket) (*domain.Segment, error) {
	// Get origin and destination stations
	originStation, ok := stations[train.OriginStation]
	if !ok {
		return nil, fmt.Errorf("origin station not found: %s", train.OriginStation)
	}

	destStation, ok := stations[train.DestStation]
	if !ok {
		return nil, fmt.Errorf("destination station not found: %s", train.DestStation)
	}

	// Convert stations to stops
	startStop, err := RzdStationToDomain(originStation)
	if err != nil {
		return nil, fmt.Errorf("error converting origin station: %w", err)
	}

	endStop, err := RzdStationToDomain(destStation)
	if err != nil {
		return nil, fmt.Errorf("error converting destination station: %w", err)
	}

	// Extract price and seat count from ticket
	price := 0.0
	seatCount := 0
	carType := "Общий"
	if ticket != nil {
		price = ticket.Price
		seatCount = ticket.AvailableSeats
		carType = ticket.CarType
	}

	// Generate segment ID
	segmentID := uuid.New().String()

	return &domain.Segment{
		ID:              segmentID,
		TransportType:   domain.TransportRail,
		Provider:        fmt.Sprintf("РЖД (%s, %s)", train.TrainNumber, carType),
		StartStop:       *startStop,
		EndStop:         *endStop,
		DepartureTime:   train.DepartureTime,
		ArrivalTime:     train.ArrivalTime,
		Price:           price,
		Duration:        train.ArrivalTime.Sub(train.DepartureTime),
		SeatCount:       seatCount,
		ReliabilityRate: 92.0, // Default reliability rate for trains
		Distance:        train.Distance,
	}, nil
}
