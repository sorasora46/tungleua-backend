package models

type UserDiscount struct {
	UserID     string `gorm:"column:user_id;not null"`
	DiscountID string `gorm:"column:discount_id;not null"`
}
