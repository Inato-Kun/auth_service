package utils

import (
	"auth-service/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string) (string, error) {
	secret := config.GetEnv("JWT_SECRET")
	expiration := time.Hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(expiration).Unix(),
	})

	return token.SignedString([]byte(secret))
}
