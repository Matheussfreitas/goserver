package auth

import (
	"goserver/internal/platform/env"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(env.LoadConfig().JWTSecret)

func GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"email":      email,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(jwtKey)
}
