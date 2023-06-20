package handlers

import (
	"github.com/gofiber/fiber/v2"
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
