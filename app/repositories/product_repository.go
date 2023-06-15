package repositories

import (
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/utils"
)

func CreateProduct(product *models.Product) error {
	result := utils.DB.Create(product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetProductById(productID string) (*models.Product, error) {
	product := new(models.Product)
	result := utils.DB.Find(&product, "id = ?", productID)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func GetProducts(storeID string) ([]models.Product, error) {
	products := make([]models.Product, 0)
	result := utils.DB.Find(&products, "store_id = ?", storeID)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func DeleteProductById(productID string) error {
	result := utils.DB.Delete(&models.Product{}, productID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateProductById(productID string, updates map[string]interface{}) error {
	result := utils.DB.Model(&models.Product{}).Where("id = ?", productID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetProductImages(storeID string) ([][]byte, error) {
	var images [][]byte

	// Retrieve images with matching storeID
	if err := utils.DB.Model(&models.Product{}).Select("image").Where("store_id = ?", storeID).Find(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}
