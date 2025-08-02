package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Role      string  // "admin" atau "cs"
	Balance   float64
	CartItems []CartItem
	Orders    []Order
	Topups    []Topup
}
