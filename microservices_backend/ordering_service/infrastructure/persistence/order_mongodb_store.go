package persistence

import (
	"context"
	"errors"
	"github.com/tamararankovic/microservices_demo/ordering_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "order"
	COLLECTION = "order"
)

type OrderMongoDBStore struct {
	orders *mongo.Collection
}

func NewOrderMongoDBStore(client *mongo.Client) domain.OrderStore {
	orders := client.Database(DATABASE).Collection(COLLECTION)
	return &OrderMongoDBStore{
		orders: orders,
	}
}

func (store *OrderMongoDBStore) Get(id primitive.ObjectID) (*domain.Order, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *OrderMongoDBStore) GetAll() ([]*domain.Order, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *OrderMongoDBStore) Insert(order *domain.Order) error {
	if order.Id.IsZero() {
		order.Id = primitive.NewObjectID()
	}
	result, err := store.orders.InsertOne(context.TODO(), order)
	if err != nil {
		return err
	}
	order.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *OrderMongoDBStore) DeleteAll() {
	store.orders.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *OrderMongoDBStore) UpdateStatus(order *domain.Order) error {
	result, err := store.orders.UpdateOne(
		context.TODO(),
		bson.M{"_id": order.Id},
		bson.D{
			{"$set", bson.D{{"status", order.Status}}},
		},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount != 1 {
		return errors.New("one document should've been updated")
	}
	return nil
}

func (store *OrderMongoDBStore) filter(filter interface{}) ([]*domain.Order, error) {
	cursor, err := store.orders.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *OrderMongoDBStore) filterOne(filter interface{}) (Order *domain.Order, err error) {
	result := store.orders.FindOne(context.TODO(), filter)
	err = result.Decode(&Order)
	return
}

func decode(cursor *mongo.Cursor) (orders []*domain.Order, err error) {
	for cursor.Next(context.TODO()) {
		var Order domain.Order
		err = cursor.Decode(&Order)
		if err != nil {
			return
		}
		orders = append(orders, &Order)
	}
	err = cursor.Err()
	return
}
