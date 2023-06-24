package repositories

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"log"

	"github.com/google/uuid"
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/utils"

	"image/png"

	pp "github.com/Frontware/promptpay"
	"github.com/skip2/go-qrcode"
)

func CreateOrder(order *models.Order, userid string) error {
	id := uuid.New()

	order.ID = id.String()
	order.UserID = userid
	order.PaymentStatus = "pending"
	result := utils.DB.Create(order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func FindOrder(userid string) (string, error) {
	cart := []models.Cart{}
	findOrderResult := utils.DB.Find(&cart, "user_id = ?", userid)
	if findOrderResult.Error != nil {
		return "", findOrderResult.Error
	}

	price := 0.0
	for _, item := range cart {
		product := new(models.Product)

		result2 := utils.DB.Select("price").Find(&product, "id = ?", item.ProductID)
		if result2.Error != nil {
			return "", result2.Error
		}
		CreateOrders(&models.OrderProducts{}, item)
		price += product.Price * float64(item.Amount)

	}

	fmt.Println(price)
	str := GeneratePromptPayQR(price)

	return str, nil
}
func CreateOrders(orderP *models.OrderProducts, item models.Cart) error {
	Order := []models.Order{}
	findOrderResult := utils.DB.Find(&Order, "user_id = ?", item.UserID)
	if findOrderResult.Error != nil {
		return findOrderResult.Error
	}
	for _, item := range Order {

		orderP.OrderID = item.ID

	}

	orderP.ProductID = item.ProductID
	orderP.Amount = item.Amount
	result := utils.DB.Create(orderP)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GeneratePromptPayQR(price float64) string {

	payment := pp.PromptPay{
		PromptPayID: "0959597702",
		Amount:      price,
	}

	qrData, _ := payment.Gen()

	qrCode, err := qrcode.Encode(qrData, qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err)
	}
	qrImage, _, err := image.Decode(bytes.NewReader(qrCode))
	if err != nil {
		log.Fatal(err)
	}

	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, qrImage)
	if err != nil {
		log.Fatal(err)
	}

	base64Str := base64.StdEncoding.EncodeToString(buffer.Bytes())

	return base64Str

}

// func GetOrderById(orderID string) (*models.OrderDetail, error) {
// 	order := new(models.Order)
// 	product := new(models.Product)

// 	order_result := utils.DB.Find(&order, "id = ?", orderID)
// 	if order_result.Error != nil {
// 		return nil, order_result.Error
// 	}

// 	product_result := utils.DB.Find(&product, "id = ?", order.ProductID)
// 	if product_result.Error != nil {
// 		return nil, product_result.Error
// 	}

// 	order_detail := models.OrderDetail{
// 		ID:            order.ID,
// 		UserID:        order.UserID,
// 		ProductID:     order.ProductID,
// 		StoreID:       product.StoreID,
// 		Title:         product.Title,
// 		Description:   product.Description,
// 		Image:         product.Image,
// 		Price:         product.Price,
// 		CreatedAt:     order.CreatedAt,
// 		PaymentStatus: order.PaymentStatus,
// 		Amount:        order.Amount,
// 	}

// 	return &order_detail, nil
// }

// func GetOrders(userID string) ([]models.OrderDetail, error) {
// 	orders := make([]models.Order, 0)
// 	orderDetails := make([]models.OrderDetail, 0)

// 	orderResult := utils.DB.Find(&orders, "user_id = ?", userID)
// 	if orderResult.Error != nil {
// 		return nil, orderResult.Error
// 	}

// 	for _, order := range orders {
// 		product := new(models.Product)
// 		productResult := utils.DB.Find(&product, "id = ?", order.ProductID)
// 		if productResult.Error != nil {
// 			return nil, productResult.Error
// 		}

// 		orderDetail := models.OrderDetail{
// 			ID:            order.ID,
// 			UserID:        order.UserID,
// 			ProductID:     order.ProductID,
// 			StoreID:       product.StoreID,
// 			Title:         product.Title,
// 			Description:   product.Description,
// 			Image:         product.Image,
// 			Price:         product.Price,
// 			CreatedAt:     order.CreatedAt,
// 			PaymentStatus: order.PaymentStatus,
// 			Amount:        order.Amount,
// 		}

// 		orderDetails = append(orderDetails, orderDetail)
// 	}

// 	return orderDetails, nil
// }

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
