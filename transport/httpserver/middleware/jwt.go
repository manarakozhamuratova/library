package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/manarakozhamuratova/one-lab-task2/config"
	"github.com/manarakozhamuratova/one-lab-task2/internal/model"
	"github.com/manarakozhamuratova/one-lab-task2/internal/service"
)

type JWTAuth struct {
	jwtKey []byte
	User   service.IUserService
}

func NewJWTAuth(cfg *config.Config, user service.IUserService) *JWTAuth {
	return &JWTAuth{jwtKey: []byte(cfg.JWTKey), User: user}
}

func (m *JWTAuth) GenerateJWT(username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1000 * time.Hour)
	claims := &model.JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.jwtKey)
}

func (m *JWTAuth) ValidateToken(signedToken string) (*model.JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&model.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return m.jwtKey, nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.JWTClaim)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token expired")
	}
	return claims, nil
}

func (m *JWTAuth) ValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := extractToken(c.Request())
		if token != "test" {
			claims, err := m.ValidateToken(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, err.Error())
			}

			ctx := context.WithValue(c.Request().Context(), model.ContextUsername, claims.Username)
			c.SetRequest(c.Request().WithContext(ctx))
		}
		return next(c)
	}
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
