package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reservation_service/domain/model"
)

type OutstandingHostMongoDBStore struct {
	outstandingHosts *mongo.Collection
}

func NewOutstandingHostMongoDBStore(client *mongo.Client) *OutstandingHostMongoDBStore {
	outstandingHosts := client.Database(DATABASE).Collection(HOSTCOLLECTION)
	return &OutstandingHostMongoDBStore{
		outstandingHosts: outstandingHosts,
	}
}

func (store *OutstandingHostMongoDBStore) Get(id primitive.ObjectID) (*model.OutstandingHost, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *OutstandingHostMongoDBStore) GetAll() ([]*model.OutstandingHost, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *OutstandingHostMongoDBStore) Insert(outstandingHost *model.OutstandingHost) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": outstandingHost.Id}
	opts := options.Update().SetUpsert(true)

	result, err := store.outstandingHosts.UpdateOne(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (store *OutstandingHostMongoDBStore) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}
	result, err := store.outstandingHosts.DeleteOne(context.TODO(), filter)
	return result, err
}

func (store *OutstandingHostMongoDBStore) filter(filter interface{}) ([]*model.OutstandingHost, error) {
	cursor, err := store.outstandingHosts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return store.decode(cursor)
}

func (store *OutstandingHostMongoDBStore) filterOne(filter interface{}) (user *model.OutstandingHost, err error) {
	result := store.outstandingHosts.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	return
}

func (store *OutstandingHostMongoDBStore) decode(cursor *mongo.Cursor) (outstandingHosts []*model.OutstandingHost, err error) {
	for cursor.Next(context.TODO()) {
		var outstandingHost model.OutstandingHost
		err = cursor.Decode(&outstandingHost)
		if err != nil {
			return
		}
		outstandingHosts = append(outstandingHosts, &outstandingHost)
	}
	err = cursor.Err()
	return
}
