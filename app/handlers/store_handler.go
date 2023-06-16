package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/repositories"
)

func GetStoreByUserId(c *fiber.Ctx) error {
	userID := c.Params("id")

	store, store_err := repositories.GetStoreByUserId(userID)
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
		"image":       string(store.Image),
	})
}

func CreateStore(c *fiber.Ctx) error {
	id := uuid.New().String()

	req := make(map[string]interface{})
	if parse_err := c.BodyParser(&req); parse_err != nil {
		return parse_err
	}

	store := models.Store{
		ID:          id,
		Name:        req["name"].(string),
		Contact:     req["contact"].(string),
		TimeOpen:    req["time_open"].(string),
		TimeClose:   req["time_close"].(string),
		Description: req["description"].(string),
		Latitude:    req["latitude"].(float64),
		Longitude:   req["longitude"].(float64),
		UserID:      req["user_id"].(string),
		Image:       []byte(req["image"].(string)),
	}

	err := repositories.CreateStore(&store)
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

func CheckDuplicateStore(c *fiber.Ctx) error {
	req := new(models.Store)
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	err := repositories.CheckDuplicateStore(req)
	if err != nil {
		return err
	}

	return nil
}

func PopulateMap(c *fiber.Ctx) error {
	offsetStr := c.Query("offset")
	centerLatStr := c.Query("center_lat")
	centerLongStr := c.Query("center_long")

	offset, err := strconv.ParseFloat(offsetStr, 64)
	if err != nil {
		return err
	}

	centerLat, err := strconv.ParseFloat(centerLatStr, 64)
	if err != nil {
		return err
	}

	centerLong, err := strconv.ParseFloat(centerLongStr, 64)
	if err != nil {
		return err
	}

	result, err := repositories.PopulateMap(offset, centerLat, centerLong)
	if err != nil {
		return err
	}

	stores := make([]map[string]interface{}, len(result))
	for i, s := range result {
		stores[i] = map[string]interface{}{
			"id":        s.ID,
			"latitude":  s.Latitude,
			"longitude": s.Longitude,
			"user_id":   s.UserID,
		}
	}

	return c.JSON(map[string]any{
		"stores": stores,
	})
}
