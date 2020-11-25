package grpc

import (
	"fmt"
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/recommendation/internal/server"
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/recommendation/pkg/proto"
	"google.golang.org/grpc"
	"net"
)

type srv struct {
	proto proto.RecommendationServer
	port  string
}

func NewServer(proto proto.RecommendationServer, port string) server.Server {
	return &srv{
		proto: proto,
		port:  port,
	}
}

func (s *srv) StartServer() error {
	grpcServer := grpc.NewServer()
	proto.RegisterRecommendationServer(grpcServer, s.proto)
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return err
	}
	return grpcServer.Serve(listen)
}
