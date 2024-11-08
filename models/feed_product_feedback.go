package models

import (
	"github.com/feedcast-io/feedcast.db/types"
	"time"
)

type FeedProductFeedback struct {
	ID             int32
	FeedProductId  int32
	LastImport     time.Time
	Source         types.ProductFeedbackSource
	ProductError   ProductError
	ProductErrorId int32
}
