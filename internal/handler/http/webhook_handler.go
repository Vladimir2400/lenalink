package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/internal/service"
)

// YooKassaWebhookEvent represents incoming webhook event
type YooKassaWebhookEvent struct {
	Type   string                   `json:"type"`
	Event  string                   `json:"event"`
	Object YooKassaPaymentObject    `json:"object"`
}

type YooKassaPaymentObject struct {
	ID       string            `json:"id"`
	Status   string            `json:"status"`
	Metadata map[string]string `json:"metadata"`
	Amount   YooKassaAmount    `json:"amount"`
}

type YooKassaAmount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

// WebhookHandler handles webhook notifications
type WebhookHandler struct {
	bookingService *service.BookingService
	paymentService *service.PaymentService
	errorHandler   *ErrorHandler
}

// NewWebhookHandler creates a new webhook handler
func NewWebhookHandler(
	bookingService *service.BookingService,
	paymentService *service.PaymentService,
) *WebhookHandler {
	return &WebhookHandler{
		bookingService: bookingService,
		paymentService: paymentService,
		errorHandler:   NewErrorHandler(),
	}
}

// HandleYooKassaWebhook processes YooKassa payment notifications
func (h *WebhookHandler) HandleYooKassaWebhook(w http.ResponseWriter, r *http.Request) {
	// 1. Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "INVALID_BODY", "Cannot read request body")
		return
	}
	defer r.Body.Close()

	// 2. Parse webhook event
	var event YooKassaWebhookEvent
	if err := json.Unmarshal(body, &event); err != nil {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "INVALID_JSON", "Cannot parse webhook event")
		return
	}

	// 3. Log webhook (for debugging)
	fmt.Printf("[YooKassa Webhook] Event: %s, Payment ID: %s, Status: %s\n",
		event.Event, event.Object.ID, event.Object.Status)

	// 4. Get order_id from metadata
	orderID, ok := event.Object.Metadata["order_id"]
	if !ok {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "MISSING_ORDER_ID", "order_id not found in metadata")
		return
	}

	// 5. Process based on event type
	switch event.Event {
	case "payment.succeeded":
		// Payment completed successfully
		err = h.handlePaymentSucceeded(r.Context(), orderID, event.Object.ID)
	case "payment.canceled":
		// Payment was canceled
		err = h.handlePaymentCanceled(r.Context(), orderID, event.Object.ID)
	case "refund.succeeded":
		// Refund completed
		err = h.handleRefundSucceeded(r.Context(), orderID)
	default:
		// Unknown event - log and ignore
		fmt.Printf("[YooKassa Webhook] Unknown event type: %s\n", event.Event)
	}

	if err != nil {
		fmt.Printf("[YooKassa Webhook] Error processing event: %v\n", err)
		h.errorHandler.RespondWithError(w, http.StatusInternalServerError, "PROCESSING_ERROR", err.Error())
		return
	}

	// 6. Return 200 OK to acknowledge receipt
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}

func (h *WebhookHandler) handlePaymentSucceeded(ctx context.Context, orderID, providerPaymentID string) error {
	// Get booking
	booking, err := h.bookingService.GetBooking(ctx, orderID)
	if err != nil {
		return fmt.Errorf("booking not found: %w", err)
	}

	// Update payment status
	if booking.Payment != nil {
		booking.Payment.Status = domain.PaymentCompleted
		now := time.Now()
		booking.Payment.CompletedAt = &now
		booking.Payment.ProviderPaymentID = providerPaymentID
	}

	// Confirm booking
	booking.MarkAsConfirmed()

	// Save to database
	return h.bookingService.UpdateBooking(ctx, booking)
}

func (h *WebhookHandler) handlePaymentCanceled(ctx context.Context, orderID, providerPaymentID string) error {
	// Get booking
	booking, err := h.bookingService.GetBooking(ctx, orderID)
	if err != nil {
		return fmt.Errorf("booking not found: %w", err)
	}

	// Mark payment as failed
	if booking.Payment != nil {
		booking.Payment.Status = domain.PaymentFailed
		booking.Payment.FailureReason = "Payment canceled by user or provider"
	}

	// Mark booking as failed
	booking.MarkAsFailed("Payment canceled by user or provider")

	// Rollback provider bookings if any
	// (in real implementation, cancel tickets with providers)

	return h.bookingService.UpdateBooking(ctx, booking)
}

func (h *WebhookHandler) handleRefundSucceeded(ctx context.Context, orderID string) error {
	// Update booking status to refunded
	booking, err := h.bookingService.GetBooking(ctx, orderID)
	if err != nil {
		return fmt.Errorf("booking not found: %w", err)
	}

	if booking.Payment != nil {
		booking.Payment.Status = domain.PaymentRefunded
	}
	booking.Status = domain.BookingRefunded

	return h.bookingService.UpdateBooking(ctx, booking)
}
