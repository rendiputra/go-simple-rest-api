package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rendiputra/go-simple-rest-api/databases"
	"github.com/rendiputra/go-simple-rest-api/models"
)

func Create(c *fiber.Ctx) error {
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	databases.DBConn.Create(&task)

	return c.Status(200).JSON(task)
}

func GetTasks(c *fiber.Ctx) error {
	tasks := []models.Task{}

	databases.DBConn.Find(&tasks)

	return c.Status(200).JSON(tasks)
}
