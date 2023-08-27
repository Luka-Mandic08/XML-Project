package model

import "time"

type APIKey struct {
	Value       string    `bson:"value" json:"value"`
	ValidTo     time.Time `bson:"validTo" json:"validTo"`
	IsPermanent bool      `bson:"isPermanent" json:"isPermanent"`
}

func (key *APIKey) IsValid() bool {
	if key.IsPermanent {
		return true
	}

	today := time.Now()
	if today.Before(key.ValidTo) {
		return true
	}

	return false
}
