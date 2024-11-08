package models

import "database/sql"

type FeedProductUrl struct {
	ID            int32
	FeedProductId int32
	Link          sql.NullString `gorm:"size:1024"`
	MobileLink    sql.NullString `gorm:"size:1024"`
	ImageLink     sql.NullString `gorm:"size:1024"`
	AdsRedirect   sql.NullString `gorm:"size:1024"`
}
