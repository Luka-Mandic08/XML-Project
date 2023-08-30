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
	APIKey   APIKey             `bson:"apikey,omitempty"`
}

type APIKey struct {
	Value       string    `bson:"value,omitempty"`
	ValidTo     time.Time `bson:"validTo,omitempty"`
	IsPermanent bool      `bson:"isPermanent,omitempty"`
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
