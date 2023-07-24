package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderStore interface {
	Get(id primitive.ObjectID) (*Order, error)
	GetAll() ([]*Order, error)
	Insert(product *Order) error
	DeleteAll()
	UpdateStatus(order *Order) error
}
