package models

type User struct {
	Email string `gorm:"column:email;not null"`

	Password string `gorm:"column:password;not null"`
}
type LoginRequest struct {
	Email    string `gorm:"column:email;not null"`
	Password string `gorm:"column:password;not null"`
}
type RegisterRequest struct {
	Email    string `gorm:"column:email;not null"`
	Password string `gorm:"column:password;not null"`
}
