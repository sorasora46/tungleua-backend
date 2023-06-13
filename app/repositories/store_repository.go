package repositories

import (
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/utils"
)

func GetStoreById(storeID string) (*models.Store, error) {
	store := new(models.Store)
	result := utils.DB.Find(&store, "id = ?", storeID)
	if result.Error != nil {
		return nil, result.Error
	}

	return store, nil
}

func CreateStore(store *models.Store) error {
	store_result := utils.DB.Create(store)
	if store_result.Error != nil {
		return store_result.Error
	}

	user_result := utils.DB.Model(&models.User{}).Where("id = ?", store.UserID).Update("is_shop", true)
	if user_result.Error != nil {
		return user_result.Error
	}

	return nil
}

func UpdateStoreById(storeID string, updates map[string]interface{}) error {
	result := utils.DB.Model(&models.Store{}).Where("id = ?", storeID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteStoreById(storeID string, userID string) error {
	store_result := utils.DB.Delete(&models.Store{}, storeID)
	if store_result.Error != nil {
		return store_result.Error
	}

	user_result := utils.DB.Model(&models.User{}).Update("is_shop", false)
	if user_result.Error != nil {
		return user_result.Error
	}

	product_result := utils.DB.Where("store_id", storeID).Delete(&models.Product{})
	if product_result.Error != nil {
		return product_result.Error
	}

	return nil
}

func CheckDuplicateStore(store *models.Store) error {
	result := utils.DB.Where("user_id = ?", store.UserID).Or("name = ?", store.Name).First(store)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
