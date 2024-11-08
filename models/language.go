package models

type Language struct {
	ID   int16
	Code string `gorm:"size:2,unique"`
	Name string `gorm:"size:255"`
}
