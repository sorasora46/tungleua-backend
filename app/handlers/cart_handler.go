package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/Tungleua-backend/app/repositories"
)

func GetCartByUserId(c *fiber.Ctx) error {
	userID := c.Params("id")

	results, err := repositories.GetCartByUserId(userID)
	if err != nil {
		return err
	}

	return c.JSON(map[string]any{
		"results": results,
	})
}
