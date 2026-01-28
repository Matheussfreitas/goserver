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
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Token não fornecido",
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Formato de token inválido",
			})
			return
		}

		tokenString := parts[1]

		token, err := utils.ValidateToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Token inválido ou expirado",
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email, ok := claims["email"].(string)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Token inválido: email não encontrado",
				})
				return
			}

			id, ok := claims["sub"].(string)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Token inválido: ID não encontrado",
				})
				return
			}
			ctx := context.WithValue(r.Context(), UserEmailKey, email)
			ctx = context.WithValue(ctx, "user_id", id)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
