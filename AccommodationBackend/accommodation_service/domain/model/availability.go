package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Availability struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	AccommodationId string             `bson:"accommodationid"`
	Date            primitive.DateTime `bson:"date"`
	Price           float32            `bson:"price"`
}
