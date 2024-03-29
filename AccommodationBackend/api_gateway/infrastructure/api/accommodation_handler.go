package api

import (
	"api_gateway/infrastructure/services"
	accommodation "common/proto/accommodation_service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"strconv"
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

func (handler *AccommodationHandler) UpdateAvailability(ctx *gin.Context) {
	var acc accommodation.UpdateAvailabilityRequest
	num, _ := ctx.GetRawData()
	err := protojson.Unmarshal(num, &acc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	response, err := handler.accommodationClient.UpdateAvailability(ctx, &acc)
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
	return
}

func (handler *AccommodationHandler) CheckAvailability(ctx *gin.Context) {
	var acc accommodation.CheckAvailabilityRequest
	num, _ := ctx.GetRawData()
	err := protojson.Unmarshal(num, &acc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	response, err := handler.accommodationClient.CheckAvailability(ctx, &acc)
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
	return
}

func (handler *AccommodationHandler) Search(ctx *gin.Context) {
	var acc accommodation.SearchRequest
	num, _ := ctx.GetRawData()
	err := protojson.Unmarshal(num, &acc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	response, err := handler.accommodationClient.Search(ctx, &acc)
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
	return
}

func (handler *AccommodationHandler) GetAllByHostId(ctx *gin.Context) {
	hostId := ctx.Param("hostId")
	if services.AuthorizeId(hostId, ctx) {
		ctx.JSON(http.StatusUnauthorized, "Not allowed")
		return
	}
	request := accommodation.GetAllByHostIdRequest{HostId: hostId}
	response, err := handler.accommodationClient.GetAllByHostId(ctx, &request)
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

func (handler *AccommodationHandler) GetAll(ctx *gin.Context) {
	pageNumber, err := strconv.ParseInt(ctx.Param("pageNumber"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad page number!")
		return
	}
	request := accommodation.GetAllRequest{PageNumber: int32(pageNumber)}
	response, err := handler.accommodationClient.GetAll(ctx, &request)
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

func (handler *AccommodationHandler) GetById(ctx *gin.Context) {

	request := accommodation.GetByIdRequest{Id: ctx.Param("id")}
	response, err := handler.accommodationClient.GetById(ctx, &request)
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

func (handler *AccommodationHandler) GetAvailabilities(ctx *gin.Context) {
	var request accommodation.GetAvailabilitiesRequest
	num, _ := ctx.GetRawData()
	err := protojson.Unmarshal(num, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	response, err := handler.accommodationClient.GetAvailabilities(ctx, &request)
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
