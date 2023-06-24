package models

type OrderProducts struct {
	OrderID   string `gorm:"column:order_id;not null"`
	ProductID string `gorm:"column:product_id;not null"`
	Amount    uint   `gorm:"column:amount; not null"`
}
