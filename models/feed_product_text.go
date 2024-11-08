package models

import "database/sql"

type FeedProductText struct {
	ID            int32
	FeedProductId int32
	Title         sql.NullString `gorm:"size:255"`
	Description   sql.NullString `gorm:"type:text"`
	Label0        sql.NullString `gorm:"size:255"`
	Label1        sql.NullString `gorm:"size:255"`
	Label2        sql.NullString `gorm:"size:255"`
	Label3        sql.NullString `gorm:"size:255"`
	Label4        sql.NullString `gorm:"size:255"`
	ProductType   sql.NullString `gorm:"size:1024"`
}
