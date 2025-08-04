package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"backend-kking/backend/models"
	"backend-kking/backend/seeders"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal load file .env:", err)
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal konek ke database:\nDSN: %s\nError: %v", dsn, err)
	}

	log.Println("Berhasil konek ke database.")

	if err != nil {
		log.Fatalf("Gagal melakukan migrasi database: %v", err)
	}
	
	DB = db
	
	db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.Transaction{},
		&models.Topup{},
	)
	seeders.SeedAdminUser(DB)
}
