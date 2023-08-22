package api

import (
	rating "common/proto/rating_service"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user_service/domain/service"

	reservation "common/proto/reservation_service"
	pb "common/proto/user_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service           *service.UserService
	reservationClient reservation.ReservationServiceClient
	ratingClient      rating.RatingServiceClient
}

func NewUserHandler(service *service.UserService, reservationClient reservation.ReservationServiceClient, ratingClient rating.RatingServiceClient) *UserHandler {
	return &UserHandler{
		service:           service,
		reservationClient: reservationClient,
		ratingClient:      ratingClient,
	}
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	user, err := handler.service.Get(objectId)
	if err == mongo.ErrNoDocuments {
		return nil, status.Error(codes.NotFound, "Unable to find user")
	}
	mapped := MapUserToGetResponse(user)
	response, _ := handler.ratingClient.GetAverageScoreForHost(ctx, &rating.IdRequest{Id: mapped.GetId()})
	mapped.Rating = response.GetScore()
	outstanding, _ := handler.reservationClient.GetOutstandingHost(ctx, &reservation.GetRequest{Id: id})
	if outstanding.GetId() != "" {
		mapped.IsOutstanding = true
	}
	return mapped, nil
}

func (handler *UserHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.GetResponse, error) {
	user := MapCreateRequestToUser(request)
	user, err := handler.service.Insert(user)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "Unable to insert user into database")
	}
	response := MapUserToGetResponse(user)
	return response, nil
}

func (handler *UserHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.GetResponse, error) {
	user := MapUpdateRequestToUser(request)
	result, err := handler.service.Update(user)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to update user")
	}
	if result.MatchedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find user")
	}
	response := MapUserToGetResponse(user)
	return response, nil
}

func (handler *UserHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	var err error
	if request.GetRole() == "Guest" {
		_, err = handler.reservationClient.CheckIfGuestHasReservations(ctx, &reservation.CheckReservationRequest{Id: request.GetId()})
	}
	if request.GetRole() == "Host" {
		_, err = handler.reservationClient.CheckIfHostHasReservations(ctx, &reservation.CheckReservationRequest{Id: request.GetId()})
	}
	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}
	result, err := handler.service.Delete(request.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to delete user")
	}
	if result.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find user")
	}
	return &pb.DeleteResponse{Message: "User successfully deleted"}, nil
}
