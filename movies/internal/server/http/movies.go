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

const (
	serviceVersionCustomHeader = "X-SERVICE-VERSION"
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

func (s *srv) StartServer(serviceVersion string) error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{serviceVersionCustomHeader},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(serviceVersionCustomHeader, serviceVersion)
			next.ServeHTTP(w, r)
		})
	})

	rest := handler{
		moviesService: s.MoviesService,
	}
	r.Get("/", rest.heartbeat)
	r.Get("/search", rest.search)
	r.Get("/recommendations/{movie_id}", rest.recommendation)

	return http.ListenAndServe(":"+s.Port, r)
}
