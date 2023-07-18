package domain

import "time"

type Admin struct {
	ID        uint      `json:"id" gorm:"primarykey;auto_increment"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	MobileNum string    `json:"mobilenum" gorm:"uniqueIndex;not null"`
	Password  string    `json:"password" gorm:"not null"`
}
