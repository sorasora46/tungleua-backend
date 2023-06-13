package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/Tungleua-backend/app/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	user := api.Group("/users")
	user.Get("/find-by-id/:id", handlers.GetUserById)
	user.Get("/exist", handlers.CheckDuplicateUser)
	user.Get("/find-by-email", handlers.GetUserByEmail)
	user.Put("/update/:id", handlers.UpdateUserById)
	user.Post("/", handlers.CreateUser)

	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)
	// auth.Post("/refresh-token", handlers.RefreshToken)

	store := api.Group("/stores")
	store.Get("/find-by-id/:id", handlers.GetStoreById)
	// store.Get("/images/:id")
	// store.Get("/exist")
	// store.Put("/update/:id")
	store.Post("/", handlers.CreateStore)
	// store.Delete("/delete/:id")
}
