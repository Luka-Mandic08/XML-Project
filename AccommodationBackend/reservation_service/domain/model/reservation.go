package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//Status types: Approved, Pending, Rejected(when host rejects), Canceled(when guest cancels)

type Reservation struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	AccommodationId string             `bson:"accommodation,omitempty"`
	Start           string             `bson:"start,omitempty"`
	End             string             `bson:"end,omitempty"`
	UserId          string             `bson:"user,omitempty"`
	NumberOfGuests  int32              `bson:"numberOfGuests,omitempty"`
	Status          string             `bson:"status,omitempty"`
	Price           float32            `bson:"price"`
}

func (r *Reservation) CalculateDuration() int32 {
	layout := "2006-01-02T15:04:05"
	StartDate, _ := time.Parse(layout, r.Start)
	EndDate, _ := time.Parse(layout, r.End)

	duration := EndDate.Sub(StartDate)
	durationHours := int32(duration.Hours()) / 24
	return durationHours
}
