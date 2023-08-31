package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"notification_service/domain/model"
)

type SelectedNotificationTypesStore interface {
	GetByUserId(userId string) (*model.SelectedNotificationTypes, error)
	Insert(selectedTypes *model.SelectedNotificationTypes) (*model.SelectedNotificationTypes, error)
	Update(userId string, selectedTypes []string) (*model.SelectedNotificationTypes, error)
	Delete(userId string) (*mongo.DeleteResult, error)
}
