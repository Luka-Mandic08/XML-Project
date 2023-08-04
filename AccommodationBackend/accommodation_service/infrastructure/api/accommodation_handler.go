package api

import (
	"accommodation_service/domain/service"
	accommodation "common/proto/accommodation_service"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccommodationHandler struct {
	accommodation.UnimplementedAccommodationServiceServer
	service *service.AccommodationService
}

func NewAccommodationHandler(service *service.AccommodationService) *AccommodationHandler {
	return &AccommodationHandler{
		service: service,
	}
}

func (handler *AccommodationHandler) Get(ctx context.Context, request *accommodation.GetRequest) (*accommodation.GetResponse, error) {
	//todo
	return nil, nil
}

func (handler *AccommodationHandler) Create(ctx context.Context, request *accommodation.CreateRequest) (*accommodation.CreateResponse, error) {
	acc := MapCreateRequestToAccommodation(request)
	acc, err := handler.service.Insert(acc)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "Unable to find user")
	}
	return &accommodation.CreateResponse{Message: "Accommodation successfully created"}, nil
}
