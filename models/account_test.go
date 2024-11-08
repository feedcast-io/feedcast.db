package models

import (
	feedcast_database "github.com/feedcast-io/feedcast.db"
	"testing"
)

func TestGetAccountCampaignReportingSummary(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var accs []Account
	if e := conn.Gorm.
		Preload("MerchantCredential").
		Where("merchant_credential_id > 0").
		Limit(100).Find(&accs).Error; e != nil {
		t.Fatal(e)
	}

	for _, acc := range accs {
		if nil == acc.MerchantCredential {
			t.Error("expected merchant_credential to be populated")
		}
	}
}
