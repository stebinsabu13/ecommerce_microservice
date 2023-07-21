package client

import (
	"context"
	"fmt"

	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/client/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/config"
	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authClient struct {
	Server pb.AuthServiceClient
}

func InitauthClient(c *config.Config) (pb.AuthServiceClient, error) {
	authcc, autherr := grpc.Dial(c.AuthService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if autherr != nil {
		return nil, autherr
	}
	return pb.NewAuthServiceClient(authcc), nil
}

func NewAuthClient(server pb.AuthServiceClient) interfaces.AuthClient {
	return &authClient{
		Server: server,
	}
}

func (c *authClient) GetAddress(addressid uint) (*pb.UserAddress, error) {
	fmt.Println("inside client", addressid)
	res, err := c.Server.GetAddress(context.Background(), &pb.GetAddressRequest{
		Addressid: uint32(addressid),
	})
	if err != nil {
		return nil, err
	}
	return res.Address[0], nil
}
