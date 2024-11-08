package models

import (
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Authentication struct {
	ID                   int32
	Email                sql.NullString `gorm:"index,unique,size:64"`
	Password             sql.NullString `gorm:"size:255"`
	Token                sql.NullString `gorm:"index,unique,size:32"`
	RefreshPasswordToken sql.NullString `gorm:"index,unique,size:32"`
	GuestToken           sql.NullString `gorm:"size:32"`
	LastConnection       sql.NullTime
	AffiliationCode      sql.NullString `gorm:"size:16"`
	Validated            sql.NullBool   `gorm:"default:true"`
	Guest                *Authentication
	GuestID              sql.NullInt32
	UserID               sql.NullInt32
	User                 *User
	ResellerID           sql.NullInt32
	Reseller             *Reseller
	MerchantUserID       sql.NullInt32
	MerchantUser         *MerchantUser
	CommercialID         sql.NullInt32
	Commercial           *Commercial
}

func GetAuthByEmail(conn *gorm.DB, email string) (*Authentication, error) {
	var auth *Authentication

	if e := conn.
		Model(&auth).
		Preload("User").
		Preload("Reseller").
		Preload("MerchantUser").
		Preload("MerchantUser.Merchant").
		Preload("Commercial").
		Where("email = ?", email).First(&auth).Error; nil != e {
		return nil, e
	}

	if auth.ResellerID.Valid && nil == auth.Reseller {
		return nil, errors.New("deleted reseller")
	}

	if auth.UserID.Valid && nil == auth.User {
		return nil, errors.New("deleted user")
	}

	if auth.MerchantUserID.Valid && nil == auth.MerchantUser {
		return nil, errors.New("deleted merchant")
	}

	if auth.CommercialID.Valid && nil == auth.Commercial {
		return nil, errors.New("deleted commercial")
	}

	return auth, nil
}

func GetAuthByRefreshPasswordToken(conn *gorm.DB, refreshToken string) (*Authentication, error) {
	var auth Authentication

	if e := conn.
		Preload("User").
		Preload("Commercial").
		Preload("Reseller").
		Preload("MerchantUser").
		Preload("MerchantUser.Merchant").
		Where("refresh_password_token = ?", fmt.Sprintf("%x", md5.Sum([]byte(refreshToken)))).
		First(&auth).Error; nil != e {
		return nil, e
	}

	return &auth, nil
}

// Get Merchant  or Reseller from Stripe customer ID
func GetStripeClient(conn *gorm.DB, stripeClientId string) (sql.NullInt32, sql.NullInt32, error) {
	var merchant Merchant
	var merchantId, resellerId sql.NullInt32

	if err := conn.
		Preload("Reseller").
		Where(Merchant{StripeCustomerId: sql.NullString{stripeClientId, true}}).
		Limit(1).
		Find(&merchant).Error; nil != err {
		return merchantId, resellerId, err
	}

	if merchant.ID > 0 {
		merchantId = sql.NullInt32{merchant.ID, true}
		if merchant.IsDirectCustomer() && merchant.ResellerId.Valid {
			resellerId = sql.NullInt32{merchant.ResellerId.Int32, true}
		}
	}

	if !merchantId.Valid {
		var reseller Reseller
		if err := conn.
			Where(Reseller{StripeCustomerId: sql.NullString{stripeClientId, true}}).
			Limit(1).
			Find(&reseller).Error; nil != err {
		}

		if reseller.ID > 0 {
			resellerId = sql.NullInt32{reseller.ID, true}
		}
	}

	return merchantId, resellerId, nil
}
