package middleware

import (
	"context"
	"encoding/json"
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
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Token não fornecido",
			})
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Formato de token inválido",
			})
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		token, err := utils.ValidateToken(tokenString)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Token inválido ou expirado",
			})
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email, ok := claims["email"].(string)
			if !ok {
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Token inválido: email não encontrado",
				})
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), UserEmailKey, email)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
