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

func (handler *NotificationHandler) GetAllNotificationsByUserIdAndType(ctx context.Context, request *pb.UserIdRequest) (*pb.GetAllNotificationsByUserIdAndTypeResponse, error) {
	notifications, err := handler.service.GetAllNotificationsByUserIdAndType(request.GetUserId())
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

func (handler *NotificationHandler) AcknowledgeNotification(ctx context.Context, request *pb.IdRequest) (*pb.Notification, error) {
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}

	result, err := handler.service.AcknowledgeNotification(objectId)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to acknowledge notification")
	}

	response := MapNotificationToResponse(result)
	return response, nil
}

func (handler *NotificationHandler) GetSelectedNotificationTypesByUserId(ctx context.Context, request *pb.UserIdRequest) (*pb.SelectedNotificationTypes, error) {
	selectedNotificationTypes, err := handler.service.GetSelectedNotificationTypesByUserId(request.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	response := MapSelectedNotificationTypesToResponse(selectedNotificationTypes)
	return response, nil
}

func (handler *NotificationHandler) InsertSelectedNotificationTypes(ctx context.Context, request *pb.SelectedNotificationTypes) (*pb.MessageResponse, error) {
	selectedNotificationTypes := MapRequestToSelectedNotificationTypes(request)
	_, err := handler.service.InsertSelectedNotificationTypes(selectedNotificationTypes)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "Unable to insert notification into database")
	}
	return &pb.MessageResponse{Message: "Successfully created SelectedNotificationTypes"}, nil
}

func (handler *NotificationHandler) UpdateSelectedNotificationTypes(ctx context.Context, request *pb.SelectedNotificationTypes) (*pb.SelectedNotificationTypes, error) {
	result, err := handler.service.UpdateSelectedNotificationTypes(request.GetUserId(), request.GetSelectedTypes())
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to acknowledge notification")
	}

	response := MapSelectedNotificationTypesToResponse(result)
	return response, nil
}

func (handler *NotificationHandler) DeleteSelectedNotificationTypes(ctx context.Context, request *pb.UserIdRequest) (*pb.MessageResponse, error) {
	result, err := handler.service.DeleteSelectedNotificationTypes(request.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to delete host rating.")
	}
	if result.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find host rating")
	}

	return &pb.MessageResponse{Message: "SelectedNotificationTypes successfully deleted"}, nil
}
