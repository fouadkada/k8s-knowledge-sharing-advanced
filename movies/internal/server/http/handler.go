package http

import (
	"bitbucket.lab.dynatrace.org/users/fouad.alkada/repos/k8s-knowledge-sharing-advanced/ratings/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
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
	title := chi.URLParam(r, "title")
	h.moviesService.Search(r.Context(), title)
	//if err != nil {
	//	respond(w, r, NewInternalServerError(err))
	//	log.Printf("error fetching title: [%s] information with error: [%#v]", title, err)
	//	return
	//}
	//respond(w, r, rating)
}
