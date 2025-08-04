package seeders

import (
	"backend-kking/backend/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAdminUser(DB *gorm.DB) {
	var count int64
	DB.Model(&models.User{}).Where("role = ?", "admin").Count(&count)

	if count == 0 {
		password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

		admin := models.User{
			Name:     "Admin Kebab",
			Email:    "admin@kebab.com",
			Password: string(password),
			Role:     "admin",
			Balance:  100000,
		}
		DB.Create(&admin)
	}
}
