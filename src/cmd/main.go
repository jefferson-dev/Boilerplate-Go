package main

import (
	"github.com/joho/godotenv"

	"boilerplate/src/infra/database"
	"boilerplate/src/infra/routes"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()

	app := routes.Routes()

	log.Fatal(app.Run(":3000"))
}
