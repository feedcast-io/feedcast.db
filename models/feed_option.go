package models

import "database/sql"

type FeedOption struct {
	ID                    int32
	FeedId                int32
	ExportComparator      bool
	AllowTrial            bool
	EnableFreeListing     bool
	BlacklistedBrandTitle sql.NullString
}
