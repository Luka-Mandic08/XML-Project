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
	service *service.ReservationService
}

func NewReservationHandler(service *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		service: service,
	}
}

func (handler *ReservationHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	reservation, err := handler.service.Get(objectId)
	if err == mongo.ErrNoDocuments {
		return nil, status.Error(codes.NotFound, "Unable to find reservation")
	}
	response := MapReservationToGetResponse(reservation)
	return response, nil
}

func (handler *ReservationHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.GetResponse, error) {
	reservation, err := MapCreateRequestToReservation(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Unable to convert String to DateTime")
	}
	reservation, err = handler.service.Insert(reservation)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "Unable to insert reservation into database")
	}
	response := MapReservationToGetResponse(reservation)
	return response, nil
}

func (handler *ReservationHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.GetResponse, error) {
	reservation, err := MapUpdateRequestToReservation(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Unable to convert String to DateTime")
	}
	result, err := handler.service.Update(reservation)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to update reservation")
	}
	if result.MatchedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find reservation")
	}
	response := MapReservationToGetResponse(reservation)
	return response, nil
}

func (handler *ReservationHandler) Delete(ctx context.Context, request *pb.GetRequest) (*pb.DeleteResponse, error) {
	result, err := handler.service.Delete(request.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to delete reservation")
	}
	if result.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find reservation")
	}
	return &pb.DeleteResponse{Message: "Reservation successfully deleted"}, nil
}

func (handler *ReservationHandler) Test(ctx context.Context, request *pb.TestRequest) (*pb.TestResponse, error) {

	return &pb.TestResponse{Id: "Okej je radi!"}, nil
}
