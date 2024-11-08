package models

type FeedProductReference struct {
	ID        int32
	Reference string `gorm:"size:96,unique"`
}
