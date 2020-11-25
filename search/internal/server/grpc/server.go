package grpc

import (
	"fmt"
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/search/internal/server"
	"github.com/fouadkada/k8s-knowledge-sharing-advanced/search/pkg/proto"
	"google.golang.org/grpc"
	"net"
)

type srv struct {
	proto proto.SearchServer
	port  string
}

func NewServer(proto proto.SearchServer, port string) server.Server {
	return &srv{
		proto: proto,
		port:  port,
	}
}

func (s *srv) StartServer() error {
	grpcServer := grpc.NewServer()
	proto.RegisterSearchServer(grpcServer, s.proto)
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return err
	}
	return grpcServer.Serve(listen)
}
