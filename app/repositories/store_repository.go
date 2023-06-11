package repositories

import (
	"github.com/google/uuid"
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

func CreateStore(store *models.Store, images [][]byte) error {
	store_result := utils.DB.Create(store)
	if store_result.Error != nil {
		return store_result.Error
	}

	for _, image := range images {
		id := uuid.New().String()
		imageData := models.StoreImage{
			ID:      id,
			StoreID: store.ID,
			Image:   image,
		}
		image_result := utils.DB.Create(&imageData)
		if image_result.Error != nil {
			return image_result.Error
		}
	}

	user_result := utils.DB.Model(&models.User{}).Where("id = ?", store.UserID).Update("is_shop", true)
	if user_result.Error != nil {
		return user_result.Error
	}

	return nil
}

func UpdateStoreById(storeID string, updates map[string]interface{}) (string, error) {
	result := utils.DB.Model(&models.Store{}).Where("id = ?", storeID).Updates(updates)
	if result.Error != nil {
		return "failed", result.Error
	}
	return "success", nil
}

func GetStoreImages(storeID string) ([]models.StoreImage, error) {
	images := make([]models.StoreImage, 0)
	result := utils.DB.Find(&images, "id = ?", storeID)
	if result.Error != nil {
		return nil, result.Error
	}

	return images, nil
}
