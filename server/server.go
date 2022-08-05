package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseURL string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
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
	}, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	log.Printf("Starting server on port %s\n", b.Config().Port)
	if err := http.ListenAndServe(b.Config().Port, b.router); err != nil {
		log.Fatal("Error while starting server: ", err)
	}
}
