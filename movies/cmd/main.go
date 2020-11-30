package main

import (
	"context"
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/movies/internal/client/grpc"
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/movies/internal/server/http"
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/movies/internal/service"
	"github.com/spf13/viper"
	"log"
)

const (
	HttpServerPort           = "HTTP_PORT"
	SearchServiceURI         = "SEARCH_SERVICE_URI"
	RecommendationServiceURI = "RECOMMENDATION_SERVICE_URI"
	ServiceVersion           = "SERVICE_VERSION"
)

func main() {
	log.SetPrefix("MOVIES")

	_ = viper.BindEnv(HttpServerPort)
	_ = viper.BindEnv(SearchServiceURI)
	_ = viper.BindEnv(RecommendationServiceURI)
	_ = viper.BindEnv(ServiceVersion)

	httpServerPort := viper.GetString(HttpServerPort)
	searchServiceURI := viper.GetString(SearchServiceURI)
	recommendationServiceURI := viper.GetString(RecommendationServiceURI)
	serviceVersion := viper.GetString(ServiceVersion)

	if httpServerPort == "" {
		log.Fatal("The moviesHttpServer server port is missing, cannot start the application")
	}

	if searchServiceURI == "" {
		log.Fatal("The search service URI is missing, cannot start the application")
	}

	if recommendationServiceURI == "" {
		log.Fatal("The recommendation service URI is missing, cannot start the application")
	}

	if serviceVersion == "" {
		serviceVersion = "1"
	}

	client, err := grpc.NewClient(context.Background(), searchServiceURI, recommendationServiceURI)
	if err != nil {
		log.Fatal("Cannot connect to the searching service, cannot start the application")
	}

	moviesService := service.NewMovieService(client)
	moviesHttpServer := http.NewMoviesServer(httpServerPort, moviesService)
	log.Fatal(moviesHttpServer.StartServer(serviceVersion))
}
