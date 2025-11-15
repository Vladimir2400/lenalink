package service

import (
	"context"
	"fmt"
	"time"

	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/internal/repository"
	"github.com/lenalink/backend/pkg/utils"
)

// RouteService implements business logic for routes
type RouteService struct {
	routeRepo repository.RouteRepository
}

// NewRouteService creates a new route service
func NewRouteService(routeRepo repository.RouteRepository) *RouteService {
	return &RouteService{
		routeRepo: routeRepo,
	}
}

// GetRouteByID retrieves a route by ID
func (s *RouteService) GetRouteByID(ctx context.Context, id string) (*domain.Route, error) {
	if id == "" {
		return nil, fmt.Errorf("route ID cannot be empty")
	}

	return s.routeRepo.FindByID(ctx, id)
}

// SearchRoutes searches for routes based on criteria
func (s *RouteService) SearchRoutes(ctx context.Context, criteria *domain.RouteSearchCriteria) (*domain.RouteSearchResult, error) {
	if criteria == nil {
		return nil, fmt.Errorf("search criteria cannot be nil")
	}

	if err := s.validateSearchCriteria(criteria); err != nil {
		return nil, err
	}

	// Find all routes matching criteria
	routes, err := s.routeRepo.FindByCriteria(ctx, criteria)
	if err != nil {
		return nil, err
	}

	if len(routes) == 0 {
		return &domain.RouteSearchResult{
			RequestID:     utils.GenerateID(),
			FromCity:      criteria.FromCity,
			ToCity:        criteria.ToCity,
			DepartureDate: criteria.DepartureDate,
			PassengerCount: criteria.PassengerCount,
			SearchedAt:    time.Now(),
		}, nil
	}

	// Select optimal, fastest, and cheapest routes
	result := &domain.RouteSearchResult{
		RequestID:      utils.GenerateID(),
		FromCity:       criteria.FromCity,
		ToCity:         criteria.ToCity,
		DepartureDate:  criteria.DepartureDate,
		PassengerCount: criteria.PassengerCount,
		SearchedAt:     time.Now(),
	}

	// Find optimal route (highest reliability score)
	optimalIdx := 0
	for i := 1; i < len(routes); i++ {
		if routes[i].ReliabilityScore > routes[optimalIdx].ReliabilityScore {
			optimalIdx = i
		}
	}
	result.OptimalRoute = &routes[optimalIdx]

	// Find fastest route (shortest duration)
	fastestIdx := 0
	for i := 1; i < len(routes); i++ {
		if routes[i].TotalDuration < routes[fastestIdx].TotalDuration {
			fastestIdx = i
		}
	}
	result.FastestRoute = &routes[fastestIdx]

	// Find cheapest route (lowest price)
	cheapestIdx := 0
	for i := 1; i < len(routes); i++ {
		if routes[i].TotalPrice < routes[cheapestIdx].TotalPrice {
			cheapestIdx = i
		}
	}
	result.CheapestRoute = &routes[cheapestIdx]

	return result, nil
}

// SaveRoute saves a new route
func (s *RouteService) SaveRoute(ctx context.Context, route *domain.Route) error {
	if route == nil {
		return fmt.Errorf("route cannot be nil")
	}

	if err := s.validateRoute(route); err != nil {
		return err
	}

	if route.ID == "" {
		route.ID = utils.GenerateID()
	}

	if route.SavedAt.IsZero() {
		route.SavedAt = time.Now()
	}

	return s.routeRepo.Save(ctx, route)
}

// UpdateRoute updates an existing route
func (s *RouteService) UpdateRoute(ctx context.Context, route *domain.Route) error {
	if route == nil {
		return fmt.Errorf("route cannot be nil")
	}

	if route.ID == "" {
		return fmt.Errorf("route ID is required for update")
	}

	if err := s.validateRoute(route); err != nil {
		return err
	}

	return s.routeRepo.Update(ctx, route)
}

// DeleteRoute removes a route
func (s *RouteService) DeleteRoute(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("route ID cannot be empty")
	}

	return s.routeRepo.Delete(ctx, id)
}

// Private validation methods

func (s *RouteService) validateSearchCriteria(criteria *domain.RouteSearchCriteria) error {
	if criteria.FromCity == "" {
		return fmt.Errorf("from_city is required")
	}

	if criteria.ToCity == "" {
		return fmt.Errorf("to_city is required")
	}

	if criteria.FromCity == criteria.ToCity {
		return fmt.Errorf("from_city and to_city cannot be the same")
	}

	if criteria.PassengerCount <= 0 {
		criteria.PassengerCount = 1
	}

	if criteria.MaxConnections <= 0 {
		criteria.MaxConnections = 3
	}

	if criteria.MaxTransferTime <= 0 {
		criteria.MaxTransferTime = 1440 // 24 hours
	}

	return nil
}

func (s *RouteService) validateRoute(route *domain.Route) error {
	if route.FromCity == "" {
		return fmt.Errorf("from_city is required")
	}

	if route.ToCity == "" {
		return fmt.Errorf("to_city is required")
	}

	if route.FromCity == route.ToCity {
		return fmt.Errorf("from_city and to_city cannot be the same")
	}

	if len(route.Segments) == 0 {
		return fmt.Errorf("route must have at least one segment")
	}

	if route.TotalPrice <= 0 {
		return fmt.Errorf("total_price must be greater than 0")
	}

	return nil
}
