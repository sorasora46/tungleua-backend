package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/repositories"
)

func GetStoreById(c *fiber.Ctx) error {
	storeID := c.Params("id")

	store, store_err := repositories.GetStoreById(storeID)
	if store_err != nil {
		return store_err
	}

	return c.JSON(map[string]any{
		"id":          store.ID,
		"name":        store.Name,
		"contact":     store.Contact,
		"time_open":   store.TimeOpen,
		"time_close":  store.TimeClose,
		"description": store.Description,
		"latitude":    store.Latitude,
		"longtitude":  store.Longitude,
		"user_id":     store.UserID,
		"image":       store.Image,
	})
}

func CreateStore(c *fiber.Ctx) error {
	id := uuid.New().String()

	req := new(models.Store)
	if parse_err := c.BodyParser(&req); parse_err != nil {
		return parse_err
	}
	req.ID = id

	err := repositories.CreateStore(req)
	if err != nil {
		return err
	}

	return nil
}

func UpdateStoreById(c *fiber.Ctx) error {
	storeID := c.Params("id")

	updates := make(map[string]interface{})
	if err := c.BodyParser(&updates); err != nil {
		return err
	}

	err := repositories.UpdateStoreById(storeID, updates)
	if err != nil {
		return err
	}

	return nil
}

func DeleteStoreById(c *fiber.Ctx) error {
	storeID := c.Query("store_id")
	userID := c.Query("user_id")

	err := repositories.DeleteStoreById(storeID, userID)
	if err != nil {
		return err
	}

	return nil
}

	}

	err := repositories.CreateStore(&store, images)
	if err != nil {
		return err
	}

	return nil
}
