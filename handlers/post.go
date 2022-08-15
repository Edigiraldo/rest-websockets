package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Edigiraldo/RestWebSockets/models"
	"github.com/Edigiraldo/RestWebSockets/repository"
	"github.com/Edigiraldo/RestWebSockets/server"
	"github.com/golang-jwt/jwt"
	"github.com/segmentio/ksuid"
)

type PostRequest struct {
	Content string `json:"content"`
}

type PostResponse struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func CreatePost(s server.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var claims *models.AppClaims
		var ok bool

		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecret), nil
			})

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}

		if claims, ok = token.Claims.(*models.AppClaims); !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body := PostRequest{}
		err = json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		post := models.Post{
			Id:      id.String(),
			Content: body.Content,
			UserId:  claims.UserId,
		}
		err = repository.InsertPost(r.Context(), &post)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)

	})
}
