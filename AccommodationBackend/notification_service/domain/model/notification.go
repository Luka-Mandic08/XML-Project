package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Notification struct {
	Id               primitive.ObjectID `bson:"_id,omitempty"`
	NotificationText string             `bson:"notificationText"`
	IsAcknowledged   bool               `bson:"isAcknowledged"`
	HostId           string             `bson:"hostId"`
	GuestId          string             `bson:"hostId"`
	DateCreated      time.Time          `bson:"dateCreated"`
}
