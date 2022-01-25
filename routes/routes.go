package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rendiputra/go-simple-rest-api/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/api/v1/tasks", controllers.GetTasks)
	app.Get("/api/v1/task/:id", controllers.GetTaskById)
	app.Post("/api/v1/task/", controllers.Create)
	app.Put("/api/v1/task/:id", controllers.Update)
}
