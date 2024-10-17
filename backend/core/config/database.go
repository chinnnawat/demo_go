package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PORT := os.Getenv("DB_PORT")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s TimeZone=Asia/Shanghai", DB_HOST, DB_USER, DB_NAME, DB_PORT, DB_PASSWORD)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Success connect DB")

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get DB from gorm: %v", err)
	}
	defer sqlDB.Close()

	// db.AutoMigrate()
	// Perform any database initialization here (e.g., migrations)
	// db.AutoMigrate(&YourModel{})
	return db
}
