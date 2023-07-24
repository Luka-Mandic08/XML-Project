package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Color struct {
	Code string `bson:"code"`
}

type Product struct {
	Id    string `bson:"id"`
	Color Color  `bson:"color"`
}

type OrderItem struct {
	Product  Product `bson:"product"`
	Quantity uint16  `bson:"quantity"`
}

type OrderStatus int8

const (
	PendingApproval OrderStatus = iota
	Approved
	Cancelled
)

func (status OrderStatus) String() string {
	switch status {
	case PendingApproval:
		return "Pending Approval"
	case Approved:
		return "Approved"
	case Cancelled:
		return "Cancelled"
	}
	return "Unknown"
}

type Order struct {
	Id        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	Status    OrderStatus        `bson:"status"`
	Items     []OrderItem        `bson:"items"`
}
