package interfaces

import (
	"context"

	"github.com/stebinsabu13/ecommerce_microservice/product_service/pkg/domain"
	"github.com/stebinsabu13/ecommerce_microservice/product_service/pkg/pb"
)

type ProductRepo interface {
	FindAllProducts(context.Context, *pb.ListAllProductRequest) ([]*pb.Product, error)
	FindProductById(context.Context, *pb.ListProductDetailsRequest) ([]*pb.ProdDetail, error)
	AddProduct(ctx context.Context, product domain.Product) error
	AddProductDetail(context.Context, domain.ProductDetails) error
	CartProductDetail(context.Context, *pb.CartProdDetailRequest) (*pb.CartProdDetailResponse, error)
	FindProductDetailById(id string) (domain.ProductDetails, int, error)
	UpdateStock(*pb.StockRequest) error
	CheckStock(*pb.StockRequest) error
}
