package database

import (
	"boilerplate/src/modules/user/domain/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnection() {
	dsn := "host=localhost user=postgres password=123456 dbname=database port=5432 sslmode=disable TimeZone=America/Fortaleza"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Error Connection Database")
	}

	DB.AutoMigrate(&entity.User{})
}
