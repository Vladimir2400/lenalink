package http

import (
	"encoding/json"
	"net/http"

	"github.com/lenalink/backend/internal/handler/http/dto"
	"github.com/lenalink/backend/internal/service"
)

// AuthHandler handles authentication HTTP endpoints
type AuthHandler struct {
	authService  *service.AuthService
	errorHandler *ErrorHandler
	validator    *Validator
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService:  authService,
		errorHandler: NewErrorHandler(),
		validator:    NewValidator(),
	}
}

// Register handles POST /api/register
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "INVALID_REQUEST", "Invalid request body")
		return
	}

	// Validate request
	if err := h.validator.ValidateStruct(&req); err != nil {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}

	// Register user
	user, err := h.authService.Register(r.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "REGISTRATION_FAILED", err.Error())
		return
	}

	// Generate token
	token, err := h.authService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		h.errorHandler.RespondWithError(w, http.StatusInternalServerError, "TOKEN_GENERATION_FAILED", "Failed to generate token")
		return
	}

	// Build response
	response := dto.AuthResponse{
		Token: token,
		User: dto.UserResponse{
			ID:          user.ID,
			Name:        user.Name,
			Email:       user.Email,
			CreatedAt:   user.CreatedAt,
			LastLoginAt: user.LastLoginAt,
		},
	}

	h.errorHandler.RespondWithJSON(w, http.StatusCreated, response)
}

// Login handles POST /api/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "INVALID_REQUEST", "Invalid request body")
		return
	}

	// Validate request
	if err := h.validator.ValidateStruct(&req); err != nil {
		h.errorHandler.RespondWithError(w, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}

	// Login user
	token, user, err := h.authService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		h.errorHandler.RespondWithError(w, http.StatusUnauthorized, "LOGIN_FAILED", err.Error())
		return
	}

	// Build response
	response := dto.AuthResponse{
		Token: token,
		User: dto.UserResponse{
			ID:          user.ID,
			Name:        user.Name,
			Email:       user.Email,
			CreatedAt:   user.CreatedAt,
			LastLoginAt: user.LastLoginAt,
		},
	}

	h.errorHandler.RespondWithJSON(w, http.StatusOK, response)
}
