package http

import (
	"bitbucket.lab.dynatrace.org/users/fouad.alkada/repos/k8s-knowledge-sharing-advanced/ratings/internal/server"
	"bitbucket.lab.dynatrace.org/users/fouad.alkada/repos/k8s-knowledge-sharing-advanced/ratings/internal/service"
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
	r.Get("/{title}", rest.search)

	return http.ListenAndServe(":"+s.Port, r)
}
