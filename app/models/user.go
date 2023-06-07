package models

type User struct {
	ID    int    `gorm:"column:id;primaryKey"`
	Phone string `gorm:"column:phone;not null"`
	Image []byte `gorm:"column:image"`
	Email string `gorm:"column:email;not null"`
	Role  int    `gorm:"column:role;not null"`
	Name  string `gorm:"column:name;not null"`
}
