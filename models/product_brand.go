package models

type ProductBrand struct {
	ID   int32
	Slug string `gorm:"size:128"`
	Name string `gorm:"size:255"`
	Hash string `gorm:"size:32,unique"`
}
