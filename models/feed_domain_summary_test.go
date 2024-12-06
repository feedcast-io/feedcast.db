package models

import (
	feedcast_database "github.com/feedcast-io/feedcast.db"
	"log"
	"testing"
)

func TestGetFeedDomainSummaryModels(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var fds []FeedDomainSummary

	if e := conn.Gorm.
		Preload("FeedDomain").
		Limit(100).Find(&fds).Error; e != nil {
		t.Error(e)
	}

	if 0 == len(fds) {
		t.Error("no feed domain summary found")
	}

	for _, s := range fds {
		log.Println(s)
	}

	t.Error("END SUMMARY")
}
