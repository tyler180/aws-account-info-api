package main

import (
	"log"
	"net/http"

	"github.com/tyler180/aws-account-info-api/pkg/api"
	"github.com/tyler180/aws-account-info-api/pkg/config"
	"github.com/tyler180/aws-account-info-api/pkg/logger"

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
