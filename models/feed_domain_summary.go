package models

import (
	"database/sql"
	"gorm.io/datatypes"
)

type FeedDomainSummary struct {
	ID int32

	LanguageId int32
	Language   Language

	FeedDomainId int32
	FeedDomain   FeedDomain

	Description   string
	Target        sql.NullString
	MainProducts  datatypes.JSONSlice[string] `gorm:"type:text"`
	SellingPoints datatypes.JSONSlice[string] `gorm:"type:text"`
}
