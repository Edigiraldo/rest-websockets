package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Edigiraldo/RestWebSockets/database"
	"github.com/Edigiraldo/RestWebSockets/repository"
	"github.com/Edigiraldo/RestWebSockets/websocket"
	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseURL string
}

type Server interface {
	Config() *Config
	Hub() *websocket.Hub
}

type Broker struct {
	config *Config
	router *mux.Router
	hub    *websocket.Hub
}

func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Hub() *websocket.Hub {
	return b.hub
}

func NewServer(ctx context.Context, config Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port must be specified in config")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("JWTSecret must be specified in config")
	}
	if config.DatabaseURL == "" {
		return nil, errors.New("database URL must be specified in config")
	}

	return &Broker{
		config: &config,
		router: mux.NewRouter(),
		hub:    websocket.NewHub(),
	}, nil
}

func (b *Broker) Start(binder func(server Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)

	postgresDatabase, err := database.NewDatabase(b.config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	postDB := database.NewPostDatabase(postgresDatabase)
	userDB := database.NewUserDatabase(postgresDatabase)

	repository.SetPostRepository(postDB)
	repository.SetUserRepository(userDB)

	go b.hub.Run()

	log.Printf("Starting server on port %s\n", b.Config().Port)

	port := fmt.Sprintf(":%s", b.Config().Port)
	if err := http.ListenAndServe(port, b.router); err != nil {
		log.Fatal("Error while starting server: ", err)
	}
}
