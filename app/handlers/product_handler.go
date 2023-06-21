package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/repositories"
)

func CreateProduct(c *fiber.Ctx) error {
	id := uuid.New().String()
	req := make(map[string]interface{})
	product := new(models.Product)

	if err := c.BodyParser(&req); err != nil {
		return err
	}
	// map yak chip hai kuay
	product.ID = id
	product.Amount = uint(req["amount"].(float64))
	product.Description = req["description"].(string)
	product.Image = []byte(req["image"].(string))
	product.Price = req["price"].(float64)
	product.StoreID = req["store_id"].(string)
	product.Title = req["title"].(string)

	create_err := repositories.CreateProduct(product)
	if create_err != nil {
		return create_err
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
		"image":       string(result.Image),
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
			"image":       string(p.Image),
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

func GetProductImages(c *fiber.Ctx) error {
	storeID := c.Params("id")

	images, err := repositories.GetProductImages(storeID)
	if err != nil {
		return err
	}

	// Convert []byte to string
	convertedImages := make([]string, len(images))
	for i, img := range images {
		convertedImages[i] = string(img)
	}

	return c.JSON(map[string]any{
		"images": convertedImages,
	})
}
