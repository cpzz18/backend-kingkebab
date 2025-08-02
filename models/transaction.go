package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID uint
	Type   string  // "topup", "purchase", "refund"
	Amount float64 // bisa negatif (purchase) atau positif (topup)
	Note   string
}
