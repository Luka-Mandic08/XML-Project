package api

import (
	notification "common/proto/notification_service"
	"github.com/golang/protobuf/ptypes"
	"notification_service/domain/model"
	"time"
)

func MapManyNotificationsToResponse(ns []*model.Notification) *notification.GetAllNotificationsByUserIdAndTypeResponse {
	var notifications []*notification.Notification

	for _, n := range ns {
		protoTimestamp, _ := ptypes.TimestampProto(n.DateCreated)
		var mappedNotification = notification.Notification{
			Id:               n.Id.Hex(),
			NotificationText: n.NotificationText,
			UserId:           n.UserId,
			IsAcknowledged:   n.IsAcknowledged,
			DateCreated:      protoTimestamp,
			Type:             n.Type,
		}
		notifications = append(notifications, &mappedNotification)
	}

	return &notification.GetAllNotificationsByUserIdAndTypeResponse{Notifications: notifications}
}

func MapNotificationToResponse(n *model.Notification) *notification.Notification {
	protoTimestamp, _ := ptypes.TimestampProto(n.DateCreated)

	var mappedNotification = notification.Notification{
		Id:               n.Id.Hex(),
		NotificationText: n.NotificationText,
		UserId:           n.UserId,
		IsAcknowledged:   n.IsAcknowledged,
		DateCreated:      protoTimestamp,
		Type:             n.Type,
	}
	return &mappedNotification
}

func MapCreateRequestToNotification(n *notification.CreateNotification) *model.Notification {
	var mappedNotification = model.Notification{
		NotificationText: n.NotificationText,
		UserId:           n.UserId,
		IsAcknowledged:   false,
		DateCreated:      time.Now(),
		Type:             n.Type,
	}
	return &mappedNotification
}

func MapSelectedNotificationTypesToResponse(n *model.SelectedNotificationTypes) *notification.SelectedNotificationTypes {
	var mappedSelectedNotificationType = notification.SelectedNotificationTypes{
		UserId:        n.UserId,
		SelectedTypes: n.SelectedTypes,
	}
	return &mappedSelectedNotificationType
}

func MapRequestToSelectedNotificationTypes(n *notification.SelectedNotificationTypes) *model.SelectedNotificationTypes {
	var mappedSelectedNotificationType = model.SelectedNotificationTypes{
		UserId:        n.UserId,
		SelectedTypes: n.SelectedTypes,
	}
	return &mappedSelectedNotificationType
}
