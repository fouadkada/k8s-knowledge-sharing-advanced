package model

import "net/http"

type Result struct {
	Movies          Movies `json:"movies"`
	Recommendations Movies `json:"recommendations"`
}

func (res Result) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
