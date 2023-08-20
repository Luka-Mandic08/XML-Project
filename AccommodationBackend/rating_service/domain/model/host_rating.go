package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type HostRating struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	GuestId string             `bson:"guestid"`
	HostId  string             `bson:"hostid"`
	Score   int32              `bson:"score"`
	Comment string             `bson:"comment"`
	Date    time.Time          `bson:"date"`
}
