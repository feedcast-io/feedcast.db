package models

import (
	feedcast_database "github.com/feedcast-io/feedcast.db"
	"testing"
)

func TestFeedProductCustomDataModel(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var fpcd []FeedProductCustomData

	if e := conn.Gorm.
		Where("data LIKE '%title%'").
		Limit(100).Find(&fpcd).Error; e != nil {
		t.Error(e)
	}

	if 0 == len(fpcd) {
		t.Error("no feed product custom data")
	}

	// Test json properly decoded
	for _, item := range fpcd {
		if title, ok := item.Data["title"]; !ok || title == "" {
			t.Error("invalid feed product custom title")
		}
	}
}
