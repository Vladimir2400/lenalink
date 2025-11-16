package service

import (
	"time"

	"github.com/lenalink/backend/internal/domain"
)

// InsuranceConfig holds insurance calculation parameters
type InsuranceConfig struct {
	BasePremiumRate         float64 // Base premium as percentage (e.g., 0.05 = 5%)
	TightConnectionSurcharge float64 // Surcharge for connections < 2 hours
	NightFlightSurcharge    float64 // Surcharge for night flights (22:00-06:00)
	RiverTransportSurcharge float64 // Surcharge for river transport (weather risks)
	MultiSegmentSurcharge   float64 // Surcharge for routes with 3+ segments
}

// DefaultInsuranceConfig returns default insurance configuration
func DefaultInsuranceConfig() InsuranceConfig {
	return InsuranceConfig{
		BasePremiumRate:         0.05,  // 5% base premium
		TightConnectionSurcharge: 0.01,  // +1% per tight connection
		NightFlightSurcharge:    0.005, // +0.5% for night flights
		RiverTransportSurcharge: 0.02,  // +2% for river transport
		MultiSegmentSurcharge:   0.01,  // +1% for 3+ segments
	}
}

// InsuranceService calculates insurance premiums for bookings
type InsuranceService struct {
	config InsuranceConfig
}

// NewInsuranceService creates a new insurance service
func NewInsuranceService(config InsuranceConfig) *InsuranceService {
	return &InsuranceService{config: config}
}

// CalculatePremium calculates insurance premium for a route
func (is *InsuranceService) CalculatePremium(route *domain.Route) float64 {
	if route == nil || len(route.Segments) == 0 {
		return 0
	}

	totalPrice := route.TotalPrice
	premiumRate := is.config.BasePremiumRate

	// Add surcharges based on risk factors
	premiumRate += is.calculateRiskSurcharges(route)

	premium := totalPrice * premiumRate
	return premium
}

// calculateRiskSurcharges calculates additional surcharges based on route risks
func (is *InsuranceService) calculateRiskSurcharges(route *domain.Route) float64 {
	surcharge := 0.0

	// Multi-segment surcharge
	if len(route.Segments) >= 3 {
		surcharge += is.config.MultiSegmentSurcharge
	}

	// Check for tight connections
	tightConnections := is.countTightConnections(route)
	surcharge += float64(tightConnections) * is.config.TightConnectionSurcharge

	// Check for night flights
	hasNightFlights := is.hasNightFlights(route)
	if hasNightFlights {
		surcharge += is.config.NightFlightSurcharge
	}

	// Check for river transport
	hasRiverTransport := is.hasRiverTransport(route)
	if hasRiverTransport {
		surcharge += is.config.RiverTransportSurcharge
	}

	return surcharge
}

// countTightConnections counts connections with less than 2 hours between segments
func (is *InsuranceService) countTightConnections(route *domain.Route) int {
	count := 0
	for i := 0; i < len(route.Segments)-1; i++ {
		currentSegment := route.Segments[i]
		nextSegment := route.Segments[i+1]

		connectionTime := nextSegment.DepartureTime.Sub(currentSegment.ArrivalTime)
		if connectionTime < 2*time.Hour {
			count++
		}
	}
	return count
}

// hasNightFlights checks if route has any night flights (22:00-06:00)
func (is *InsuranceService) hasNightFlights(route *domain.Route) bool {
	for _, segment := range route.Segments {
		if segment.TransportType != domain.TransportAir {
			continue
		}

		hour := segment.DepartureTime.Hour()
		if hour >= 22 || hour < 6 {
			return true
		}
	}
	return false
}

// hasRiverTransport checks if route includes river transport
func (is *InsuranceService) hasRiverTransport(route *domain.Route) bool {
	for _, segment := range route.Segments {
		if segment.TransportType == domain.TransportRiver {
			return true
		}
	}
	return false
}

// GetPremiumBreakdown returns detailed premium calculation breakdown
func (is *InsuranceService) GetPremiumBreakdown(route *domain.Route) map[string]float64 {
	breakdown := make(map[string]float64)

	totalPrice := route.TotalPrice
	breakdown["base_price"] = totalPrice
	breakdown["base_premium_rate"] = is.config.BasePremiumRate
	breakdown["base_premium"] = totalPrice * is.config.BasePremiumRate

	tightConnections := is.countTightConnections(route)
	if tightConnections > 0 {
		breakdown["tight_connections_count"] = float64(tightConnections)
		breakdown["tight_connections_surcharge"] = totalPrice * float64(tightConnections) * is.config.TightConnectionSurcharge
	}

	if is.hasNightFlights(route) {
		breakdown["night_flight_surcharge"] = totalPrice * is.config.NightFlightSurcharge
	}

	if is.hasRiverTransport(route) {
		breakdown["river_transport_surcharge"] = totalPrice * is.config.RiverTransportSurcharge
	}

	if len(route.Segments) >= 3 {
		breakdown["multi_segment_surcharge"] = totalPrice * is.config.MultiSegmentSurcharge
	}

	breakdown["total_premium"] = is.CalculatePremium(route)

	return breakdown
}
