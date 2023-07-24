package startup

import (
	"github.com/tamararankovic/microservices_demo/ordering_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var orders = []*domain.Order{
	{
		Id:        getObjectId("623b0cc336a1d6fd8c1cf0f6"),
		Status:    domain.Approved,
		CreatedAt: time.Now(),
		Items: []domain.OrderItem{
			{
				Product: domain.Product{
					Id:    "623b0cc3a34d25d8567f9f82",
					Color: domain.Color{Code: "R"},
				},
				Quantity: 5,
			},
			{
				Product: domain.Product{
					Id:    "623b0cc3a34d25d8567f9f83",
					Color: domain.Color{Code: "G"},
				},
				Quantity: 3,
			},
		},
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
