package sync

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/internal/repository"
	"github.com/lenalink/backend/pkg/sync/api/aviasales"
	"github.com/lenalink/backend/pkg/sync/api/gars"
	"github.com/lenalink/backend/pkg/sync/api/rzd"
	"github.com/lenalink/backend/pkg/sync/internal/mapper"
)

// service handles synchronization of data from external providers to the database.
type service struct {
	garsClient      *gars.Client
	aviasalesClient *aviasales.Client
	rzdClient       *rzd.MockClient
	stopRepo        repository.StopRepository
	segmentRepo     repository.SegmentRepository
}

// Ensure service implements Syncer interface.
var _ Syncer = (*service)(nil)

// SyncAll synchronizes data from all providers.
func (s *service) SyncAll(ctx context.Context) error {
	log.Println("Starting full synchronization...")

	// Sync GARS (АвиБус) data
	if err := s.syncGarsData(ctx); err != nil {
		log.Printf("Error syncing GARS data: %v", err)
		// Continue with other providers even if one fails
	}

	// Sync Aviasales data
	if err := s.syncAviasalesData(ctx); err != nil {
		log.Printf("Error syncing Aviasales data: %v", err)
	}

	// Sync RZD data
	if err := s.syncRzdData(ctx); err != nil {
		log.Printf("Error syncing RZD data: %v", err)
	}

	// Clean up old segments (older than 7 days)
	cutoffDate := time.Now().AddDate(0, 0, -7)
	if err := s.segmentRepo.DeleteOldSegments(ctx, cutoffDate); err != nil {
		log.Printf("Error deleting old segments: %v", err)
	}

	log.Println("Full synchronization completed")
	return nil
}

// SyncProvider synchronizes data from a specific provider.
func (s *service) SyncProvider(ctx context.Context, provider Provider) error {
	switch provider {
	case ProviderGARS:
		return s.syncGarsData(ctx)
	case ProviderAviasales:
		return s.syncAviasalesData(ctx)
	case ProviderRZD:
		return s.syncRzdData(ctx)
	default:
		return fmt.Errorf("unknown provider: %s", provider)
	}
}

// StartPeriodicSync starts periodic synchronization with the specified interval.
func (s *service) StartPeriodicSync(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	log.Printf("Starting periodic sync with interval: %s", interval)

	// Run initial sync
	if err := s.SyncAll(ctx); err != nil {
		log.Printf("Error in initial sync: %v", err)
	}

	// Run periodic sync
	for {
		select {
		case <-ticker.C:
			if err := s.SyncAll(ctx); err != nil {
				log.Printf("Error in periodic sync: %v", err)
			}
		case <-ctx.Done():
			log.Println("Stopping periodic sync")
			return
		}
	}
}

// syncGarsData synchronizes data from GARS (АвиБус) API.
func (s *service) syncGarsData(ctx context.Context) error {
	log.Println("Syncing GARS (АвиБус) data...")

	// Fetch stops from GARS
	garsService := gars.NewService(s.garsClient)
	stops, _, err := garsService.Stops(ctx)
	if err != nil {
		return fmt.Errorf("error fetching GARS stops: %w", err)
	}

	log.Printf("Fetched %d stops from GARS", len(stops))

	// Convert and save stops
	stopsCount := 0
	for _, garsStop := range stops {
		domainStop, err := mapper.GarsStopToDomain(garsStop)
		if err != nil {
			log.Printf("Error converting GARS stop %s: %v", garsStop.RefKey, err)
			continue
		}

		if err := s.stopRepo.Upsert(ctx, domainStop); err != nil {
			log.Printf("Error saving stop %s: %v", domainStop.ID, err)
			continue
		}
		stopsCount++
	}

	log.Printf("Saved %d stops from GARS", stopsCount)

	// Fetch schedules for next 30 days
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 30)

	log.Printf("Fetching GARS schedules from %s to %s", startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	// Fetch trip schedules
	tripSchedules, _, err := garsService.TripSchedules(ctx)
	if err != nil {
		log.Printf("Warning: Error fetching trip schedules: %v", err)
		log.Println("GARS data sync completed (stops only)")
		return nil
	}

	log.Printf("Fetched %d trip schedules from GARS", len(tripSchedules))

	// Create stop map for quick lookup
	stopMap := make(map[string]gars.Stop)
	for _, stop := range stops {
		stopMap[stop.RefKey] = stop
	}

	// Convert trip schedules to segments
	segmentsCount := 0
	for _, schedule := range tripSchedules {
		// Fetch stops for this trip
		tripStops, _, err := garsService.TripScheduleStops(ctx, gars.WithFilter(fmt.Sprintf("TripScheduleKey eq '%s'", schedule.RefKey)))
		if err != nil {
			log.Printf("Warning: Error fetching stops for schedule %s: %v", schedule.RefKey, err)
			continue
		}

		if len(tripStops) < 2 {
			continue // Need at least origin and destination
		}

		// For each trip schedule, create segments for upcoming days
		for d := 0; d < 7; d++ { // Next 7 days
			tripDate := startDate.AddDate(0, 0, d)

			// Try to fetch fare information (optional)
			var fare *gars.Fare
			fares, _, err := garsService.Fares(ctx, gars.WithFilter(fmt.Sprintf("TripScheduleKey eq '%s'", schedule.RefKey)), gars.WithTop(1))
			if err == nil && len(fares) > 0 {
				fare = &fares[0]
			}

			// Convert to segment
			segment, err := mapper.GarsScheduleToSegment(schedule, tripStops, stopMap, fare, nil, tripDate)
			if err != nil {
				log.Printf("Warning: Error converting schedule %s to segment: %v", schedule.RefKey, err)
				continue
			}

			// Save segment
			if err := s.segmentRepo.Save(ctx, segment); err != nil {
				log.Printf("Warning: Error saving segment %s: %v", segment.ID, err)
				continue
			}
			segmentsCount++
		}
	}

	log.Printf("Saved %d segments from GARS", segmentsCount)
	log.Println("GARS data sync completed")
	return nil
}

