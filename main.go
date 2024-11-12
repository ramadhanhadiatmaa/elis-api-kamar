package main

import (
	"kamar/models"
	"kamar/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	models.ConnectionDatabase()

	app := fiber.New()

	app.Use(cors.New())
	
	routes.Route(app)

	app.Listen(":8113")
}

