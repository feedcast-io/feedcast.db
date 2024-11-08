package enums

type FeedSynchroStatus int16

const (
	FeedSynchroStatusAlways FeedSynchroStatus = iota + 1
	FeedSynchroStatusOnSubscription
	FeedSynchroStatusNever
)
