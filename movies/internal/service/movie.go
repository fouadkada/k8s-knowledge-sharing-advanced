package service

import (
	"context"
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/movies/internal/client/grpc"
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/movies/internal/model"
)

type Movie interface {
	Search(ctx context.Context, title string) (model.Movies, error)
	Recommendations(ctx context.Context, mID string) (model.Movies, error)
}

type movie struct {
	client grpc.Client
}

func NewMovieService(client grpc.Client) Movie {
	return &movie{
		client: client,
	}
}

func (m *movie) Search(ctx context.Context, title string) (model.Movies, error) {
	res, err := m.client.Search(ctx, title)
	if err != nil {
		return model.Movies{}, err
	}
	return res, nil
}

func (m *movie) Recommendations(ctx context.Context, mID string) (model.Movies, error) {
	res, err := m.client.GetRecommendations(ctx, mID)
	if err != nil {
		return model.Movies{}, err
	}
	return res, nil
}
