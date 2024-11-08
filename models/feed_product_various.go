package models

import "database/sql"

type FeedProductVarious struct {
	ID            int32
	FeedProductId int32
	ItemGroupId   sql.NullString `gorm:"size:255"`
	Mpn           sql.NullString `gorm:"size:255"`
	Gtin          sql.NullString `gorm:"size:255"`
	AvailableDate sql.NullString `gorm:"size:255"`
	SalePriceDate sql.NullString `gorm:"size:255"`
	Color         sql.NullString `gorm:"size:255"`
	Material      sql.NullString `gorm:"size:255"`
	Size          sql.NullString `gorm:"size:255"`
}
