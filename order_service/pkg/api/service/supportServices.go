package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/config"
	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/domain"
	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/pb"
	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/support"
)

func (c *orderService) Checkproductstock(userid uint) (*pb.CartResponse, error) {
	cart, err := c.CartClient.FindCartById(uint(userid))
	if err != nil {
		return nil, err
	}
	log.Println("validate cart:", cart)
	cartitems, err1 := c.CartClient.Findcartitems(uint(cart.Id))
	if err1 != nil {
		return nil, err1
	}
	for _, v := range cartitems {
		if err2 := c.ProdClient.CheckStock(uint(v.ProductDetailID), uint(v.Quantity)); err2 != nil {
			return nil, err2
		}
	}
	return cart, nil
}

func (c *orderService) ValidateCoupon(couponcode string, cart *pb.CartResponse) (*uint, error) {
	log.Println("couponcode:", couponcode)
	var found bool
	if couponcode != "" {
		cartitems, err1 := c.CartClient.Findcartitems(uint(cart.Id))
		if err1 != nil {
			return nil, err1
		}
		coupon, err2 := c.Repo.FindCoupon(couponcode)
		if err2 != nil {
			return nil, err2
		}
		if coupon.MinimumOrderAmount != nil && *coupon.MinimumOrderAmount > uint(cart.Grandtotal) {
			return nil, errors.New("requires a minimum amount for the coupon to apply")
		}
		if time.Now().After(coupon.ExpirationDate) {
			return nil, errors.New("the coupon had expired")
		}
		usecount, err3 := c.Repo.CouponUsage(coupon.ID, cart.Userid)
		if err3 != nil {
			return nil, err3
		}
		if usecount >= coupon.UsageLimit {
			return nil, errors.New("coupon usage limit exceeds")
		}
		if coupon.ProductID != nil {
			for _, v := range cartitems {
				if usecount >= coupon.UsageLimit {
					break
				}
				res, err := c.ProdClient.Prodetail(uint(v.ProductDetailID))
				if err != nil {
					return nil, err
				}
				if *coupon.ProductID == int(res.Productid) {
					found = true
					if coupon.CouponType == 1 {
						discount := (uint(res.Price) * coupon.Discount) / 100
						cart.Grandtotal = cart.Grandtotal - int32(discount)
					} else if coupon.CouponType == 2 {
						cart.Grandtotal = cart.Grandtotal - int32(coupon.Discount)
					}
					usecount++
				}
			}
			if !found {
				return nil, errors.New("this coupon can't be aplied for these products")
			}
		} else {
			for usecount < coupon.UsageLimit {
				if coupon.CouponType == 1 {
					discount := (uint(cart.Grandtotal) * coupon.Discount) / 100
					cart.Grandtotal = cart.Grandtotal - int32(discount)
				} else if coupon.CouponType == 2 {
					cart.Grandtotal = cart.Grandtotal - int32(coupon.Discount)
				}
				usecount++
			}
		}
		if err := c.Repo.UpdateCouponUsage(coupon.ID, usecount); err != nil {
			return nil, err
		}
		res, err := c.CartClient.UpdateCart(cart)
		if err != nil {
			return nil, err
		}
		log.Println("response after coupon applying, grandtotal=", res.Grandtotal)
		return &coupon.ID, nil
	}
	return nil, nil
}

func (c *orderService) ServiceAddtoOrders(addressid uint, paymentid uint, cart *pb.CartResponse, couponid *uint) (*pb.AddOrderResponse, error) {
	cartitems, err := c.CartClient.Findcartitems(uint(cart.Id))
	if err != nil {
		return nil, err
	}
	log.Println("Find Cartitems:", cartitems)
	order := domain.Order{
		UserID:     uint(cart.Userid),
		PlacedDate: time.Now(),
		AddressID:  addressid,
		PaymentID:  paymentid,
		GrandTotal: uint(cart.Grandtotal),
		CouponID:   couponid,
	}
	orderid, err1 := c.Repo.AddtoOrders(order)
	if err1 != nil {
		// panic(err1)
		return nil, err1
	}
	defer c.DeferredActions(orderid)
	for _, v := range cartitems {
		orderitem := domain.OrderDetails{
			OrderID:         orderid,
			OrderStatusID:   3,
			DeliveredDate:   nil,
			CancelledDate:   nil,
			ProductDetailID: uint(v.ProductDetailID),
			Quantity:        uint(v.Quantity),
		}
		if err := c.ProdClient.UpdateStock(uint(v.ProductDetailID), uint(v.Quantity)); err != nil {
			panic(err)
		}
		if err := c.Repo.AddOrderItem(orderitem); err != nil {
			panic(err)
		}
	}
	res, err2 := c.CartClient.DeleteCart(cart.Id)
	if err2 != nil {
		panic(err2)
	}
	log.Println("Total after deleting cart", res.Grandtotal)
	return &pb.AddOrderResponse{
		Status:   http.StatusOK,
		Response: "Product added",
	}, nil
}

func (c *orderService) DeferredActions(orderid uint) (*pb.AddOrderResponse, error) {
	err := recover()
	for err != nil {
		log.Println("inside  defer", err)
		err = c.Repo.CleanUp(orderid)
	}
	return &pb.AddOrderResponse{}, fmt.Errorf("errror")
}

func (c *orderService) Razorpayment(cart *pb.CartResponse, couponid *uint) (*pb.AddOrderResponse, error) {
	log.Println("Inside razorpayment in orderservice", cart)
	var razorpayOrder *pb.AddOrderResponse
	// generate razorpay order
	//razorpay amount is caluculate on pisa for india so make the actual price into paisa
	razorPayAmount := cart.Grandtotal * 100
	razopayOrderId, err := support.GenerateRazorpayOrder(int(razorPayAmount), "test reciept")
	if err != nil {
		return razorpayOrder, err
	}
	// set all details on razopay order
	razorpayOrder.AmountToPay = uint32(cart.Grandtotal)
	razorpayOrder.RazorpayAmount = razorPayAmount

	razorpayOrder.RazorpayKey = config.GetCofig().RAZORPAYKEY

	razorpayOrder.RazorpayOrderID = razopayOrderId
	razorpayOrder.UserID = cart.Userid
	razorpayOrder.Status = http.StatusOK

	return razorpayOrder, nil
}
