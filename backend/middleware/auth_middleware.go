package middleware

import (
	"context"
	"net/http"

	"github.com/shuxinqiao/event-attendance-tracker/backend/utils"
)

// Define a custom type for the context key to avoid accidental key collisions.
type contextKey string

const userKey = contextKey("claims")

// AuthMiddleware checks for a valid JWT in the request
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorizaed: No token provided", http.StatusUnauthorized)
			return
		}

		claims, err := utils.VerifyJWT(cookie.Value)
		if err != nil {
			http.Error(w, "Unauthorizaed: Invalid token", http.StatusUnauthorized)
			return
		}

		// Add claims to request context if needed
		r = r.WithContext(context.WithValue(r.Context(), userKey, claims))
		next.ServeHTTP(w, r)
	})
}
