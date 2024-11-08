package models

import (
	"github.com/feedcast-io/feedcast.db/types"
	"gorm.io/gorm"
)

type Credential struct {
	ID    int32
	Code  types.CredentialCode `gorm:"size:32,unique"`
	Email string               `gorm:"size:255,unique"`
	Data  types.CredentialData `gorm:"type:text"`
}

func GetAdminCredential(conn *gorm.DB, credentialCode types.CredentialCode) (*Credential, error) {
	var credential Credential

	if e := conn.
		Where(Credential{Code: credentialCode}).
		First(&credential).
		Error; nil != e {
		return nil, e
	}

	return &credential, nil
}
