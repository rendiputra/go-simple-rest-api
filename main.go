package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rendiputra/go-simple-rest-api/databases"
	"github.com/rendiputra/go-simple-rest-api/routes"
)

func main() {
	databases.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Use(cors.New())

	app.Listen(":8000")
}
