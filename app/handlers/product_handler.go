package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/repositories"
)

func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	return nil
}

func GetProductById(c *fiber.Ctx) error {
	productID := c.Params("id")

	result, err := repositories.GetProductById(productID)
	if err != nil {
		return err
	}

	return c.JSON(map[string]any{
		"id":          result.ID,
		"title":       result.Title,
		"description": result.Description,
		"price":       result.Price,
		"store_id":    result.StoreID,
		"image":       result.Image,
		"amount":      result.Amount,
	})
}

func GetProducts(c *fiber.Ctx) error {
	storeID := c.Params("id")

	result, err := repositories.GetProducts(storeID)
	if err != nil {
		return err
	}

	products := make([]map[string]interface{}, len(result))
	for i, p := range result {
		products[i] = map[string]interface{}{
			"id":          p.ID,
			"title":       p.Title,
			"description": p.Description,
			"price":       p.Price,
			"store_id":    p.StoreID,
			"image":       p.Image,
			"amount":      p.Amount,
		}
	}

	return c.JSON(map[string]any{
		"products": products,
	})
}

func DeleteProductById(c *fiber.Ctx) error {
	productID := c.Params("id")

	err := repositories.DeleteProductById(productID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProductById(c *fiber.Ctx) error {
	storeID := c.Params("id")

	updates := make(map[string]interface{})
	if err := c.BodyParser(&updates); err != nil {
		return err
	}

	err := repositories.UpdateProductById(storeID, updates)
	if err != nil {
		return err
	}

	return nil
}
