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

func GetTaskById(c *fiber.Ctx) error {
	task := models.Task{}
	id := c.Params("id")

	databases.DBConn.Find(&task, id)

	return c.Status(200).JSON(task)
}

func Update(c *fiber.Ctx) error {
	task := []models.Task{}
	id := c.Params("id")
	data := new(models.Task)

	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	databases.DBConn.Model(&task).Where("id = ?", id).Updates(map[string]interface{}{"Title": data.Title, "Description": data.Description})

	return c.Status(200).JSON("Update berhasil!")
}

func Delete(c *fiber.Ctx) error {
	task := []models.Task{}
	id := c.Params("id")
	
	databases.DBConn.Where("id = ?", id).Delete(&task)

	return c.Status(200).JSON("Deleted")
}
