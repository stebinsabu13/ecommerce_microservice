package client

import (
	"context"

	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/client/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/pb"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/service"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/utils"
)

type prodClient struct {
	Server pb.ProductServiceClient
}

func NewProductClient(service service.Clients) interfaces.ProductClient {
	return &prodClient{
		Server: service.Prodcli,
	}
}

func (c *prodClient) FindAllProducts(ctx context.Context, pagination utils.Pagination) (*pb.ListAllProductResponse, error) {
	res, err := c.Server.ListAllProduct(ctx, &pb.ListAllProductRequest{
		Page:  int32(pagination.Offset),
		Limit: int32(pagination.Limit),
	})
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *prodClient) AddProduct(ctx context.Context, body utils.AddProduct) (*pb.AddProductResponse, error) {
	res, err := c.Server.AddProduct(ctx, &pb.AddProductRequest{
		ModelName:  body.ModelName,
		Image:      body.Image,
		BrandID:    uint32(body.BrandID),
		CategoryID: uint32(body.CategoryID),
	})
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *prodClient) FindProductById(ctx context.Context, id string, pagination utils.Pagination) (*pb.ListProductDetailResponse, error) {
	res, err := c.Server.ListProductDetails(ctx, &pb.ListProductDetailsRequest{
		Page:      int32(pagination.Offset),
		Limit:     int32(pagination.Limit),
		Productid: id,
	})
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *prodClient) AddProductDetail(ctx context.Context, body utils.AddProductDetail) (*pb.AddProductResponse, error) {
	res, err := c.Server.AddProductDetails(ctx, &pb.AddProdDetailRequest{
		Price:             uint32(body.Price),
		Stock:             uint32(body.Stock),
		AvailableSizeID:   uint32(body.AvailableSizeID),
		AvailableColourID: uint32(body.AvailableColourID),
		ProductID:         uint32(body.ProductID),
		DiscountID:        uint32(body.DiscountID),
	})
	if err != nil {
		return res, err
	}
	return res, nil
}
