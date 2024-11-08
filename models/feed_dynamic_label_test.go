package models

import (
	feedcast_database "github.com/feedcast-io/feedcast.db"
	"testing"
)

func TestFeedDynamicLabelModel(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var labels []FeedDynamicLabel

	if e := conn.Gorm.
		Limit(100).
		Preload("Feed").
		Where("LENGTH(config) > 20").
		Find(&labels).Error; e != nil {
		t.Error(e)
	}

	if 0 == len(labels) {
		t.Error("no labels found for test")
	}

	for _, label := range labels {
		if len(label.Config.Rules) < 1 {
			t.Errorf("No rules found for %s (%d)", label.Feed.Name.String, label.Feed.ID)
		}
		for i, rule := range label.Config.Rules {
			if 0 == len(rule.LabelValue) {
				t.Errorf("No label value found for rule #%d for %s (%d)", i, label.Feed.Name.String, label.Feed.ID)
			}
		}
	}
}
