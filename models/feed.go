package models

import (
	"cmp"
	"database/sql"
	"github.com/feedcast-io/feedcast.db/enums"
	"github.com/feedcast-io/feedcast.db/types"
	"gorm.io/gorm"
	"math/rand/v2"
	"slices"
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

	Option  *FeedOption
	Options types.FeedOptionMap
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
	entity := FeedScoreDate{
		FeedcastScore:    score.Feedcast,
		GoogleScore:      score.Google,
		BingScore:        score.Bing,
		FreeListingScore: score.FreeListing,
		MetaScore:        score.Meta,
	}
	var newFeed Feed

	now := time.Now()

	e := cmp.Or(
		conn.
			Where("feed_id = ? AND date = ?", feed.ID, now.Format(time.DateOnly)).
			Attrs(FeedScoreDate{
				FeedId: feed.ID,
				Date:   now,
			}).
			Assign(entity).
			FirstOrCreate(&entity).
			Error,
		conn.Model(Feed{}).
			Where("id = ?", feed.ID).
			Assign(&Feed{
				LastScoreId: sql.NullInt32{entity.ID, true},
			}).
			FirstOrCreate(&newFeed).
			Error,
	)

	return &entity, e
}

func SaveFeedImport(conn *gorm.DB, feed *Feed, imported int32, startTime, endTime time.Time) (*FeedStatDate, error) {
	entity := FeedStatDate{
		Imported:            sql.NullInt32{imported, true},
		DateStartLastImport: sql.NullTime{startTime, true},
		DateEndLastImport:   sql.NullTime{endTime, true},
	}

	date := time.Now()

	e := cmp.Or(
		conn.
			Where("feed_id = ? AND date = ?", feed.ID, date.Format(time.DateOnly)).
			Attrs(FeedStatDate{
				FeedId: feed.ID,
				Date:   date,
			}).
			Assign(entity).
			FirstOrCreate(&entity).
			Error,
		conn.Where("id = ?", feed.ID).
			Assign(&Feed{
				LastStatId: sql.NullInt32{entity.ID, true},
			}).
			FirstOrCreate(feed).
			Error,
	)

	return &entity, e
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
				Where("feed_id = ? AND id > ? AND reference_id > 0", feedId, startId).
				Order("id ASC").
				Limit(batchSize).
				Find(&products).Error; nil != e {
				err <- e
			}

			if foundInBatch := len(products); foundInBatch >= 0 {
				ch <- products

				if foundInBatch != batchSize {
					break
				}
				startId = products[foundInBatch-1].ID
			}
		}

		close(ch)
		close(err)
	}(500)

	return ch, err
}

func GetFeedTaskToDo(conn *gorm.DB, taskType types.FeedTasks, maxLastLaunch time.Time) (int32, error) {
	var feeds []int32

	if err := conn.
		Model(Feed{}).
		Select("feed.id").
		Joins("LEFT JOIN feed_task ON feed.id = feed_task.feed_id AND feed_task.type = ?", taskType).
		Where("feed_task.last_launch IS NULL").
		Limit(100).
		Scan(&feeds).Error; nil != err {
		return 0, err
	}

	if 0 == len(feeds) {
		if err := conn.
			Model(Feed{}).
			Select("feed.id").
			Joins("JOIN feed_task ON feed.id = feed_task.feed_id AND feed_task.type = ?", taskType).
			Where("feed_task.last_launch < ?", maxLastLaunch).
			Order("feed_task.last_launch").
			Limit(10).
			Scan(&feeds).Error; nil != err {
			return 0, err
		}
	}

	var feedTask FeedTask

	if len(feeds) > 0 {
		slices.SortFunc(feeds, func(a, b int32) int {
			odd := 1 == rand.Int()%2
			if odd {
				return -1 * rand.Int()
			} else {
				return rand.Int()
			}
		})
		if err := conn.
			Where(FeedTask{
				FeedId: feeds[0],
				Type:   taskType,
			}).
			Assign(FeedTask{LastLaunch: time.Now()}).
			FirstOrCreate(&feedTask).Error; nil != err {
			return 0, err
		}

		return feeds[0], nil
	}

	return 0, nil
}
