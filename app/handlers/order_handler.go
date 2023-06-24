package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/repositories"
)

func CreateOrder(c *fiber.Ctx) error {
	order := new(models.Order)
	userID := c.Params("id")
	err := repositories.CreateOrder(order, userID)
	if err != nil {
		return err
	}

	result, errs := repositories.FindOrder(userID)
	if errs != nil {
		return errs
	} else {
		c.SendString(result)
	}
	return nil
}

// func GetOrderById(c *fiber.Ctx) error {
// 	orderID := c.Params("id")

// 	order, err := repositories.GetOrderById(orderID)
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(map[string]any{
// 		"order_id":       order.ID,
// 		"user_id":        order.UserID,
// 		"product_id":     order.ProductID,
// 		"store_id":       order.StoreID,
// 		"title":          order.Title,
// 		"description":    order.Description,
// 		"image":          order.Image,
// 		"created_at":     order.CreatedAt,
// 		"amount":         order.Amount,
// 		"price":          order.Price,
// 		"payment_status": order.PaymentStatus,
// 	})
// }

// func GetOrders(c *fiber.Ctx) error {
// 	userID := c.Params("id")

// 	orders, err := repositories.GetOrders(userID)
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(map[string][]models.OrderDetail{
// 		"orders": orders,
// 	})
// }

func DeleteOrderById(c *fiber.Ctx) error {
	orderID := c.Params("id")

	err := repositories.DeleteOrderById(orderID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateOrderStatus(c *fiber.Ctx) error {
	userID := c.Params("id")
	status := c.Query("status")

	err := repositories.UpdateOrderStatus(userID, status)
	if err != nil {
		return err
	}

	return nil
}
