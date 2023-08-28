package api

import (
	notification "common/proto/notification_service"
	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"notification_service/domain/model"
	"time"
)

func MapHostRatingToResponse(r *model.Notification) *notification.Notification {
	var mappedNotification = notification.Notification{
		Id:               r.Id.Hex(),
		NotificationText: r.NotificationText,
		HostId:           r.HostId,
		IsAcknowledged:   r.IsAcknowledged,
	}
	return &mappedNotification
}

func MapManyNotificationsToResponse(rs []*model.Notification) *notification.GetAllNotificationsForHostResponse {
	var notifications []*notification.Notification

	for _, r := range rs {
		protoTimestamp, _ := ptypes.TimestampProto(r.DateCreated)
		var mappedHostRating = notification.Notification{
			Id:               r.Id.Hex(),
			NotificationText: r.NotificationText,
			HostId:           r.HostId,
			IsAcknowledged:   r.IsAcknowledged,
			DateCreated:      protoTimestamp,
		}
		notifications = append(notifications, &mappedHostRating)
	}

	return &notification.GetAllNotificationsForHostResponse{Notifications: notifications}
}

func MapNotificationToResponse(r *model.Notification) *notification.Notification {
	protoTimestamp, _ := ptypes.TimestampProto(r.DateCreated)

	var mappedNotification = notification.Notification{
		Id:               r.Id.Hex(),
		NotificationText: r.NotificationText,
		HostId:           r.HostId,
		IsAcknowledged:   r.IsAcknowledged,
		DateCreated:      protoTimestamp,
	}
	return &mappedNotification
}

func MapCreateRequestToNotification(r *notification.CreateNotification) *model.Notification {
	var mappedNotification = model.Notification{
		Id:               primitive.NewObjectID(),
		NotificationText: r.NotificationText,
		HostId:           r.HostId,
		IsAcknowledged:   false,
		DateCreated:      time.Now(),
	}
	return &mappedNotification
}

func MapToNotification(r *notification.Notification, objectId primitive.ObjectID) *model.Notification {
	var mappedAccommodationRating = model.Notification{
		Id:               objectId,
		NotificationText: r.NotificationText,
		HostId:           r.HostId,
		IsAcknowledged:   false,
		DateCreated:      r.DateCreated.AsTime(),
	}
	return &mappedAccommodationRating
}
