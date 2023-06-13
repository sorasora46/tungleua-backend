package models

import "time"

type Order struct {
	ID            string    `gorm:"column:id;primaryKey"`
	UserID        string    `gorm:"column:user_id;not null"`
	ProductID     string    `gorm:"column:product_id;not null"`
	CreatedAt     time.Time `gorm:"column:created_at;not null"`
	PaymentStatus string    `gorm:"column:payment_status;not null"`
	Amount        uint      `gorm:"column:amount; not null"`
}
