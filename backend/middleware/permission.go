package middleware

import (
	"net/http"
)

func RoleAuthorization(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract user's role from session or context here
			role := r.Context().Value("Role").(string)

			// Check if user's role is allowed
			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					next.ServeHTTP(w, r)
					return
				}
			}

			// If not authorized, return a 403 forbidden
			http.Error(w, "Forbidden", http.StatusForbidden)
		})
	}
}
