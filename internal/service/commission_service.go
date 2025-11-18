package service

import (
	"github.com/lenalink/backend/internal/domain"
)

// CommissionConfig holds commission rates for different transport types
type CommissionConfig struct {
	AirCommissionRate     float64 // Percentage (e.g., 0.07 = 7%)
	RailCommissionRate    float64
	BusCommissionRate     float64
	RiverCommissionRate   float64
	TaxiCommissionRate    float64
	WalkCommissionRate    float64
	DefaultCommissionRate float64
}

// DefaultCommissionConfig returns default commission rates
func DefaultCommissionConfig() CommissionConfig {
	return CommissionConfig{
		AirCommissionRate:     0.07, // 7% for flights
		RailCommissionRate:    0.05, // 5% for trains
		BusCommissionRate:     0.08, // 8% for buses
		RiverCommissionRate:   0.10, // 10% for river transport
		TaxiCommissionRate:    0.15, // 15% for taxi
		WalkCommissionRate:    0.00, // 0% for walking (free)
		DefaultCommissionRate: 0.07, // 7% default
	}
}

// CommissionService calculates commission/markup on tickets
type CommissionService struct {
	config CommissionConfig
}

// NewCommissionService creates a new commission service
func NewCommissionService(config CommissionConfig) *CommissionService {
	return &CommissionService{config: config}
}

// CalculateCommission calculates commission for a segment
func (cs *CommissionService) CalculateCommission(transportType domain.TransportType, basePrice float64) float64 {
	rate := cs.getCommissionRate(transportType)
	return basePrice * rate
}

// CalculateTotalPrice calculates total price including commission
func (cs *CommissionService) CalculateTotalPrice(transportType domain.TransportType, basePrice float64) float64 {
	commission := cs.CalculateCommission(transportType, basePrice)
	return basePrice + commission
}

// CalculateCommissionForSegment calculates commission for a route segment
func (cs *CommissionService) CalculateCommissionForSegment(segment *domain.Segment) (basePrice, commission, totalPrice float64) {
	basePrice = segment.Price
	commission = cs.CalculateCommission(segment.TransportType, basePrice)
	totalPrice = basePrice + commission
	return basePrice, commission, totalPrice
}

// CalculateRouteCommission calculates total commission for entire route
func (cs *CommissionService) CalculateRouteCommission(route *domain.Route) (basePrice, totalCommission, grandTotal float64) {
	for i := range route.Segments {
		segmentBase, segmentCommission, _ := cs.CalculateCommissionForSegment(&route.Segments[i])
		basePrice += segmentBase
		totalCommission += segmentCommission
	}
	grandTotal = basePrice + totalCommission
	return basePrice, totalCommission, grandTotal
}

// getCommissionRate returns commission rate for transport type
func (cs *CommissionService) getCommissionRate(transportType domain.TransportType) float64 {
	switch transportType {
	case domain.TransportAir:
		return cs.config.AirCommissionRate
	case domain.TransportRail:
		return cs.config.RailCommissionRate
	case domain.TransportBus:
		return cs.config.BusCommissionRate
	case domain.TransportRiver:
		return cs.config.RiverCommissionRate
	case domain.TransportTaxi:
		return cs.config.TaxiCommissionRate
	case domain.TransportWalk:
		return cs.config.WalkCommissionRate
	default:
		return cs.config.DefaultCommissionRate
	}
}

// GetCommissionRate returns the commission rate for a transport type (for transparency)
func (cs *CommissionService) GetCommissionRate(transportType domain.TransportType) float64 {
	return cs.getCommissionRate(transportType)
}

// GetCommissionRatePercentage returns commission rate as percentage (e.g., 7.0 for 7%)
func (cs *CommissionService) GetCommissionRatePercentage(transportType domain.TransportType) float64 {
	return cs.getCommissionRate(transportType) * 100
}
