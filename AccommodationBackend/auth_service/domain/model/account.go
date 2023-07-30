package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username" unique`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
	UserID   string             `bson:"userid"`
}
