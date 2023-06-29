package repositories

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"log"
	"strconv"

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
func TopUp(userid string, amount string) (string, error) {
	order := new(models.Order)
	id := uuid.New()

	order.ID = id.String()
	order.UserID = userid
	order.PaymentStatus = "pending"
	result := utils.DB.Create(order)
	if result.Error != nil {
		return "", result.Error
	}

	amountValue, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return "", err
	}
	user := models.User{}
	if err := utils.DB.First(&user, "id = ?", userid).Error; err != nil {
		return "", err
	}

	if order.PaymentStatus == "pending" {
		GeneratePromptPayQR(amountValue)
		fmt.Print("pending")
	} else {

		user.Balance += amountValue

	}

	if err := utils.DB.Save(&user).Error; err != nil {
		return "", err
	}

	return "", nil
}

func FindOrder2(userid string, couponID string) (string, error) {
	cart := []models.Cart{}
	Order := models.Order{}
	findOrderResult := utils.DB.Find(&cart, "user_id = ?", userid)
	if findOrderResult.Error != nil {
		return "", findOrderResult.Error
	}
	findUserOrder := utils.DB.Find(&Order, "user_id = ?", userid)
	if findUserOrder.Error != nil {
		return "", findUserOrder.Error
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

	balance := 0.0

	User := new(models.User)

	result2 := utils.DB.Select("balance").Find(&User, "id = ?", Order.UserID)
	if result2.Error != nil {
		return "", result2.Error
	}

	balance = User.Balance

	discount := 0.0

	if couponID != "" {
		discounts := []models.Discount{}
		findDiscount := utils.DB.Select("discount").Find(&discounts, "id = ?", couponID)
		if findDiscount.Error != nil {
			return "", findDiscount.Error
		}
		for _, item := range discounts {
			discount += item.Discount
		}
	}
	fmt.Println(discount)

	if balance >= price-discount*price {
		fmt.Println(balance - (price - discount*price))
		User := new(models.User)
		result2 := utils.DB.Model(&User).Where("id = ?", userid).Update("balance", balance-(price-discount*price))
		if result2.Error != nil {
			return "", result2.Error
		}

		result := utils.DB.Model(&Order).Where("id = ?", Order.ID).Update("payment_status", "Success")
		if result.Error != nil {
			return "", result.Error
		}

		message := "Payment was success"
		Balance := "Your Balance " + strconv.Itoa(int(balance-(price-discount*price))) + " left"
		responseMap := map[string]string{
			"message": message,
			"Balance": Balance,
		}
		jsonData, err := json.Marshal(responseMap)
		if err != nil {
			log.Fatal(err)
		}

		return string(jsonData), nil

	} else {
		fin := price - discount*price
		fin2 := balance - fin
		if fin2 < 0 {
			fin2 = -fin2
		}
		fmt.Print(fin2)
		User := new(models.User)

		message := "Not enough balance Please topup"
		Balance := strconv.Itoa(int(fin2))
		image := GeneratePromptPayQR(fin2)

		responseMap := map[string]string{
			"message": message,
			"Balance": Balance,
			"image":   image,
		}
		jsonData, err := json.Marshal(responseMap)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(User.Balance)

		return string(jsonData), nil

	}

}

func FindOrder(userid string, couponID string) (string, error) {
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
	dis := 0.0
	discount := 0.0
	if couponID != "" {
		discounts := []models.Discount{}
		findDiscount := utils.DB.Select("discount").Find(&discounts, "id = ?", couponID)
		if findDiscount.Error != nil {
			return "", findDiscount.Error
		}
		for _, item := range discounts {
			fmt.Print(item.Discount)
			discount += item.Discount
		}
		dis += (price * discount)

		str := GeneratePromptPayQR(price - dis)
		responseMap := map[string]string{
			"message": "Your Order is " + strconv.Itoa(int(price-dis)) + " Bath",
			"image":   str,
		}
		jsonData, err := json.Marshal(responseMap)
		if err != nil {
			log.Fatal(err)
		}
		return string(jsonData), nil

	}
	str := GeneratePromptPayQR(price)
	responseMap := map[string]string{
		"message": "Your Order is " + strconv.Itoa(int(price)) + " Bath",
		"image":   str,
	}
	jsonData, err := json.Marshal(responseMap)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonData), nil

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

// func GetOrders2() ([]*models.Order, error) {
// 	orders := []models.Order{}

// 	orderResult := utils.DB.Find(&orders)
// 	if orderResult.Error != nil {
// 		return nil, orderResult.Error
// 	}

// 	allOrders := make([]*models.Order, len(orders))

// 	for i, order := range orders {
// 		user := new(models.User)
// 		allOrders[i] = &models.Order{
// 			ID:            order.ID,
// 			UserID:        order.UserID,
// 			CreatedAt:     order.CreatedAt,
// 			PaymentStatus: order.PaymentStatus,
// 		}
// 	}

// 	return allOrders, nil
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
