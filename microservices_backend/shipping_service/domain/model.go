package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderStatus int8

const (
	Scheduled OrderStatus = iota
	InTransport
	Delivered
	Cancelled
)

func (status OrderStatus) String() string {
	switch status {
	case Scheduled:
		return "Pending Approval"
	case InTransport:
		return "Approved"
	case Delivered:
		return "Delivered"
	case Cancelled:
		return "Cancelled"
	}
	return "Unknown"
}

type Order struct {
	Id              primitive.ObjectID `bson:"_id"`
	Status          OrderStatus        `bson:"status"`
	ShippingAddress string             `bson:"shipping_address"`
}
