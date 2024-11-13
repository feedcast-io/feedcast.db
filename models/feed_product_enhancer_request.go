package models

import (
	"github.com/feedcast-io/feedcast.db/types"
	"gorm.io/gorm"
)

type FeedProductEnhancerRequest struct {
	ID        int32
	FeedId    int32
	Feed      Feed
	Pending   bool
	Products  types.ArrayString
	DeletedAt gorm.DeletedAt
}
