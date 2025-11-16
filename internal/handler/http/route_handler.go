package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/internal/handler/http/dto"
	"github.com/lenalink/backend/internal/service"
)

// RouteHandler handles route-related HTTP endpoints
type RouteHandler struct {
	routeService     *service.RouteService
	errorHandler     *ErrorHandler
	validator        *Validator
}

// NewRouteHandler creates a new route handler
func NewRouteHandler(routeService *service.RouteService) *RouteHandler {
	return &RouteHandler{
		routeService: routeService,
		errorHandler: NewErrorHandler(),
		validator:    NewValidator(),
	}
}

// SearchRoutes handles POST /api/v1/routes/search
func (h *RouteHandler) SearchRoutes(w http.ResponseWriter, r *http.Request) {
	var req dto.SearchRouteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "INVALID_REQUEST", "Invalid request body")
		return
	}

	// Validate request
	if err := h.validator.ValidateSearchRouteRequest(&req); err != nil {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}

	// Parse departure date
	departureDate, err := time.Parse("2006-01-02", req.DepartureDate)
	if err != nil {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "INVALID_DATE", "Invalid departure date format (use YYYY-MM-DD)")
		return
	}

	// Create search criteria
	criteria := &domain.RouteSearchCriteria{
		FromCity:        req.From,
		ToCity:          req.To,
		DepartureDate:   departureDate,
		PassengerCount:  req.Passengers,
		MaxConnections:  3,
		MaxTransferTime: 1440, // 24 hours
	}

	// Search routes
	result, err := h.routeService.SearchRoutes(r.Context(), criteria)
	if err != nil {
		h.errorHandler.RespondWithDomainError(w, err)
		return
	}

	// Build response with individual route responses
	var routes []dto.RouteResponse
	if result.OptimalRoute != nil {
		routes = append(routes, ToRouteResponse(result.OptimalRoute, "optimal"))
	}
	if result.FastestRoute != nil {
		routes = append(routes, ToRouteResponse(result.FastestRoute, "fastest"))
	}
	if result.CheapestRoute != nil {
		routes = append(routes, ToRouteResponse(result.CheapestRoute, "cheapest"))
	}

	searchResp := dto.RouteSearchResponse{
		Routes:         routes,
		SearchCriteria: req,
	}

	h.errorHandler.RespondWithJSON(w, http.StatusOK, searchResp)
}

// GetRouteByID handles GET /api/v1/routes/{id}
func (h *RouteHandler) GetRouteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	routeID := vars["id"]

	if routeID == "" {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "INVALID_ROUTE_ID", "Route ID is required")
		return
	}

	// Fetch route
	route, err := h.routeService.GetRouteByID(r.Context(), routeID)
	if err != nil {
		h.errorHandler.RespondWithDomainError(w, err)
		return
	}

	resp := dto.RouteDetailsResponse{
		Route:              ToRouteResponse(route, "details"),
		InsuranceAvailable: true,
		InsurancePremium:   route.InsurancePremium,
	}

	h.errorHandler.RespondWithJSON(w, http.StatusOK, resp)
}
