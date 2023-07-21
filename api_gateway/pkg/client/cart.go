package client

import (
	"context"

	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/client/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/pb"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/service"
)

type cartClient struct {
	Server pb.CartServiceClient
}

func NewCartClient(service service.Clients) interfaces.CartClient {
	return &cartClient{
		Server: service.Cartcli,
	}
}

func (c *cartClient) ViewCart(userid uint) (*pb.ViewCartResponse, error) {
	res, err := c.Server.ViewCart(context.Background(), &pb.ViewCartRequest{
		Userid: int32(userid),
	})
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *cartClient) AdditemtoCart(productid string, userid uint) (*pb.AddCartResponse, error) {
	res, err := c.Server.AddtoCart(context.Background(), &pb.AddCartRequest{
		Productid: productid,
		Userid:    uint32(userid),
	})
	if err != nil {
		return res, err
	}
	return res, nil
}
