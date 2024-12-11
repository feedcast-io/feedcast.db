package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type FeedProductEnhancerRequest struct {
	ID        int32
	FeedId    int32
	Feed      Feed
	Pending   bool
	Products  datatypes.JSONSlice[int]
	DeletedAt gorm.DeletedAt
}
