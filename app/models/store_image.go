package models

type StoreImage struct {
	ID      string `gorm:"column:id;primaryKey;"` // image id
	StoreID string `gorm:"column:store_id;not null"`
	Image   []byte `gorm:"column:image;not null"`
}
