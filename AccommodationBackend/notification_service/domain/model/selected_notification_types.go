package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NotificationTypes Type - ReservationCreated, ReservationCanceled, HostRated, AccommodationRated, OutstandingHostStatus, ReservationApprovedOrDenied

type SelectedNotificationTypes struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	UserId        string             `bson:"userId"`
	SelectedTypes []string           `bson:"selectedTypes"`
}
