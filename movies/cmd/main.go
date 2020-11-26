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
	HttpServerPort = "HTTP_PORT"
)

func main() {
	log.SetPrefix("MOVIES")

	_ = viper.BindEnv(HttpServerPort)

	httpServerPort := viper.GetString(HttpServerPort)

	if httpServerPort == "" {
		log.Fatal("The moviesHttpServer server port is missing, cannot start the application")
	}

	client, err := grpc.NewClient(context.Background(), "localhost:3000", "localhost:4000")
	if err != nil {
		log.Fatal("Cannot connect to the searching service, cannot start the application")
	}

	moviesService := service.NewMovieService(client)
	moviesHttpServer := http.NewMoviesServer(httpServerPort, moviesService)
	log.Fatal(moviesHttpServer.StartServer())
}
