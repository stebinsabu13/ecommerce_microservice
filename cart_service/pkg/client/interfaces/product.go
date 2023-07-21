package interfaces

import "github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/pb"

type ProductClient interface {
	FindProductDetail(string) (*pb.ProductDetailResponse, error)
	FindCartProductDetail(string) (*pb.CartProdDetailResponse, error)
}
