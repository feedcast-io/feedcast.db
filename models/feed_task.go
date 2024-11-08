package models

import (
	"github.com/feedcast-io/feedcast.db/types"
	"time"
)

type FeedTask struct {
	ID         int32
	Feed       Feed
	FeedId     int32
	Type       types.FeedTasks
	LastLaunch time.Time
}
