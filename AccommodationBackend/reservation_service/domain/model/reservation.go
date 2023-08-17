package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Reservation struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	AccommodationId string             `bson:"accommodation,omitempty"`
	Start           time.Time          `bson:"start,omitempty"`
	End             time.Time          `bson:"end,omitempty"`
	UserId          string             `bson:"user,omitempty"`
	NumberOfGuests  int32              `bson:"numberOfGuests,omitempty"`
	Status          string             `bson:"status,omitempty"`
	Price           float32            `bson:"price"`
}