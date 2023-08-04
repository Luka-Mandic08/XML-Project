package api

import (
	"api_gateway/infrastructure/services"
	accommodation "common/proto/accommodation_service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type AccommodationHandler struct {
	accommodationClient accommodation.AccommodationServiceClient
}

func NewAccommodationHandler(accommodationClient accommodation.AccommodationServiceClient) *AccommodationHandler {
	return &AccommodationHandler{accommodationClient: accommodationClient}
}

func (handler *AccommodationHandler) Create(ctx *gin.Context) {
	var acc accommodation.CreateRequest
	err := ctx.ShouldBindJSON(&acc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if services.AuthorizeId(acc.HostId, ctx) {
		ctx.JSON(http.StatusUnauthorized, "Not allowed")
		return
	}

	response, err := handler.accommodationClient.Create(ctx, &acc)
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
	ctx.JSON(http.StatusCreated, response)
	return
}
