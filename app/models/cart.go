package models

type Cart struct {
	UserID    string `gorm:"column:user_id;not null"`
	ProductID string `gorm:"column:product_id;not null"`
	Amount    uint   `gorm:"column:amount;not null"`
}
