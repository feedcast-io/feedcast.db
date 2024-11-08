package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID                   int32
	Code                 string `gorm:"size:50"`
	Name                 string `gorm:"size:100"`
	IsChild              bool
	LastCheck            sql.NullTime
	MinReportDate        sql.NullTime
	MaxReportDate        sql.NullTime
	Credential           Credential
	CredentialId         int32
	MerchantCredential   *MerchantUserCredential
	MerchantCredentialId sql.NullInt32
	DeletedAt            gorm.DeletedAt
}

type CampaignReport struct {
	CampaignId      string `gorm:"column:code"`
	DateMin         time.Time
	DateMax         time.Time
	Impressions     int32
	Conversions     int32
	ConversionValue int32
	Spent           int32
}

func GetAccountCampaignReportingSummary(conn *gorm.DB, accountId int32, feedId int32, from time.Time, to time.Time) ([]CampaignReport, error) {
	var reports []CampaignReport

	q := conn.
		Model(AccountCampaign{}).
		Joins("INNER JOIN account_campaign_metric acm ON acm.campaign_id = account_campaign.id AND account_campaign.feed_creation_id = ? AND acm.date >= ? AND acm.date <= ?", feedId, from.Format(time.DateOnly), to.Format(time.DateOnly)).
		Select("code, MIN(acm.date) AS date_min, MAX(acm.date) AS date_max, SUM(acm.impressions) AS impressions, SUM(acm.conversions) AS conversions, SUM(acm.spent) AS spent, SUM(acm.conversion_value) AS conversion_value").
		Where(AccountCampaign{
			AccountId: accountId,
		}).
		Group("code").
		Scan(&reports)

	return reports, q.Error
}
