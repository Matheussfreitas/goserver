package middleware

import (
	"context"
	"goserver/internal/utils"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const UserEmailKey ContextKey = "user_email"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token não fornecido", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Formato de token inválido", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		token, err := utils.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Token inválido ou expirado", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["email"].(string)
			ctx := context.WithValue(r.Context(), UserEmailKey, email)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
