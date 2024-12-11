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

func TestGetFeedDomainSummaryArrayOfStrings(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var feedDomain FeedDomain
	if e := conn.Gorm.Last(&feedDomain).Error; nil != e {
		t.Fatal(e)
	}

	tx := conn.Gorm.Begin()

	var summary FeedDomainSummary
	if e := tx.
		Where(&FeedDomainSummary{
			FeedDomainId: feedDomain.ID,
			LanguageId:   1,
		}).
		Attrs(&FeedDomainSummary{
			FeedDomainId: feedDomain.ID,
			LanguageId:   1,
		}).
		Assign(&FeedDomainSummary{
			MainProducts:  []string{"product type 1", "product type 2"},
			SellingPoints: []string{"selling point 1", "selling point 2"},
		}).FirstOrCreate(&summary).Error; e != nil {
		t.Error(e)
	}

	tx.Rollback()
}
