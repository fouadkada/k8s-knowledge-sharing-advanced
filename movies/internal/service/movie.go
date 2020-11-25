package service

import (
	"bitbucket.lab.dynatrace.org/users/fouad.alkada/repos/k8s-knowledge-sharing-advanced/ratings/internal/model"
	"context"
)

type Movie interface {
	Search(ctx context.Context, title string) (model.Result, error)
}

type movie struct {
}

func NewMovieService() Movie {
	return movie{}
}

func (m movie) Search(ctx context.Context, title string) {
	panic("implement me")
}
