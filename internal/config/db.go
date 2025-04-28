package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (DB *gorm.DB) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get generic DB object: ", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	fmt.Println("Database connection established")

	MigrateDB(DB)

	return
}

func MigrateDB(DB *gorm.DB) {
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
	fmt.Println("Database migrated successfully")
}
