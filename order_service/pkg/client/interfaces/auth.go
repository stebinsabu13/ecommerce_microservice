package interfaces

import "github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/pb"

type AuthClient interface {
	GetAddress(uint) (*pb.UserAddress, error)
}
