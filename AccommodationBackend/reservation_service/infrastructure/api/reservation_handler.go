package api

import (
	pb "common/proto/reservation_service"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reservation_service/domain/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationHandler struct {
	pb.UnimplementedReservationServiceServer
	reservationService *service.ReservationService
}

func NewReservationHandler(reservationService *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		reservationService: reservationService,
	}
}

func (handler *ReservationHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	reservation, err := handler.reservationService.Get(objectId)
	if err == mongo.ErrNoDocuments {
		return nil, status.Error(codes.NotFound, "Unable to find reservation: id = "+request.Id)
	}
	response := MapReservationToGetResponse(reservation)
	return response, nil
}

func (handler *ReservationHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	reservation, err := MapCreateRequestToReservation(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Unable to convert String to DateTime")
	}
	reservation, err = handler.reservationService.Insert(reservation)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "Unable to insert reservation into database")
	}

	// OVDE SE POZIVA SAGA

	response := MapReservationToCreateResponse(reservation)
	return response, nil
}

func (handler *ReservationHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	reservation, err := MapUpdateRequestToReservation(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Unable to convert String to DateTime")
	}
	result, err := handler.reservationService.Update(reservation)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to update reservation: id = "+request.Id)
	}
	if result.MatchedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find reservation: id = "+request.Id)
	}
	response := MapReservationToUpdateResponse(reservation)
	return response, nil
}

func (handler *ReservationHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	result, err := handler.reservationService.Delete(request.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to delete reservation: id = "+request.Id)
	}
	if result.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find reservation: id = "+request.Id)
	}
	response := MapReservationToDeleteResponse()
	return response, nil
}

func (handler *ReservationHandler) GetAllByUserId(ctx context.Context, request *pb.GetAllByUserIdRequest) (*pb.GetAllByUserIdResponse, error) {
	userId := request.UserId
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	reservations, err := handler.reservationService.GetAllByUserId(objectId)
	if err == mongo.ErrNoDocuments {
		return nil, status.Error(codes.NotFound, "Unable to find reservations for user: id = "+request.UserId)
	}
	response := MapReservationsToGetAllByUserIdResponse(reservations)
	return response, nil
}

func (handler *ReservationHandler) Request(ctx context.Context, request *pb.RequestRequest) (*pb.RequestResponse, error) {
	reservationRequest, err := MapRequestRequestToReservation(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Unable to convert String to DateTime")
	}
	reservationRequest, err = handler.reservationService.Insert(reservationRequest)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "Unable to insert request into database")
	}
	response := MapReservationToRequestResponse(reservationRequest)
	return response, nil
}
