package main

import (
	"log"
	"os"

	"example/rest-api-fiber/database"
	"example/rest-api-fiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/healthcheck", routes.HealthCheck)
	app.Get("/todos", routes.GetTodos)
	app.Post("/todos", routes.AddTodo)
}

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	database.ConnectDb()
}

func main() {

	app := fiber.New()

	setUpRoutes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(os.Getenv("PORT")))
}