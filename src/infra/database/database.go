package database

import (
	"boilerplate/src/modules/user/domain/entity"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() {
	connectionString := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"

	dsn := fmt.Sprintf(
		connectionString,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("SSL_MODE"),
		os.Getenv("TIMEZONE"),
	)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Error Connection Database")
	}

	DB.AutoMigrate(&entity.User{})
}
