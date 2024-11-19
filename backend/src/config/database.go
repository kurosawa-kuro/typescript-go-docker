package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	dsn := "host=db user=postgres password=postgres dbname=myapp port=5432 sslmode=disable TimeZone=Asia/Tokyo"

	fmt.Printf("Attempting to connect with DSN: %s\n", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Connection error: %v\n", err)
		panic("Failed to connect database")
	}

	fmt.Println("Successfully connected to database!")
	return db
}
