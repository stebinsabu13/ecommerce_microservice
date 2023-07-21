package client

import (
	"context"

	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/client/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/config"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type productClient struct {
	Server pb.ProductServiceClient
}

func InitClient(c *config.Config) (pb.ProductServiceClient, error) {
	prodcc, proderr := grpc.Dial(c.ProductService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if proderr != nil {
		return nil, proderr
	}
	return pb.NewProductServiceClient(prodcc), nil

}

func NewProdClient(server pb.ProductServiceClient) interfaces.ProductClient {
	return &productClient{
		Server: server,
	}
}

func (c *productClient) FindProductDetail(prodid string) (*pb.ProductDetailResponse, error) {
	res, err := c.Server.ProductDetail(context.Background(), &pb.CartProdDetailRequest{
		Productdetailid: prodid,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *productClient) FindCartProductDetail(prodid string) (*pb.CartProdDetailResponse, error) {
	res, err := c.Server.CartProductDetails(context.Background(), &pb.CartProdDetailRequest{
		Productdetailid: prodid,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
