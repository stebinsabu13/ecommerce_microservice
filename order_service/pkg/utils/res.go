package utils

import "time"

type ResOrders struct {
	ID         uint      `json:"id"`
	PlacedDate time.Time `json:"placeddate"`
	AddressID  uint      `json:"addressid"`
	Mode       string    `json:"mode"`
	GrandTotal uint      `json:"grandtotal"`
}
