package api

import (
	"auth_service/domain/service"
	"context"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "common/proto/auth_service"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (handler *AuthHandler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	account, err := handler.service.GetByUsername(request.Username)
	if err != nil {
		return nil, err
	}
	if account != nil {
		err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(request.Password))
		if err == nil {
			return LoginMapper(account), nil
		}
	}
	return nil, status.Error(codes.Unauthenticated, "Wrong username or password")

}

func (handler *AuthHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	existingAccount, _ := handler.service.GetByUsername(request.Username)
	if existingAccount != nil {
		return nil, status.Error(codes.AlreadyExists, "An account with this username already exists")
	}
	account := RegisterMapper(request)
	account, err := handler.service.Insert(account)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "Unable to insert account into database")
	}
	return &pb.RegisterResponse{Id: account.Id.String()}, nil
}

func (handler *AuthHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	account := UpdateMapper(request)
	result, err := handler.service.Update(account)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to update account")
	}
	if result.MatchedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find account")
	}
	return &pb.UpdateResponse{Message: "Account successfully updated"}, nil
}

func (handler *AuthHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	result, err := handler.service.Delete(request.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to delete account")
	}
	if result.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find account")
	}
	return &pb.DeleteResponse{Message: "Account successfully deleted"}, nil
}
