package client

import (
	"context"

	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/client/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/config"
	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type cartClient struct {
	Server pb.CartServiceClient
}

func InitClient(c *config.Config) (pb.CartServiceClient, error) {
	cartcc, carterr := grpc.Dial(c.CartServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

func (c *cartClient) FindCartById(userid uint) (*pb.CartResponse, error) {
	res, err := c.Server.FindCart(context.Background(), &pb.CartRequest{
		Userid: uint32(userid),
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *cartClient) Findcartitems(id uint) ([]*pb.CartItem, error) {
	res, err := c.Server.CartItems(context.Background(), &pb.CartItemRequest{
		Cartid: uint32(id),
	})
	if err != nil {
		return nil, err
	}
	return res.Items, nil
}

func (c *cartClient) UpdateCart(req *pb.CartResponse) (*pb.UpdateResponse, error) {
	res, err := c.Server.UpdateCart(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *cartClient) DeleteCart(cartid uint32) (*pb.UpdateResponse, error) {
	res, err := c.Server.DeleteCart(context.Background(), &pb.CartItemRequest{
		Cartid: cartid,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
