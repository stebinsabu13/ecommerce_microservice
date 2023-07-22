package interfaces

import "github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/pb"

type ProductClient interface {
	Prodetail(uint) (*pb.ProductDetailResponse, error)
	UpdateStock(uint, uint) error
	CheckStock(uint, uint) error
}
