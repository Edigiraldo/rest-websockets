package middleware

import (
	"net/http"
	"strings"

	"github.com/Edigiraldo/RestWebSockets/models"
	"github.com/Edigiraldo/RestWebSockets/server"
	"github.com/golang-jwt/jwt"
)

var NO_AUTH_NEEDED = []string{
	"login",
	"signup",
}

func shouldCheckToken(route string) bool {
	for _, w := range NO_AUTH_NEEDED {
		if strings.Contains(route, w) {
			return false
		}
	}

	return true
}

func CheckAuthMiddleware(s server.Server) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}

			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
			_, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecret), nil
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
