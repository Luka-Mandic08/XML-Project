package api

import (
	"accommodation_service/domain/service"
	accommodation "common/proto/accommodation_service"
	"context"
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

func (a AccommodationHandler) Get(ctx context.Context, request *accommodation.GetRequest) (*accommodation.GetResponse, error) {
	//todo
	return nil, nil
}
