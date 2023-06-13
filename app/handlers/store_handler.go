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
	}

	imagesInterface := req["images"].([]interface{})
	images := make([][]byte, len(imagesInterface))
	for i, img := range imagesInterface {
		imgStr := img.(string)
		decodedImg, err := base64.StdEncoding.DecodeString(imgStr)
		if err != nil {
			return err
		}
		images[i] = decodedImg
	}

	err := repositories.CreateStore(&store, images)
	if err != nil {
		return err
	}

	return nil
}
