package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OutstandingHost struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
}
