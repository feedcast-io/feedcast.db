package models

import (
	"database/sql"
	feedcast_database "github.com/feedcast-io/feedcast.db"
	"github.com/feedcast-io/feedcast.db/types"
	"slices"
	"testing"
	"time"
)

func TestSaveFeedScore(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var feed Feed

	// Test on last active feed
	conn.Gorm.Last(&feed)

	feedScore := types.FeedScore{
		Feedcast:    sql.NullFloat64{12.33, true},
		Google:      sql.NullFloat64{78.567, true},
		FreeListing: sql.NullFloat64{89.01, true},
	}

	tx := conn.Gorm.Begin()

	fsd, err := SaveFeedScore(tx, &feed, &feedScore)
	if err != nil {
		t.Error(err)
	}

	if fsd.FeedcastScore.Float64 != feedScore.Feedcast.Float64 {
		t.Error("saved feed score does not match feedcast score")
	}

	tx.Rollback()
}

func TestUpdateFeedScore(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var feed Feed
	conn.Gorm.Last(&feed)

	tx := conn.Gorm.Begin()
	fsd, err := SaveFeedImport(tx, &feed, 4567, time.Now().Add(time.Hour), time.Now())
	if err != nil {
		t.Error(err)
	}
	if fsd.Imported.Int32 != 4567 {
		t.Error("saved feed import does not match import")
	}

	fsd2, _ := SaveFeedImport(tx, &feed, 1234, time.Now().Add(time.Hour), time.Now())
	if fsd2.Imported.Int32 != 1234 {
		t.Error("saved feed import does not match import")
	}

	if fsd.ID != fsd2.ID {
		t.Errorf("saved feed id does not match import (%d/%d)", fsd.ID, fsd2.ID)
	}

	tx.Rollback()
}

func TestGetFeedAllProducts(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var err error
	var batch []FeedProduct
	var found int32

	ch, e := GetFeedAllProducts(conn.Gorm, 10956)

	for ok1, ok2 := true, true; ok1 || ok2; {
		select {
		case batch, ok1 = <-ch:
			if ok1 {
				for _, p := range batch {
					found++
					if nil != p.CustomData {
						if !p.ReferenceId.Valid {
							t.Error("missing reference")
						}
						if nil == p.Text {
							t.Error("missing text")
						}
						if nil == p.Url {
							t.Error("missing url")
						}
						if nil == p.Feed.Language {
							t.Error("missing language")
						}
						if p.CategoryId.Valid && nil == p.Category {
							t.Error("missing category")
						}
						if p.ProductBrandId.Valid && nil == p.ProductBrand {
							t.Error("missing brand")
						}
					}
				}
			}
			break

		case err, ok2 = <-e:
			if ok2 {
				t.Error(err)
			}
			break
		}
	}

	if 0 == found {
		t.Error("no products found")
	}
}

func TestGetFeedInvoicePackCodes(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var subItem SubscriptionItem
	if e := conn.Gorm.Last(&subItem).Error; nil != e {
		t.Error(e)
	}

	codes, e := GetFeedInvoicePackCodes(conn.Gorm, subItem.FeedId.Int32)
	if e != nil {
		t.Error(e)
	}
	if len(codes) == 0 {
		t.Error("no code found")
	}

	if !slices.Contains(codes, types.InvoiceProductPackStarter) && !slices.Contains(codes, types.InvoiceProductPackPremium) {
		t.Error("invalid pack codes")
	}
}
