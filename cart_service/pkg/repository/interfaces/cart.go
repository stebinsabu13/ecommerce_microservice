package interfaces

import (
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/domain"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/pb"
)

type CartRepo interface {
	ViewCart(userid uint) ([]*pb.CartItems, error)
	FindCartById(userid uint) (domain.Cart, error)
	AddNewitem(item domain.CartItem) error
	FindProductExsist(id string, cartid uint) (domain.CartItem, error)
	UpdateCartitem(exsistitem domain.CartItem) error
	CreateCart(uint) error
	CartItemsOrder(*pb.CartItemRequest) ([]*pb.CartItem, error)
	UpdateCart(domain.Cart) (int, error)
	DelectCart(*pb.CartItemRequest) (int, error)
}
