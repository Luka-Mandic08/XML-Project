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

func (handler *AccommodationHandler) Create(ctx context.Context, request *accommodation.CreateRequest) (*accommodation.Response, error) {
	acc := MapCreateRequestToAccommodation(request)
	acc, err := handler.service.Insert(acc)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "An accommodation already exists at this location")
	}
	return &accommodation.Response{Message: "Accommodation successfully created"}, nil
}

func (handler *AccommodationHandler) UpdateAvailability(ctx context.Context, request *accommodation.UpdateAvailabilityRequest) (*accommodation.Response, error) {
	err := handler.service.UpdateAvailability(*request)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}
	return &accommodation.Response{Message: "Accommodation availability successfully updated"}, nil
}

func (handler *AccommodationHandler) CheckAvailability(ctx context.Context, request *accommodation.CheckAvailabilityRequest) (*accommodation.CheckAvailabilityResponse, error) {
	response, err := handler.service.CheckAccommodationAvailability(request)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return response, nil
}

func (handler *AccommodationHandler) Search(ctx context.Context, request *accommodation.SearchRequest) (*accommodation.SearchResponse, error) {
	accommodations, prices, numberofDays, err := handler.service.Search(request)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return MapAccommodationsToSearchRequest(accommodations, prices, numberofDays), nil
}
