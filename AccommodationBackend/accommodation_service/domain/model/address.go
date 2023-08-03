package model

type Address struct {
	Street  string `bson:"street"`
	City    string `bson:"city"`
	Country string `bson:"country"`
}
