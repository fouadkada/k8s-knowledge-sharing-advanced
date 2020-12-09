package main

import (
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/search/internal/server/grpc"
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/search/internal/service"
	"github.com/spf13/viper"
	"log"
)

const (
	APIKey = "TMDB_API_KEY"
)

func main() {
	_ = viper.BindEnv(APIKey)
	apiKey := viper.GetString(APIKey)

	if apiKey == "" {
		log.Fatal("TMDB API Key is missing, cannot start the application")
	}

	tmdbService := service.NewTmdbService(apiKey)
	server := grpc.NewServer(tmdbService, "3000")
	log.Fatal(server.StartServer())
}
