package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Accommodation struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Address   Address            `bson:"address"`
	Amenities []string           `bson:"amenities"`
	Images    []string           `bson:"image"`
	minGuests int                `bson:"minGuests"`
	maxGuests int                `bson:"maxGuests"`
	HostId    string             `bson:"hostid"`
}
