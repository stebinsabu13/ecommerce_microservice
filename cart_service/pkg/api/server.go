package api

import (
	"fmt"
	"net"

	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/config"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/pb"
	"google.golang.org/grpc"
)

type Server struct {
	gs   *grpc.Server
	Lis  net.Listener
	Port string
}

func NewgrpcServe(c *config.Config, service pb.CartServiceServer) (*Server, error) {
	grpcserver := grpc.NewServer()
	pb.RegisterCartServiceServer(grpcserver, service)
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		return nil, err
	}
	return &Server{
		gs:   grpcserver,
		Lis:  lis,
		Port: c.Port,
	}, nil
}

func (s *Server) Start() error {
	fmt.Println("Authentication service on:", s.Port)
	return s.gs.Serve(s.Lis)
}
