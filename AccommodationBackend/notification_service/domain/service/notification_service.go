package service

import (
	"notification_service/domain/model"
	"notification_service/domain/repository"
)

type NotificationService struct {
	notificationStore repository.NotificationStore
}

func NewNotificationService(notificationStore repository.NotificationStore) *NotificationService {
	return &NotificationService{
		notificationStore: notificationStore,
	}
}

func (service *NotificationService) GetAllNotificationsForHost(hostId string) ([]*model.Notification, error) {
	return service.notificationStore.GetAllNotificationsForHost(hostId)
}

func (service *NotificationService) InsertNotification(notification *model.Notification) (*model.Notification, error) {
	return service.notificationStore.CreateNotification(notification)
}

func (service *NotificationService) AcknowledgeNotification(notification *model.Notification) (*model.Notification, error) {
	return service.notificationStore.AcknowledgeNotification(notification)
}
