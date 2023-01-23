package routes

import (
	"example/rest-api-fiber/database"
	"example/rest-api-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func GetTodos(c *fiber.Ctx) error {
	todos := []models.Todo{}

	database.DB.Db.Find(&todos)

	return c.Status(200).JSON(todos)
}

func AddTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	
	err := c.BodyParser(todo)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&todo)

	return c.Status(201).JSON(todo)
}