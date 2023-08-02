package api

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user_service/domain/service"

	pb "common/proto/user_service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
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
	response := MapUserToGetResponse(user)
	return response, nil
}

func (handler *UserHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.GetResponse, error) {
	user := MapCreateRequestToUser(request)
	user, error := handler.service.Insert(user)
	if error != nil {
		return nil, status.Error(codes.AlreadyExists, "Unable to insert user into database")
	}
	response := MapUserToGetResponse(user)
	return response, nil
}

func (handler *UserHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.GetResponse, error) {
	user := MapUpdateRequestToUser(request)
	result, error := handler.service.Update(user)
	if result.MatchedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find user")
	}
	if error != nil {
		return nil, status.Error(codes.Unknown, "Unable to update user")
	}
	response := MapUserToGetResponse(user)
	return response, nil
}

func (handler *UserHandler) Delete(ctx context.Context, request *pb.GetRequest) (*pb.DeleteResponse, error) {
	result, error := handler.service.Delete(request.Id)
	if result.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find user")
	}
	if error != nil {
		return nil, status.Error(codes.Unknown, "Unable to delete user")
	}
	return &pb.DeleteResponse{Message: "User successfully deleted"}, nil
}
