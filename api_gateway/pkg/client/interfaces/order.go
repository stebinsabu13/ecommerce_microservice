package interfaces

import (
	"context"

	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/pb"
)

type OrderClient interface {
	AddtoOrders(context.Context, uint, uint, uint, string) (*pb.AddOrderResponse, error)
	Orders(context.Context, uint) (*pb.ShowOrdersResponse, error)
	OrderDetail(uint) (*pb.ShowOrderDetailResponse, error)
}
