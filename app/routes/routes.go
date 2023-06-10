package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/Tungleua-backend/app/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	user := api.Group("/users")
	user.Get("/find-by-id/:id", handlers.GetUserById)
	user.Get("/exist", handlers.CheckIsUserExist)
	user.Get("/find-by-email", handlers.GetUserByEmail)
	user.Put("/update/:id", handlers.UpdateUserById)
	user.Post("/", handlers.CreateUser)

	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)
	// auth.Post("/refresh-token", handlers.RefreshToken)
}
