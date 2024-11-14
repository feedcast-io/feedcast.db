package models

import (
	"database/sql"
	"github.com/feedcast-io/feedcast.db/enums"
	"github.com/feedcast-io/feedcast.db/types"
	"gorm.io/gorm"
	"time"
)

type Feed struct {
	ID            int32
	Platform      *FeedPlatform
	Merchant      Merchant
	MerchantId    int32
	Name          sql.NullString `gorm:"type:text"`
	Url           sql.NullString
	Source        int16
	ItemLimit     sql.NullInt32
	SynchroStatus enums.FeedSynchroStatus

	Language   *Language
	LanguageId sql.NullInt32

	Country   *Country
	CountryId sql.NullInt32

	Domain   *FeedDomain
	DomainId *int32

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	DynamicLabel *FeedDynamicLabel

	SourceCredential   *MerchantUserCredential
	SourceCredentialId sql.NullInt32

	LastStat   *FeedStatDate
	LastStatId sql.NullInt32

	LastScore   *FeedScoreDate
	LastScoreId sql.NullInt32

	CurrentSubscription   *FeedSubscriptionHistory
	CurrentSubscriptionId sql.NullInt32

	SubscriptionItems []SubscriptionItem

	FeedTask []FeedTask
	Objects  []FeedObject

	Option *FeedOption
}

func (f *Feed) CanSynchro() bool {
	return f.SynchroStatus == enums.FeedSynchroStatusAlways ||
		f.SynchroStatus == enums.FeedSynchroStatusOnSubscription && f.CurrentSubscriptionId.Valid
}

func (f *Feed) GetObjectByType(objectType enums.FeedObjects) *FeedObject {
	for _, obj := range f.Objects {
		if obj.Type == objectType {
			return &obj
		}
	}

	return nil
}

func GetFeedInvoicePackCodes(conn *gorm.DB, feedId int32) ([]types.InvoiceProduct, error) {
	var feed Feed
	packs := make([]types.InvoiceProduct, 0)

	e := conn.
		Preload("SubscriptionItems").
		Preload("SubscriptionItems.InvoicePackPrice").
		Preload("SubscriptionItems.InvoicePackPrice.InvoicePack").
		Find(&feed, feedId).
		Error

	for _, it := range feed.SubscriptionItems {
		packs = append(packs, it.InvoicePackPrice.InvoicePack.Code)
	}

	return packs, e
}

func SaveFeedScore(conn *gorm.DB, feed *Feed, score *types.FeedScore) (*FeedScoreDate, error) {
	var entity FeedScoreDate

	date := time.Now()

	err := conn.
		Where("feed_id = ? AND date = ?", feed.ID, date.Format(time.DateOnly)).
		Attrs(FeedScoreDate{
			FeedId: feed.ID,
			Date:   date,
		}).FirstOrInit(&entity).
		Error

	if nil != err {
		return nil, err
	}

	if score.Feedcast.Valid {
		entity.FeedcastScore = score.Feedcast
	}
	if score.Google.Valid {
		entity.GoogleScore = score.Google
	}
	if score.Bing.Valid {
		entity.BingScore = score.Bing
	}
	if score.FreeListing.Valid {
		entity.FreeListingScore = score.FreeListing
	}
	if score.Meta.Valid {
		entity.MetaScore = score.Meta
	}

	if err := conn.Save(&entity).Error; nil != err {
		return nil, err
	}

	if feed.LastScoreId.Int32 != entity.ID {
		feed.LastScoreId.Int32 = entity.ID
		feed.LastScoreId.Valid = true

		if err := conn.Select("last_score_id").Updates(&feed).Error; nil != err {
			return nil, err
		}
	}

	return &entity, nil
}

func SaveFeedImport(conn *gorm.DB, feed *Feed, imported int32, startTime, endTime time.Time) (*FeedStatDate, error) {
	var entity FeedStatDate

	date := time.Now()

	err := conn.
		Where("feed_id = ? AND date = ?", feed.ID, date.Format(time.DateOnly)).
		Attrs(FeedStatDate{
			FeedId:              feed.ID,
			Date:                date,
			Imported:            sql.NullInt32{Valid: true},
			DateStartLastImport: sql.NullTime{Valid: true},
			DateEndLastImport:   sql.NullTime{Valid: true},
		}).FirstOrInit(&entity).
		Error

	if nil != err {
		return nil, err
	}

	entity.Imported.Int32 = imported
	entity.DateStartLastImport.Time = startTime
	entity.DateEndLastImport.Time = endTime

	if err := conn.Save(&entity).Error; nil != err {
		return nil, err
	}

	if feed.LastStatId.Int32 != entity.ID {
		feed.LastStatId.Int32 = entity.ID
		feed.LastStatId.Valid = true

		if err := conn.Select("last_stat_id").Updates(&feed).Error; nil != err {
			return nil, err
		}
	}

	return &entity, nil
}

func GetFeedAllProducts(conn *gorm.DB, feedId int32) (chan []FeedProduct, chan error) {
	ch, err := make(chan []FeedProduct), make(chan error)

	go func(batchSize int) {
		startId := int32(0)

		for {
			var products []FeedProduct

			if e := conn.
				Preload("Reference").
				Preload("CustomData").
				Preload("Text").
				Preload("Feed").
				Preload("Feed.Merchant").
				Preload("Feed.Language").
				Preload("Feed.Country").
				Preload("Category").
				Preload("Category.Texts").
				Preload("ProductBrand").
				Preload("Shipping").
				Preload("Currency").
				Preload("Various").
				Preload("Url").
				Where("feed_id = ? AND id > ?", feedId, startId).
				Order("id ASC").
				Limit(batchSize).
				Find(&products).Error; nil != e {
				err <- e
			}

			if foundInBatch := len(products); foundInBatch > 0 {
				ch <- products
				startId = products[foundInBatch-1].ID

				if foundInBatch != batchSize {
					break
				}
			}
		}

		close(ch)
		close(err)
	}(500)

	return ch, err
}
