package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Edigiraldo/RestWebSockets/models"
	"github.com/Edigiraldo/RestWebSockets/repository"
	"github.com/Edigiraldo/RestWebSockets/server"
	"github.com/segmentio/ksuid"
)

var (
	ErrInvalidId = errors.New("invalid id")
)

type PostRequest struct {
	Content string `json:"content"`
}

type PostResponse struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

func ListPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := CheckAuthentication(r.Header.Get("Authorization"), s.Config().JWTSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		pageStr := r.URL.Query().Get("page")
		var page = uint64(0)
		if pageStr != "" {
			page, err = strconv.ParseUint(pageStr, 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
		posts, err := repository.ListPosts(r.Context(), claims.UserId, page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	}
}

func CreatePost(s server.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := CheckAuthentication(r.Header.Get("Authorization"), s.Config().JWTSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
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

func GetPostById(s server.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := CheckAuthentication(r.Header.Get("Authorization"), s.Config().JWTSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		params := mux.Vars(r)
		id, ok := params["id"]
		if !ok {
			http.Error(w, ErrInvalidId.Error(), http.StatusBadRequest)
		}

		post, err := repository.GetPostById(r.Context(), id, claims.UserId)
		if err != nil {
			log.Println(id, claims.UserId)
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		response := PostResponse{
			Id:      post.Id,
			Content: post.Content,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	})
}

func UpdatePostById(s server.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := CheckAuthentication(r.Header.Get("Authorization"), s.Config().JWTSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		params := mux.Vars(r)
		id, ok := params["id"]
		if !ok {
			http.Error(w, ErrInvalidId.Error(), http.StatusBadRequest)
		}

		newPost := PostRequest{}
		err = json.NewDecoder(r.Body).Decode(&newPost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		post, err := repository.UpdatePostById(r.Context(), newPost.Content, id, claims.UserId)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		response := PostResponse{
			Id:      post.Id,
			Content: post.Content,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})
}

func DeletePostById(s server.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := CheckAuthentication(r.Header.Get("Authorization"), s.Config().JWTSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		params := mux.Vars(r)
		id, ok := params["id"]
		if !ok {
			http.Error(w, ErrInvalidId.Error(), http.StatusBadRequest)
		}

		err = repository.DeletePostById(r.Context(), id, claims.UserId)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})
}
