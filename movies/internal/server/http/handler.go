package http

import (
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/movies/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

type handler struct {
	moviesService service.Movie
}

func (h *handler) heartbeat(w http.ResponseWriter, r *http.Request) {
	var response = make(render.M)
	response["message"] = "I am alive 😀"

	render.JSON(w, r, response)
}

func (h *handler) search(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title == "" {
		respond(w, r, NewBadRequestError("The title parameter is required, how should I know what to look for 😏"))
		return
	}

	result, err := h.moviesService.Search(r.Context(), title)
	if err != nil {
		respond(w, r, NewInternalServerError())
		log.Printf("error fetching title: [%s] information with error: [%v]", title, err)
		return
	}
	respond(w, r, result)
}

func (h *handler) recommendation(w http.ResponseWriter, r *http.Request) {
	mID := chi.URLParam(r, "movie_id")
	if mID == "" {
		respond(w, r, NewBadRequestError("The movie_id parameter is required 😏"))
		return
	}

	result, err := h.moviesService.Recommendations(r.Context(), mID)
	if err != nil {
		respond(w, r, NewInternalServerError())
		log.Printf("error fetching recommendation for movie_id: [%s] with error: [%v]", mID, err)
		return
	}
	respond(w, r, result)
}
