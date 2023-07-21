package interfaces

import (
	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/domain"
	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/utils"
)

type OderRepo interface {
	FindCoupon(string) (domain.Coupon, error)
	CouponUsage(uint, uint32) (uint, error)
	UpdateCouponUsage(uint, uint) error
	AddtoOrders(domain.Order) (uint, error)
	AddOrderItem(domain.OrderDetails) error
	CleanUp(uint) error
	Orders(uint) ([]utils.ResOrders, error)
}
