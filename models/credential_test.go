package models

import (
	feedcast_database "github.com/feedcast-io/feedcast.db"
	"github.com/feedcast-io/feedcast.db/types"
	"testing"
)

func TestCredentialModel(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var creds []Credential

	if e := conn.Gorm.
		Where("code IN ?", []types.CredentialCode{
			types.CredentialCodeAdsFeedcast,
			types.CredentialCodeAdsMcc,
			types.CredentialCodeFacebookFeedcast,
		}).
		Find(&creds).Error; e != nil {
		t.Fatal(e)
	}

	// Non-empty Access or Refresh tokens should be present for FB / Google
	for _, cred := range creds {
		if len(cred.Data.RefreshToken) < 100 && len(cred.Data.AccessToken) < 100 {
			t.Errorf("missing refresh token or access token for %s", cred.Code)
		}
	}
}

func TestGetAdminCredential(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	c, e := GetAdminCredential(conn.Gorm, types.CredentialCodeAdsMcc)
	if nil != e {
		t.Error(e)
	}
	if nil == c {
		t.Errorf("nil Credential")
	}

	if len(c.Data.RefreshToken) == 0 && len(c.Data.AccessToken) == 0 && len(c.Data.CustomerId) == 0 {
		t.Errorf("missing refresh token or access token")
	}
}
