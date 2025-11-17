package http

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lenalink/backend/internal/handler/http/middleware"
	"github.com/lenalink/backend/internal/service"
)

// Router sets up all HTTP routes
type Router struct {
	*mux.Router
	healthHandler  *HealthHandler
	routeHandler   *RouteHandler
	bookingHandler *BookingHandler
	webhookHandler *WebhookHandler
}

// NewRouter creates and configures the HTTP router
func NewRouter(
	routeService *service.RouteService,
	bookingService *service.BookingService,
	paymentService *service.PaymentService,
) *Router {
	r := mux.NewRouter()

	// Create handlers
	healthHandler := NewHealthHandler()
	routeHandler := NewRouteHandler(routeService)
	bookingHandler := NewBookingHandler(bookingService)
	webhookHandler := NewWebhookHandler(bookingService, paymentService)

	// Global middleware (applied to all routes)
	r.Use(middleware.Recovery)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logging)
	r.Use(middleware.CORS)

	// Health check endpoints
	r.HandleFunc("/health", healthHandler.Health).Methods("GET")
	r.HandleFunc("/ready", healthHandler.Ready).Methods("GET")

	// API v1 routes
	api := r.PathPrefix("/api/v1").Subrouter()

	// Route endpoints
	api.HandleFunc("/routes/search", routeHandler.SearchRoutes).Methods("POST")
	api.HandleFunc("/routes/{id}", routeHandler.GetRouteByID).Methods("GET")

	// Booking endpoints
	api.HandleFunc("/bookings", bookingHandler.CreateBooking).Methods("POST")
	api.HandleFunc("/bookings", bookingHandler.ListBookings).Methods("GET")
	api.HandleFunc("/bookings/{id}", bookingHandler.GetBooking).Methods("GET")
	api.HandleFunc("/bookings/{id}/cancel", bookingHandler.CancelBooking).Methods("POST")

	// Webhook endpoints (no auth required for payment provider callbacks)
	api.HandleFunc("/webhooks/yookassa", webhookHandler.HandleYooKassaWebhook).Methods("POST")

	// 404 handler
	r.NotFoundHandler = r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		eh := NewErrorHandler()
		eh.RespondWithError(w, http.StatusNotFound, "NOT_FOUND", "Endpoint not found")
	}).GetHandler()

	return &Router{
		Router:         r,
		healthHandler:  healthHandler,
		routeHandler:   routeHandler,
		bookingHandler: bookingHandler,
		webhookHandler: webhookHandler,
	}
}
