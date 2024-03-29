package repository

import (
	"accommodation_service/domain/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
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

func (store *AvailabilityStore) FindAndGroupAvailableDates(dateFrom, dateTo time.Time, numberOfDays int) ([]string, []float64, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"date":        bson.M{"$gte": dateFrom, "$lte": dateTo},
				"isAvailable": true,
			},
		},
		{
			"$group": bson.M{
				"_id":         "$accommodationid",
				"sumQuantity": bson.M{"$sum": 1},
				"totalPrice":  bson.M{"$sum": "$price"},
			},
		},
		{
			"$match": bson.M{
				"sumQuantity": numberOfDays,
			},
		},
		{
			"$sort": bson.M{
				"_id": 1, // Sorting by _id in ascending order
			},
		},
	}

	cursor, err := store.availabilities.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, nil, err
	}
	results := []bson.M{}
	cursor.All(context.TODO(), &results)
	var ids []string
	var prices []float64
	for _, result := range results {
		ids = append(ids, result["_id"].(string))
		prices = append(prices, result["totalPrice"].(float64))
	}
	return ids, prices, nil
}

func (store *AvailabilityStore) GetAvailabilitiesForAccommodation(dateFrom, dateTo time.Time, accommodationid string) ([]*model.Availability, error) {
	filter := bson.M{
		"date":            bson.M{"$gte": dateFrom, "$lte": dateTo},
		"accommodationid": accommodationid,
	}

	sort := bson.D{{"date", 1}} // Sort by date in descending order (latest to newest)

	return store.filter(filter, sort)
}

func (store *AvailabilityStore) filter(filter interface{}, sort interface{}) ([]*model.Availability, error) {
	findOptions := options.Find().SetSort(sort) // Set the sort options

	cursor, err := store.availabilities.Find(context.TODO(), filter, findOptions)
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

func decodeAvailabilities(cursor *mongo.Cursor) (availabilities []*model.Availability, err error) {
	for cursor.Next(context.TODO()) {
		var availability model.Availability
		err = cursor.Decode(&availability)
		if err != nil {
			return
		}
		availabilities = append(availabilities, &availability)
	}
	err = cursor.Err()
	return
}

func (store *AvailabilityStore) CheckAvailabilityExists(id string, date time.Time) (*model.Availability, error) {
	filter := bson.M{"accommodationid": id, "date": date}
	return store.filterOne(filter)
}

func (store *AvailabilityStore) DeleteAllForAccommodation(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{"accommodationid": id}
	return store.availabilities.DeleteMany(context.TODO(), filter)
}

func (store *AvailabilityStore) GetByDateAndAccommodationAllToCancel(id string, date time.Time) (*model.Availability, error) {
	filter := bson.M{"accommodationid": id, "date": date, "isAvailable": false}
	return store.filterOne(filter)
}

func (store *AvailabilityStore) GetAllAvailabilitiesForRevert(id string, date time.Time) (*model.Availability, error) {
	filter := bson.M{"accommodationid": id, "date": date}
	return store.filterOne(filter)
}
