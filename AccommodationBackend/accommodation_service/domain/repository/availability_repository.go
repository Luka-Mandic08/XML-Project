package repository

import (
	"accommodation_service/domain/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type AvailabilityStore struct {
	availabilities *mongo.Collection
}

func NewAvailabilityStore(client *mongo.Client) *AvailabilityStore {
	availabilities := client.Database(DATABASE).Collection("availability")
	return &AvailabilityStore{
		availabilities: availabilities,
	}
}

func (store *AvailabilityStore) GetByDateAndAccommodation(id string, date time.Time) (*model.Availability, error) {
	filter := bson.M{"accommodationid": id, "date": date, "isAvailable": true}
	return store.filterOne(filter)
}

func (store *AvailabilityStore) Upsert(availability *model.Availability) error {
	filter := bson.M{"date": availability.Date, "accommodationid": availability.AccommodationId}
	update := bson.D{{"$set",
		bson.D{
			{"date", availability.Date},
			{"accommodationid", availability.AccommodationId},
			{"price", availability.Price},
			{"isAvailable", availability.IsAvailable},
		},
	}}
	opts := options.Update().SetUpsert(true)
	_, err := store.availabilities.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return err
	}
	return nil
}

func (store *AvailabilityStore) GetByDate(date primitive.DateTime) (*model.Availability, error) {
	filter := bson.M{"date": date}
	return store.filterOne(filter)
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
