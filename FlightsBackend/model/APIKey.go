package model

import "time"

type APIKey struct {
	Value       string    `bson:"value"`
	ValidTo     time.Time `bson:"validTo"`
	IsPermanent bool      `bson:"isPermanent"`
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
