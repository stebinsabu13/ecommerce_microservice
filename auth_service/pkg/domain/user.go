package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint      `json:"id" gorm:"primarykey;auto_increment"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"not null"`
	FirstName   string    `json:"firstname" gorm:"not null"`
	LastName    string    `json:"lastname" gorm:"not null"`
	Email       string    `json:"email" gorm:"uniqueIndex;not null"`
	MobileNum   string    `json:"mobilenum" gorm:"uniqueIndex;not null"`
	Password    string    `json:"password" gorm:"not null"`
	Block       bool      `json:"block" gorm:"default:false"`
	Verified    bool      `json:"verified" gorm:"default:false"`
	ReferalCode string    `json:"referalcode" gorm:"uniqueIndex"`
}

type Address struct {
	gorm.Model
	HouseName string `json:"housename" gorm:"not null"`
	Street    string `json:"street" gorm:"not null"`
	City      string `json:"city" gorm:"not null"`
	State     string `json:"state" gorm:"not null"`
	Country   string `json:"country" gorm:"not null"`
	Pincode   string `json:"pincode" gorm:"not null"`
	UserID    uint   `json:"userid"`
}
