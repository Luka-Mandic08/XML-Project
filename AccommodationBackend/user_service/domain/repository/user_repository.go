package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"user_service/domain/model"
)

const (
	DATABASE   = "user"
	COLLECTION = "user"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &UserMongoDBStore{
		users: users,
	}
}

func (store *UserMongoDBStore) Get(id primitive.ObjectID) (*model.User, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) GetAll() ([]*model.User, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *UserMongoDBStore) Insert(user *model.User) (*model.User, error) {
	result, err := store.users.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (store *UserMongoDBStore) Delete(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}
	result, error := store.users.DeleteOne(context.TODO(), filter)
	return result, error
}

func (store *UserMongoDBStore) Update(user *model.User) (*mongo.UpdateResult, error) {
	result, err := store.users.UpdateByID(context.TODO(), user.Id, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (store *UserMongoDBStore) filter(filter interface{}) ([]*model.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *UserMongoDBStore) filterOne(filter interface{}) (user *model.User, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	return
}

func decode(cursor *mongo.Cursor) (users []*model.User, err error) {
	for cursor.Next(context.TODO()) {
		var user model.User
		err = cursor.Decode(&user)
		if err != nil {
			return
		}
		users = append(users, &user)
	}
	err = cursor.Err()
	return
}
