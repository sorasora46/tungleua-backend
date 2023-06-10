package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/Tungleua-backend/app/routes"
	"github.com/sorasora46/Tungleua-backend/app/utils"
	"github.com/sorasora46/Tungleua-backend/config"
)

func main() {
	config.LoadConfig()

	utils.ConnectDatabase()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
