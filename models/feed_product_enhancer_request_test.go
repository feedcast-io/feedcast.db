package models

import (
	feedcast_database "github.com/feedcast-io/feedcast.db"
	"testing"
)

func TestFeedProductEnhancerRequestModel(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var results []FeedProductEnhancerRequest
	if e := conn.Gorm.Unscoped().
		Preload("Feed").
		Limit(100).
		Find(&results).Error; e != nil {
		t.Error(e)
	}

	if len(results) == 0 {
		t.Error("no results found")
	}

	for _, result := range results {
		if 0 == len(result.Products) {
			t.Error("no products found")
		}

		if result.Feed.ID != result.FeedId {
			t.Error("no feed found")
		}
	}
}
