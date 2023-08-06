package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Availability struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	AccommodationId string             `bson:"accommodationid"`
	Date            time.Time          `bson:"date"`
	Price           float32            `bson:"price"`
	IsAvailable     bool               `bson:"isAvailable"`
}
