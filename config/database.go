package config

import (
	"fmt"
	"log"
	"os"

	"Hotel-Management/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
	fmt.Println("Database connected successfully")

	err = DB.AutoMigrate(
		&models.Hotel{},
		&models.Room{},
		&models.Guest{},
		&models.Booking{},
		&models.Payment{},
		&models.Staff{},
		&models.Amenity{},
		&models.RoomService{},
		&models.Review{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	fmt.Println("Database migrated successfully")
}
