package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID     uint
	TotalPrice float64
	Status     string // "pending", "paid", "delivered", "cancelled"
	PaidAt     *string
	OrderItems []OrderItem
}
