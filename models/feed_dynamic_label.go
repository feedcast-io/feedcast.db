package models

import "github.com/feedcast-io/feedcast.db/types"

type FeedDynamicLabel struct {
	ID     int32
	Feed   Feed
	FeedId int32
	Config types.LabelRules `gorm:"type:text"`
}
