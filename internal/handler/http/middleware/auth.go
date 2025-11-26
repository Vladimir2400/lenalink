package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/lenalink/backend/internal/service"
)

// AuthMiddleware validates JWT tokens from Authorization header
func AuthMiddleware(authService *service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, `{"error":"UNAUTHORIZED","message":"Authorization header required"}`, http.StatusUnauthorized)
				return
			}

			// Extract token from "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, `{"error":"UNAUTHORIZED","message":"Invalid Authorization header format. Use: Bearer <token>"}`, http.StatusUnauthorized)
				return
			}

			token := parts[1]

			// Validate token
			userID, err := authService.ValidateToken(token)
			if err != nil {
				http.Error(w, `{"error":"UNAUTHORIZED","message":"Invalid or expired token"}`, http.StatusUnauthorized)
				return
			}

			// Add user_id to request context
			ctx := context.WithValue(r.Context(), "user_id", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserIDFromContext extracts user_id from request context
func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value("user_id").(string)
	return userID, ok
}
