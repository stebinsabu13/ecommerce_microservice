package domain

import "time"

type Order struct {
	ID         uint      `json:"id" gorm:"primarykey;auto_increment"`
	UserID     uint      `json:"userid" gorm:"not null"`
	PlacedDate time.Time `json:"placeddate" gorm:"not null"`
	AddressID  uint      `json:"addressid" gorm:"not null"`
	PaymentID  uint      `json:"paymentid" gorm:"not null"`
	GrandTotal uint      `json:"grandtotal" gorm:"not null"`
	CouponID   *uint     `json:"couponid"`
}

type OrderDetails struct {
	ID               uint       `json:"id" gorm:"primarykey;auto_increment"`
	OrderID          uint       `json:"orderid" gorm:"not null"`
	OrderStatusID    uint       `json:"orderstatusid" gorm:"not null"`
	DeliveredDate    *time.Time `json:"delivereddate"`
	CancelledDate    *time.Time `json:"cancelleddate"`
	ReturnSubmitDate *time.Time `json:"returnsubmitdate"`
	ProductDetailID  uint       `json:"productdetailsid"`
	Quantity         uint       `json:"quantity" gorm:"not null"`
}

type OrderStatus struct {
	ID     uint   `json:"id" gorm:"primarykey;auto_increment"`
	Status string `json:"status" gorm:"not null"`
}

type Coupon struct {
	ID                 uint      `json:"id" gorm:"primarykey;auto_increment"`
	CouponCode         string    `json:"couponcode" gorm:"uniqueIndex;not null"`
	CouponType         uint      `json:"coupontype" gorm:"not null"`
	Discount           uint      `json:"discount" gorm:"not null"`
	UsageLimit         uint      `json:"usagelimit" gorm:"default:1"`
	ExpirationDate     time.Time `json:"expirationdate" gorm:"not null"`
	MinimumOrderAmount *uint     `json:"minimumorderamount"`
	ProductID          *int      `json:"productid"`
}

type CouponType struct {
	ID   uint   `json:"id" gorm:"primarykey;auto_increment"`
	Type string `json:"type" gorm:"not null"`
}

type CouponUsage struct {
	ID       uint `json:"id" gorm:"primarykey;auto_increment"`
	UserID   uint `json:"userid" gorm:"not null"`
	CouponID uint `json:"couponid" gorm:"not null"`
	Usage    uint `json:"usage" gorm:"not null"`
}

type PaymentMode struct {
	ID   uint   `json:"id" gorm:"primarykey;auto_increment"`
	Mode string `json:"mode" gorm:"not null"`
}

type Wallet struct {
	ID           uint       `json:"id" gorm:"primarykey;auto_increment"`
	UserID       uint       `json:"userid" gorm:"not null"`
	CreditedDate *time.Time `json:"crediteddate"`
	DebitedDate  *time.Time `json:"debiteddate"`
	Amount       int        `json:"amount" gorm:"not null"`
}
