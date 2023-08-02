package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Surname string             `bson:"surname"`
	Email   string             `bson:"email"`
	Address Address            `bson:"address"`
}
