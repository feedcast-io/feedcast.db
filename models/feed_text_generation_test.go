package models

import (
	feedcast_database "github.com/feedcast-io/feedcast.db"
	"testing"
	"time"
)

func TestFeedTextGenerationModel(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var ftgs []FeedTextGeneration

	if e := conn.Gorm.
		Model(&FeedTextGeneration{}).
		Preload("Feed").
		Order("date DESC").
		Limit(1000).
		Scan(&ftgs).Error; e != nil {
		t.Error(e)
	}

	if 0 == len(ftgs) {
		t.Error("no results found")
	}

	for _, ftg := range ftgs {
		if ftg.Date.Before(time.Now().AddDate(-1, 0, 0)) {
			t.Error("invalid date")
		}
	}

	oldestGen := ftgs[len(ftgs)-1]
	newestGen := ftgs[0]

	found, e := CountFeedTextGeneration(conn.Gorm, oldestGen.FeedId, newestGen.Date)
	if e != nil {
		t.Error(e)
	}
	if 0 != found {
		t.Errorf("no generations for feed %d expected, got: %d", oldestGen.FeedId, found)
	}

	found, e = CountFeedTextGeneration(conn.Gorm, newestGen.FeedId, oldestGen.Date)
	if 0 == found {
		t.Errorf("no generations for feed %d expected at least 1", oldestGen.FeedId)
	}
}
