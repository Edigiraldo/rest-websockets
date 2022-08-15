package handlers

import (
	"errors"
	"strings"

	"github.com/Edigiraldo/RestWebSockets/models"
	"github.com/golang-jwt/jwt"
)

var (
	ErrorNonAuthorized = errors.New("invalid credentials")
)

func CheckAuthentication(authToken string, JWTSecret string) (claims *models.AppClaims, err error) {
	var ok bool

	tokenString := strings.TrimSpace(authToken)
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JWTSecret), nil
		})

	if err != nil {
		return nil, ErrorNonAuthorized
	}

	if claims, ok = token.Claims.(*models.AppClaims); !ok || !token.Valid {
		return nil, ErrorNonAuthorized
	}

	return claims, nil
}
