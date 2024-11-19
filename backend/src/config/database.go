package config

import (
	"fmt"
	"os"

	"backend/src/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	host := "db"
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := "5432"

	// Debug all environment variables
	fmt.Printf("Environment Variables:\n")
	fmt.Printf("Host: %s\n", host)
	fmt.Printf("POSTGRES_USER: %s\n", user)
	fmt.Printf("POSTGRES_PASSWORD: %s\n", password)
	fmt.Printf("POSTGRES_DB: %s\n", dbname)

	// Add debug logging
	fmt.Printf("Connecting to PostgreSQL at %s:%s\n", host, port)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	// マイグレーション
	db.AutoMigrate(&model.Micropost{})

	return db
}
