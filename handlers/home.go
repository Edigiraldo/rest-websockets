package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Edigiraldo/RestWebSockets/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
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
