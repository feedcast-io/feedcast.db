package models

type FeedDomain struct {
	ID     int32
	Domain string `gorm:"size:96,unique"`
}
