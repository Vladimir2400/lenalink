package domain

import "fmt"

// DomainError represents a domain-level error
type DomainError struct {
	Code    string
	Message string
}

func (e DomainError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Domain errors
var (
	ErrRouteNotFound      = DomainError{Code: "ROUTE_NOT_FOUND", Message: "Route not found"}
	ErrBookingNotFound    = DomainError{Code: "BOOKING_NOT_FOUND", Message: "Booking not found"}
	ErrInvalidRoute       = DomainError{Code: "INVALID_ROUTE", Message: "Invalid route"}
	ErrInvalidBooking     = DomainError{Code: "INVALID_BOOKING", Message: "Invalid booking"}
	ErrSegmentNotFound    = DomainError{Code: "SEGMENT_NOT_FOUND", Message: "Segment not found"}
	ErrInvalidSegment     = DomainError{Code: "INVALID_SEGMENT", Message: "Invalid segment"}
	ErrBookingFailed      = DomainError{Code: "BOOKING_FAILED", Message: "Booking failed"}
	ErrTransactionFailed  = DomainError{Code: "TRANSACTION_FAILED", Message: "Transaction failed"}
	ErrInvalidConnection  = DomainError{Code: "INVALID_CONNECTION", Message: "Invalid connection between segments"}
	ErrSearchFailed       = DomainError{Code: "SEARCH_FAILED", Message: "Route search failed"}
	ErrDatabaseError      = DomainError{Code: "DATABASE_ERROR", Message: "Database error"}
	ErrValidationFailed   = DomainError{Code: "VALIDATION_FAILED", Message: "Validation failed"}
)

// NewDomainError creates a new domain error
func NewDomainError(code, message string) DomainError {
	return DomainError{Code: code, Message: message}
}
