package main

import (
	"context"
	"log"
	"os"

	"github.com/Edigiraldo/RestWebSockets/handlers"
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
	r.HandleFunc("/", handlers.HomeHandler(s))
}
