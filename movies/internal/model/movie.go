package model

import "net/http"

type Movie struct {
	VoteAverage      float64 `json:"vote_average"`
	ID               int     `json:"id"`
	VoteCount        int     `json:"vote_count"`
	ReleaseDate      string  `json:"release_date"`
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	Title            string  `json:"title"`
	GenreIds         []int   `json:"genre_ids"`
	Popularity       float64 `json:"popularity"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	PosterPath       string  `json:"poster_path"`
	Overview         string  `json:"overview"`
	Video            bool    `json:"video"`
}

func (m Movie) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Movies []Movie

func (m Movies) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
