package config

import (
	"fmt"
	"os"

	"backend/src/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	// Add error checking for required environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("USER_NAME")
	password := os.Getenv("USER_PASSWORD")
	dbname := os.Getenv("DATABASE")
	port := os.Getenv("DB_PORT")

	// Verify that all required variables are present
	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		panic("Missing required database environment variables")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tokyo",
		host, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	// マイグレーション
	db.AutoMigrate(&model.Micropost{})

	return db
}
