package client

import (
	"context"
	"fmt"

	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/client/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/pb"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/service"
)

type orderClient struct {
	Server pb.OrderServiceClient
}

func NewOrderClient(services service.Clients) interfaces.OrderClient {
	return &orderClient{
		Server: services.Ordercli,
	}
}

func (c *orderClient) AddtoOrders(ctx context.Context, addressid uint, paymentid uint, userid uint, code string) (*pb.AddOrderResponse, error) {
	fmt.Println("Inside client of add to order:req=", addressid, paymentid, userid, code)
	res, err := c.Server.AddtoOrders(ctx, &pb.AddOrderRequest{
		Addressid:  uint32(addressid),
		Paymentid:  uint32(paymentid),
		Userid:     uint32(userid),
		Couponcode: code,
	})
	fmt.Println("inside client of addtoorder:", res)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (c *orderClient) Orders(ctx context.Context, userid uint) (*pb.ShowOrdersResponse, error) {
	res, err := c.Server.GetOrders(ctx, &pb.ShowOrdersRequest{
		Userid: uint32(userid),
	})
	fmt.Println("inside client of showorder:", res)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (c *orderClient) OrderDetail(orderid uint) (*pb.ShowOrderDetailResponse, error) {
	res, err := c.Server.GetOrderDetails(context.Background(), &pb.ViewOrderDetailResquest{
		Orderid: uint32(orderid),
	})
	fmt.Println("inside client of showorderdetail:", res)
	if err != nil {
		return nil, err
	}
	return res, err
}
