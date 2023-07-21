package interfaces

import "github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/pb"

type CartClient interface {
	FindCartById(userid uint) (*pb.CartResponse, error)
	Findcartitems(id uint) ([]*pb.CartItem, error)
	UpdateCart(*pb.CartResponse) (*pb.UpdateResponse, error)
	DeleteCart(uint32) (*pb.UpdateResponse, error)
}
