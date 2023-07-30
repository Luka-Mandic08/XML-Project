package startup

import (
	"auth_service/domain/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var accounts = []*model.Account{}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
