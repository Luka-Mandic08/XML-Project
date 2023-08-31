package api

import (
	notification "common/proto/notification_service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
)

type NotificationHandler struct {
	client notification.NotificationServiceClient
}

func NewNotificationHandler(client notification.NotificationServiceClient) *NotificationHandler {
	return &NotificationHandler{client: client}
}

func (handler *NotificationHandler) GetAllNotificationsByUserIdAndType(ctx *gin.Context) {
	id := ctx.Param("userId")
	request := notification.UserIdRequest{UserId: id}
	response, err := handler.client.GetAllNotificationsByUserIdAndType(ctx, &request)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.AlreadyExists:
				ctx.JSON(http.StatusConflict, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (handler *NotificationHandler) CreateNotification(ctx *gin.Context) {
	var notificationReq notification.CreateNotification
	raw, _ := ctx.GetRawData()
	err := protojson.Unmarshal(raw, &notificationReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response, err := handler.client.InsertNotification(ctx, &notificationReq)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.Unauthenticated:
				ctx.JSON(http.StatusConflict, grpcError.Message())
				return
			default:
				ctx.JSON(http.StatusBadRequest, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (handler *NotificationHandler) AcknowledgeNotification(ctx *gin.Context) {
	var notificationReq notification.IdRequest
	raw, _ := ctx.GetRawData()
	err := protojson.Unmarshal(raw, &notificationReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response, err := handler.client.AcknowledgeNotification(ctx, &notificationReq)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			ctx.JSON(http.StatusBadRequest, grpcError.Message())
			return
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (handler *NotificationHandler) GetSelectedNotificationTypesByUserId(ctx *gin.Context) {
	id := ctx.Param("userId")
	request := notification.UserIdRequest{UserId: id}
	response, err := handler.client.GetSelectedNotificationTypesByUserId(ctx, &request)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.AlreadyExists:
				ctx.JSON(http.StatusConflict, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (handler *NotificationHandler) InsertSelectedNotificationTypes(ctx *gin.Context) {
	var request notification.SelectedNotificationTypes
	raw, _ := ctx.GetRawData()
	err := protojson.Unmarshal(raw, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response, err := handler.client.InsertSelectedNotificationTypes(ctx, &request)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.Unauthenticated:
				ctx.JSON(http.StatusConflict, grpcError.Message())
				return
			default:
				ctx.JSON(http.StatusBadRequest, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (handler *NotificationHandler) UpdateSelectedNotificationTypes(ctx *gin.Context) {
	var request notification.SelectedNotificationTypes
	raw, _ := ctx.GetRawData()
	err := protojson.Unmarshal(raw, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response, err := handler.client.UpdateSelectedNotificationTypes(ctx, &request)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.Unauthenticated:
				ctx.JSON(http.StatusConflict, grpcError.Message())
				return
			default:
				ctx.JSON(http.StatusBadRequest, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (handler *NotificationHandler) DeleteSelectedNotificationTypes(ctx *gin.Context) {
	id := ctx.Param("userId")
	request := notification.UserIdRequest{UserId: id}
	response, err := handler.client.DeleteSelectedNotificationTypes(ctx, &request)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.AlreadyExists:
				ctx.JSON(http.StatusConflict, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, response)
}
