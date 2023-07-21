package service

import (
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/config"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	Authcli  pb.AuthServiceClient
	Prodcli  pb.ProductServiceClient
	Cartcli  pb.CartServiceClient
	Ordercli pb.OrderServiceClient
}

func InitClient(c *config.Config) (Clients, error) {
	authcc, autherr := grpc.Dial(c.AuthService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if autherr != nil {
		return Clients{}, autherr
	}
	prodcc, proderr := grpc.Dial(c.ProductService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if proderr != nil {
		return Clients{}, proderr
	}
	cartcc, carterr := grpc.Dial(c.CartServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if carterr != nil {
		return Clients{}, carterr
	}
	ordercc, ordererr := grpc.Dial(c.OrderServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if autherr != nil {
		return Clients{}, ordererr
	}
	orderclient := pb.NewOrderServiceClient(ordercc)
	cartclient := pb.NewCartServiceClient(cartcc)
	authclient := pb.NewAuthServiceClient(authcc)
	productclient := pb.NewProductServiceClient(prodcc)
	return Clients{
		Authcli:  authclient,
		Prodcli:  productclient,
		Cartcli:  cartclient,
		Ordercli: orderclient,
	}, nil
}
