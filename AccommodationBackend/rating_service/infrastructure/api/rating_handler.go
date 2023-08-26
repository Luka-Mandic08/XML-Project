package api

import (
	reservation "common/proto/reservation_service"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"rating_service/domain/service"

	pb "common/proto/rating_service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingHandler struct {
	pb.UnimplementedRatingServiceServer
	service           *service.RatingService
	reservationClient reservation.ReservationServiceClient
}

func NewRatingHandler(service *service.RatingService, reservationClient reservation.ReservationServiceClient) *RatingHandler {
	return &RatingHandler{
		service:           service,
		reservationClient: reservationClient,
	}
}

func (handler *RatingHandler) GetHostRatingById(ctx context.Context, request *pb.IdRequest) (*pb.HostRating, error) {
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}

	hostRating, err := handler.service.GetHostRatingById(objectId)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, status.Error(codes.NotFound, "Unable to find host rating.")
	}

	response := MapHostRatingToResponse(hostRating)
	return response, nil
}

func (handler *RatingHandler) GetAllRatingsForHost(ctx context.Context, request *pb.IdRequest) (*pb.GetAllRatingsForHostResponse, error) {
	hostRatings, err := handler.service.GetAllForHost(request.Id)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	response := MapManyHostRatingsToResponse(hostRatings)
	return response, nil
}

func (handler *RatingHandler) GetAverageScoreForHost(ctx context.Context, request *pb.IdRequest) (*pb.GetAverageScoreForHostResponse, error) {
	averageScore, err := handler.service.GetAverageScoreForHost(request.Id)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	response := pb.GetAverageScoreForHostResponse{
		HostId: request.Id,
		Score:  averageScore,
	}

	return &response, nil
}

func (handler *RatingHandler) CreateHostRating(ctx context.Context, request *pb.CreateHostRatingRequest) (*pb.HostRating, error) {
	_, err := handler.reservationClient.CheckIfGuestVisitedHost(ctx, &reservation.CheckPreviousReservationRequest{
		Id:      request.HostId,
		GuestId: request.GuestId,
	})
	if err != nil {
		return nil, err
	}
	hostRating := MapCreateRequestToHostRating(request)
	oldRating, _ := handler.service.GetAverageScoreForHost(hostRating.HostId)
	hostRating, err = handler.service.InsertHostRating(hostRating)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "Unable to insert host rating into database")
	}
	newRating, _ := handler.service.GetAverageScoreForHost(hostRating.HostId)
	shouldSendRequest, newStatus := CompareAverageRatings(oldRating, newRating)
	if shouldSendRequest {
		handler.reservationClient.UpdateOutstandingHostStatus(ctx, &reservation.UpdateOutstandingHostStatusRequest{HostId: request.GetHostId(), ShouldUpdate: newStatus})
	}
	response := MapHostRatingToResponse(hostRating)
	return response, nil
}

func (handler *RatingHandler) UpdateHostRating(ctx context.Context, request *pb.HostRating) (*pb.HostRating, error) {
	objectId, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		return nil, err
	}

	oldRating, _ := handler.service.GetAverageScoreForHost(request.GetId())
	hostRating := MapToHostRating(request, objectId)
	result, err := handler.service.UpdateHostRating(hostRating)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to update host rating")
	}
	if result.MatchedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find host rating")
	}

	updatedHostRating, err := handler.service.GetHostRatingById(objectId)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to find host rating")
	}
	newRating, _ := handler.service.GetAverageScoreForHost(hostRating.HostId)
	shouldSendRequest, newStatus := CompareAverageRatings(oldRating, newRating)
	if shouldSendRequest {
		handler.reservationClient.UpdateOutstandingHostStatus(ctx, &reservation.UpdateOutstandingHostStatusRequest{HostId: request.GetHostId(), ShouldUpdate: newStatus})
	}
	response := MapHostRatingToResponse(updatedHostRating)
	return response, nil
}

