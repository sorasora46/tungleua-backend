package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/repositories"
)

func GetDiscounts(c *fiber.Ctx) error {
	userID := c.Params("id")

	discounts, err := repositories.GetDiscounts(userID)
	if err != nil {
		return err
	}

	return c.JSON(map[string][]models.Discount{
		"discounts": discounts,
	})
}

func GetDiscountById(c *fiber.Ctx) error {
	discountID := c.Params("id")

	discount, err := repositories.GetDiscountById(discountID)
	if err != nil {
		return err
	}

	return c.JSON(map[string]any{
		"discount_id": discount.ID,
		"titile":      discount.Title,
		"discount":    discount.Discount,
		"expire_date": discount.ExpireDate,
	})
}

func DeleteCouponById(c *fiber.Ctx) error {
	discountID := c.Params("id")

	if err := repositories.DeleteCouponById(discountID); err != nil {
		return err
	}

	return nil
}

func RemoveCouponFromUser(c *fiber.Ctx) error {
	discountID := c.Query("discount_id")
	userID := c.Query("user_id")

	if err := repositories.RemoveCouponFromUser(discountID, userID); err != nil {
		return err
	}

	return nil
}

func UpdateCouponById(c *fiber.Ctx) error {
	discountID := c.Params("id")

	updates := make(map[string]interface{})
	if err := c.BodyParser(&updates); err != nil {
		return err
	}

	if err := repositories.UpdateCouponById(discountID, updates); err != nil {
		return err
	}

	return nil
}

func CreateCoupon(c *fiber.Ctx) error {
	id := uuid.New().String()
	req := make(map[string]interface{})
	discount := new(models.Discount)

	if err := c.BodyParser(&req); err != nil {
		return err
	}
	discount.ID = id
	discount.Title = req["title"].(string)
	discount.Discount = req["discount"].(float64)

	// Convert string to time.Time
	expireDateStr := req["expire_date"].(string)
	expireDate, err := time.Parse("2006-01-02T15:04:05-0700", expireDateStr)
	if err != nil {
		// Handle the error if the string cannot be parsed as a time.Time
		return err
	}
	discount.ExpireDate = expireDate

	if err := repositories.CreateCoupon(discount); err != nil {
		return err
	}

	return nil
}
