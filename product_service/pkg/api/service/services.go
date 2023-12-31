package service

import (
	"context"
	"net/http"

	"github.com/stebinsabu13/ecommerce_microservice/product_service/pkg/domain"
	"github.com/stebinsabu13/ecommerce_microservice/product_service/pkg/pb"
	"github.com/stebinsabu13/ecommerce_microservice/product_service/pkg/repository/interfaces"
)

type productService struct {
	Repo interfaces.ProductRepo
	pb.UnimplementedProductServiceServer
}

func NewProductService(repo interfaces.ProductRepo) pb.ProductServiceServer {
	return &productService{
		Repo: repo,
	}
}

func (cr *productService) ListAllProduct(ctx context.Context, req *pb.ListAllProductRequest) (*pb.ListAllProductResponse, error) {
	res, err := cr.Repo.FindAllProducts(ctx, req)
	if err != nil {
		return &pb.ListAllProductResponse{
			Status: http.StatusBadRequest,
		}, err
	}
	return &pb.ListAllProductResponse{
		Status:  http.StatusOK,
		Product: res,
	}, nil
}

func (cr *productService) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	product := domain.Product{
		ModelName:  req.ModelName,
		Image:      req.Image,
		BrandID:    uint(req.BrandID),
		CategoryID: uint(req.CategoryID),
	}
	if err := cr.Repo.AddProduct(ctx, product); err != nil {
		return &pb.AddProductResponse{
			Status: http.StatusInternalServerError,
		}, err
	}
	return &pb.AddProductResponse{
		Status:  http.StatusOK,
		Reponse: "product created",
	}, nil
}

func (cr *productService) AddProductDetails(ctx context.Context, req *pb.AddProdDetailRequest) (*pb.AddProductResponse, error) {
	prodetail := domain.ProductDetails{
		Price:             uint(req.Price),
		Stock:             uint(req.Stock),
		AvailableSizeID:   uint(req.AvailableSizeID),
		AvailableColourID: uint(req.AvailableColourID),
		ProductID:         uint(req.ProductID),
		DiscountID:        uint(req.DiscountID),
	}
	if err := cr.Repo.AddProductDetail(ctx, prodetail); err != nil {
		return &pb.AddProductResponse{
			Status: http.StatusInternalServerError,
		}, err
	}
	return &pb.AddProductResponse{
		Status:  http.StatusOK,
		Reponse: "product details created",
	}, nil
}

func (cr *productService) ListProductDetails(ctx context.Context, req *pb.ListProductDetailsRequest) (*pb.ListProductDetailResponse, error) {
	res, err := cr.Repo.FindProductById(ctx, req)
	if err != nil {
		return &pb.ListProductDetailResponse{
			Status: http.StatusBadRequest,
		}, err
	}
	return &pb.ListProductDetailResponse{
		Status:        http.StatusOK,
		ProductDetail: res,
	}, nil
}

func (cr *productService) CartProductDetails(ctx context.Context, req *pb.CartProdDetailRequest) (*pb.CartProdDetailResponse, error) {
	res, err := cr.Repo.CartProductDetail(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (cr *productService) ProductDetail(ctx context.Context, req *pb.CartProdDetailRequest) (*pb.ProductDetailResponse, error) {
	prodetail, discount, err := cr.Repo.FindProductDetailById(req.Productdetailid)
	if err != nil {
		return nil, err
	}
	return &pb.ProductDetailResponse{
		Id:       uint32(prodetail.ID),
		Stock:    int32(prodetail.Stock),
		Price:    uint32(prodetail.Price),
		Discount: int32(discount),
	}, nil
}

func (cr *productService) UpdateStock(ctx context.Context, req *pb.StockRequest) (*pb.StockResponse, error) {
	err := cr.Repo.UpdateStock(req)
	if err != nil {
		return nil, err
	}
	return &pb.StockResponse{
		Response: "stock updated",
	}, nil
}

func (cr *productService) CheckStock(ctx context.Context, req *pb.StockRequest) (*pb.StockResponse, error) {
	err := cr.Repo.CheckStock(req)
	if err != nil {
		return nil, err
	}
	return &pb.StockResponse{
		Response: "insufficient stock",
	}, nil
}
