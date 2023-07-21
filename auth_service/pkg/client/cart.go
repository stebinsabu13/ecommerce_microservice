package client

import (
	"context"

	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/client/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/config"
	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type cartClient struct {
	Server pb.CartServiceClient
}

func InitCartClient(c *config.Config) (pb.CartServiceClient, error) {
	cartcc, carterr := grpc.Dial(c.Cart_service, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if carterr != nil {
		return nil, carterr
	}
	return pb.NewCartServiceClient(cartcc), nil
}

func NewCartClient(server pb.CartServiceClient) interfaces.CartClient {
	return &cartClient{
		Server: server,
	}
}

func (c *cartClient) CreateCart(userid uint) error {
	_, err := c.Server.CreateCart(context.Background(), &pb.CreateCartRequest{
		Userid: int32(userid),
	})
	if err != nil {
		return err
	}
	return nil
}