// syncAviasalesData synchronizes data from Aviasales API.
func (s *service) syncAviasalesData(ctx context.Context) error {
	log.Println("Syncing Aviasales data...")

	// Fetch airports from Aviasales
	airports, err := s.aviasalesClient.GetAirports(ctx)
	if err != nil {
		return fmt.Errorf("error fetching airports: %w", err)
	}

	log.Printf("Fetched %d airports from Aviasales", len(airports))

	// Filter only Russian airports for now
	russianAirports := []aviasales.Airport{}
	for _, airport := range airports {
		if airport.CountryCode == "RU" && airport.Flightable && airport.IataType == "airport" {
			russianAirports = append(russianAirports, airport)
		}
	}

	log.Printf("Found %d active Russian airports", len(russianAirports))

	// Convert and save airports as stops
	// Build map of all airports (not just Russian) for flight conversion
	// Use CityCode as key since flight API uses city codes
	airportMap := make(map[string]aviasales.Airport)
	for _, airport := range airports {
		// Store the first airport for each city code
		if _, exists := airportMap[airport.CityCode]; !exists {
			airportMap[airport.CityCode] = airport
		}
	}

	// Save Russian airports to database
	airportsCount := 0
	for _, airport := range russianAirports {
		domainStop, err := mapper.AviasalesAirportToDomain(airport)
		if err != nil {
			log.Printf("Error converting airport %s: %v", airport.Code, err)
			continue
		}

		if err := s.stopRepo.Upsert(ctx, domainStop); err != nil {
			log.Printf("Error saving airport %s: %v", airport.Code, err)
			continue
		}
		airportsCount++
	}

	log.Printf("Saved %d airports from Aviasales", airportsCount)
	log.Printf("Built airport map with %d cities for flight conversion", len(airportMap))

	// Fetch flight prices for Yakutia routes
	// Focus on routes connecting Yakutia with major cities and internal Yakutia routes
	yakutiaRoutes := []struct {
		origin      string
		destination string
		description string
	}{
		// Якутск - крупные города России
		{"MOW", "YKS", "Москва - Якутск"},
		{"YKS", "MOW", "Якутск - Москва"},
		{"LED", "YKS", "Санкт-Петербург - Якутск"},
		{"YKS", "LED", "Якутск - Санкт-Петербург"},
		{"SVX", "YKS", "Екатеринбург - Якутск"},
		{"YKS", "SVX", "Якутск - Екатеринбург"},
		{"KJA", "YKS", "Красноярск - Якутск"},
		{"YKS", "KJA", "Якутск - Красноярск"},
		{"IKT", "YKS", "Иркутск - Якутск"},
		{"YKS", "IKT", "Якутск - Иркутск"},
		{"NSK", "YKS", "Новосибирск - Якутск"},
		{"YKS", "NSK", "Якутск - Новосибирск"},

		// Внутренние маршруты Якутии (основные города)
		{"YKS", "MJZ", "Якутск - Мирный"},
		{"MJZ", "YKS", "Мирный - Якутск"},
		{"YKS", "ULK", "Якутск - Ленск"},
		{"ULK", "YKS", "Ленск - Якутск"},
		{"YKS", "NER", "Якутск - Нерюнгри"},
		{"NER", "YKS", "Нерюнгри - Якутск"},
		{"YKS", "CKH", "Якутск - Чокурдах"},
		{"CKH", "YKS", "Чокурдах - Якутск"},
		{"YKS", "SUK", "Якутск - Саккырыр"},
		{"SUK", "YKS", "Саккырыр - Якутск"},
		{"YKS", "UMS", "Якутск - Усть-Мая"},
		{"UMS", "YKS", "Усть-Мая - Якутск"},
		{"YKS", "VYI", "Якутск - Вилюйск"},
		{"VYI", "YKS", "Вилюйск - Якутск"},
		{"YKS", "DKS", "Якутск - Диксон"},
		{"DKS", "YKS", "Диксон - Якутск"},

		// Дальний Восток - Якутск
		{"VVO", "YKS", "Владивосток - Якутск"},
		{"YKS", "VVO", "Якутск - Владивосток"},
		{"HBR", "YKS", "Хабаровск - Якутск"},
		{"YKS", "HBR", "Якутск - Хабаровск"},
		{"PKC", "YKS", "Петропавловск-Камчатский - Якутск"},
		{"YKS", "PKC", "Якутск - Петропавловск-Камчатский"},
		{"UUS", "YKS", "Южно-Сахалинск - Якутск"},
		{"YKS", "UUS", "Якутск - Южно-Сахалинск"},
		{"GDX", "YKS", "Магадан - Якутск"},
		{"YKS", "GDX", "Якутск - Магадан"},

		// Дополнительные маршруты из Москвы (для транзита)
		{"MOW", "KJA", "Москва - Красноярск"},
		{"MOW", "IKT", "Москва - Иркутск"},
		{"MOW", "NSK", "Москва - Новосибирск"},
		{"MOW", "VVO", "Москва - Владивосток"},
		{"MOW", "HBR", "Москва - Хабаровск"},
	}

	segmentsCount := 0
	for _, route := range yakutiaRoutes {
		// Get flights for the next month
		departureDate := time.Now().Format("2006-01")

		log.Printf("Fetching flights for %s (%s → %s)", route.description, route.origin, route.destination)
		flights, err := s.aviasalesClient.GetPrices(ctx, route.origin, route.destination, departureDate)
		if err != nil {
			log.Printf("Warning: Error fetching flights for %s: %v", route.description, err)
			continue
		}

		log.Printf("Fetched %d flights for %s-%s", len(flights), route.origin, route.destination)

		// Convert flights to segments
		segments := []domain.Segment{}
		for _, flight := range flights {
			segment, err := mapper.AviasalesFlightToSegment(flight, airportMap)
			if err != nil {
				log.Printf("Error converting flight %s→%s on %s: %v", flight.Origin, flight.Destination, flight.DepartDate, err)
				continue
			}
			segments = append(segments, *segment)
		}

		// Batch save segments
		if len(segments) > 0 {
			if err := s.segmentRepo.BatchSave(ctx, segments); err != nil {
				log.Printf("Error saving segments for %s-%s: %v", route.origin, route.destination, err)
				continue
			}
			segmentsCount += len(segments)
		}
	}

	log.Printf("Saved %d flight segments from Aviasales", segmentsCount)
	log.Println("Aviasales data sync completed")
	return nil
}

