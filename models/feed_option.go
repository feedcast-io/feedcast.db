package models

type FeedOption struct {
	ID                int32
	FeedId            int32
	ExportComparator  bool
	AllowTrial        bool
	EnableFreeListing bool
}
