package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("my_secret_key")

type Claims struct {
	UserLogin string `json:"login"`
	jwt.RegisteredClaims
}

func GenerateJWT(userLogin string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserLogin: userLogin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
