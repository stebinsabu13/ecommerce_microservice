package client

import (
	"context"

	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/client/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/config"
	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type productClient struct {
	Server pb.ProductServiceClient
}

func InitProdClient(c *config.Config) (pb.ProductServiceClient, error) {
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

func (c *productClient) Prodetail(prodetailid uint) (*pb.ProductDetailResponse, error) {
	res, err := c.Server.OrderProductDetail(context.Background(), &pb.OrderProdDetailRequest{
		Productdetailid: uint32(prodetailid),
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *productClient) UpdateStock(prodetailid uint, quantity uint) error {
	if _, err := c.Server.UpdateStock(context.Background(), &pb.UpdateStockRequest{
		Prodetailid: uint32(prodetailid),
		Quantity:    uint32(quantity),
	}); err != nil {
		return err
	}
	return nil
}
