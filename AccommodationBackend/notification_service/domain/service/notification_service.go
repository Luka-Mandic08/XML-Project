package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"notification_service/domain/model"
	"notification_service/domain/repository"
)

type NotificationService struct {
	notificationStore              repository.NotificationStore
	selectedNotificationTypesStore repository.SelectedNotificationTypesStore
}

func NewNotificationService(notificationStore repository.NotificationStore, selectedNotificationTypesStore repository.SelectedNotificationTypesStore) *NotificationService {
	return &NotificationService{
		notificationStore:              notificationStore,
		selectedNotificationTypesStore: selectedNotificationTypesStore,
	}
}

func (service *NotificationService) GetAllNotificationsByUserIdAndType(userId string) ([]*model.Notification, error) {
	selectedNotificationTypes, err := service.selectedNotificationTypesStore.GetByUserId(userId)
	if err != nil {
		return nil, err
	}
	return service.notificationStore.GetAllNotificationsByUserIdAndType(userId, selectedNotificationTypes.SelectedTypes)
}

func (service *NotificationService) InsertNotification(notification *model.Notification) (*model.Notification, error) {
	return service.notificationStore.CreateNotification(notification)
}

func (service *NotificationService) AcknowledgeNotification(notificationId primitive.ObjectID) (*model.Notification, error) {
	return service.notificationStore.AcknowledgeNotification(notificationId)
}

func (service *NotificationService) GetSelectedNotificationTypesByUserId(userId string) (*model.SelectedNotificationTypes, error) {
	return service.selectedNotificationTypesStore.GetByUserId(userId)
}

func (service *NotificationService) InsertSelectedNotificationTypes(selectedNotificationTypes *model.SelectedNotificationTypes) (*model.SelectedNotificationTypes, error) {
	return service.selectedNotificationTypesStore.Insert(selectedNotificationTypes)
}

func (service *NotificationService) UpdateSelectedNotificationTypes(userId string, selectedTypes []string) (*model.SelectedNotificationTypes, error) {
	return service.selectedNotificationTypesStore.Update(userId, selectedTypes)
}

func (service *NotificationService) DeleteSelectedNotificationTypes(userId string) (*mongo.DeleteResult, error) {
	return service.selectedNotificationTypesStore.Delete(userId)
}
