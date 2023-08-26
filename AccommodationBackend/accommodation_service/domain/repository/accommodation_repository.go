package repository

import (
	"accommodation_service/domain/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE   = "accommodation"
	COLLECTION = "accommodation"
)

type AccommodationMongoDBStore struct {
	accommodations *mongo.Collection
}

func NewAccommodationMongoDBStore(client *mongo.Client) AccommodationStore {
	accommodations := client.Database(DATABASE).Collection(COLLECTION)
	return &AccommodationMongoDBStore{
		accommodations: accommodations,
	}
}

func (store *AccommodationMongoDBStore) GetById(id primitive.ObjectID) (*model.Accommodation, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *AccommodationMongoDBStore) Insert(accommodation *model.Accommodation) (*model.Accommodation, error) {
	result, err := store.accommodations.InsertOne(context.TODO(), accommodation)
	if err != nil {
		return nil, err
	}
	accommodation.Id = result.InsertedID.(primitive.ObjectID)
	return accommodation, nil
}

func (store *AccommodationMongoDBStore) GetByAddress(address model.Address) (*model.Accommodation, error) {
	filter := bson.M{"address": address}
	return store.filterOne(filter)
}

func (store *AccommodationMongoDBStore) Update(accommodation *model.Accommodation) (*mongo.UpdateResult, error) {
	filter := bson.M{"userid": accommodation.HostId}
	update := bson.D{{"$set",
		bson.D{
			{"username", accommodation.Name},
		},
	}}
	result, err := store.accommodations.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (store *AccommodationMongoDBStore) Delete(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{"userid": id}
	result, err := store.accommodations.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (store *AccommodationMongoDBStore) DeleteAllForHost(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{"userid": id}
	result, err := store.accommodations.DeleteMany(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (store *AccommodationMongoDBStore) filter(filter interface{}) ([]*model.Accommodation, error) {
	cursor, err := store.accommodations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *AccommodationMongoDBStore) filterOne(filter interface{}) (accommodation *model.Accommodation, err error) {
	result := store.accommodations.FindOne(context.TODO(), filter)
	err = result.Decode(&accommodation)
	return
}

func decode(cursor *mongo.Cursor) (accommodations []*model.Accommodation, err error) {
	for cursor.Next(context.TODO()) {
		var accommodation model.Accommodation
		err = cursor.Decode(&accommodation)
		if err != nil {
			return
		}
		accommodations = append(accommodations, &accommodation)
	}
	err = cursor.Err()
	return
}

func (store *AccommodationMongoDBStore) GetAllByHostId(hostId string) ([]*model.Accommodation, error) {
	filter := bson.M{"hostid": hostId}
	result, err := store.accommodations.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	accommodations, err := decode(result)
	if err != nil {
		return nil, err
	}
	return accommodations, nil
}

func (store *AccommodationMongoDBStore) GetAll(page int) ([]*model.Accommodation, error) {
	skip := (page - 1) * 9
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(9))

	filter := bson.D{}
	result, err := store.accommodations.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}

	accommodations, err := decode(result)
	if err != nil {
		return nil, err
	}

	return accommodations, nil
}

func (store *AccommodationMongoDBStore) GetAllForHostByAccommodationId(id primitive.ObjectID) ([]string, string, error) {
	accommodation, err := store.GetById(id)
	if err != nil {
		return nil, "", err
	}
	allAccommodations, err := store.GetAllByHostId(accommodation.HostId)
	if err != nil {
		return nil, "", err
	}

	var results []string

	for _, accomm := range allAccommodations {
		results = append(results, accomm.Id.Hex())
	}

	return results, accommodation.HostId, nil
}
