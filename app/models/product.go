package models

type Product struct {
	ID          string  `gorm:"column:id;primaryKey;"`
	Title       string  `gorm:"column:title;not null"`
	Description string  `gorm:"column:description;not null"`
	Price       float64 `gorm:"column:price;not null"`
	StoreID     string  `gorm:"column:store_id;not null"`
	Image       []byte  `gorm:"column:image;not null"`
	Amount      uint    `gorm:"column:amount;not null"`
}
