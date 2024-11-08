package models

import "github.com/feedcast-io/feedcast.db/enums"

type FeedObject struct {
	ID         int32
	FeedId     int32
	Type       enums.FeedObjects
	Identifier string `gorm:"size:128"`
	Name       string `gorm:"size:255"`
}
