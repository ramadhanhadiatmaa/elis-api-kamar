package routes

import (
	"kamar/controllers"
	"kamar/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {
	auth := r.Group("/api")

	auth.Get("/", middlewares.Auth, controllers.Show)
	auth.Post("/", middlewares.Auth, controllers.Create)
	auth.Put("/:id", middlewares.Auth, controllers.Update)
	auth.Delete("/:id", middlewares.Auth, controllers.Delete)
}
