package api

import (
	"api_gateway/infrastructure/services"
	reservation "common/proto/reservation_service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
)

type ReservationHandler struct {
	client reservation.ReservationServiceClient
}

func NewReservationHandler(client reservation.ReservationServiceClient) *ReservationHandler {
	return &ReservationHandler{client: client}
}

func (handler *ReservationHandler) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	request := reservation.GetRequest{Id: id}
	response, err := handler.client.Get(ctx, &request)
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

func (handler *ReservationHandler) Create(ctx *gin.Context) {
	var reservation reservation.CreateRequest
	err := ctx.ShouldBindJSON(&reservation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	response, err := handler.client.Create(ctx, &reservation)
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

func (handler *ReservationHandler) Update(ctx *gin.Context) {
	var reservation reservation.UpdateRequest
	err := ctx.ShouldBindJSON(&reservation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	response, err := handler.client.Update(ctx, &reservation)
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

func (handler *ReservationHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	request := reservation.DeleteRequest{Id: id}
	response, err := handler.client.Delete(ctx, &request)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, grpcError.Message())
				return
			}
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (handler *ReservationHandler) GetAllByUserId(ctx *gin.Context) {
	userId := ctx.Param("id")
	request := reservation.GetAllByUserIdRequest{UserId: userId}

	if services.AuthorizeId(userId, ctx) {
		ctx.JSON(http.StatusUnauthorized, "Not allowed")
		return
	}

	response, err := handler.client.GetAllByUserId(ctx, &request)
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

func (handler *ReservationHandler) Request(ctx *gin.Context) {
	var reservation reservation.RequestRequest
	num, _ := ctx.GetRawData()
	err := protojson.Unmarshal(num, &reservation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if services.AuthorizeId(reservation.UserId, ctx) {
		ctx.JSON(http.StatusUnauthorized, "Not allowed")
		return
	}

	response, err := handler.client.Request(ctx, &reservation)
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

func (handler *ReservationHandler) Approve(ctx *gin.Context) {
	id := ctx.Param("id")
	request := reservation.ApproveRequest{Id: id}
	response, err := handler.client.Approve(ctx, &request)
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

func (handler *ReservationHandler) Deny(ctx *gin.Context) {
	id := ctx.Param("id")
	request := reservation.DenyRequest{Id: id}
	response, err := handler.client.Deny(ctx, &request)
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

func (handler *ReservationHandler) Cancel(ctx *gin.Context) {
	id := ctx.Param("id")
	request := reservation.CancelRequest{Id: id}
	response, err := handler.client.Cancel(ctx, &request)
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

func (handler *ReservationHandler) GetAllByAccommodationId(ctx *gin.Context) {
	accommodationId := ctx.Param("id")
	request := reservation.GetRequest{Id: accommodationId}

	response, err := handler.client.GetAllByAccommodationId(ctx, &request)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.AlreadyExists:
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
