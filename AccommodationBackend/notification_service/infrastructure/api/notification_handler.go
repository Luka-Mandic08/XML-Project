package api

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"notification_service/domain/service"

	pb "common/proto/notification_service"
)

type NotificationHandler struct {
	pb.UnimplementedNotificationServiceServer
	service *service.NotificationService
}

func NewNotificationHandler(service *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		service: service,
	}
}

func (handler *NotificationHandler) GetAllNotificationsForHost(ctx context.Context, request *pb.IdRequestHost) (*pb.GetAllNotificationsForHostResponse, error) {
	notifications, err := handler.service.GetAllNotificationsForHost(request.HostId)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	response := MapManyNotificationsToResponse(notifications)
	return response, nil
}

func (handler *NotificationHandler) GetAllNotificationsForGuest(ctx context.Context, request *pb.IdRequestGuest) (*pb.GetAllNotificationsForHostResponse, error) {
	notifications, err := handler.service.GetAllNotificationsForGuest(request.GuestId)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	response := MapManyNotificationsToResponse(notifications)
	return response, nil
}

func (handler *NotificationHandler) InsertNotification(ctx context.Context, request *pb.CreateNotification) (*pb.Notification, error) {
	notification := MapCreateRequestToNotification(request)
	notification, err := handler.service.InsertNotification(notification)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "Unable to insert notification into database")
	}
	response := MapNotificationToResponse(notification)
	return response, nil
}

func (handler *NotificationHandler) AcknowledgeNotification(ctx context.Context, request *pb.Notification) (*pb.Notification, error) {
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}

	accommodationRating := MapToNotification(request, objectId)
	result, err := handler.service.AcknowledgeNotification(accommodationRating)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to update accommodation rating")
	}

	response := MapNotificationToResponse(result)
	return response, nil
}
