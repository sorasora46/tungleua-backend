package models

import "time"

type OrderDetail struct {

	// Order Data
	ID            string    `gorm:"column:id;primaryKey"`
	UserID        string    `gorm:"column:user_id;not null"`
	ProductID     string    `gorm:"column:product_id;not null"`
	CreatedAt     time.Time `gorm:"column:created_at;not null"`
	PaymentStatus string    `gorm:"column:payment_status;not null"`
	Amount        uint      `gorm:"column:amount; not null"`

	// Product Data
	Title       string `gorm:"column:title;not null"`
	Description string `gorm:"column:description;not null"`
	Price       uint   `gorm:"column:price;not null"`
	StoreID     string `gorm:"column:store_id;not null"`
	Image       []byte `gorm:"column:image;not null"`
}
