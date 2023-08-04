package repository

import (
	"accommodation_service/domain/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AvailabilityStore struct {
	availabilities *mongo.Collection
}

func NewAvailabilityStoreStore(client *mongo.Client) *AvailabilityStore {
	availabilitys := client.Database(DATABASE).Collection("availability")
	return &AvailabilityStore{
		availabilities: availabilitys,
	}
}

func (store *AvailabilityStore) GetById(id primitive.ObjectID) (*model.Availability, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *AvailabilityStore) Insert(availability *model.Availability) (*model.Availability, error) {
	result, err := store.availabilities.InsertOne(context.TODO(), availability)
	if err != nil {
		return nil, err
	}
	availability.Id = result.InsertedID.(primitive.ObjectID)
	return availability, nil
}

func (store *AvailabilityStore) GetByDate(date primitive.DateTime) (*model.Availability, error) {
	filter := bson.M{"date": date}
	return store.filterOne(filter)
}

func (store *AvailabilityStore) Update(availability *model.Availability) (*mongo.UpdateResult, error) {
	filter := bson.M{"date": availability.Date}
	update := bson.D{{"$set",
		bson.D{
			{"price", availability.Price},
		},
	}}
	result, err := store.availabilities.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (store *AvailabilityStore) Delete(date primitive.DateTime) (*mongo.DeleteResult, error) {
	filter := bson.M{"date": date}
	result, err := store.availabilities.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (store *AvailabilityStore) filter(filter interface{}) ([]*model.Availability, error) {
	cursor, err := store.availabilities.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeAvailabilities(cursor)
}

func (store *AvailabilityStore) filterOne(filter interface{}) (availability *model.Availability, err error) {
	result := store.availabilities.FindOne(context.TODO(), filter)
	err = result.Decode(&availability)
	return
}

func decodeAvailabilities(cursor *mongo.Cursor) (availabilitys []*model.Availability, err error) {
	for cursor.Next(context.TODO()) {
		var availability model.Availability
		err = cursor.Decode(&availability)
		if err != nil {
			return
		}
		availabilitys = append(availabilitys, &availability)
	}
	err = cursor.Err()
	return
}
