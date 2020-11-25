package main

import (
	"bitbucket.lab.dynatrace.org/users/fouad.alkada/repos/k8s-knowledge-sharing-advanced/ratings/internal/server/http"
	"bitbucket.lab.dynatrace.org/users/fouad.alkada/repos/k8s-knowledge-sharing-advanced/ratings/internal/service"
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

	moviesService := service.NewMovieService()
	moviesHttpServer := http.NewMoviesServer(httpServerPort, moviesService)
	log.Fatal(moviesHttpServer.StartServer())
}