// syncRzdData synchronizes mock data from RZD client.
func (s *service) syncRzdData(ctx context.Context) error {
	log.Println("Syncing RZD mock data...")

	// Fetch stations from RZD
	stations, err := s.rzdClient.GetStations(ctx)
	if err != nil {
		return fmt.Errorf("error fetching RZD stations: %w", err)
	}

	log.Printf("Fetched %d stations from RZD", len(stations))

	// Convert and save stations
	stationsCount := 0
	stationMap := make(map[string]rzd.Station)
	for _, station := range stations {
		stationMap[station.Code] = station

		domainStop, err := mapper.RzdStationToDomain(station)
		if err != nil {
			log.Printf("Error converting station %s: %v", station.Code, err)
			continue
		}

		if err := s.stopRepo.Upsert(ctx, domainStop); err != nil {
			log.Printf("Error saving station %s: %v", station.Code, err)
			continue
		}
		stationsCount++
	}

	log.Printf("Saved %d stations from RZD", stationsCount)

	// Fetch trains for next 7 days
	segmentsCount := 0
	for i := 0; i < 7; i++ {
		date := time.Now().AddDate(0, 0, i)

		trains, err := s.rzdClient.GetTrains(ctx, "", "", date)
		if err != nil {
			log.Printf("Error fetching trains for %s: %v", date.Format("2006-01-02"), err)
			continue
		}

		log.Printf("Fetched %d trains for %s", len(trains), date.Format("2006-01-02"))

		// Convert trains to segments
		segments := []domain.Segment{}
		for _, train := range trains {
			// Get tickets for the train
			tickets, err := s.rzdClient.GetTickets(ctx, train.TrainNumber)
			if err != nil {
				log.Printf("Error fetching tickets for train %s: %v", train.TrainNumber, err)
				continue
			}

			// Create segment for each ticket class
			for _, ticket := range tickets {
				segment, err := mapper.RzdTrainToSegment(train, stationMap, &ticket)
				if err != nil {
					log.Printf("Error converting train %s: %v", train.TrainNumber, err)
					continue
				}
				segments = append(segments, *segment)
			}
		}

		// Batch save segments
		if len(segments) > 0 {
			if err := s.segmentRepo.BatchSave(ctx, segments); err != nil {
				log.Printf("Error saving segments for %s: %v", date.Format("2006-01-02"), err)
				continue
			}
			segmentsCount += len(segments)
		}
	}

	log.Printf("Saved %d train segments from RZD", segmentsCount)
	log.Println("RZD data sync completed")
	return nil
}
