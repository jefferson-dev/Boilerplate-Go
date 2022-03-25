package main

import (
	"boilerplate/src/infra/database"
	"boilerplate/src/infra/routes"
	"log"
)

func main() {
	database.DbConnection()

	app := routes.Routes()

	log.Fatal(app.Run(":3000"))
}
