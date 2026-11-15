package middleware

import (
	"context"
	"event-api/internal/platform/auth"
	"net/http"
	"strings"
)

// UserClaimsKey is the context key for user claims.
var UserClaimsKey = "userClaims"

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func HasRole(role string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(UserClaimsKey).(*auth.Claims)
		if !ok {
			http.Error(w, "Could not retrieve user claims", http.StatusInternalServerError)
			return
		}

		if claims.Role != role && claims.Role != "administrator" { // Admins can do anything
			http.Error(w, "Forbidden: insufficient privileges", http.StatusForbidden)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}
