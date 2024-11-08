package models

type CountryVat struct {
	ID        int32
	Code      string `gorm:"size:10"`
	Name      string `gorm:"size:255"`
	CountryId int32
}
