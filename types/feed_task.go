package types

type FeedTasks int16

const (
	FeedTaskCheckAccountLinked FeedTasks = iota + 1
	FeedTaskCheckAccountPaiment
	FeedTaskCheckCampaigns
	FeedTaskCheckCampaignPerfs
	FeedTaskDownload
	FeedImportProduct
	FeedCatalogSynchro
	// Import recent account metrics
	FeedImportAccountMetrics
	// Import old account metrics
	FeedImportAccountMetricsOld
	FeedTaskFeedSynchro
)
