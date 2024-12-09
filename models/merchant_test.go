package models

import (
	"database/sql"
	feedcast_database "github.com/feedcast-io/feedcast.db"
	"testing"
)

func TestMerchantSpam(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var merchants []Merchant

	if e := conn.Gorm.Unscoped().Where(Merchant{
		IsSpamSuspicion: sql.NullBool{true, true},
	}).Limit(10).Find(&merchants).Error; nil != e {
		t.Error(e)
	}
}
