package repositories

import (
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/utils"
)

func CreateOrder(order *models.Order) error {
	result := utils.DB.Create(order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetOrderById(orderID string) (*models.OrderDetail, error) {
	order := new(models.Order)
	product := new(models.Product)

	order_result := utils.DB.Find(&order, "id = ?", orderID)
	if order_result.Error != nil {
		return nil, order_result.Error
	}

	product_result := utils.DB.Find(&product, "id = ?", order.ProductID)
	if product_result.Error != nil {
		return nil, product_result.Error
	}

	order_detail := models.OrderDetail{
		ID:            order.ID,
		UserID:        order.UserID,
		ProductID:     order.ProductID,
		StoreID:       product.StoreID,
		Title:         product.Title,
		Description:   product.Description,
		Image:         product.Image,
		Price:         product.Price,
		CreatedAt:     order.CreatedAt,
		PaymentStatus: order.PaymentStatus,
		Amount:        order.Amount,
	}

	return &order_detail, nil
}

// TODO return json with product info
func GetOrders(userID string) ([]models.Order, error) {
	orders := make([]models.Order, 0)

	result := utils.DB.Find(&orders, "user_id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func DeleteOrderById(userID string) error {
	result := utils.DB.Delete(&models.Order{}, userID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateOrderStatus(userID string, status string) error {
	result := utils.DB.Model(&models.Order{}).Where("user_id", userID).Update("payment_status", status)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
