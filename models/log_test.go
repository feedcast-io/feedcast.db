package models

import (
	feedcast_database "github.com/feedcast-io/feedcast.db"
	"github.com/feedcast-io/feedcast.db/types"
	"testing"
)

func TestLogModels(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var feed Feed

	conn.Gorm.Find(&feed, 10956)
	tx := conn.Gorm.Begin()

	newLog, e := AddFeedLog(tx, types.LogTypeProductAiSuccess, &feed, map[string]interface{}{})
	if nil != e {
		t.Error(e)
	}

	if !newLog.MerchantUserId.Valid {
		t.Error("MerchantUserId not valid")
	}

	tx.Rollback()
}