func (handler *RatingHandler) DeleteHostRating(ctx context.Context, request *pb.DeleteRequest) (*pb.DeletedResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(request.RatingId)
	if err != nil {
		return nil, err
	}

	rating, err := handler.service.GetHostRatingById(objectId)
	if err != nil {
		return nil, err
	}

	oldRating, _ := handler.service.GetAverageScoreForHost(rating.HostId)
	result, err := handler.service.DeleteHostRating(objectId)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to delete host rating.")
	}
	if result.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find host rating")
	}
	newRating, _ := handler.service.GetAverageScoreForHost(rating.HostId)
	shouldSendRequest, newStatus := CompareAverageRatings(oldRating, newRating)
	if shouldSendRequest {
		handler.reservationClient.UpdateOutstandingHostStatus(ctx, &reservation.UpdateOutstandingHostStatusRequest{HostId: rating.HostId, ShouldUpdate: newStatus})
	}
	return &pb.DeletedResponse{Message: "Host rating successfully deleted"}, nil
}

func (handler *RatingHandler) GetAccommodationRatingById(ctx context.Context, request *pb.IdRequest) (*pb.AccommodationRating, error) {
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}

	accommodationRating, err := handler.service.GetAccommodationRatingById(objectId)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, status.Error(codes.NotFound, "Unable to find accommodation rating.")
	}

	response := MapAccommodationRatingToResponse(accommodationRating)
	return response, nil
}

func (handler *RatingHandler) GetAllRatingsForAccommodation(ctx context.Context, request *pb.IdRequest) (*pb.GetAllRatingsForAccommodationResponse, error) {
	accommodationRatings, err := handler.service.GetAllForAccommodation(request.Id)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	response := MapManyAccommodationRatingsToResponse(accommodationRatings)
	return response, nil
}

func (handler *RatingHandler) GetAverageScoreForAccommodation(ctx context.Context, request *pb.IdRequest) (*pb.GetAverageScoreForAccommodationResponse, error) {
	averageScore, err := handler.service.GetAverageScoreForAccommodation(request.Id)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	response := pb.GetAverageScoreForAccommodationResponse{
		AccommodationId: request.Id,
		Score:           averageScore,
	}

	return &response, nil
}

func (handler *RatingHandler) CreateAccommodationRating(ctx context.Context, request *pb.CreateAccommodationRatingRequest) (*pb.AccommodationRating, error) {
	_, err := handler.reservationClient.CheckIfGuestVisitedAccommodation(ctx, &reservation.CheckPreviousReservationRequest{
		Id:      request.AccommodationId,
		GuestId: request.GuestId,
	})
	if err != nil {
		return nil, err
	}

	accommodationRating := MapCreateRequestToAccommodationRating(request)
	accommodationRating, err = handler.service.InsertAccommodationRating(accommodationRating)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "Unable to insert accommodation rating into database")
	}

	response := MapAccommodationRatingToResponse(accommodationRating)
	return response, nil
}

func (handler *RatingHandler) UpdateAccommodationRating(ctx context.Context, request *pb.AccommodationRating) (*pb.AccommodationRating, error) {
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}

	accommodationRating := MapToAccommodationRating(request, objectId)
	result, err := handler.service.UpdateAccommodationRating(accommodationRating)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to update accommodation rating")
	}
	if result.MatchedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find accommodation rating")
	}

	updatedAccommodationRating, err := handler.service.GetAccommodationRatingById(objectId)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to find host rating")
	}

	response := MapAccommodationRatingToResponse(updatedAccommodationRating)
	return response, nil
}

func (handler *RatingHandler) DeleteAccommodationRating(ctx context.Context, request *pb.DeleteRequest) (*pb.DeletedResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(request.RatingId)
	if err != nil {
		return nil, err
	}

	result, err := handler.service.DeleteAccommodationRating(objectId)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to delete host rating.")
	}
	if result.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find host rating")
	}
	return &pb.DeletedResponse{Message: "Accommodation rating successfully deleted"}, nil
}
