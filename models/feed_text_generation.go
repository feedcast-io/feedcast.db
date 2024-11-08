package models

import (
	"github.com/feedcast-io/feedcast.db/types"
	"gorm.io/gorm"
	"time"
)

type FeedTextGeneration struct {
	ID              int32
	FeedId          int32
	Type            types.FeedTextGenerationTypes
	Date            time.Time
	NbTextGenerated int16
}

func CountFeedTextGeneration(conn *gorm.DB, feedId int32, from time.Time) (int64, error) {
	var count int64
	if e := conn.Model(FeedTextGeneration{}).
		Where("feed_id = ? AND date >= ?", feedId, from).
		Count(&count).Error; nil != e {
		return 0, e
	}

	return count, nil
}
