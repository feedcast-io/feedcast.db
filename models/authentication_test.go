package models

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	feedcast_database "github.com/feedcast-io/feedcast.db"
	"testing"
)

func TestGetAuthByEmail(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	email := "romain@orixa-media.com"
	a, e := GetAuthByEmail(conn.Gorm, email)
	if e != nil {
		t.Error(e)
	}

	if a.Email.String != email {
		t.Error("Email does not match")
	}

	email = "jfklslkdskl@jklfdjklsdjflks.com"
	a, e = GetAuthByEmail(conn.Gorm, email)
	if e == nil {
		t.Error("GetAuthByEmail should return an error on invalid email")
	}

	var deletedEmail string
	if e := conn.Gorm.
		Model(&Authentication{}).
		Unscoped().
		Joins("INNER JOIN user u ON authentication.user_id = u.id AND u.deleted_at IS NOT NULL").
		Select("authentication.email").
		Where("authentication.email LIKE '%@%'").
		Limit(1).
		Find(&deletedEmail).Error; e != nil {
		t.Error(e)
	}

	a, e = GetAuthByEmail(conn.Gorm, deletedEmail)
	if nil == e {
		t.Error("GetAuthByEmail should return an error on invalid email")
	}

	a, e = GetAuthByEmail(conn.Gorm.Unscoped(), deletedEmail)
	if nil != e {
		t.Error("GetAuthByEmail should not return an error on invalid email but unscoped connection")
	}
	if a.Email.String != deletedEmail {
		t.Error("Email does not match for unscoped connection")
	}
}

func TestGetAuthByRefreshPasswordToken(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	tx := conn.Gorm.Begin()
	var lastCustomer MerchantUser
	if e := tx.
		Preload("Authentication").
		Last(&lastCustomer).Error; nil != e {
		t.Error(e)
	}

	token := md5.Sum([]byte(lastCustomer.Authentication.Email.String))
	if e := tx.Model(&lastCustomer.Authentication).Updates(Authentication{
		RefreshPasswordToken: sql.NullString{fmt.Sprintf("%x", token), true},
	}).Error; nil != e {
		t.Error(e)
	}

	a, e := GetAuthByRefreshPasswordToken(tx, lastCustomer.Authentication.Email.String)
	if nil != e {
		t.Error(e)
	}

	if a.ID != lastCustomer.Authentication.ID {
		t.Error("ID does not match")
	}

	tx.Rollback()
}

func TestGetStripeClient(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var reseller, merchant sql.NullInt32

	merchant, reseller, e := GetStripeClient(conn.Gorm, "fjddsjlfdsjsfld")
	if nil != e {
		t.Error(e)
	}

	if reseller.Valid || merchant.Valid {
		t.Error("GetStripeClient should return an error on invalid connection")
	}

	var m Merchant
	if e := conn.Gorm.
		Where("stripe_customer_id IS NOT NULL").
		Last(&m).Error; nil != e {
		t.Error(e)
	}

	merchant, reseller, e = GetStripeClient(conn.Gorm, m.StripeCustomerId.String)
	if nil != e {
		t.Error(e)
	}

	if !merchant.Valid {
		t.Error("GetStripeClient should return a merchant ID")
	}
	if reseller.Valid {
		t.Error("GetStripeClient should return an empty reseller")
	}
}
