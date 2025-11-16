package http

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"time"

	"github.com/lenalink/backend/internal/handler/http/dto"
)

// Validator handles request validation
type Validator struct{}

// NewValidator creates a new validator
func NewValidator() *Validator {
	return &Validator{}
}

// ValidateSearchRouteRequest validates search route request
func (v *Validator) ValidateSearchRouteRequest(req *dto.SearchRouteRequest) error {
	if strings.TrimSpace(req.From) == "" {
		return errors.New("'from' field is required")
	}

	if strings.TrimSpace(req.To) == "" {
		return errors.New("'to' field is required")
	}

	if req.From == req.To {
		return errors.New("'from' and 'to' must be different")
	}

	if strings.TrimSpace(req.DepartureDate) == "" {
		return errors.New("'departure_date' field is required")
	}

	// Validate date format (YYYY-MM-DD)
	_, err := time.Parse("2006-01-02", req.DepartureDate)
	if err != nil {
		return errors.New("'departure_date' must be in YYYY-MM-DD format")
	}

	if req.Passengers < 0 {
		return errors.New("'passengers' must be a positive number")
	}

	if req.Passengers > 10 {
		return errors.New("'passengers' cannot exceed 10")
	}

	return nil
}

// ValidateCreateBookingRequest validates create booking request
func (v *Validator) ValidateCreateBookingRequest(req *dto.CreateBookingRequest) error {
	if strings.TrimSpace(req.RouteID) == "" {
		return errors.New("'route_id' field is required")
	}

	// Validate passenger
	if err := v.validatePassengerRequest(&req.Passenger); err != nil {
		return fmt.Errorf("passenger validation failed: %w", err)
	}

	// Validate payment method
	validMethods := map[string]bool{
		"card":      true,
		"yookassa":  true,
		"cloudpay":  true,
		"sberpay":   true,
	}

	if !validMethods[req.PaymentMethod] {
		return errors.New("'payment_method' must be one of: card, yookassa, cloudpay, sberpay")
	}

	return nil
}

// validatePassengerRequest validates passenger information
func (v *Validator) validatePassengerRequest(req *dto.PassengerRequest) error {
	if strings.TrimSpace(req.FirstName) == "" {
		return errors.New("'first_name' is required")
	}

	if strings.TrimSpace(req.LastName) == "" {
		return errors.New("'last_name' is required")
	}

	if strings.TrimSpace(req.DateOfBirth) == "" {
		return errors.New("'date_of_birth' is required")
	}

	// Validate date format
	dob, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		return errors.New("'date_of_birth' must be in YYYY-MM-DD format")
	}

	// Check passenger age (must be 18+)
	age := time.Since(dob).Hours() / 24 / 365
	if age < 18 {
		return errors.New("passenger must be at least 18 years old")
	}

	if age > 120 {
		return errors.New("invalid date of birth")
	}

	if strings.TrimSpace(req.PassportNumber) == "" {
		return errors.New("'passport_number' is required")
	}

	// Validate passport number format (Russian passport: 4 digits space 6 digits)
	passportPattern := regexp.MustCompile(`^\d{4}\s?\d{6}$`)
	if !passportPattern.MatchString(strings.ReplaceAll(req.PassportNumber, " ", "")) {
		return errors.New("'passport_number' must be in format: XXXX XXXXXX (10 digits)")
	}

	// Validate email
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("'email' is required")
	}

	_, err = mail.ParseAddress(req.Email)
	if err != nil {
		return errors.New("'email' must be a valid email address")
	}

	// Validate phone
	if strings.TrimSpace(req.Phone) == "" {
		return errors.New("'phone' is required")
	}

	// Russian phone format: +7XXXXXXXXXX
	phonePattern := regexp.MustCompile(`^\+?7\d{10}$`)
	cleanPhone := strings.ReplaceAll(strings.ReplaceAll(req.Phone, " ", ""), "-", "")
	if !phonePattern.MatchString(cleanPhone) {
		return errors.New("'phone' must be a valid Russian phone number (+7XXXXXXXXXX)")
	}

	return nil
}

// ValidateCancelBookingRequest validates cancel booking request
func (v *Validator) ValidateCancelBookingRequest(req *dto.CancelBookingRequest) error {
	if strings.TrimSpace(req.Reason) == "" {
		return errors.New("'reason' field is required")
	}

	if len(req.Reason) < 10 {
		return errors.New("'reason' must be at least 10 characters")
	}

	if len(req.Reason) > 500 {
		return errors.New("'reason' cannot exceed 500 characters")
	}

	return nil
}
