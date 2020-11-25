package service

import (
	"bitbucket.lab.dynatrace.org/users/fouad.alkada/repos/k8s-knowledge-sharing-advanced/search/pkg/proto"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type tmdb struct {
	apiKey        string
	pbUnmarshaler jsonpb.Unmarshaler
}

func NewTmdbService(apiKey string) *tmdb {
	return &tmdb{
		apiKey: apiKey,
		pbUnmarshaler: jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		},
	}
}

func (t *tmdb) Search(ctx context.Context, movie *proto.Movie) (*proto.Results, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	endpoint := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=%s&language=en-US&query=%s&page=1", t.apiKey, url.QueryEscape(movie.Title))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not create request for the 3rd party movies API")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, status.Error(codes.Internal, "got an error when querying 3rd party movies API")
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, status.Error(codes.Internal, "error reading response from 3rd party movies API")
	}
	tmdbResResp := tmdbResultResponse{}
	err = json.Unmarshal(b, &tmdbResResp)
	if err != nil {
		return nil, status.Error(codes.Internal, "error transforming response from 3rd party movies API to object")
	}

	res := proto.Results{}
	err = t.pbUnmarshaler.Unmarshal(bytes.NewBuffer(b), &res)
	if err != nil {
		return nil, status.Error(codes.Internal, "error transforming go struct to protobuf")
	}

	return &res, nil
}

type tmdbResultResponse struct {
	Page       int `json:"-"`
	TotalPages int `json:"-"`
	Results    []struct {
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
	} `json:"results"`
	TotalResults int `json:"-"`
}
