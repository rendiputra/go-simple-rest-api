package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rendiputra/go-simple-rest-api/databases"
	"github.com/rendiputra/go-simple-rest-api/routes"
)

func main() {
	databases.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
