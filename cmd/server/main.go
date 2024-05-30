package main

import (
	"log"
	"my-aws-api/pkg/api"
	"my-aws-api/pkg/config"
	"my-aws-api/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log := logger.NewLogger(cfg.LogLevel)

	r := mux.NewRouter()
	api.RegisterRoutes(r)

	log.Infof("Starting server on %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
