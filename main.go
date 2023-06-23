package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/Tungleua-backend/app/handlers"
	"github.com/sorasora46/Tungleua-backend/app/routes"
	"github.com/sorasora46/Tungleua-backend/app/utils"
	"github.com/sorasora46/Tungleua-backend/config"
)

func main() {
	config.LoadConfig()

	utils.ConnectDatabase()

	app := fiber.New(fiber.Config{
		StrictRouting: true,
	})

	routes.SetupRoutes(app)
	app.Get("/p", handlers.GetCartByUserId)

	log.Fatal(app.Listen(":3000"))
}
