package types

import (
	"database/sql/driver"
	"encoding/json"
)

type FeedOptionMap struct {
	BlacklistedBrandTitle string `json:"blacklisted_brand_title"`
}

func (r *FeedOptionMap) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *FeedOptionMap) Scan(value interface{}) error {
	bytes, _ := value.([]byte)

	var res FeedOptionMap
	err := json.Unmarshal(bytes, &res)

	if nil == err {
		*r = res
	}

	return nil
}
