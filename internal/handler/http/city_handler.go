package http

import (
	"net/http"

	"github.com/lenalink/backend/internal/handler/http/dto"
	"github.com/lenalink/backend/internal/service"
)

// CityHandler handles city-related HTTP endpoints
type CityHandler struct {
	cityService  *service.CityService
	errorHandler *ErrorHandler
}

// NewCityHandler creates a new city handler
func NewCityHandler(cityService *service.CityService) *CityHandler {
	return &CityHandler{
		cityService:  cityService,
		errorHandler: NewErrorHandler(),
	}
}

// SearchCities handles GET /api/v1/cities?name=prefix
func (h *CityHandler) SearchCities(w http.ResponseWriter, r *http.Request) {
	namePrefix := r.URL.Query().Get("name")

	if namePrefix == "" {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "MISSING_PARAMETER", "Query parameter 'name' is required")
		return
	}

	if len(namePrefix) < 2 {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "INVALID_PARAMETER", "Parameter 'name' must be at least 2 characters long")
		return
	}

	// Call service
	stops, err := h.cityService.SearchCitiesByName(r.Context(), namePrefix)
	if err != nil {
		h.errorHandler.RespondWithError(w, http.StatusInternalServerError, "SEARCH_ERROR", err.Error())
		return
	}

	// Convert to response
	cities := make([]dto.CityResponse, 0, len(stops))
	for _, stop := range stops {
		cities = append(cities, dto.CityResponse{
			Name:      stop.CityDisplayName,
			Latitude:  stop.Latitude,
			Longitude: stop.Longitude,
		})
	}

	response := dto.CitiesResponse{
		Cities: cities,
	}

	h.errorHandler.RespondWithJSON(w, http.StatusOK, response)
}
