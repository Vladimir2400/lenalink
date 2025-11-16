package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/internal/handler/http/dto"
	"github.com/lenalink/backend/internal/service"
)

// BookingHandler handles booking-related HTTP endpoints
type BookingHandler struct {
	bookingService *service.BookingService
	errorHandler   *ErrorHandler
	validator      *Validator
}

// NewBookingHandler creates a new booking handler
func NewBookingHandler(bookingService *service.BookingService) *BookingHandler {
	return &BookingHandler{
		bookingService: bookingService,
		errorHandler:   NewErrorHandler(),
		validator:      NewValidator(),
	}
}

// CreateBooking handles POST /api/v1/bookings
func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "INVALID_REQUEST", "Invalid request body")
		return
	}

	// Validate request
	if err := h.validator.ValidateCreateBookingRequest(&req); err != nil {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}

	// Convert passenger request to domain
	passenger, err := ToDomainPassenger(&req.Passenger)
	if err != nil {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "INVALID_PASSENGER", "Invalid passenger data: "+err.Error())
		return
	}

	// Parse payment method
	paymentMethod := domain.PaymentMethod(req.PaymentMethod)
	validMethods := map[domain.PaymentMethod]bool{
		domain.PaymentCard:     true,
		domain.PaymentYooKassa: true,
		domain.PaymentCloudPay: true,
		domain.PaymentSberPay:  true,
	}
	if !validMethods[paymentMethod] {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "INVALID_PAYMENT_METHOD", "Payment method must be 'card', 'yookassa', 'cloudpay', or 'sberpay'")
		return
	}

	// Create booking
	booking, err := h.bookingService.CreateBooking(
		r.Context(),
		req.RouteID,
		passenger,
		req.IncludeInsurance,
		paymentMethod,
	)
	if err != nil {
		h.errorHandler.RespondWithDomainError(w, err)
		return
	}

	resp := ToBookingResponse(booking)
	h.errorHandler.RespondWithJSON(w, http.StatusCreated, resp)
}

// GetBooking handles GET /api/v1/bookings/{id}
func (h *BookingHandler) GetBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingID := vars["id"]

	if bookingID == "" {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "INVALID_BOOKING_ID", "Booking ID is required")
		return
	}

	// Fetch booking
	booking, err := h.bookingService.GetBooking(r.Context(), bookingID)
	if err != nil {
		h.errorHandler.RespondWithDomainError(w, err)
		return
	}

	resp := ToBookingResponse(booking)
	h.errorHandler.RespondWithJSON(w, http.StatusOK, resp)
}

// CancelBooking handles POST /api/v1/bookings/{id}/cancel
func (h *BookingHandler) CancelBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingID := vars["id"]

	if bookingID == "" {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "INVALID_BOOKING_ID", "Booking ID is required")
		return
	}

	var req dto.CancelBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "INVALID_REQUEST", "Invalid request body")
		return
	}

	// Validate cancellation request
	if err := h.validator.ValidateCancelBookingRequest(&req); err != nil {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}

	// Cancel booking
	if err := h.bookingService.CancelBooking(r.Context(), bookingID, req.Reason); err != nil {
		h.errorHandler.RespondWithDomainError(w, err)
		return
	}

	// Fetch updated booking
	booking, err := h.bookingService.GetBooking(r.Context(), bookingID)
	if err != nil {
		h.errorHandler.RespondWithDomainError(w, err)
		return
	}

	resp := map[string]interface{}{
		"message": "Booking cancelled successfully",
		"booking": ToBookingResponse(booking),
	}

	h.errorHandler.RespondWithJSON(w, http.StatusOK, resp)
}

// ListBookings handles GET /api/v1/bookings (admin endpoint)
func (h *BookingHandler) ListBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := h.bookingService.ListBookings(r.Context())
	if err != nil {
		h.errorHandler.RespondWithDomainError(w, err)
		return
	}

	summaries := make([]dto.BookingSummaryResponse, len(bookings))
	for i, booking := range bookings {
		summaries[i] = dto.BookingSummaryResponse{
			ID:             booking.ID,
			Status:         string(booking.Status),
			RouteID:        booking.RouteID,
			PassengerEmail: booking.Passenger.Email,
			GrandTotal:     booking.GrandTotal,
			CreatedAt:      booking.CreatedAt,
			ConfirmedAt:    booking.ConfirmedAt,
		}
	}

	resp := map[string]interface{}{
		"total":    len(bookings),
		"bookings": summaries,
	}

	h.errorHandler.RespondWithJSON(w, http.StatusOK, resp)
}
