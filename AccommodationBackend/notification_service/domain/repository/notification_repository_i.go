package repository

import (
	"notification_service/domain/model"
)

type NotificationStore interface {
	GetAllNotificationsForHost(hostId string) ([]*model.Notification, error)
	AcknowledgeNotification(notification *model.Notification) (*model.Notification, error)
	CreateNotification(notification *model.Notification) (*model.Notification, error)
}
