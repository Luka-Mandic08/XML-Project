package repository

import (
	"accommodation_service/domain/model"
	accommodation "common/proto/accommodation_service"
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

func (store *AccommodationMongoDBStore) GetForSearch(id primitive.ObjectID, req *accommodation.SearchRequest, hostIds []string) (*model.Accommodation, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"_id":             id,
				"address.city":    bson.M{"$regex": req.GetCity(), "$options": "i"},
				"address.country": bson.M{"$regex": req.GetCountry(), "$options": "i"},
				"minGuests":       bson.M{"$lte": req.GetNumberOfGuests()},
				"maxGuests":       bson.M{"$gte": req.GetNumberOfGuests()},
			},
		},
	}
	//TODO: Make this work with for loop probably
	/*if len(req.GetAmenities()) > 0 {
		amenityRegexes := []bson.M{}
		for _, amenity := range req.GetAmenities() {
			println(amenity)
			pattern := fmt.Sprintf(".%s.", regexp.QuoteMeta(amenity))
			amenityRegexes = append(amenityRegexes, bson.M{"$regex": pattern, "$options": "i"})
		}
		pipeline = append(pipeline, bson.M{"$match": bson.M{"$all": amenityRegexes}})
	}*/

	if len(hostIds) > 0 {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"hostid": bson.M{"$in": hostIds}}})
	}

	cursor, err := store.accommodations.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}

	var results []*model.Accommodation
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}

	return results[0], nil
}
