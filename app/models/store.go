package models

type Store struct {
	ID          string  `gorm:"column:id;primaryKey;"`
	Name        string  `gorm:"column:name;not null"`
	Contact     string  `gorm:"column:contact;not null"`
	TimeOpen    string  `gorm:"column:time_open;not null"`
	TimeClose   string  `gorm:"column:time_close;not null"`
	Description string  `gorm:"column:description;not null"`
	Latitude    float64 `gorm:"column:latitude;not null"`
	Longitude   float64 `gorm:"column:longitude;not null"`
	UserID      string  `gorm:"column:user_id;not null;"`
}
