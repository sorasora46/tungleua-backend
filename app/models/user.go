package models

type User struct {
	ID     string `gorm:"column:id;primaryKey;"`
	Phone  string `gorm:"column:phone;not null"`
	Image  []byte `gorm:"column:image"`
	Email  string `gorm:"column:email;not null"`
	IsShop bool   `gorm:"column:is_shop;not null"`
	Name   string `gorm:"column:name;not null"`
}

// For Create User only!
type UserWithPassword struct {
	User
	Password string
}

type Password struct {
	UserID         string `gorm:"column:user_id;primaryKey;"`
	HashedPassword string `gorm:"column:hashed_password;not null"`
}
