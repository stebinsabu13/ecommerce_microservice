package service

import (
	"context"
	"fmt"
	"net/http"

	client "github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/client/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/pb"
	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/repository/interfaces"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type orderService struct {
	Repo       interfaces.OderRepo
	CartClient client.CartClient
	ProdClient client.ProductClient
	AuthClient client.AuthClient
	pb.UnimplementedOrderServiceServer
}

func NewOrderService(repo interfaces.OderRepo, cartclient client.CartClient, prodclient client.ProductClient,
	authclient client.AuthClient) pb.OrderServiceServer {
	return &orderService{
		Repo:       repo,
		CartClient: cartclient,
		ProdClient: prodclient,
		AuthClient: authclient,
	}
}

func (c *orderService) AddtoOrders(ctx context.Context, req *pb.AddOrderRequest) (*pb.AddOrderResponse, error) {
	fmt.Println("inside service addtoorder:", req)
	couponid, cart, err := c.ValidateCoupon(uint(req.Userid), req.Couponcode)
	if err != nil {
		return nil, err
	}
	if req.Paymentid == 2 {
		body, err1 := c.Razorpayment(cart, couponid)
		if err1 != nil {
			return nil, err1
		}
		return body, nil
	} else {
		body, err1 := c.ServiceAddtoOrders(uint(req.Addressid), uint(req.Paymentid), cart, couponid)
		if err1 != nil {
			return nil, err1
		}
		return body, nil
	}
}

func (c *orderService) GetOrders(ctx context.Context, req *pb.ShowOrdersRequest) (*pb.ShowOrdersResponse, error) {
	var orderresponse []*pb.ShowOrders
	orders, err := c.Repo.Orders(uint(req.Userid))
	if err != nil {
		return nil, err
	}
	fmt.Println("orders:", orders)
	for i := 0; i < len(orders); i++ {
		addres, err := c.AuthClient.GetAddress(orders[i].AddressID)
		fmt.Println("addres:", addres)
		if err != nil {
			return nil, err
		}
		timestamp := timestamppb.New(orders[i].PlacedDate)
		orderresponse = append(orderresponse, &pb.ShowOrders{
			ID:         uint32(orders[i].ID),
			HouseName:  addres.HouseName,
			Street:     addres.Street,
			City:       addres.City,
			State:      addres.State,
			Country:    addres.Country,
			Pincode:    addres.Pincode,
			PlacedDate: timestamp,
			Mode:       orders[i].Mode,
			GrandTotal: uint32(orders[i].GrandTotal),
		})
		fmt.Println("inside for loop:", orderresponse)
	}
	fmt.Println("outside for loop:", orderresponse)
	return &pb.ShowOrdersResponse{
		Status: http.StatusOK,
		Orders: orderresponse,
	}, err
}
