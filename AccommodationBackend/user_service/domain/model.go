package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id      primitive.ObjectID `bson:"_id"`
	Name    string             `bson:"name"`
	Surname string             `bson:"surname"`
}
