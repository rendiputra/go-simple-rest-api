package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rendiputra/go-simple-rest-api/controllers"
)

func Setup(app *fiber.App) {
	//app.Get("/api/v1/task", task.HelloWorld)
	//app.Get("/api/v1/task/:id", task.HelloWorld)
	app.Post("/api/v1/task/", controllers.Create)
	//app.Delete("/api/v1/task/:id", task.HelloWorld)
}
