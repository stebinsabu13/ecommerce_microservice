package interfaces

import (
	"context"

	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/pb"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/utils"
)

type ProductClient interface {
	FindAllProducts(context.Context, utils.Pagination) (*pb.ListAllProductResponse, error)
	FindProductById(context.Context, string, utils.Pagination) (*pb.ListProductDetailResponse, error)
	AddProduct(context.Context, utils.AddProduct) (*pb.AddProductResponse, error)
	AddProductDetail(context.Context, utils.AddProductDetail) (*pb.AddProductResponse, error)
}
