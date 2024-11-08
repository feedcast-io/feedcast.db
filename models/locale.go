package models

type Locale struct {
	ID         int32
	Locale     string `gorm:"size:5,unique"`
	Currency   Currency
	CurrencyId int32
}
