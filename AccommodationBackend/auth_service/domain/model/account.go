package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
	UserId   string             `bson:"userid"`
	APIKey   string             `bson:"apikey"`
}
