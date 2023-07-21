package repository

import (
	"errors"
	"fmt"

	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/domain"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/pb"
	"github.com/stebinsabu13/ecommerce_microservice/cart_service/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type cartRepo struct {
	DB *gorm.DB
}

func NewCartRepo(db *gorm.DB) interfaces.CartRepo {
	return &cartRepo{
		DB: db,
	}
}

func (c *cartRepo) CreateCart(userid uint) error {
	if err := c.DB.Create(&domain.Cart{UserID: userid}).Error; err != nil {
		return err
	}
	return nil
}

func (c *cartRepo) ViewCart(cartid uint) ([]*pb.CartItems, error) {
	var cartdetail []*pb.CartItems
	query := `select product_detail_id,quantity,total from cart_items where cart_id=?`
	err := c.DB.Raw(query, cartid).Scan(&cartdetail).Error
	if err != nil {
		return cartdetail, err
	}
	return cartdetail, nil
}

func (c *cartRepo) FindCartById(userid uint) (domain.Cart, error) {
	var cart domain.Cart
	if err := c.DB.Where("user_id=?", userid).Find(&cart); err.RowsAffected == 0 {
		return cart, errors.New("no row to return")
	}
	return cart, nil
}

func (c *cartRepo) FindProductExsist(id string, cartid uint) (domain.CartItem, error) {
	var exsistitem domain.CartItem
	result := c.DB.Where("product_detail_id=$1 and cart_id=$2", id, cartid).Find(&exsistitem)
	if result.Error != nil {
		return exsistitem, result.Error
	}
	return exsistitem, nil
}

func (c *cartRepo) UpdateCartitem(exsistitem domain.CartItem) error {
	var grandtotal int
	tx := c.DB.Begin()
	if err := tx.Model(&domain.CartItem{}).Where("id=?", exsistitem.ID).UpdateColumns(&exsistitem).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&domain.CartItem{}).Where("cart_id=?", exsistitem.CartID).Select("COALESCE(SUM(total), 0)").Scan(&grandtotal).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&domain.Cart{}).Where("id=?", exsistitem.CartID).UpdateColumn("grand_total", grandtotal).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *cartRepo) AddNewitem(item domain.CartItem) error {
	var grandtotal int
	tx := c.DB.Begin()
	if err := tx.Create(&item).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&domain.CartItem{}).Where("cart_id=?", item.CartID).Select("SUM(total)").Scan(&grandtotal).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&domain.Cart{}).Where("id=?", item.CartID).UpdateColumn("grand_total", grandtotal).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *cartRepo) CartItemsOrder(req *pb.CartItemRequest) ([]*pb.CartItem, error) {
	var res []*pb.CartItem
	if err := c.DB.Model(&domain.CartItem{}).Where("cart_id=?", req.Cartid).Select("cart_id,product_detail_id,quantity").Scan(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (c *cartRepo) UpdateCart(cart domain.Cart) (int, error) {
	if err := c.DB.Save(&cart).Error; err != nil {
		return cart.GrandTotal, err
	}
	return cart.GrandTotal, nil
}

func (c *cartRepo) DelectCart(req *pb.CartItemRequest) (int, error) {
	tx := c.DB.Begin()
	query := `delete from cart_items where cart_id=$1`
	if err := tx.Exec(query, req.Cartid).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Model(&domain.Cart{}).Where("id=?", req.Cartid).UpdateColumn("grand_total", 0).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	fmt.Println("inside cart repo delete:")
	return 0, nil
}
