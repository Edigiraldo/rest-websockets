package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Edigiraldo/RestWebSockets/models"
	"github.com/Edigiraldo/RestWebSockets/repository"
	"github.com/Edigiraldo/RestWebSockets/server"
	"github.com/golang-jwt/jwt"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func Home(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(
			HomeResponse{
				Message: "Welcome to the server!!!",
				Status:  200,
			},
		)
	}
}

func Me(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecret), nil
			})

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}

		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			user, err := repository.GetUserById(r.Context(), claims.UserId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
