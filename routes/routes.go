package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rendiputra/go-simple-rest-api/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/api/v1/task", controllers.HelloWorld)
	app.Get("/api/v1/task/:id", controllers.HelloWorld)
	app.Post("/api/v1/task/", controllers.HelloWorld)
	app.Delete("/api/v1/task/:id", controllers.HelloWorld)
}
