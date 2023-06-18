package repositories

import (
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/utils"
)

func GetCartByUserId(userID string) ([]models.CartDetail, error) {
	cartItems := make([]models.Cart, 0)
	results := make([]models.CartDetail, 0)

	if result := utils.DB.Find(&cartItems, "user_id = ?", userID); result.Error != nil {
		return nil, result.Error
	}

	for _, item := range cartItems {
		product := new(models.Product)

		if result := utils.DB.Find(&product, "id = ?", item.ProductID); result.Error != nil {
			return nil, result.Error
		}

		results = append(results, models.CartDetail{
			UserID:    item.UserID,
			ProductID: item.ProductID,
			Amount:    item.Amount,
			Title:     product.Title,
			Price:     product.Price,
			Image:     product.Image,
		})

	}

	return results, nil
}
