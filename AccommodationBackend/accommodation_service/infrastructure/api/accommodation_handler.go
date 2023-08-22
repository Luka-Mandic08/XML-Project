package api

import (
	"accommodation_service/domain/service"
	accommodation "common/proto/accommodation_service"
	rating "common/proto/rating_service"
	reservation "common/proto/reservation_service"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccommodationHandler struct {
	accommodation.UnimplementedAccommodationServiceServer
	service           *service.AccommodationService
	ratingClient      rating.RatingServiceClient
	reservationClient reservation.ReservationServiceClient
}

func NewAccommodationHandler(service *service.AccommodationService, ratingClient rating.RatingServiceClient, reservationClient reservation.ReservationServiceClient) *AccommodationHandler {
	return &AccommodationHandler{
		service:           service,
		ratingClient:      ratingClient,
		reservationClient: reservationClient,
	}
}

func (handler *AccommodationHandler) GetById(ctx context.Context, request *accommodation.GetByIdRequest) (*accommodation.GetByIdResponse, error) {
	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid ID format")
	}

	accommodation, err := handler.service.GetById(id)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	mapped := MapAccommodation(accommodation)

	return mapped, nil
}

func (handler *AccommodationHandler) Create(ctx context.Context, request *accommodation.CreateRequest) (*accommodation.Response, error) {
	acc := MapCreateRequestToAccommodation(request)
	acc, err := handler.service.Insert(acc)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "An accommodation already exists at this location")
	}
	return &accommodation.Response{
		Message: "Accommodation successfully created",
	}, nil
}

func (handler *AccommodationHandler) UpdateAvailability(ctx context.Context, request *accommodation.UpdateAvailabilityRequest) (*accommodation.Response, error) {
	response, err := handler.reservationClient.GetAllForDateRange(ctx, &reservation.GetAllForDateRangeRequest{
		From:            request.DateFrom,
		To:              request.DateTo,
		AccommodationId: request.Accommodationid,
	})
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	if len(response.GetReservations()) > 0 {
		return nil, status.Error(codes.Canceled, "Unable to update availability because there are already existing reservations for this period")
	}
	err = handler.service.UpdateAvailability(*request)
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
	accommodations, prices, numberOfDays, err := handler.service.Search(request)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	mapped := MapAccommodationsToSearchRequest(accommodations, prices, numberOfDays, int(request.NumberOfGuests), int(request.GetPageNumber()))
	for _, acc := range mapped.GetAccommodations() {
		response, _ := handler.ratingClient.GetAverageScoreForAccommodation(ctx, &rating.IdRequest{Id: acc.GetId()})
		acc.Rating = response.GetScore()
	}
	return mapped, nil
}

func (handler *AccommodationHandler) GetAllByHostId(ctx context.Context, request *accommodation.GetAllByHostIdRequest) (*accommodation.GetAllByHostIdResponse, error) {
	accommodations, err := handler.service.GetAllByHostId(request.HostId)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	mapped, _ := MapAccommodations(accommodations)

	return mapped, nil
}

func (handler *AccommodationHandler) GetAll(ctx context.Context, request *accommodation.GetAllRequest) (*accommodation.GetAllResponse, error) {
	accommodations, err := handler.service.GetAll(int(request.GetPageNumber()))
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	_, mapped := MapAccommodations(accommodations)
	for _, acc := range mapped.GetAccommodations() {
		response, _ := handler.ratingClient.GetAverageScoreForAccommodation(ctx, &rating.IdRequest{Id: acc.GetId()})
		acc.Rating = response.GetScore()
	}

	return mapped, nil
}

func (handler *AccommodationHandler) GetAvailabilities(ctx context.Context, request *accommodation.GetAvailabilitiesRequest) (*accommodation.GetAvailabilitiesResponse, error) {
	availabilities, err := handler.service.GetAvailabilitiesForAccommodation(request)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return MapAvailabilities(availabilities), nil
}

func (handler *AccommodationHandler) DeleteAllForHost(ctx context.Context, request *accommodation.GetByIdRequest) (*accommodation.Response, error) {
	_, err := handler.service.DeleteAllForHost(request.GetId())
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &accommodation.Response{Message: "Success"}, nil
}

func (handler *AccommodationHandler) CheckCanApprove(ctx context.Context, request *accommodation.CheckCanApproveRequest) (*accommodation.CheckCanApproveResponse, error) {
	response, err := handler.service.CheckCanApprove(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (handler *AccommodationHandler) GetAndCancelAllAvailabilitiesToCancel(ctx context.Context, request *accommodation.GetAndCancelAllAvailabilitiesToCancelRequest) (*accommodation.GetAndCancelAllAvailabilitiesToCancelResponse, error) {
	response, err := handler.service.GetAndCancelAllAvailabilitiesToCancel(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (handler *AccommodationHandler) GetAllForHostByAccommodationId(ctx context.Context, request *accommodation.GetByIdRequest) (*accommodation.GetAllForHostByAccommodationIdResponse, error) {
	response, err := handler.service.GetAllForHostByAccommodationId(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
