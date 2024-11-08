package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type CredentialData struct {
	ApiKey       string `json:"api_key"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	CustomerId   string `json:"customer_id"`
	Expiration   string `json:"expiration"`
}

func (r *CredentialData) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *CredentialData) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var res CredentialData
	err := json.Unmarshal(bytes, &res)

	if nil == err {
		*r = res
	}

	// Silent error if unable to decode, map will just be empty
	return nil
}
