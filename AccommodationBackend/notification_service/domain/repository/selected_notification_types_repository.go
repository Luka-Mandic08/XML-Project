package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"notification_service/domain/model"
)

type SelectedNotificationTypesMongoDBStore struct {
	selectedTypes *mongo.Collection
}

func NewSelectedNotificationTypesMongoDBStore(client *mongo.Client) SelectedNotificationTypesStore {
	selectedTypes := client.Database(DATABASE).Collection(NOTIFICATIONCOLLECTION)
	return &SelectedNotificationTypesMongoDBStore{
		selectedTypes: selectedTypes,
	}
}

func (store *SelectedNotificationTypesMongoDBStore) GetByUserId(userId string) (*model.SelectedNotificationTypes, error) {
	filter := bson.M{"userId": userId}
	return store.filterOne(filter)
}

func (store *SelectedNotificationTypesMongoDBStore) Insert(selectedTypes *model.SelectedNotificationTypes) (*model.SelectedNotificationTypes, error) {
	result, err := store.selectedTypes.InsertOne(context.TODO(), selectedTypes)
	if err != nil {
		return nil, err
	}
	selectedTypes.Id = result.InsertedID.(primitive.ObjectID)
	return selectedTypes, nil
}

func (store *SelectedNotificationTypesMongoDBStore) Update(userId string, selectedTypes []string) (*model.SelectedNotificationTypes, error) {
	filter := bson.M{"userId": userId}

	update := bson.D{{"$set",
		bson.D{
			{"selectedTypes", selectedTypes},
		},
	}}
	_, err := store.selectedTypes.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	selectedNotificationTypes, _ := store.GetByUserId(userId)
	return selectedNotificationTypes, nil
}

func (store *SelectedNotificationTypesMongoDBStore) Delete(userId string) (*mongo.DeleteResult, error) {
	filter := bson.M{"userId": userId}
	result, err := store.selectedTypes.DeleteOne(context.TODO(), filter)
	return result, err
}

func (store *SelectedNotificationTypesMongoDBStore) filter(filter interface{}) ([]*model.SelectedNotificationTypes, error) {
	cursor, err := store.selectedTypes.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return store.decode(cursor)
}

func (store *SelectedNotificationTypesMongoDBStore) filterOne(filter interface{}) (SelectedNotificationTypes *model.SelectedNotificationTypes, err error) {
	result := store.selectedTypes.FindOne(context.TODO(), filter)
	err = result.Decode(&SelectedNotificationTypes)
	return
}

func (store *SelectedNotificationTypesMongoDBStore) decode(cursor *mongo.Cursor) (selectedTypes []*model.SelectedNotificationTypes, err error) {
	for cursor.Next(context.TODO()) {
		var SelectedNotificationTypes model.SelectedNotificationTypes
		err = cursor.Decode(&SelectedNotificationTypes)
		if err != nil {
			return
		}
		selectedTypes = append(selectedTypes, &SelectedNotificationTypes)
	}
	err = cursor.Err()
	return
}
