package middleware

import (
	"context"
	"fmt"
	"monly-login-api/utils"
	"net/http"
	"strings"
)

// alias
type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("AuthMiddleware called")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.ResponseWithAppropriateError(w, utils.AuthError{"Authorization header required"})
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.ResponseWithAppropriateError(w, utils.AuthError{"Invalid Authorization header format"})
			return
		}
		claims, err := utils.ParseJWT(parts[1], utils.SecretKey)
		if err != nil {
			utils.ResponseWithAppropriateError(w, utils.AuthError{"Invalid token: " + err.Error()})
			return
		}
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
