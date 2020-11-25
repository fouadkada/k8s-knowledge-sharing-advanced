package http

import (
	"github.com/go-chi/render"
	"net/http"
)

func respond(w http.ResponseWriter, r *http.Request, obj render.Renderer) {
	render.Render(w, r, obj)
}
