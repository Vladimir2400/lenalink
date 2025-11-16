package http

import (
	"encoding/json"
	"net/http"

	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/internal/handler/http/dto"
)

// ErrorHandler provides error response helpers
type ErrorHandler struct{}

// NewErrorHandler creates a new error handler
func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

// RespondWithError sends a standardized error response
func (eh *ErrorHandler) RespondWithError(w http.ResponseWriter, statusCode int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := dto.ErrorResponse{
		Error: dto.ErrorDetail{
			Code:    code,
			Message: message,
		},
	}

	json.NewEncoder(w).Encode(resp)
}

// RespondWithJSON sends a successful JSON response
func (eh *ErrorHandler) RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := map[string]interface{}{
		"success": true,
		"data":    data,
	}

	json.NewEncoder(w).Encode(resp)
}

// MapDomainErrorToHTTP maps domain errors to HTTP status codes
func (eh *ErrorHandler) MapDomainErrorToHTTP(err error) (int, string, string) {
	if domainErr, ok := err.(domain.DomainError); ok {
		switch domainErr.Code {
		case "VALIDATION_FAILED", "INVALID_ROUTE", "INVALID_BOOKING", "INVALID_SEGMENT", "INVALID_CONNECTION":
			return http.StatusBadRequest, domainErr.Code, domainErr.Message
		case "ROUTE_NOT_FOUND", "BOOKING_NOT_FOUND", "SEGMENT_NOT_FOUND":
			return http.StatusNotFound, domainErr.Code, domainErr.Message
		case "BOOKING_FAILED", "SEARCH_FAILED", "TRANSACTION_FAILED":
			return http.StatusConflict, domainErr.Code, domainErr.Message
		case "DATABASE_ERROR":
			return http.StatusInternalServerError, domainErr.Code, domainErr.Message
		default:
			return http.StatusInternalServerError, "INTERNAL_ERROR", domainErr.Message
		}
	}

	// Generic error handling
	return http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "An unexpected error occurred"
}

// RespondWithDomainError handles domain errors and sends appropriate HTTP response
func (eh *ErrorHandler) RespondWithDomainError(w http.ResponseWriter, err error) {
	statusCode, code, message := eh.MapDomainErrorToHTTP(err)
	eh.RespondWithError(w, statusCode, code, message)
}
