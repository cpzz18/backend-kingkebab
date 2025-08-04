package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Role      string `gorm:"default:customer"`
	Balance   float64
	CartItems []CartItem
	Orders    []Order
	Topups    []Topup
	CreatedAt time.Time
	UpdatedAt time.Time
}
