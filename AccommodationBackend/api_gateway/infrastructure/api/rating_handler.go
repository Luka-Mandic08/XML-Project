package api

import (
	"api_gateway/infrastructure/services"
	rating "common/proto/rating_service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
)

type RatingHandler struct {
	client rating.RatingServiceClient
}

func NewRatingHandler(client rating.RatingServiceClient) *RatingHandler {
	return &RatingHandler{client: client}
}

func (handler *RatingHandler) GetHostRatingById(ctx *gin.Context) {
	id := ctx.Param("ratingId")
	request := rating.IdRequest{Id: id}
	response, err := handler.client.GetHostRatingById(ctx, &request)
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

func (handler *RatingHandler) GetAllRatingsForHost(ctx *gin.Context) {
	id := ctx.Param("hostId")
	request := rating.IdRequest{Id: id}
	response, err := handler.client.GetAllRatingsForHost(ctx, &request)
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

func (handler *RatingHandler) GetAverageScoreForHost(ctx *gin.Context) {
	id := ctx.Param("hostId")
	request := rating.IdRequest{Id: id}
	response, err := handler.client.GetAverageScoreForHost(ctx, &request)
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

func (handler *RatingHandler) CreateHostRating(ctx *gin.Context) {
	var hostRating rating.CreateHostRatingRequest
	raw, _ := ctx.GetRawData()
	err := protojson.Unmarshal(raw, &hostRating)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response, err := handler.client.CreateHostRating(ctx, &hostRating)
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

func (handler *RatingHandler) UpdateHostRating(ctx *gin.Context) {
	var hostRating rating.HostRating
	raw, _ := ctx.GetRawData()
	err := protojson.Unmarshal(raw, &hostRating)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if services.AuthorizeId(hostRating.GuestId, ctx) {
		ctx.JSON(http.StatusUnauthorized, "Not allowed")
		return
	}

	response, err := handler.client.UpdateHostRating(ctx, &hostRating)
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

func (handler *RatingHandler) DeleteHostRating(ctx *gin.Context) {
	var request rating.DeleteRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if services.AuthorizeId(request.GuestId, ctx) {
		ctx.JSON(http.StatusUnauthorized, "Can not delete other guests ratings!")
		return
	}

	response, err := handler.client.DeleteHostRating(ctx, &request)
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

func (handler *RatingHandler) GetAccommodationRatingById(ctx *gin.Context) {
	id := ctx.Param("ratingId")
	request := rating.IdRequest{Id: id}
	response, err := handler.client.GetAccommodationRatingById(ctx, &request)
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

func (handler *RatingHandler) GetAllRatingsForAccommodation(ctx *gin.Context) {
	id := ctx.Param("accommodationId")
	request := rating.IdRequest{Id: id}
	response, err := handler.client.GetAllRatingsForAccommodation(ctx, &request)
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

func (handler *RatingHandler) GetAverageScoreForAccommodation(ctx *gin.Context) {
	id := ctx.Param("accommodationId")
	request := rating.IdRequest{Id: id}
	response, err := handler.client.GetAverageScoreForAccommodation(ctx, &request)
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

func (handler *RatingHandler) CreateAccommodationRating(ctx *gin.Context) {
	var accommodationRating rating.CreateAccommodationRatingRequest
	raw, _ := ctx.GetRawData()
	err := protojson.Unmarshal(raw, &accommodationRating)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response, err := handler.client.CreateAccommodationRating(ctx, &accommodationRating)
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

func (handler *RatingHandler) UpdateAccommodationRating(ctx *gin.Context) {
	var accommodationRating rating.AccommodationRating
	raw, _ := ctx.GetRawData()
	err := protojson.Unmarshal(raw, &accommodationRating)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if services.AuthorizeId(accommodationRating.GuestId, ctx) {
		ctx.JSON(http.StatusUnauthorized, "Not allowed")
		return
	}

	response, err := handler.client.UpdateAccommodationRating(ctx, &accommodationRating)
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

func (handler *RatingHandler) DeleteAccommodationRating(ctx *gin.Context) {
	var request rating.DeleteRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if services.AuthorizeId(request.GuestId, ctx) {
		ctx.JSON(http.StatusUnauthorized, "Can not delete other guests ratings!")
		return
	}

	response, err := handler.client.DeleteAccommodationRating(ctx, &request)
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
