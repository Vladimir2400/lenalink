package memory

import (
	"context"
	"fmt"
	"sync"
	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/internal/repository"
)

// RouteRepository implements RouteRepository interface using in-memory storage
type RouteRepository struct {
	mu     sync.RWMutex
	routes map[string]*domain.Route
}

// NewRouteRepository creates a new in-memory route repository
func NewRouteRepository() repository.RouteRepository {
	return &RouteRepository{
		routes: make(map[string]*domain.Route),
	}
}

// FindByID retrieves a route by ID
func (r *RouteRepository) FindByID(ctx context.Context, id string) (*domain.Route, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	route, exists := r.routes[id]
	if !exists {
		return nil, domain.ErrRouteNotFound
	}

	return route, nil
}

// FindAll retrieves all routes
func (r *RouteRepository) FindAll(ctx context.Context) ([]domain.Route, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	routes := make([]domain.Route, 0, len(r.routes))
	for _, route := range r.routes {
		routes = append(routes, *route)
	}

	return routes, nil
}

// Save stores a new route
func (r *RouteRepository) Save(ctx context.Context, route *domain.Route) error {
	if route == nil {
		return domain.ErrInvalidRoute
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.routes[route.ID]; exists {
		return fmt.Errorf("route with ID %s already exists", route.ID)
	}

	r.routes[route.ID] = route

	return nil
}

// Update modifies an existing route
func (r *RouteRepository) Update(ctx context.Context, route *domain.Route) error {
	if route == nil {
		return domain.ErrInvalidRoute
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.routes[route.ID]; !exists {
		return domain.ErrRouteNotFound
	}

	r.routes[route.ID] = route

	return nil
}

// Delete removes a route
func (r *RouteRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.routes[id]; !exists {
		return domain.ErrRouteNotFound
	}

	delete(r.routes, id)

	return nil
}

// FindByCriteria searches routes by criteria
func (r *RouteRepository) FindByCriteria(ctx context.Context, criteria *domain.RouteSearchCriteria) ([]domain.Route, error) {
	if criteria == nil {
		return nil, fmt.Errorf("criteria cannot be nil")
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	var results []domain.Route

	for _, route := range r.routes {
		// Filter by cities
		if route.FromCity != criteria.FromCity || route.ToCity != criteria.ToCity {
			continue
		}

		// Filter by departure date (same day)
		if route.DepartureTime.Year() != criteria.DepartureDate.Year() ||
			route.DepartureTime.Month() != criteria.DepartureDate.Month() ||
			route.DepartureTime.Day() != criteria.DepartureDate.Day() {
			continue
		}

		// Filter by budget if specified
		if criteria.BudgetMax > 0 && route.TotalPrice > criteria.BudgetMax {
			continue
		}
		if criteria.BudgetMin > 0 && route.TotalPrice < criteria.BudgetMin {
			continue
		}

		results = append(results, *route)
	}

	return results, nil
}
