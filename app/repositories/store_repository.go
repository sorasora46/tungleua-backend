package repositories

import (
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/utils"
)

func GetStoreById(storeID string) (*models.Store, error) {
	store := new(models.Store)
	result := utils.DB.Find(store, "id = ?", storeID)
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

func GetStoreImages(storeID string) ([]models.StoreImage, error) {
	images := make([]models.StoreImage, 0)
	result := utils.DB.Find(&images, "id = ?", storeID)
	if result.Error != nil {
		return nil, result.Error
	}

	return images, nil
}
