package models

import (
	"database/sql"
	"github.com/feedcast-io/feedcast.db/types"
)

type MerchantUserCredential struct {
	ID             int32
	MerchantUser   MerchantUser
	MerchantUserId int32
	Credential     Credential
	CredentialId   int32
	Data           types.CredentialData
	Title          string `gorm:"default:''"`
	LastCheck      sql.NullTime
	LastCheckValid sql.NullBool
}
