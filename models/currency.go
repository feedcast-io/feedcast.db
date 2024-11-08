package models

type Currency struct {
	ID     int32
	Code   string `gorm:"size:3,unique"`
	Symbol string `gorm:"size:8"`
}
