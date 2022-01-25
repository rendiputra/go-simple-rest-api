package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func create(c *fiber.Ctx) error {
	task := new(Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	res, err := db

	return c.SendString("Hello, World 👋!")
}
