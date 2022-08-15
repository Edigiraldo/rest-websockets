package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Edigiraldo/RestWebSockets/repository"
	"github.com/Edigiraldo/RestWebSockets/server"
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
		claims, err := CheckAuthentication(r.Header.Get("Authorization"), s.Config().JWTSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		user, err := repository.GetUserById(r.Context(), claims.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}
