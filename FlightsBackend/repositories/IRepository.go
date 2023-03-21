package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type IRepository interface {
	Disconnect(ctx context.Context) error
	Ping()
	getCollection() *mongo.Collection
}
