package models

type FeedDomainSummary struct {
	ID           int32
	FeedDomainId int32
	FeedDomain   FeedDomain
	Description  string
}
