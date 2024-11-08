package models

import (
	"github.com/feedcast-io/feedcast.db/types"
	"time"
)

type AccountCampaignMetric struct {
	ID              int32
	CampaignId      int32
	Campaign        AccountCampaign
	Date            time.Time
	Device          types.DeviceType
	Impressions     int32
	Clicks          int32
	Conversions     int32
	Spent           int32
	ConversionValue int64
}
