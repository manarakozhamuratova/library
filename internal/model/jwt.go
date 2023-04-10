package model

import "github.com/golang-jwt/jwt"

type JWTClaim struct {
	Username string
	jwt.StandardClaims
}

type AuthResponse struct {
	Token string `json:"token"`
}
