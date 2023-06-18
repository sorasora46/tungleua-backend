package models

type CartDetail struct {
	UserID    string
	ProductID string
	Amount    uint
	Title     string
	Price     uint
	Image     []byte
	MaxAmount uint
}
