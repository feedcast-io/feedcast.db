package models

import "database/sql"

type ProductCategoryText struct {
	ID                int32
	ProductCategory   *ProductCategory
	ProductCategoryId sql.NullInt32
	Lang              string `gorm:"size:2"`
	Title             string `gorm:"size:2048"`
}
