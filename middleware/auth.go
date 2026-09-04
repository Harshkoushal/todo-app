package middleware

import (
	"context"
	"net/http"
	"strings"

	"todo-app/utils"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
