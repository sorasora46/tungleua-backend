package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/Tungleua-backend/app/repositories"
)

func GetCartByUserId(c *fiber.Ctx) error {
	userID := c.Params("id")

	results, err := repositories.GetCartByUserId(userID)
	if err != nil {
		return err
	}

	response := make([]map[string]interface{}, len(results))

	for i, item := range results {
		response[i] = map[string]interface{}{
			"user_id":    item.UserID,
			"product_id": item.ProductID,
			"amount":     item.Amount,
			"title":      item.Title,
			"price":      item.Price,
			"image":      string(item.Image),
			"max_amount": item.MaxAmount,
		}
	}

	return c.JSON(map[string]any{
		"results": response,
	})
}

func DeleteItemFromCart(c *fiber.Ctx) error {
	userID := c.Query("user_id")
	productID := c.Query("product_id")

	if err := repositories.DeleteItemFromCart(userID, productID); err != nil {
		return err
	}

	return nil
}

func UpdateItemAmount(c *fiber.Ctx) error {
	userID := c.Query("user_id")
	productID := c.Query("product_id")
	newAmountStr := c.Query("new_amount")

	newAmount, err := strconv.ParseUint(newAmountStr, 10, 64)
	if err != nil {
		return err
	}

	if err := repositories.UpdateItemAmount(userID, productID, uint(newAmount)); err != nil {
		return err
	}

	return nil
}

func AddItemToCart(c *fiber.Ctx) error {
	userID := c.Query("user_id")
	productID := c.Query("product_id")
	amountStr := c.Query("amount")

	amount, err := strconv.ParseUint(amountStr, 10, 64)
	if err != nil {
		return err
	}

	if err := repositories.AddItemToCart(userID, productID, uint(amount)); err != nil {
		return err
	}

	return nil
}
