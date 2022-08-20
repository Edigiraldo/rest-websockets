package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Edigiraldo/RestWebSockets/handlers"
	"github.com/Edigiraldo/RestWebSockets/middleware"
	"github.com/Edigiraldo/RestWebSockets/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("There was an error loading .env file: ", err)
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("Unable to get PORT environment variable")
	}
	JWT_SECRET := os.Getenv("JWT_SECRET")
	if JWT_SECRET == "" {
		log.Fatal("Unable to get JWT_SECRET environment variable")
	}
	DATABASE_URL := os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		log.Fatal("Unable to get DATABASE_URL environment variable")
	}

	config := server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseURL: DATABASE_URL,
	}

	s, err := server.NewServer(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.Use(middleware.CheckAuthMiddleware(s))

	r.HandleFunc("/api/v1/ws", s.Hub().HandleWebSocket)

	r.HandleFunc("/api/v1/", handlers.Home(s)).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/signup", handlers.SingUp(s)).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/login", handlers.Login(s)).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/me", handlers.Me(s)).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/posts", handlers.ListPostHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/posts", handlers.CreatePost(s)).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/posts/{id}", handlers.GetPostById(s)).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/posts/{id}", handlers.UpdatePostById(s)).Methods(http.MethodPatch)
	r.HandleFunc("/api/v1/posts/{id}", handlers.DeletePostById(s)).Methods(http.MethodDelete)
}
