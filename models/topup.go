package models

import "gorm.io/gorm"

type Topup struct {
	gorm.Model
	UserID     uint
	Amount     float64
	Status     string // "pending", "approved", "rejected"
	ApprovedAt *string
}
