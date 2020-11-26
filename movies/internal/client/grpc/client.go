package grpc

import (
	"context"
	"encoding/json"
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/movies/internal/model"
	rProto "github.com/fouadkada/k8s-knowledge-sharing-advanced/recommendation/pkg/proto"
	sProto "github.com/fouadkada/k8s-knowledge-sharing-advanced/search/pkg/proto"
	"github.com/gogo/protobuf/jsonpb"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

type Client interface {
	Search(ctx context.Context, title string) (model.Movies, error)
	GetRecommendations(ctx context.Context, movieID string) (model.Movies, error)
}

type client struct {
	searchClient         sProto.SearchClient
	recommendationClient rProto.RecommendationClient
	marshaler            jsonpb.Marshaler
}

func NewClient(ctx context.Context, searchURI, recommendationURI string) (Client, error) {
	sConn, err := connect(ctx, searchURI)
	if err != nil {
		return &client{}, err
	}

	rConn, err := connect(ctx, recommendationURI)
	if err != nil {
		return &client{}, err
	}

	return &client{
		searchClient:         sProto.NewSearchClient(sConn),
		recommendationClient: rProto.NewRecommendationClient(rConn),
		marshaler: jsonpb.Marshaler{
			EmitDefaults: false,
		}}, nil
}

func (c *client) Search(ctx context.Context, title string) (model.Movies, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resp, err := c.searchClient.Search(ctx, &sProto.Movie{Title: title})
	if err != nil {
		return nil, err
	}

	if len(resp.Results) == 0 {
		return model.Movies{}, nil
	}

	str, err := c.marshaler.MarshalToString(resp)
	if err != nil {
		return nil, err
	}

	var res response
	err = json.Unmarshal([]byte(str), &res)
	if err != nil {
		return nil, err
	}
	return res.Results, nil
}

func (c *client) GetRecommendations(ctx context.Context, movieID string) (model.Movies, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	mID, err := strconv.Atoi(movieID)
	if err != nil {
		return nil, err
	}
	resp, err := c.recommendationClient.GetRecommendations(ctx, &rProto.Movie{MovieID: int32(mID)})
	if err != nil {
		return nil, err
	}

	if len(resp.Results) == 0 {
		return model.Movies{}, nil
	}

	str, err := c.marshaler.MarshalToString(resp)
	if err != nil {
		return nil, err
	}

	var res response
	err = json.Unmarshal([]byte(str), &res)
	if err != nil {
		return nil, err
	}
	return res.Results, nil
}

func connect(ctx context.Context, uri string) (conn *grpc.ClientConn, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return grpc.DialContext(ctx, uri, grpc.WithInsecure())
}

type response struct {
	Results model.Movies `json:"results"`
}
