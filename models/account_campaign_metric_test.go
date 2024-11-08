package models

import (
	feedcast_database "github.com/feedcast-io/feedcast.db"
	"testing"
)

func TestAccountCampaignMetric(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var metrics []AccountCampaignMetric

	if e := conn.Gorm.
		Order("id DESC").
		Preload("Campaign").
		Preload("Campaign.Account").
		Limit(1000).
		Find(&metrics).Error; e != nil {
		t.Fatal(e)
	}

	for _, m := range metrics {
		if 0 == m.Campaign.Account.ID {
			t.Error("AccountCampaignMetric should have a Campaign.Account")
		}
	}
}
