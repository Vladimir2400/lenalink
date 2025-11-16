package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/lenalink/backend/internal/handler/http/dto"
)

// HealthHandler handles health check endpoints
type HealthHandler struct{}

// NewHealthHandler creates a new health handler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Health returns health status
func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := dto.HealthResponse{
		Status:    "healthy",
		Version:   "0.4.0",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Services: map[string]string{
			"route_service":   "ready",
			"booking_service": "ready",
			"database":        "ready",
			"cache":           "ready",
		},
	}

	json.NewEncoder(w).Encode(response)
}

// Ready returns readiness status
func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := dto.HealthResponse{
		Status:    "ready",
		Version:   "0.4.0",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
