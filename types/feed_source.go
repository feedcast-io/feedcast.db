package types

type FeedSources int16

const (
	FeedSourceManual FeedSources = iota + 1
	FeedSourceShopify
	FeedSourcePrestashop
	FeedSourceWoocommerce
	FeedSourceGoogleSheets
	FeedSourceWebflow
	FeedSourceWizishop
	FeedSourceAutoGen
)
