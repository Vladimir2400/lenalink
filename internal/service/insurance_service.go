package service

import (
	"context"
	"time"

	"github.com/lenalink/backend/internal/domain"
)

// InsuranceInfo holds insurance calculation result
type InsuranceInfo struct {
	Premium        float64
	Level          string
	CoverageAmount float64
}

// InsuranceService implements business logic for insurance calculations
type InsuranceService struct {
	basePremiumRate      float64 // 5% base premium
	nightFlightSurcharge float64 // Additional % for night flights
	tightConnectionSurcharge float64 // Additional % for connections < 2 hours
}

// NewInsuranceService creates a new insurance service
func NewInsuranceService() *InsuranceService {
	return &InsuranceService{
		basePremiumRate:       0.05,  // 5%
		nightFlightSurcharge:  0.02,  // +2% for night flights
		tightConnectionSurcharge: 0.03, // +3% for tight connections
	}
}

// CalculateInsurance calculates insurance premium and coverage
func (s *InsuranceService) CalculateInsurance(ctx context.Context, route *domain.Route, booking *domain.Booking) (*InsuranceInfo, error) {
	info := &InsuranceInfo{
		Level: "standard",
	}

	// Calculate base premium (5% of route price)
	totalPrice := booking.TotalPrice
	basePremium := totalPrice * s.basePremiumRate

	surchargeRate := 0.0

	// Add surcharge for night flights (between 22:00 and 06:00)
	if s.hasNightFlight(route) {
		surchargeRate += s.nightFlightSurcharge
	}

	// Add surcharge for tight connections (< 2 hours)
	if s.hasTightConnection(route) {
		surchargeRate += s.tightConnectionSurcharge
	}

	// Calculate total premium
	surcharge := totalPrice * surchargeRate
	info.Premium = basePremium + surcharge

	// Determine coverage level and amount based on route characteristics
	if surchargeRate >= 0.05 {
		info.Level = "premium"
		info.CoverageAmount = totalPrice * 1.5 // 150% coverage
	} else if surchargeRate >= 0.03 {
		info.Level = "standard"
		info.CoverageAmount = totalPrice * 1.2 // 120% coverage
	} else {
		info.Level = "basic"
		info.CoverageAmount = totalPrice * 1.0 // 100% coverage
	}

	return info, nil
}

// Private helper methods

func (s *InsuranceService) hasNightFlight(route *domain.Route) bool {
	for _, segment := range route.Segments {
		hour := segment.DepartureTime.Hour()
		// Night flight: between 22:00 and 06:00
		if hour >= 22 || hour < 6 {
			return true
		}
	}
	return false
}

func (s *InsuranceService) hasTightConnection(route *domain.Route) bool {
	for _, conn := range route.Connections {
		if conn.TransferDuration < 2*time.Hour {
			return true
		}
	}
	return false
}
