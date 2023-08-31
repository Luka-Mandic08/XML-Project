package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"notification_service/domain/model"
)

type NotificationStore interface {
	GetById(notificationId primitive.ObjectID) (*model.Notification, error)
	GetAllNotificationsByUserIdAndType(userId string, selectedTypes []string) ([]*model.Notification, error)
	AcknowledgeNotification(notificationId primitive.ObjectID) (*model.Notification, error)
	CreateNotification(notification *model.Notification) (*model.Notification, error)
}
