package repository

import (
	"errors"
	"time"

	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/domain"
	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/repository/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/order_service/pkg/utils"
	"gorm.io/gorm"
)

type orderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo(db *gorm.DB) interfaces.OderRepo {
	return &orderRepo{
		DB: db,
	}
}

func (c *orderRepo) FindCoupon(code string) (domain.Coupon, error) {
	var coupon domain.Coupon
	result := c.DB.Model(&domain.Coupon{}).Where("coupon_code=?", code).Find(&coupon)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return coupon, errors.New("coupon doesn't exsist")
		}
	}
	return coupon, nil
}

func (c *orderRepo) CouponUsage(couponid uint, userid uint32) (uint, error) {
	tx := c.DB.Begin()
	var useCount uint
	if err := tx.Model(&domain.CouponUsage{}).Where("coupon_id=? and user_id=?", couponid, userid).Select("usage").Scan(&useCount); err.Error != nil {
		tx.Rollback()
		return useCount, err.Error
	} else if err.RowsAffected == 0 {
		couponusage := domain.CouponUsage{
			UserID:   uint(userid),
			CouponID: couponid,
			Usage:    0,
		}
		if err1 := tx.Create(&couponusage).Error; err1 != nil {
			tx.Rollback()
			return useCount, err1
		}
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return useCount, err
	}
	return useCount, nil
}

func (c *orderRepo) UpdateCouponUsage(couponid, usecount uint) error {
	if err := c.DB.Model(&domain.CouponUsage{}).Where("coupon_id=?", couponid).UpdateColumn("usage", usecount).Error; err != nil {
		return err
	}
	return nil
}

func (c *orderRepo) AddtoOrders(order domain.Order) (uint, error) {
	var balance int
	tx := c.DB.Begin()
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return order.ID, err
	}
	if order.PaymentID == 3 {
		if err := tx.Model(&domain.Wallet{}).Select("sum(amount)").Where("user_id=?", order.UserID).Scan(&balance).Error; err != nil {
			tx.Rollback()
			return order.ID, err
		}
		if balance < int(order.GrandTotal) {
			tx.Rollback()
			return order.ID, errors.New("insufficient balance in wallet,choose different payment option")
		}
		current := time.Now()
		wallet := domain.Wallet{
			UserID:       order.UserID,
			CreditedDate: nil,
			DebitedDate:  &current,
			Amount:       -1 * int(order.GrandTotal),
		}
		if err := tx.Create(&wallet).Error; err != nil {
			tx.Rollback()
			return order.ID, err
		}
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return order.ID, err
	}
	return order.ID, nil
}

func (c *orderRepo) AddOrderItem(orderitem domain.OrderDetails) error {
	if err := c.DB.Create(&orderitem).Error; err != nil {
		return err
	}
	return nil
}

func (c *orderRepo) CleanUp(orderid uint) error {
	tx := c.DB.Begin()
	if err := tx.Where("id=?", orderid).Delete(&domain.Order{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("order_id=?", orderid).Delete(&domain.OrderDetails{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *orderRepo) Orders(userid uint) ([]utils.ResOrders, error) {
	var orders []utils.ResOrders
	query := `SELECT o.id,o.placed_date,o.grand_total,o.address_id,pm.mode from orders o
	inner join payment_modes pm on pm.id=o.payment_id where o.user_id=?`
	err := c.DB.Raw(query, userid).Scan(&orders).Error
	if err != nil {
		return orders, err
	}
	return orders, nil
}
