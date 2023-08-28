package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"notification_service/domain/model"
)

const (
	DATABASE               = "notification"
	NOTIFICATIONCOLLECTION = "notification"
)

type NotificationMongoDBStore struct {
	notifications *mongo.Collection
}

func NewNotificationMongoDBStore(client *mongo.Client) NotificationStore {
	notifications := client.Database(DATABASE).Collection(NOTIFICATIONCOLLECTION)
	return &NotificationMongoDBStore{
		notifications: notifications,
	}
}

func (store NotificationMongoDBStore) GetAllNotificationsForHost(hostId string) ([]*model.Notification, error) {
	filter := bson.M{"hostId": hostId}
	return store.filter(filter)
}

func (store NotificationMongoDBStore) AcknowledgeNotification(notification *model.Notification) (*mongo.UpdateResult, error) {
	update := bson.D{{"$set",
		bson.D{
			{"notificationText", notification.NotificationText},
			{"hostId", notification.HostId},
			{"isAcknowledged", true},
			{"dateCreated", notification.NotificationText},
		},
	}}
	result, err := store.notifications.UpdateByID(context.TODO(), notification.Id, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (store NotificationMongoDBStore) CreateNotification(notification *model.Notification) (*model.Notification, error) {
	result, err := store.notifications.InsertOne(context.TODO(), notification)
	if err != nil {
		return nil, err
	}
	notification.Id = result.InsertedID.(primitive.ObjectID)
	return notification, nil
}

func (store *NotificationMongoDBStore) filter(filter interface{}) ([]*model.Notification, error) {
	cursor, err := store.notifications.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return store.decode(cursor)
}

func (store *NotificationMongoDBStore) decode(cursor *mongo.Cursor) (notifications []*model.Notification, err error) {
	for cursor.Next(context.TODO()) {
		var notification model.Notification
		err = cursor.Decode(&notification)
		if err != nil {
			return
		}
		notifications = append(notifications, &notification)
	}
	err = cursor.Err()
	return
}
