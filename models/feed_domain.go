package models

import "github.com/feedcast-io/feedcast.db/types"

type FeedDomain struct {
	ID     int32
	Domain string `gorm:"size:96,unique"`
	Status types.FeedDomainStatus
}
