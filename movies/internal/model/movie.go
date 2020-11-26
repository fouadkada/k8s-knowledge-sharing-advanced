package model

import "net/http"

type Movie struct {
	VoteAverage  float32 `json:"vote_average"`
	ID           string  `json:"id"`
	ReleaseDate  string  `json:"release_date"`
	BackdropPath string  `json:"backdrop_path"`
	Title        string  `json:"title"`
	PosterPath   string  `json:"poster_path"`
	Overview     string  `json:"overview"`
}

func (m Movie) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Movies []Movie

func (m Movies) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
