package http

import (
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/movies/internal/server"
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/movies/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

type srv struct {
	Port          string
	MoviesService service.Movie
}

func NewMoviesServer(port string, moviesService service.Movie) server.Server {
	return &srv{
		Port:          port,
		MoviesService: moviesService,
	}
}

func (s *srv) StartServer() error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	rest := handler{
		moviesService: s.MoviesService,
	}
	r.Get("/", rest.heartbeat)
	r.Get("/search", rest.search)
	r.Get("/recommendation/{movie_id}", rest.recommendation)

	return http.ListenAndServe(":"+s.Port, r)
}
