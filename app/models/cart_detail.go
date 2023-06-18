package models

type CartDetail struct {
	UserID    string
	ProductID string
	Amount    string
	Title     string
	Price     uint
	Image     []byte
}
