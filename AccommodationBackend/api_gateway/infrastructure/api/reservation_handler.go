package api

import (
	"api_gateway/infrastructure/services"
	reservation "common/proto/reservation_service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	/*if services.AuthorizeId(reservation.Id, ctx) {
		ctx.JSON(http.StatusUnauthorized, "Not allowed")
		return
	}*/
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
	if services.AuthorizeId(reservation.Id, ctx) {
		ctx.JSON(http.StatusUnauthorized, "Not allowed")
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

func (handler *ReservationHandler) Test(ctx *gin.Context) {
	var reservation reservation.TestRequest

	// DOVDE SVE PROLAZI ALI KAD OCE U HANDLER, ONDA PUCA
	response, err := handler.client.Test(ctx, &reservation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response)
	}

	ctx.JSON(http.StatusOK, response)
	return
}
