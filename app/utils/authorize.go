package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type LoginCredentials struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type Claims struct {
	Username string `json:"username"`
	Sub      int    `json:"sub"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("my_secret_key")

func GenerateJWT(email string, id int) (*LoginCredentials, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	refreshExpirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Username: email,
		Sub:      id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	refresh_claims := &Claims{
		Username: email,
		Sub:      id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpirationTime),
		},
	}

	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claims)
	refreshTokenString, err := refresh_token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return &LoginCredentials{
		Access:  tokenString,
		Refresh: refreshTokenString,
	}, nil
}
