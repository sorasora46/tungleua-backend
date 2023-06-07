package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/utils"
	"github.com/sorasora46/Tungleua-backend/config"
)

func main() {
	config.LoadConfig()

	utils.ConnectDatabase()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		user := new(models.User)
		utils.DB.First(&user)
		fmt.Println(user)
		return c.JSON(user)
	})

	log.Fatal(app.Listen(":3000"))
}
