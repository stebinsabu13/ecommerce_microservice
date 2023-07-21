package interfaces

import "github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/pb"

type CartClient interface {
	ViewCart(uint) (*pb.ViewCartResponse, error)
	AdditemtoCart(string, uint) (*pb.AddCartResponse, error)
}
