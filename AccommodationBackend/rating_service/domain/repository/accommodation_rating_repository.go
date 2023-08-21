package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"rating_service/domain/model"
)

type AccommodationMongoDBStore struct {
	accommodationRatings *mongo.Collection
}

func NewAccommodationMongoDBStore(client *mongo.Client) AccommodationRatingStore {
	accommodationRatings := client.Database(DATABASE).Collection(ACCOMMODATIONCOLLECTION)
	return &AccommodationMongoDBStore{
		accommodationRatings: accommodationRatings,
	}
}

func (a AccommodationMongoDBStore) Get(id primitive.ObjectID) (*model.AccommodationRating, error) {
	filter := bson.M{"_id": id}
	return a.filterOne(filter)
}

func (a AccommodationMongoDBStore) GetAllForAccommodation(accommodationId string) ([]*model.AccommodationRating, error) {
	filter := bson.M{"accommodationid": accommodationId}
	return a.filter(filter)
}

func (a AccommodationMongoDBStore) GetAverageScoreForAccommodation(accommodationId string) (float32, error) {
	pipeline := bson.A{
		bson.D{
			{"$match", bson.D{
				{"accommodationid", accommodationId},
			}},
		},
		bson.D{
			{"$group", bson.D{
				{"_id", nil},
				{"averageScore", bson.D{
					{"$avg", "$score"},
				}},
			}},
		},
	}
	cursor, err := a.accommodationRatings.Aggregate(context.TODO(), pipeline)
	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		return 0, err
	}
	if len(result) > 0 {
		return result[0]["averageScore"].(float32), nil
	}
	return 0.0, nil
}

func (a AccommodationMongoDBStore) Insert(accommodationRating *model.AccommodationRating) (*model.AccommodationRating, error) {
	result, err := a.accommodationRatings.InsertOne(context.TODO(), accommodationRating)
	if err != nil {
		return nil, err
	}
	accommodationRating.Id = result.InsertedID.(primitive.ObjectID)
	return accommodationRating, nil
}

func (a AccommodationMongoDBStore) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}
	result, err := a.accommodationRatings.DeleteOne(context.TODO(), filter)
	return result, err
}

func (a AccommodationMongoDBStore) Update(accommodationRating *model.AccommodationRating) (*mongo.UpdateResult, error) {
	update := bson.D{{"$set",
		bson.D{
			{"guestid", accommodationRating.GuestId},
			{"accommodationid", accommodationRating.AccommodationId},
			{"score", accommodationRating.Score},
			{"comment", accommodationRating.Comment},
		},
	}}
	result, err := a.accommodationRatings.UpdateByID(context.TODO(), accommodationRating.Id, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a AccommodationMongoDBStore) filter(filter interface{}) ([]*model.AccommodationRating, error) {
	cursor, err := a.accommodationRatings.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	return a.decode(cursor)
}

func (a AccommodationMongoDBStore) filterOne(filter interface{}) (rating *model.AccommodationRating, err error) {
	result := a.accommodationRatings.FindOne(context.TODO(), filter)
	err = result.Decode(&rating)
	return
}

func (a AccommodationMongoDBStore) decode(cursor *mongo.Cursor) (accommodationRatings []*model.AccommodationRating, err error) {
	for cursor.Next(context.TODO()) {
		var accommodationRating model.AccommodationRating
		err = cursor.Decode(&accommodationRating)
		if err != nil {
			return
		}
		accommodationRatings = append(accommodationRatings, &accommodationRating)
	}
	err = cursor.Err()
	return
}
