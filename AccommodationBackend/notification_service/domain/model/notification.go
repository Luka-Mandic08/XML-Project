package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// NotificationTypes Type - ReservationCreated, ReservationCanceled, HostRated, AccommodationRated, ReservationApprovedOrDenied

type Notification struct {
	Id               primitive.ObjectID `bson:"_id,omitempty"`
	NotificationText string             `bson:"notificationText"`
	IsAcknowledged   bool               `bson:"isAcknowledged"`
	UserId           string             `bson:"userId"`
	DateCreated      time.Time          `bson:"dateCreated"`
	Type             string             `bson:"type"`
}
