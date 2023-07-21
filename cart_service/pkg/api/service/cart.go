package service

import (
	"context"
	"errors"
	"net/http"

	client "github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/client/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/domain"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/pb"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/repository/interfaces"
)

type CartService struct {
	Repo    interfaces.CartRepo
	ProdCli client.ProductClient
	pb.UnimplementedCartServiceServer
}

func NewCartService(repo interfaces.CartRepo, prodcli client.ProductClient) pb.CartServiceServer {
	return &CartService{
		Repo:    repo,
		ProdCli: prodcli,
	}
}

func (c *CartService) CreateCart(ctx context.Context, req *pb.CreateCartRequest) (*pb.CreateCartResponse, error) {
	err := c.Repo.CreateCart(uint(req.Userid))
	if err != nil {
		return nil, err
	}
	return &pb.CreateCartResponse{
		Status: http.StatusCreated,
	}, nil
}

func (c *CartService) ViewCart(ctx context.Context, req *pb.ViewCartRequest) (*pb.ViewCartResponse, error) {
	cart, err2 := c.Repo.FindCartById(uint(req.Userid))
	cartdetail, err1 := c.Repo.ViewCart(uint(cart.ID))
	err := errors.Join(err1, err2)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(cartdetail); i++ {
		cartprodetail, err := c.ProdCli.FindCartProductDetail(cartdetail[i].ProductDetailID)
		if err != nil {
			return nil, err
		}
		cartdetail[i].BrandName = cartprodetail.BrandName
		cartdetail[i].Colour = cartprodetail.Colour
		cartdetail[i].Image = cartprodetail.Image
		cartdetail[i].ModelName = cartprodetail.ModelName
		cartdetail[i].Percentage = cartprodetail.Percentage
		cartdetail[i].Price = cartprodetail.Price
		cartdetail[i].Size = cartprodetail.Size
	}
	return &pb.ViewCartResponse{
		Status:     http.StatusOK,
		Grandtotal: int32(cart.GrandTotal),
		Items:      cartdetail,
	}, nil
}

func (c *CartService) AddtoCart(ctx context.Context, req *pb.AddCartRequest) (*pb.AddCartResponse, error) {
	cart, err := c.Repo.FindCartById(uint(req.Userid))
	if err != nil {
		return nil, err
	}
	productdetail, err1 := c.ProdCli.FindProductDetail(req.Productid)
	if err1 != nil {
		return nil, err1
	}
	if productdetail.Stock <= 0 {
		return nil, errors.New("out of stock")
	} else {
		exsistitem, err2 := c.Repo.FindProductExsist(req.Productid, cart.ID)
		if err2 != nil {
			return nil, err2
		}
		discount := (int(productdetail.Discount) * int(productdetail.Price)) / 100
		if exsistitem.ID != 0 {
			exsistitem.Quantity += 1
			exsistitem.Total = exsistitem.Quantity * (uint(productdetail.Price) - uint(discount))
			if err := c.Repo.UpdateCartitem(exsistitem); err != nil {
				return nil, err
			}
		} else {
			item := domain.CartItem{
				CartID:          cart.ID,
				ProductDetailID: uint(productdetail.Id),
				Quantity:        1,
				Total:           uint(productdetail.Price) - uint(discount),
			}
			if err := c.Repo.AddNewitem(item); err != nil {
				return nil, err
			}
		}
	}
	return &pb.AddCartResponse{
		Status:   http.StatusOK,
		Response: "item added to cart",
	}, nil
}

func (c *CartService) FindCart(ctx context.Context, req *pb.CartRequest) (*pb.CartResponse, error) {
	cart, err := c.Repo.FindCartById(uint(req.Userid))
	if err != nil {
		return nil, err
	}
	return &pb.CartResponse{
		Id:         uint32(cart.ID),
		Userid:     uint32(cart.UserID),
		Grandtotal: int32(cart.GrandTotal),
	}, nil
}

func (c *CartService) CartItems(ctx context.Context, req *pb.CartItemRequest) (*pb.CartItemResponse, error) {
	items, err := c.Repo.CartItemsOrder(req)
	if err != nil {
		return nil, err
	}
	return &pb.CartItemResponse{
		Items: items,
	}, nil
}

func (c *CartService) UpdateCart(ctx context.Context, req *pb.CartResponse) (*pb.UpdateResponse, error) {
	cart := domain.Cart{
		ID:         uint(req.Id),
		UserID:     uint(req.Userid),
		GrandTotal: int(req.Grandtotal),
	}
	res, err := c.Repo.UpdateCart(cart)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{
		Grandtotal: int32(res),
	}, nil
}

func (c *CartService) DeleteCart(ctx context.Context, req *pb.CartItemRequest) (*pb.UpdateResponse, error) {
	res, err := c.Repo.DelectCart(req)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{
		Grandtotal: int32(res),
	}, nil
}
