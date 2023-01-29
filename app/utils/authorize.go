package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AccessCredentials struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type Claims struct {
	Username string `json:"username"`
	Sub      int    `json:"sub"`
	jwt.RegisteredClaims
}

var (
	ErrWrongSignature = errors.New("wrong signature")
	ErrTokenIsInvalid = errors.New("token is invalid")
	ErrTokenIsExpired = errors.New("token is expired")
)

var jwtKey = []byte("my_secret_key")

func GenerateJWT(email string, id int) (*AccessCredentials, error) {
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

	return &AccessCredentials{
		Access:  tokenString,
		Refresh: refreshTokenString,
	}, nil
}

func RefreshJWT(refresh string) (*AccessCredentials, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(refresh, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, ErrWrongSignature
		}
		return nil, err
	}
	if !tkn.Valid {
		return nil, ErrTokenIsInvalid
	}

	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		return nil, ErrTokenIsExpired
	}

	return GenerateJWT(claims.Username, claims.Sub)
}
