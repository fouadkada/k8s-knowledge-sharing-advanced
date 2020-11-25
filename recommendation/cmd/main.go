package main

import (
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/recommendation/internal/server/grpc"
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/recommendation/internal/service"
	"github.com/spf13/viper"
	"log"
)

const (
	APIKey         = "TMDB_API_KEY"
	GrpcServerPort = "GRPC_PORT"
)

func main() {
	log.SetPrefix("RECOMMENDATION")

	_ = viper.BindEnv(APIKey)
	_ = viper.BindEnv(GrpcServerPort)

	grpcServerPort := viper.GetString(GrpcServerPort)
	apiKey := viper.GetString(APIKey)

	if apiKey == "" {
		log.Fatal("TMDB API Key is missing, cannot start the application")
	}

	if grpcServerPort == "" {
		log.Fatal("The grpc server port is missing, cannot start the application")
	}

	tmdbService := service.NewTmdbService(apiKey)
	server := grpc.NewServer(tmdbService, grpcServerPort)
	log.Fatal(server.StartServer())
}
