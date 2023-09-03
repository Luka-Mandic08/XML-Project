package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AccommodationRating struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	GuestId         string             `bson:"guestid"`
	AccommodationId string             `bson:"accommodationid"`
	Score           int32              `bson:"score"`
	Comment         string             `bson:"comment"`
	Date            time.Time          `bson:"date"`
}

type RecommendedAccommodation struct {
	AccommodationID string
	AverageScore    float32
}
