package model

import "github.com/golang-jwt/jwt"

type JWTClaim struct {
	ID uint
	jwt.StandardClaims
}

type AuthResponse struct {
	Token string `json:"token"`
}
