package service

import (
	"context"
	"fmt"

	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/internal/repository"
)

// CityService implements business logic for city search
type CityService struct {
	stopRepo repository.StopRepository
}

// NewCityService creates a new city service
func NewCityService(stopRepo repository.StopRepository) *CityService {
	return &CityService{
		stopRepo: stopRepo,
	}
}

// SearchCitiesByName searches for cities by partial name match
func (s *CityService) SearchCitiesByName(ctx context.Context, namePrefix string) ([]domain.Stop, error) {
	if namePrefix == "" {
		return nil, fmt.Errorf("city name prefix cannot be empty")
	}

	if len(namePrefix) < 2 {
		return nil, fmt.Errorf("city name prefix must be at least 2 characters")
	}

	return s.stopRepo.FindCitiesByName(ctx, namePrefix)
}
