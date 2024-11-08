package models

import "database/sql"

type AccountCampaign struct {
	ID             int32
	Code           string
	FeedCreationId sql.NullInt32
	AccountId      int32
	Account        Account
}
