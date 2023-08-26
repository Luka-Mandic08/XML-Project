package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Account struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
	UserId   string             `bson:"userid"`
	APIKey   APIKey             `bson:"apikey"`
}

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
