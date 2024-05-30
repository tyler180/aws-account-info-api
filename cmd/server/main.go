package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gorilla/mux"
	"github.com/tyler180/aws-account-info-api/pkg/api"
	"github.com/tyler180/aws-account-info-api/pkg/config"
	"github.com/tyler180/aws-account-info-api/pkg/logger"
)

var (
	r *mux.Router
)

func init() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log := logger.NewLogger(cfg.LogLevel)
	r = mux.NewRouter()
	api.RegisterRoutes(r)
	log.Infof("Routes registered")
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	request, err := http.NewRequest(req.HTTPMethod, req.Path, nil)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	responseRecorder := httptest.NewRecorder()
	r.ServeHTTP(responseRecorder, request)

	return events.APIGatewayProxyResponse{
		StatusCode: responseRecorder.Code,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       responseRecorder.Body.String(),
	}, nil
}

func main() {
	lambda.Start(handler)
}
