package models

type CartDetail struct {
	UserID    string
	ProductID string
	Amount    uint
	Title     string
	Price     float64
	Image     []byte
	MaxAmount uint
}
