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
	authHandler    *AuthHandler
	routeHandler   *RouteHandler
	cityHandler    *CityHandler
	bookingHandler *BookingHandler
	webhookHandler *WebhookHandler
}

// NewRouter creates and configures the HTTP router
func NewRouter(
	authService *service.AuthService,
	routeService *service.RouteService,
	cityService *service.CityService,
	bookingService *service.BookingService,
	paymentService *service.PaymentService,
) *Router {
	r := mux.NewRouter()

	// Create handlers
	healthHandler := NewHealthHandler()
	authHandler := NewAuthHandler(authService)
	routeHandler := NewRouteHandler(routeService)
	cityHandler := NewCityHandler(cityService)
	bookingHandler := NewBookingHandler(bookingService)
	webhookHandler := NewWebhookHandler(bookingService, paymentService)

	// Global middleware (applied to all routes)
	r.Use(middleware.Recovery)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logging)
	r.Use(middleware.CORS)

	// API routes (no version prefix for auth)
	api := r.PathPrefix("/api").Subrouter()
	auth := api.PathPrefix("/api").Subrouter()

	// Authentication endpoints (no auth required)
	auth.HandleFunc("/register", authHandler.Register).Methods("POST")
	auth.HandleFunc("/login", authHandler.Login).Methods("POST")

	// Protected endpoints (require authentication)
	api.HandleFunc("/my_routes", middleware.AuthMiddleware(authService)(http.HandlerFunc(bookingHandler.GetMyRoutes)).ServeHTTP).Methods("GET")

	// API v1 routes
	v1 := api.PathPrefix("/v1").Subrouter()

	// Health check endpoints (no auth required)
	v1.HandleFunc("/health", healthHandler.Health).Methods("GET")
	v1.HandleFunc("/ready", healthHandler.Ready).Methods("GET")

	// Route endpoints (no auth required for search/view)
	v1.HandleFunc("/routes/search", routeHandler.SearchRoutes).Methods("POST")
	v1.HandleFunc("/routes/{id}", routeHandler.GetRouteByID).Methods("GET")

	// City endpoints (no auth required)
	v1.HandleFunc("/cities", cityHandler.SearchCities).Methods("GET")

	// Booking endpoints (auth required)
	v1.Handle("/bookings", middleware.AuthMiddleware(authService)(http.HandlerFunc(bookingHandler.CreateBooking))).Methods("POST")
	v1.Handle("/bookings", middleware.AuthMiddleware(authService)(http.HandlerFunc(bookingHandler.ListBookings))).Methods("GET")
	v1.Handle("/bookings/{id}", middleware.AuthMiddleware(authService)(http.HandlerFunc(bookingHandler.GetBooking))).Methods("GET")
	v1.Handle("/bookings/{id}/cancel", middleware.AuthMiddleware(authService)(http.HandlerFunc(bookingHandler.CancelBooking))).Methods("POST")

	// Webhook endpoints (no auth required for payment provider callbacks)
	v1.HandleFunc("/webhooks/yookassa", webhookHandler.HandleYooKassaWebhook).Methods("POST")

	// 404 handler
	r.NotFoundHandler = r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		eh := NewErrorHandler()
		eh.RespondWithError(w, http.StatusNotFound, "NOT_FOUND", "Endpoint not found")
	}).GetHandler()

	return &Router{
		Router:         r,
		healthHandler:  healthHandler,
		authHandler:    authHandler,
		routeHandler:   routeHandler,
		cityHandler:    cityHandler,
		bookingHandler: bookingHandler,
		webhookHandler: webhookHandler,
	}
}
