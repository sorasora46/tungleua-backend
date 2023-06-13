package models

type Discount struct {
	ID       string `gorm:"column:id;primaryKey"`
	Title    string `gorm:"column:title;not null"`
	Discount uint   `gorm:"column:discount;not null"`
}
