package types

import (
	"database/sql/driver"
	"encoding/json"
)

type ArrayString []string

func (r *ArrayString) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *ArrayString) Scan(value interface{}) error {
	bytes, _ := value.([]byte)

	var res []string
	err := json.Unmarshal(bytes, &res)

	if nil == err {
		*r = res
	}

	return nil
}
