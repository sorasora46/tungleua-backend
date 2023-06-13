package repositories

import (
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/utils"
)

func GetDiscounts(userID string) ([]models.Discount, error) {
	user_discounts := make([]models.UserDiscount, 0)
	discounts := make([]models.Discount, 0)

	result1 := utils.DB.Find(&user_discounts, "user_id = ?", userID)
	if result1.Error != nil {
		return nil, result1.Error
	}

	for _, userDiscount := range user_discounts {
		discount := new(models.Discount)
		result2 := utils.DB.Find(&discount, "id = ?", userDiscount.DiscountID)
		if result2.Error != nil {
			return nil, result2.Error
		}

		discounts = append(discounts, *discount)
	}

	return discounts, nil
}
