package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"rating_service/domain/model"
)

const (
	DATABASE                = "rating"
	HOSTCOLLECTION          = "hostRating"
	ACCOMMODATIONCOLLECTION = "accomodationRating"
)

type HostRatingMongoDBStore struct {
	hostRatings *mongo.Collection
}

func NewHostRatingMongoDBStore(client *mongo.Client) HostRatingStore {
	hostRatings := client.Database(DATABASE).Collection(HOSTCOLLECTION)
	return &HostRatingMongoDBStore{
		hostRatings: hostRatings,
	}
}

func (store *HostRatingMongoDBStore) Get(id primitive.ObjectID) (*model.HostRating, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *HostRatingMongoDBStore) GetAllForHost(hostId string) ([]*model.HostRating, error) {
	filter := bson.M{"hostid": hostId}
	return store.filter(filter)
}

func (store *HostRatingMongoDBStore) GetAverageScoreForHost(hostId string) (float32, error) {
	pipeline := bson.A{
		bson.D{
			{"$match", bson.D{
				{"hostid", hostId},
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
	cursor, err := store.hostRatings.Aggregate(context.TODO(), pipeline)
	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		return 0, err
	}
	if len(result) > 0 {
		return result[0]["averageScore"].(float32), nil
	}
	return 0.0, nil
}

func (store *HostRatingMongoDBStore) Insert(hostRating *model.HostRating) (*model.HostRating, error) {
	result, err := store.hostRatings.InsertOne(context.TODO(), hostRating)
	if err != nil {
		return nil, err
	}
	hostRating.Id = result.InsertedID.(primitive.ObjectID)
	return hostRating, nil
}

func (store *HostRatingMongoDBStore) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}
	result, err := store.hostRatings.DeleteOne(context.TODO(), filter)
	return result, err
}

func (store *HostRatingMongoDBStore) Update(hostRating *model.HostRating) (*mongo.UpdateResult, error) {
	update := bson.D{{"$set",
		bson.D{
			{"guestid", hostRating.GuestId},
			{"hostid", hostRating.HostId},
			{"score", hostRating.Score},
			{"comment", hostRating.Comment},
		},
	}}
	result, err := store.hostRatings.UpdateByID(context.TODO(), hostRating.Id, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (store *HostRatingMongoDBStore) filter(filter interface{}) ([]*model.HostRating, error) {
	cursor, err := store.hostRatings.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return store.decode(cursor)
}

func (store *HostRatingMongoDBStore) filterOne(filter interface{}) (user *model.HostRating, err error) {
	result := store.hostRatings.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	return
}

func (store *HostRatingMongoDBStore) decode(cursor *mongo.Cursor) (hostRatings []*model.HostRating, err error) {
	for cursor.Next(context.TODO()) {
		var hostRating model.HostRating
		err = cursor.Decode(&hostRating)
		if err != nil {
			return
		}
		hostRatings = append(hostRatings, &hostRating)
	}
	err = cursor.Err()
	return
}
