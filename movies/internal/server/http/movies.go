package http

import (
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/movies/internal/server"
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/movies/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
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
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	rest := handler{
		moviesService: s.MoviesService,
	}
	r.Get("/", rest.heartbeat)
	r.Get("/search", rest.search)
	r.Get("/recommendations/{movie_id}", rest.recommendation)

	return http.ListenAndServe(":"+s.Port, r)
}
