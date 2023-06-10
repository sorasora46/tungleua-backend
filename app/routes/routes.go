package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/Tungleua-backend/app/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	user := api.Group("/users")
	user.Get("/:id", handlers.GetUserById)
	user.Post("/", handlers.CreateUser)

	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)
	// auth.Post("/refresh-token", handlers.RefreshToken)
}
