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

func DeleteItemFromCart(userID string, productID string) error {
	if err := utils.DB.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.Cart{}); err.Error != nil {
		return err.Error
	}

	return nil
}

// newAmount was calculated from the frontend
// This function is used for both decrement and increment
func UpdateItemAmount(userID string, productID string, newAmount uint) error {
	if err := utils.DB.Model(&models.Cart{}).Where("user_id = ? AND product_id = ?", userID, productID).Update("amount", newAmount); err.Error != nil {
		return err.Error
	}

	return nil
}

func AddItemToCart(userID string, productID string, amount uint) error {
	cart := models.Cart{
		UserID:    userID,
		ProductID: productID,
		Amount:    amount,
	}

	oldCart := new(models.Cart)
	if err := utils.DB.Find(&oldCart, "user_id = ? AND product_id = ?", userID, productID); err.Error != nil {
		return err.Error
	}

	// not found
	if oldCart.UserID == "" {
		if create_err := utils.DB.Create(cart); create_err.Error != nil {
			return create_err.Error
		}
		return nil
	}

	// found
	if err := utils.DB.Model(&models.Cart{}).Where("user_id = ? AND product_id = ?", userID, productID).Update("amount", oldCart.Amount+amount); err.Error != nil {
		return err.Error
	}

	return nil
}
